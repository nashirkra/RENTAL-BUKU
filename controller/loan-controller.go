package controller

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nashirkra/RENTAL-BUKU/dto"
	"github.com/nashirkra/RENTAL-BUKU/entity"
	"github.com/nashirkra/RENTAL-BUKU/helper"
	"github.com/nashirkra/RENTAL-BUKU/service"
)

type LoanController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	ReturnBook(context *gin.Context)
}

type loanController struct {
	loanService service.LoanService
	jwtService  service.JWTService
}

func NewLoanController(loanServ service.LoanService, jwtServ service.JWTService) LoanController {
	return &loanController{
		loanService: loanServ,
		jwtService:  jwtServ,
	}
}

func (c *loanController) All(context *gin.Context) {
	var loans []entity.Loan = c.loanService.All()
	res := helper.BuildResponse(true, "OK", loans)
	context.JSON(http.StatusOK, res)
}
func (c *loanController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var loan entity.Loan = c.loanService.FindByID(id)
	if (loan == entity.Loan{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", loan)
		context.JSON(http.StatusOK, res)
	}
}
func (c *loanController) Insert(context *gin.Context) {
	var loanCreate dto.LoanCreate
	errDTO := context.ShouldBind(&loanCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		token, errToken := c.jwtService.ValidateToken(authHeader)
		if errToken != nil {
			panic(errToken.Error())
		}
		claims := token.Claims.(jwt.MapClaims)
		userID := fmt.Sprintf("%v", claims["user_id"])
		if c.loanService.UserRole(userID) != "admin" {
			response := helper.BuildErrorResponse("You dont have permission", "You are not Administrator", helper.EmptyObj{})
			context.JSON(http.StatusForbidden, response)
		} else {
			if c.loanService.CheckBookStock(loanCreate.BookID) {
				borrowDate := time.Now()
				// determine the time according to the policy
				// for example 7 days
				dueDate := time.Now().AddDate(0, 0, 7)
				loanCreate.DueDate = &dueDate
				loanCreate.BorrowedDate = &borrowDate
				result := c.loanService.Insert(loanCreate)
				response := helper.BuildResponse(true, "OK", result)
				context.JSON(http.StatusCreated, response)
			} else {
				res := helper.BuildErrorResponse("Failed to process request", "Out of Stock", helper.EmptyObj{})
				context.JSON(http.StatusBadRequest, res)
			}
		}
	}
}
func (c *loanController) Update(context *gin.Context) {
	var loanToUpdate dto.LoanUpdate
	errDTO := context.ShouldBind(&loanToUpdate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.loanService.UserRole(userID) != "admin" {
		response := helper.BuildErrorResponse("You dont have permission", "You are not Administrator", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	} else {
		result := c.loanService.Update(loanToUpdate)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}
func (c *loanController) Delete(context *gin.Context) {
	var loan entity.Loan
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	loan.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.loanService.UserRole(userID) != "admin" {
		response := helper.BuildErrorResponse("You dont have permission", "You are not Administrator", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	} else {
		c.loanService.Delete(loan)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
}

func (c *loanController) ReturnBook(context *gin.Context) {
	var loanReturn dto.LoanUpdate
	errDTO := context.ShouldBind(&loanReturn)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.loanService.UserRole(userID) != "admin" {
		response := helper.BuildErrorResponse("You dont have permission", "You are not Administrator", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	} else {
		var loan entity.Loan = c.loanService.FindByID(loanReturn.ID)

		dateNow := time.Now()
		loanReturn.ReturnDate = &dateNow
		loanReturn.DueDate = loan.DueDate
		loanReturn.BorrowedDate = loan.BorrowedDate
		loanReturn.BookID = loan.BookID
		loanReturn.UserID = loan.UserID
		// give tolerance 1 day from due date
		dueDate := loan.DueDate.AddDate(0, 0, 1)
		// if the book has not been returned
		if nil == loan.ReturnDate {

			if dateNow.After(dueDate) {
				// logic for late
				xdays := math.Round(dateNow.Sub(*loanReturn.DueDate).Hours() / 24)
				if c.loanService.FinePayment(loanReturn.ID, xdays) {
					result := c.loanService.ReturnBook(loanReturn)
					response := helper.BuildResponse(true, "OK", result)
					context.JSON(http.StatusOK, response)
				} else {

					response := helper.BuildErrorResponse("Failed pay", "Failed Fine Payment", helper.EmptyObj{})
					context.JSON(http.StatusBadRequest, response)
				}
			} else {
				// normal return date
				result := c.loanService.ReturnBook(loanReturn)
				response := helper.BuildResponse(true, "OK", result)
				context.JSON(http.StatusOK, response)
			}
		}
	}
}
