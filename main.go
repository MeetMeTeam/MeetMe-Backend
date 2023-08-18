package main

import (
	"context"
	"fmt"
	"meetme/be/actions/handlers"
	"meetme/be/actions/repositories"
	"meetme/be/actions/services"
	"net/http"
	"strings"
	"time"

	_ "meetme/be/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"

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

// @title Meet Me API
// @version 1.0
// @description This is a API for Meet Me.

// @host 3f9d-202-28-7-128.ngrok-free.app
// @BasePath /api
func main() {

	initConfig()
	initTimeZone()
	db := initDB()

	userRepository := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.CORS())

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://localhost:" + viper.GetString("app.port"), viper.GetString("app.ngrok")}, // กำหนดโดเมนที่ยอมรับ
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	//e.Validator = &utils.CustomValidator{Validator: validator.New()}

	e.GET("/migrate", func(c echo.Context) error {
		db.AutoMigrate(User{})
		return c.String(http.StatusOK, "Migrate DB success !")
	})

	api := e.Group("/api")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.GET("/users", userHandler.GetAllUser)
	// api.POST("/points", userHandler.AddPoints)
	// api.GET("/rewards", rewardHandler.GetRewards)
	// api.GET("/reward/:rewardID", rewardHandler.GetDetailReward)
	// api.POST("/redemption", redeemHandler.Redeem)

	e.Logger.Fatal(e.Start(":" + viper.GetString("app.port")))
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
	Firstname string    `gorm:"size:255;not null"`
	Lastname  string    `gorm:"size:255"`
	Birthday  time.Time `gorm:"type:date;not null"`
	Email     string    `gorm:"size:255;not null"`
	Password  string    `gorm:"not null"`
	Image     string    `gorm:"not null"`
}
