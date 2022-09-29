package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Car struct {
	CarID string `json:"card_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var cars = []*Car{}

func GetCar(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, cars)
}

func AddCar(ctx *gin.Context) {
	newCar := Car{}
	err := ctx.ShouldBindJSON(&newCar)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			struct{ Code, Message string }{Code: "231", Message: err.Error()})
		return
	}

	newCar.CarID = uuid.NewString()
	cars = append(cars, &newCar)
	ctx.JSON(http.StatusOK, newCar)
}

func EditCar(ctx *gin.Context) {
	carID := ctx.Param("carId")

	car, getCarErr := getCarById(carID)
	if getCarErr != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			struct{ Code, Message string }{"234-434", getCarErr.Error()})
		return
	}

	car.Brand = "yakuza yakuza"

	bindErr := ctx.ShouldBindJSON(&car)
	if bindErr != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			struct{ Code, Message string }{Code: "231", Message: bindErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, car)
}

func getCarById(id string) (*Car, error) {
	for _, v := range cars {
		if v.CarID == id {
			return v, nil
		}
	}

	return &Car{}, errors.New("Car not found")
}

func DeleteCar(ctx *gin.Context) {
}
