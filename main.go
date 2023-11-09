package main

import (
	"bytes"
	"context"
	"fmt"
	header "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html/template"
	"log"
	"meetme/be/actions/handlers"
	"meetme/be/actions/repositories"
	"meetme/be/actions/services"
	_ "meetme/be/docs"
	"net/smtp"
	"os"
	"time"
)

var (
	router *mux.Router
	Server *gosocketio.Server
)

var auth smtp.Auth

//	@title			Meet Me API
//	@version		1.0
//	@description	This is a API for Meet Me.

// @host		meetme-backend.com
// @BasePath	/api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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
	//initMail()
	db := initDB()

	userRepository := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	//friendRepository := repositories.NewFriendshipRepositoryDB(db)

	friendRepository := repositories.NewFriendRepositoryDB(db)
	friendService := services.NewFriendService(friendRepository, userRepository)
	friendHandler := handlers.NewFriendHandler(friendService)
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

	userApi := api.Group("/users")
	userApi.GET("", userHandler.GetAllUser)
	userApi.POST("/register", userHandler.Register)
	userApi.POST("/login", userHandler.Login)

	inviteApi := api.Group("/invitations")
	inviteApi.POST("", friendHandler.InviteFriend)
	inviteApi.GET("", friendHandler.CheckFriendInvite)
	inviteApi.DELETE("/:inviteId", friendHandler.RejectFriend)
	inviteApi.DELETE("", friendHandler.RejectAllFriend)
	inviteApi.PUT("/:inviteId", friendHandler.AcceptFriend)
	inviteApi.PUT("", friendHandler.AcceptAllFriend)

	friendApi := api.Group("/friends")
	friendApi.GET("", friendHandler.FriendList)

	e.Logger.Fatal(e.Start(":"+os.Getenv("APP_PORT")), header.CORS(headers, methods, origins)(e))
}

func initConfig() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load env file", err)
	}
	log.Print("env successfully loaded.")

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
	opts := options.Client().ApplyURI("mongodb+srv://" + os.Getenv("MONGO_USERNAME") + ":" + os.Getenv("MONGO_PASSWORD") + ".@cluster0.salidj6.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	//opts := options.Client().ApplyURI("mongodb+srv://" + url.QueryEscape(viper.GetString("mongodb.username")) + ":" + url.QueryEscape(viper.GetString("mongodb.password")) + "@meetme.wlhqxcx.mongodb.net/?maxPoolSize=100").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("MeetMe").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client.Database("MeetMe")
}

func initMail() {
	//config.ConnectMailer(
	//	os.Getenv("MAILER_HOST"),
	//	os.Getenv("MAILER_USERNAME"),
	//	os.Getenv("MAILER_PASSWORD"),
	//)
	//
	//m := services.Mailer{}
	//message := gomail.NewMessage()
	//message.SetHeader("To", "kanyapat.winnerkypt@mail.kmutt.ac.th")
	//message.SetHeader("Subject", "Hello! Saharak")
	//message.SetBody("text/html", "ทดสอบการส่ง Email ด้วย Golang <br> สวัสดี Saharak Manoo!")
	//m.Send(message)

	auth = smtp.PlainAuth("", os.Getenv("MAILER_USERNAME"), os.Getenv("MAILER_PASSWORD"), os.Getenv("MAILER_HOST"))
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "Dhanush",
		URL:  "http://geektrust.in",
	}
	r := NewRequest([]string{"kanyapat.winnerkypt@mail.kmutt.ac.th"}, "Hello Junk!", "Hello, World!")
	err := r.ParseTemplate("verifyFile.html", templateData)
	if err == nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	}
}

// Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := os.Getenv("MAILER_HOST") + ":" + os.Getenv("MAILER_PORT")

	if err := smtp.SendMail(addr, auth, os.Getenv("MAILER_USERNAME"), r.to, msg); err != nil {
		log.Print(err)
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
