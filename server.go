package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nashirkra/RENTAL-BUKU/conf"
	"github.com/nashirkra/RENTAL-BUKU/controller"
	"github.com/nashirkra/RENTAL-BUKU/middleware"
	"github.com/nashirkra/RENTAL-BUKU/repository"
	"github.com/nashirkra/RENTAL-BUKU/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = conf.SetupDBConn()
	userRep        repository.UserRepository = repository.NewUserRepository(db)
	jwtServ        service.JWTService        = service.NewJWTService()
	userServ       service.UserService       = service.NewUserService(userRep)
	authServ       service.AuthService       = service.NewAuthService(userRep)
	authController controller.AuthController = controller.NewAuthController(authServ, jwtServ)
	userController controller.UserController = controller.NewUserController(userServ, jwtServ)
)

func main() {
	r := gin.Default()

	//	authRoutes := r.Group("api/auth", middleware.AuthorizeJWT(jwtServ))
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtServ))
	{
		userRoutes.PUT("/profile", userController.Update)
		userRoutes.GET("/profile", userController.Profile)
	}
	r.Run()
}
