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
	// setup Database Connection
	db        *gorm.DB                         = conf.SetupDBConn()
	userRepo  repository.UserRepository        = repository.NewUserRepository(db)
	catRepo   repository.CategoryRepository    = repository.NewCategoryRepository(db)
	bookRepo  repository.BookRepository        = repository.NewBookRepository(db)
	loanRepo  repository.LoanRepository        = repository.NewLoanRepository(db)
	finePRepo repository.FinePaymentRepository = repository.NewFinePaymentRepository(db)
	//setup services
	jwtServ   service.JWTService         = service.NewJWTService()
	userServ  service.UserService        = service.NewUserService(userRepo)
	catServ   service.CategoryService    = service.NewCategoryService(catRepo)
	bookServ  service.BookService        = service.NewBookService(bookRepo)
	loanServ  service.LoanService        = service.NewLoanService(loanRepo)
	finePServ service.FinePaymentService = service.NewFinePaymentService(finePRepo)
	authServ  service.AuthService        = service.NewAuthService(userRepo)
	//setup controller
	authController  controller.AuthController        = controller.NewAuthController(authServ, jwtServ)
	userController  controller.UserController        = controller.NewUserController(userServ, jwtServ)
	catController   controller.CategoryController    = controller.NewCategoryController(catServ, jwtServ)
	bookController  controller.BookController        = controller.NewBookController(bookServ, jwtServ)
	loanController  controller.LoanController        = controller.NewLoanController(loanServ, jwtServ)
	finePController controller.FinePaymentController = controller.NewFinePaymentController(finePServ, jwtServ)
)

func main() {
	r := gin.Default()

	// AuthRoutes := r.Group("api/auth", middleware.AuthorizeJWT(jwtServ))
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	// General Routes
	anyRoutes := r.Group("api")
	{
		anyRoutes.GET("/book", bookController.All)
	}
	// User Routes
	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtServ))
	{
		userRoutes.PUT("/profile", userController.Update)
		userRoutes.GET("/profile", userController.Profile)
	}
	// Admin Routes
	adminRoutes := r.Group("api/admin", middleware.AuthorizeJWT(jwtServ))
	{
		adminRoutes.PUT("/profile", userController.Update)
		adminRoutes.GET("/profile", userController.Profile)
		adminRoutes.GET("/users", userController.All)

		// Category handler
		adminRoutes.GET("/cat", catController.All)
		adminRoutes.GET("/cat/:id", catController.FindByID)
		adminRoutes.POST("/cat", catController.Insert)
		adminRoutes.PUT("/cat", catController.Update)
		adminRoutes.DELETE("/cat/:id", catController.Delete)

		// Book handler
		adminRoutes.GET("/book", bookController.All)
		adminRoutes.GET("/book/:id", bookController.FindByID)
		adminRoutes.POST("/book", bookController.Insert)
		adminRoutes.PUT("/book", bookController.Update)
		adminRoutes.DELETE("/book/:id", bookController.Delete)

		// Loan handler
		adminRoutes.GET("/loan", loanController.All)
		adminRoutes.GET("/loan/:id", loanController.FindByID)
		adminRoutes.POST("/loan", loanController.Insert)
		adminRoutes.PUT("/loan", loanController.Update)
		adminRoutes.PUT("/return", loanController.ReturnBook)
		adminRoutes.DELETE("/loan/:id", loanController.Delete)

		// FinePayment handler
		adminRoutes.GET("/pay", finePController.All)
		adminRoutes.GET("/pay/:id", finePController.FindByID)
		adminRoutes.POST("/pay", finePController.Insert)
		adminRoutes.PUT("/pay", finePController.Update)
		adminRoutes.DELETE("/pay/:id", finePController.Delete)
	}
	r.Run()
}
