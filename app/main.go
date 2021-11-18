package main

import (
	"log"
	"time"
	"twilux/app/routes"
	userUsecase "twilux/business/users"
	userController "twilux/controllers/users"
	userRepo "twilux/drivers/databases/users"

	snippetRepo "twilux/drivers/databases/snippets"
	"twilux/drivers/mysql"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile("app/config/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&userRepo.User{})
	db.AutoMigrate(&snippetRepo.Snippet{})
}

func main() {
	configDb := mysql.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDb.InitialDb()
	dbMigrate(db)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()
	userRepoInterface := userRepo.NewUserRepository(db)
	userUseCaseInterface := userUsecase.NewUsecase(userRepoInterface, timeoutContext)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController: *userControllerInterface,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}
