package main

import (
	"context"
	"fmt"
	header "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	"meetme/be/actions/handlers"
	"meetme/be/actions/repositories"
	"meetme/be/actions/services"
	_ "meetme/be/docs"
	"net/http"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n ======================================\n", sql)
}

var (
	router *mux.Router
	Server *gosocketio.Server

	connectedUsers = make(map[string]ConnectedUser)
)

type Message struct {
	Text string `json:"message"`
}

func init() {
	Server = gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
	fmt.Println("Socket Inititalize...")
}

func LoadSocket() {
	// socket connection
	Server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {

		fmt.Println("Connected", c.Id())
		//fmt.Println("Connected", c)
		addNewConnectedUser(c.Id())
		c.Join("Room")
	})

	// socket disconnection
	Server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		//fmt.Println("Disconnected", c.Id())
		removeConnectedUser(c.Id())
		// handles when someone closes the tab
		c.Leave("Room")
	})

	// chat socket
	Server.On("/chat", func(c *gosocketio.Channel, message Message) string {
		fmt.Println(message.Text)
		c.BroadcastTo("Room", "/message", message.Text)
		return "message sent successfully."
	})
}

func CreateRouter() {
	router = mux.NewRouter()
}

func InititalizeRoutes() {
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})
	router.Handle("/socket.io/", Server)
}

func StartServer() {
	fmt.Println("Server Started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", header.CORS(header.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}), header.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), header.AllowedOrigins([]string{"*"}))(router)))
}

type ConnectedUser struct {
	UserID string
}

func addNewConnectedUser(userID string) {
	newUser := ConnectedUser{UserID: userID}
	connectedUsers[userID] = newUser
	fmt.Println("new connected users")
	fmt.Println(connectedUsers)
}

func removeConnectedUser(userID string) {
	if _, ok := connectedUsers[userID]; ok {
		delete(connectedUsers, userID)

	}
	fmt.Println("users disconnect")
	fmt.Println(connectedUsers)
}

// @title Meet Me API
// @version 1.0
// @description This is a API for Meet Me.

// @host 3f9d-202-28-7-128.ngrok-free.app
// @BasePath /api
func main() {

	e := echo.New()

	headers := header.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := header.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := header.AllowedOrigins([]string{"*"})

	e.Use(middleware.CORS())

	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"http://localhost:3000"}, // กำหนดโดเมนที่ยอมรับ
	//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//}))

	//server := socketio.NewServer(nil)
	//
	//server.OnConnect("/", func(s socketio.Conn) error {
	//	s.SetContext("")
	//	fmt.Println("connected:")
	//	fmt.Println("connected:", s.ID())
	//	return nil
	//})
	//
	//go server.Serve()
	//defer server.Close()
	//

	//connectedUsers := make(map[string]ConnectedUser) // สร้าง map โดยกำหนด key เป็น string และ value เป็น ConnectedUser struct

	// เพิ่มข้อมูลเข้าไปใน map
	//addNewConnectedUser(connectedUsers, "socket1", "user1")
	//addNewConnectedUser(connectedUsers, "socket2", "user2")

	// แสดงผลค่าใน map
	//fmt.Println(connectedUsers)
	//LoadSocket()
	//CreateRouter()
	//InititalizeRoutes()
	//StartServer()

	//log.Fatal(http.ListenAndServe(":8080", header.CORS(headers, methods, origins)(e)))

	initConfig()
	initTimeZone()
	db := initDB()

	userRepository := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	friendRepository := repositories.NewFriendshipRepositoryDB(db)

	inviteRepository := repositories.NewFriendInvitationRepositoryDB(db)
	inviteService := services.NewFriendInvitationService(inviteRepository, userRepository, friendRepository)
	inviteHandler := handlers.NewFriendInvitationHandler(inviteService)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	//e.Validator = &utils.CustomValidator{Validator: validator.New()}

	e.GET("/migrate", func(c echo.Context) error {
		db.AutoMigrate(User{}, FriendInvitation{}, Friendship{})
		return c.String(http.StatusOK, "Migrate DB success !")
	})

	api := e.Group("/api")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.GET("/users", userHandler.GetAllUser)

	inviteApi := api.Group("/invitation")
	inviteApi.POST("/add", inviteHandler.InviteFriend)
	inviteApi.GET("/check/:receiverId", inviteHandler.CheckFriendInvite)
	inviteApi.POST("/rejected", inviteHandler.RejectFriend)
	inviteApi.POST("/accept", inviteHandler.AcceptFriend)
	// api.GET("/rewards", rewardHandler.GetRewards)
	// api.GET("/reward/:rewardID", rewardHandler.GetDetailReward)
	// api.POST("/redemption", redeemHandler.Redeem)
	//e.Logger.Fatal(http.ListenAndServe(":"+viper.GetString("app.port"), header.CORS(headers, methods, origins)(e)))
	e.Logger.Fatal(e.Start(":"+viper.GetString("app.port")), header.CORS(headers, methods, origins)(e))
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	//dsn := "root:root@tcp(127.0.0.1:8889)/erc?parseTime=true"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}

	return db
}

type User struct {
	gorm.Model
	Username  string    `gorm:"size:255;not null"`
	Firstname string    `gorm:"size:255;not null"`
	Lastname  string    `gorm:"size:255"`
	Birthday  time.Time `gorm:"type:date;not null"`
	Email     string    `gorm:"size:255;not null"`
	Password  string    `gorm:"not null"`
	Image     string    `gorm:"not null"`
}

type FriendInvitation struct {
	gorm.Model
	SenderId   int `gorm:"not null"`
	ReceiverId int `gorm:"not null"`
}

type Friendship struct {
	ID       int            `gorm:"autoIncrement"`
	UserId1  int            `gorm:"not null"`
	UserID2  int            `gorm:"not null"`
	DateAdd  time.Time      `gorm:"autoCreateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}
