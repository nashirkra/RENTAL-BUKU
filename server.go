package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nashirkra/RENTAL-BUKU/conf"
	"github.com/nashirkra/RENTAL-BUKU/controller"
	"github.com/nashirkra/RENTAL-BUKU/repository"
	"github.com/nashirkra/RENTAL-BUKU/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = conf.SetupDBConn()
	userRep        repository.UserRepository = repository.NewUserRepository(db)
	jwtServ        service.JWTService        = service.NewJWTService()
	authServ       service.AuthService       = service.NewAuthService(userRep)
	authController controller.AuthController = controller.NewAuthController(authServ, jwtServ)
)

func main() {
	r := gin.Default()

	//	authRoutes := r.Group("api/auth", middleware.AuthorizeJWT(jwtServ))
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run()
}
