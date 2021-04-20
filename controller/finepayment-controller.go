package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nashirkra/RENTAL-BUKU/dto"
	"github.com/nashirkra/RENTAL-BUKU/entity"
	"github.com/nashirkra/RENTAL-BUKU/helper"
	"github.com/nashirkra/RENTAL-BUKU/service"
)

type FinePaymentController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type finePaymentController struct {
	finePService service.FinePaymentService
	jwtService   service.JWTService
}

func NewFinePaymentController(finePServ service.FinePaymentService, jwtServ service.JWTService) FinePaymentController {
	return &finePaymentController{
		finePService: finePServ,
		jwtService:   jwtServ,
	}
}

func (c *finePaymentController) All(context *gin.Context) {
	var loans []entity.FinePayment = c.finePService.All()
	res := helper.BuildResponse(true, "OK", loans)
	context.JSON(http.StatusOK, res)
}
func (c *finePaymentController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var fineP entity.FinePayment = c.finePService.FindByID(id)
	if (fineP == entity.FinePayment{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", fineP)
		context.JSON(http.StatusOK, res)
	}
}
func (c *finePaymentController) Insert(context *gin.Context) {
	var finePCreate dto.FinePaymentCreate
	errDTO := context.ShouldBind(&finePCreate)
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
		if c.finePService.UserRole(userID) != "admin" {
			response := helper.BuildErrorResponse("You dont have permission", "You are not Administrator", helper.EmptyObj{})
			context.JSON(http.StatusForbidden, response)
		} else {
			result := c.finePService.Insert(finePCreate)
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusCreated, response)
		}
	}
}
func (c *finePaymentController) Update(context *gin.Context) {
	var finePUpdate dto.FinePaymentUpdate
	errDTO := context.ShouldBind(&finePUpdate)
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
	if c.finePService.UserRole(userID) != "admin" {
		response := helper.BuildErrorResponse("You dont have permission", "You are not Administrator", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	} else {
		result := c.finePService.Update(finePUpdate)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}
func (c *finePaymentController) Delete(context *gin.Context) {
	var fineP entity.FinePayment
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	fineP.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.finePService.UserRole(userID) != "admin" {
		response := helper.BuildErrorResponse("You dont have permission", "You are not Administrator", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	} else {
		c.finePService.Delete(fineP)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
}
