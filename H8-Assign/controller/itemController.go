package controller

import (
	"Hactiv8_Bootcamp/H8-Assign/db"
	"Hactiv8_Bootcamp/H8-Assign/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Orders = []models.Order{}

// type Item struct {
// 	ItemId      string    `json:"item_id"`
// 	ItemCode    string    `json:"item_code"`
// 	Description string    `json:"description"`
// 	Quantity    uint      `json:"quantity"`
// 	OrderId     uint      `json:"order_id"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at"`
// }

func CreateOrder(ctx *gin.Context) {
	db := db.GetDB()

	var tOrder models.Order
	ctx.ShouldBindJSON(&tOrder)
	err := db.Create(&tOrder).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't create order",
			"error_detail":  err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"created_order": tOrder,
	})
}

func GetOrder(ctx *gin.Context) {
	db := db.GetDB()

	err := db.Find(&Orders).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Data not found")
			return
		}
		print("Error finding Order: ", err)
	}
	for Order := range Orders {
		fmt.Printf("Order Data : %v \n ", Order)
	}
	ctx.JSON(http.StatusOK, Orders)
}

func DeleteOrder(ctx *gin.Context) {
	db := db.GetDB()
	var orderId string
	orderId = ctx.Param("OrderId")

	if err := db.First(&models.Order{}, "order_id = ?", orderId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v Not Found", orderId),
		})
		return
	}
	if err := db.Where("order_id = ?", orderId).Delete(&models.Order{}).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Deleting Order: %v", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprintf("Order with id %v Has Been Successfully Deleted", orderId),
	})

}

func UpdateOrder(ctx *gin.Context) {
	db := db.GetDB()
	var orderId string
	orderId = ctx.Param("OrderId")

	var updatedOrder models.Order

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v Not Found", orderId),
		})
		return
	}

	existingOrder := models.Order{}
	if err := db.Preload("OrderList").First(&existingOrder, "order_id = ?", orderId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v Not Found", orderId),
		})
		return
	}

	if err := db.Model(&existingOrder).Updates(updatedOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Updating Order Data: %v", err.Error()),
		})
		return
	}

	for i, updatedItem := range updatedOrder.OrderList {
		if i < len(existingOrder.OrderList) {
			existingOrder.OrderList[i].ItemCode = updatedItem.ItemCode
			existingOrder.OrderList[i].Description = updatedItem.Description
			existingOrder.OrderList[i].Quantity = updatedItem.Quantity
			db.Save(&existingOrder.OrderList[i])
		} else {
			updatedItem.OrderId = existingOrder.OrderId
			db.Create(&updatedItem)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprintf("Order with id %v Has Been Successfully Updated", orderId),
	})
}
