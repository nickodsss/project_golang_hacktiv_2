package controller

import (
	"assignment_2_golang/config"
	"assignment_2_golang/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(c *gin.Context) {
	var orders []models.Order
	config.ConnectDatabase().Preload("Items").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// for _, item := range order.Items {
	// 	order.Items = append(order.Items, models.Item{
	// 		ItemCode:    item.ItemCode,
	// 		Description: item.Description,
	// 		Quantity:    item.Quantity,
	// 		CreatedAt:   time.Now(),
	// 		UpdatedAt:   time.Now(),
	// 	})
	// }

	if err := config.ConnectDatabase().Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	var order models.Order

	if err := config.ConnectDatabase().Preload("Items").First(&order, c.Param("order_id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := config.ConnectDatabase().Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Data Success",
		"success": true,
	})
}

func UpdateOrder(c *gin.Context) {
	var order models.Order
	var req models.Order
	if err := config.ConnectDatabase().Preload("Items").First(&order, c.Param("order_id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.CustomerName = req.CustomerName
	order.OrderedAt = req.OrderedAt
	order.UpdatedAt = time.Now()

	for _, newItem := range req.Items {
		for i, existingItem := range order.Items {
			if existingItem.OrderID == order.OrderID {
				order.Items[i].ItemCode = newItem.ItemCode
				order.Items[i].Description = newItem.Description
				order.Items[i].Quantity = newItem.Quantity
				order.Items[i].UpdatedAt = time.Now()
				break
			}
		}
	}

	config.ConnectDatabase().Save(&order)
	c.JSON(http.StatusOK, order)
}

// func UpdateOrder(c *gin.Context) {
// 	var order models.Order
// 	var req models.Order

// 	// Ambil order yang ada berdasarkan ID
// 	if err := config.ConnectDatabase().Preload("Items").First(&order, c.Param("order_id")).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
// 		return
// 	}

// 	// Bind JSON dari request ke struct
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Perbarui informasi order
// 	order.CustomerName = req.CustomerName
// 	order.OrderedAt = req.OrderedAt
// 	order.UpdatedAt = time.Now()

// 	// Perbarui item yang ada
// 	existingItemIDs := make(map[uint]bool)
// 	for _, item := range req.Items {
// 		// Temukan item yang sesuai
// 		found := false
// 		for i := range order.Items {
// 			if order.Items[i].ItemID == item.ItemID {
// 				order.Items[i].ItemCode = item.ItemCode
// 				order.Items[i].Description = item.Description
// 				order.Items[i].Quantity = item.Quantity
// 				order.Items[i].UpdatedAt = time.Now()
// 				found = true
// 				existingItemIDs[item.ItemID] = true
// 				break
// 			}
// 		}
// 		// Jika item baru, tambahkan ke order.Items
// 		if !found {
// 			order.Items = append(order.Items, models.Item{
// 				ItemCode:    item.ItemCode,
// 				Description: item.Description,
// 				Quantity:    item.Quantity,
// 				CreatedAt:   time.Now(),
// 				UpdatedAt:   time.Now(),
// 			})
// 		}
// 	}

// 	// Hapus item yang tidak ada dalam request
// 	for i := len(order.Items) - 1; i >= 0; i-- {
// 		if _, exists := existingItemIDs[order.Items[i].ItemID]; !exists {
// 			order.Items = append(order.Items[:i], order.Items[i+1:]...)
// 		}
// 	}

// 	// Simpan perubahan ke database
// 	if err := config.ConnectDatabase().Save(&order).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, order)
// }
