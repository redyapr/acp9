package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/middlewares"
	"acp9-redy-gigih/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddCartController(e echo.Context) error {
	userId := int(middlewares.ExtractToken(e))
	input := models.Cart{}
	e.Bind(&input)
	cartDB := models.Cart{}
	cartDB.UserId = userId
	cartDB.ProductID = input.ProductID
	err := config.DB.Debug().Model(&models.Cart{}).Where("user_id = ? AND product_id = ?", userId, input.ProductID).Find(&cartDB).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.CartResponse{
			false, "Check same item failed", nil,
		})
	}
	count := config.DB.Debug().Model(&models.Cart{}).Where("user_id = ? AND product_id = ?", userId, input.ProductID).Find(&cartDB).RowsAffected
	if count == 0 {
		cartDB.Qty = input.Qty
		_ = config.DB.Create(&cartDB).Error
		// if err != nil {
		// 	return e.JSON(http.StatusInternalServerError, models.CartResponse{
		// 		false, "Add to cart failed", nil,
		// 	})
		// }
	} else {
		_ = config.DB.Debug().Model(&cartDB).Where("user_id = ? AND product_id = ?", userId, input.ProductID).Update("qty", cartDB.Qty+input.Qty).Error
		// if err != nil {
		// 	return e.JSON(http.StatusInternalServerError, models.CartResponse{
		// 		false, "Update cart failed", nil,
		// 	})
		// }
	}
	return e.JSON(http.StatusOK, models.CartResponseSingle{
		true, "Add to cart success", cartDB,
	})
}

func GetCartController(e echo.Context) error {
	userId := int(middlewares.ExtractToken(e))
	carts := []models.Cart{}
	err := config.DB.Debug().Model(&models.Cart{}).Where("user_id = ?", userId).Find(&carts).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.CartResponse{
			false, "Get cart failed", nil,
		})
	}
	return e.JSON(http.StatusOK, models.CartResponse{
		true, "Get cart success", carts,
	})
}

func UpdateCartController(e echo.Context) error {
	userId := int(middlewares.ExtractToken(e))
	fmt.Println(userId)
	productId := e.Param("productId")
	input := models.Cart{}
	e.Bind(&input)
	carts := models.Cart{}
	err := config.DB.Debug().Model(&carts).Where("user_id = ? AND product_id = ?", userId, productId).Update("qty", input.Qty).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.CartResponse{
			false, "Update cart failed", nil,
		})
	}
	count := config.DB.Debug().Model(&carts).Where("user_id = ? AND product_id = ?", userId, productId).Update("qty", input.Qty).RowsAffected
	if count == 0 {
		return e.JSON(http.StatusInternalServerError, models.UserResponse{
			false, "Nothing updated", nil,
		})
	}
	return e.JSON(http.StatusOK, models.CartResponse{
		true, "Update cart success", nil,
	})
}

func DeleteCartController(e echo.Context) error {
	userId := int(middlewares.ExtractToken(e))
	productId := e.Param("productId")
	carts := []models.Cart{}
	err := config.DB.Debug().Model(&models.Cart{}).Where("user_id = ? AND product_id = ?", userId, productId).Delete(&carts).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.CartResponse{
			false, "Delete cart failed", nil,
		})
	}
	return e.JSON(http.StatusOK, models.CartResponse{
		true, "Delete cart success", carts,
	})
}
