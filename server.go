package main

import (
	"log"

	"github.com/KevinBalanta/VenturitMovieAPI/libs"
	"github.com/KevinBalanta/VenturitMovieAPI/routes"
	"github.com/KevinBalanta/VenturitMovieAPI/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	config "github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

func init() {

	config.AddConfigPath("./configs")
	config.SetConfigName("mysql")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading file, %s", err)
	}

	dbConfig := libs.DBConfig{}
	dbConfig.Host = config.GetString("default.host")
	dbConfig.Port = config.GetString("default.port")
	dbConfig.Database = config.GetString("default.database")
	dbConfig.User = config.GetString("default.user")
	dbConfig.Password = config.GetString("default.password")
	dbConfig.Charset = config.GetString("default.charset")
	dbConfig.MaxIdleConns = config.GetInt("default.MaxIdleConns")
	dbConfig.MaxOpenConns = config.GetInt("default.MaxOpenConns")

	libs.DB = dbConfig.InitDB()
	if config.GetBool("default.sql_log") {
		libs.DB.Logger.LogMode(logger.Info)
	} else {
		libs.DB.Logger.LogMode(logger.Silent)
	}

	config.AddConfigPath("./")
	config.SetConfigFile(".env")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading env file, %s", err)
	}

	services.ImdbServiceObj.InitImdb(config.GetString("IMDB_API_KEY"))

}

func main() {
	e := echo.New()

	//Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Routes
	routes.MovieRouterObj.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
