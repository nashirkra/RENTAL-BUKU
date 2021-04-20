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

type CategoryController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type categoryController struct {
	catService service.CategoryService
	jwtService service.JWTService
}

func NewCategoryController(catServ service.CategoryService, jwtServ service.JWTService) CategoryController {
	return &categoryController{
		catService: catServ,
		jwtService: jwtServ,
	}
}

func (c *categoryController) All(context *gin.Context) {
	var cats []entity.Category = c.catService.All()
	res := helper.BuildResponse(true, "OK", cats)
	context.JSON(http.StatusOK, res)
}
func (c *categoryController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var cat entity.Category = c.catService.FindByID(id)
	if (cat == entity.Category{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", cat)
		context.JSON(http.StatusOK, res)
	}
}
func (c *categoryController) Insert(context *gin.Context) {
	var catCreate dto.CategoryCreate
	errDTO := context.ShouldBind(&catCreate)
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
		if c.catService.UserRole(userID) != "admin" {
			response := helper.BuildErrorResponse("You dont have permission", "You are not Administrator", helper.EmptyObj{})
			context.JSON(http.StatusForbidden, response)
		} else {
			result := c.catService.Insert(catCreate)
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusCreated, response)
		}
	}
}
func (c *categoryController) Update(context *gin.Context) {
	var catUpdate dto.CategoryUpdate
	errDTO := context.ShouldBind(&catUpdate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	} else {
		authHeader := context.GetHeader("Authorization")
		token, errToken := c.jwtService.ValidateToken(authHeader)
		if errToken != nil {
			panic(errToken.Error())
		}
		claims := token.Claims.(jwt.MapClaims)
		userID := fmt.Sprintf("%v", claims["user_id"])
		if c.catService.UserRole(userID) != "admin" {
			response := helper.BuildErrorResponse("You dont have permission", "You are not Administrator", helper.EmptyObj{})
			context.JSON(http.StatusForbidden, response)
		} else {
			result := c.catService.Update(catUpdate)
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusOK, response)
		}
	}
}
func (c *categoryController) Delete(context *gin.Context) {
	var cat entity.Category
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	} else {
		cat.ID = id
		authHeader := context.GetHeader("Authorization")
		token, errToken := c.jwtService.ValidateToken(authHeader)
		if errToken != nil {
			panic(errToken.Error())
		}
		claims := token.Claims.(jwt.MapClaims)
		userID := fmt.Sprintf("%v", claims["user_id"])
		if c.catService.UserRole(userID) != "admin" {
			response := helper.BuildErrorResponse("You dont have permission", "You are not Administrator", helper.EmptyObj{})
			context.JSON(http.StatusForbidden, response)
		} else {
			c.catService.Delete(cat)
			res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
			context.JSON(http.StatusOK, res)
		}
	}
}
