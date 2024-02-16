package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarId string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarDatas = []Car{}

// Create a new Car
// @Summary Postdetails for a given Id
// @Description Create a new Car
// @Tags cars
// @Accept  json
// @Produce  json
// @Param models.Car body models.Car true "create a new car"
// @Success 200 {object} models.Car
// @Router /cars/ [post]
func CreateCar(ctx *gin.Context) {
	var newCar Car
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.CarId = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

// Update Car
// @Summary Update a car identified by the given Id
// @Description Update a Car
// @Tags cars
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the car to be updated"
// @Success 200 {object} models.Car
// @Router /cars/{id} [patch]
func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false

	var updateCar Car

	if err := ctx.ShouldBindJSON(&updateCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas {
		if carID == car.CarId {
			condition = true
			CarDatas[i] = updateCar
			CarDatas[i].CarId = carID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been updated", carID),
	})
}

// Get Car
// @Summary Get all cars
// @Description Get a car
// @Tags cars
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Car
// @Router /cars/getCars [get]
func GetCar(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"car": CarDatas,
	})
}

// Delete Car
// @Summary Delet cars by Id
// @Description Delete a Car
// @Tags cars
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the car to be deleted"
// @Success 204 "No Context"
// @Router /cars/{id} [delete]
func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carIndex int

	for i, car := range CarDatas {
		if carID == car.CarId {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been deleted", carID),
	})
}
