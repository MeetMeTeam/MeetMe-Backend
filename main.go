package main

import (
	"context"
	"fmt"
	header "github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"meetme/be/actions/handlers"
	"meetme/be/actions/repositories"
	"meetme/be/actions/services"

	_ "meetme/be/docs"
	"os"
	"time"
)

//	@title			Meet Me API
//	@version		1.0
//	@description	This is a API for Meet Me.

// @host		localhost:8080
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
	db := initDB()

	avatarRepo := repositories.NewAvatarRepositoryDB(db)
	bgRepo := repositories.NewBgRepositoryDB(db)
	themeRepo := repositories.NewThemeRepositoryDB(db)
	inventoryRepo := repositories.NewInventoryRepositoryDB(db)
	userRepository := repositories.NewUserRepositoryDB(db)
	friendRepository := repositories.NewFriendRepositoryDB(db)
	favoriteRepository := repositories.NewFavoriteRepositoryDB(db)
	questionRepository := repositories.NewQuestionRepositoryDB(db)

	avatarService := services.NewAvatarService(avatarRepo, userRepository, inventoryRepo)
	bgService := services.NewBgService(bgRepo, userRepository, inventoryRepo)
	themeService := services.NewThemeService(themeRepo, userRepository, inventoryRepo)
	inventoryService := services.NewInventoryService(inventoryRepo, userRepository, avatarRepo, themeRepo, bgRepo)
	userService := services.NewUserService(userRepository, inventoryRepo, avatarRepo, favoriteRepository, bgRepo)
	friendService := services.NewFriendService(friendRepository, userRepository, avatarRepo, inventoryRepo)
	favoriteService := services.NewFavoriteService(userRepository, favoriteRepository)
	questionService := services.NewQuestionService(questionRepository)

	avatarHandler := handlers.NewAvatarShopHandler(avatarService)
	bgHandler := handlers.NewBgShopHandler(bgService)
	themeHandler := handlers.NewThemeShopHandler(themeService)
	inventoryHandler := handlers.NewInventoryHandler(inventoryService)
	userHandler := handlers.NewUserHandler(userService)
	friendHandler := handlers.NewFriendHandler(friendService)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteService)
	questionHandler := handlers.NewQuestionHandler(questionService)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	api := e.Group("/api")
	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/refresh", userHandler.RefreshToken)
	api.POST("/verify-mail", userHandler.SendVerifyMail)

	avatarApi := api.Group("/avatars")
	avatarApi.GET("", avatarHandler.GetAvatarShop)
	avatarApi.POST("", avatarHandler.AddAvatarToShop)

	bgApi := api.Group("/backgrounds")
	bgApi.GET("", bgHandler.GetBgShop)
	bgApi.POST("", bgHandler.AddBgToShop)

	themeApi := api.Group("/themes")
	themeApi.POST("", themeHandler.AddThemeToShop)
	themeApi.GET("", themeHandler.GetThemeShop)

	inventoryApi := api.Group("/inventories")
	inventoryApi.GET("", inventoryHandler.GetInventory)
	inventoryApi.POST("", inventoryHandler.AddItem)

	userApi := api.Group("/users")

	userApi.PUT("", userHandler.EditUserInfo)

	userApi.PUT("/forgot-password", userHandler.SendMailForResetPassword)
	userApi.PUT("/reset-password", userHandler.ChangePassword)
	userApi.GET("/coins", userHandler.GetCoins)
	userApi.GET("/avatars/:userId", userHandler.GetAvatarsByUserId)
	userApi.PUT("/avatars/:itemId", userHandler.ChangeAvatar)
	userApi.GET("/backgrounds/:userId", userHandler.GetBgByUserId)
	userApi.PUT("/backgrounds/:itemId", userHandler.ChangeBg)

	inviteApi := api.Group("/invitations")
	inviteApi.POST("", friendHandler.InviteFriend)
	inviteApi.GET("", friendHandler.CheckFriendInvite)
	inviteApi.DELETE("/:inviteId", friendHandler.RejectFriend)
	inviteApi.DELETE("", friendHandler.RejectAllFriend)
	inviteApi.PUT("/:inviteId", friendHandler.AcceptFriend)
	inviteApi.PUT("", friendHandler.AcceptAllFriend)

	friendApi := api.Group("/friends")
	friendApi.GET("", friendHandler.FriendList)
	friendApi.DELETE("/:friendId", friendHandler.RemoveFriend)

	favApi := userApi.Group("/favorites")
	favApi.POST("/:userId", favoriteHandler.FavUser)
	favApi.DELETE("/:userId", favoriteHandler.UnFavUser)
	favApi.GET("", favoriteHandler.CountFavUser)

	quesApi := api.Group("/questions")
	quesApi.GET("", questionHandler.GetQuestions)
	quesApi.POST("", questionHandler.CreateQuestion)
	quesApi.GET("/categories", questionHandler.GetCategories)

	e.Logger.Fatal(e.Start(":"+os.Getenv("APP_PORT")), header.CORS(headers, methods, origins)(e))
}

func initConfig() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error load env file", err)
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
	opts := options.Client().ApplyURI("mongodb+srv://" +
		os.Getenv("MONGO_USERNAME") + ":" +
		os.Getenv("MONGO_PASSWORD") +
		".@cluster0.salidj6.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database(os.Getenv("MONGO_DATABASE")).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client.Database(os.Getenv("MONGO_DATABASE"))
}
