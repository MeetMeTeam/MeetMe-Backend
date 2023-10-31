package main

import (
	"context"
	"fmt"
	header "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"meetme/be/actions/handlers"
	"meetme/be/actions/repositories"
	"meetme/be/actions/services"
	_ "meetme/be/docs"
	"strings"
	"time"
)

var (
	router *mux.Router
	Server *gosocketio.Server
)

// @title Meet Me API
// @version 1.0
// @description This is a API for Meet Me.

// @host  c001-202-28-7-5.ngrok-free.app
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

	initConfig()
	initTimeZone()
	db := initDB()

	userRepository := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	//friendRepository := repositories.NewFriendshipRepositoryDB(db)
	//
	//inviteRepository := repositories.NewFriendInvitationRepositoryDB(db)
	//inviteService := services.NewFriendInvitationService(inviteRepository, userRepository, friendRepository)
	//inviteHandler := handlers.NewFriendInvitationHandler(inviteService)
	//
	//friendService := services.NewFriendShipService(friendRepository, userRepository)
	//friendHandler := handlers.NewFriendShipHandler(friendService)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	//e.Validator = &utils.CustomValidator{Validator: validator.New()}

	//e.GET("/migrate", func(c echo.Context) error {
	//	db.AutoMigrate(User{}, FriendInvitation{}, Friendship{})
	//	return c.String(http.StatusOK, "Migrate DB success !")
	//})

	api := e.Group("/api")

	api.POST("/register", userHandler.Register)
	//api.POST("/login", userHandler.Login)
	api.GET("/users", userHandler.GetAllUser)

	//api.GET("/friends", friendHandler.FriendList)
	//
	//inviteApi := api.Group("/invitation")
	//inviteApi.POST("/add", inviteHandler.InviteFriend)
	//inviteApi.GET("/check", inviteHandler.CheckFriendInvite)
	//inviteApi.DELETE("/rejected/:inviteId", inviteHandler.RejectFriend)
	//inviteApi.POST("/accept/:inviteId", inviteHandler.AcceptFriend)

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

func initDB() *mongo.Database {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://MeetMeUser:Ntw171044.@cluster0.salidj6.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	//opts := options.Client().ApplyURI("mongodb+srv://" + url.QueryEscape(viper.GetString("mongodb.username")) + ":" + url.QueryEscape(viper.GetString("mongodb.password")) + "@meetme.wlhqxcx.mongodb.net/?maxPoolSize=100").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("meet-me").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client.Database("meet-me")
}
