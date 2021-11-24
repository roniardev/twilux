package main

import (
	"log"
	"time"
	"twilux/app/routes"
	"twilux/middlewares"

	userUsecase "twilux/business/users"
	userController "twilux/controllers/users"
	userRepo "twilux/drivers/databases/users"

	snippetUsecase "twilux/business/snippets"
	snippetController "twilux/controllers/snippets"
	snippetRepo "twilux/drivers/databases/snippets"

	savedUsecase "twilux/business/saved"
	savedController "twilux/controllers/saved"
	savedRepo "twilux/drivers/databases/saved"

	commentUsecase "twilux/business/comments"
	commentController "twilux/controllers/comments"
	commentRepo "twilux/drivers/databases/comments"

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
	db.AutoMigrate(&savedRepo.Saved{})
	db.AutoMigrate(&commentRepo.Comment{})
}

func main() {
	configDb := mysql.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	db := configDb.InitialDb()
	dbMigrate(db)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	userRepoInterface := userRepo.NewUserRepository(db)
	userUseCaseInterface := userUsecase.NewUsecase(configJWT, userRepoInterface, timeoutContext)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	snippetRepoInterface := snippetRepo.NewSnippetRepository(db)
	snippetUseCaseInterface := snippetUsecase.NewUsecase(snippetRepoInterface, timeoutContext)
	snippetControllerInterface := snippetController.NewSnippetController(snippetUseCaseInterface)

	savedRepoInterface := savedRepo.NewSavedRepository(db)
	savedUseCaseInterface := savedUsecase.NewUsecase(savedRepoInterface, timeoutContext)
	savedControllerInterface := savedController.NewSavedController(savedUseCaseInterface)

	commentRepoInterface := commentRepo.NewCommentRepository(db)
	commentUseCaseInterface := commentUsecase.NewUsecase(commentRepoInterface, timeoutContext)
	commentControllerInterface := commentController.NewCommentController(commentUseCaseInterface)

	routesInit := routes.RouteControllerList{
		JwtConfig:         configJWT.Init(),
		UserController:    *userControllerInterface,
		SnippetController: *snippetControllerInterface,
		SavedController:   *savedControllerInterface,
		CommentController: *commentControllerInterface,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}
