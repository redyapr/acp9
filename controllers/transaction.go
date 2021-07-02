package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/middlewares"
	"acp9-redy-gigih/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTransactionsController(e echo.Context) error {
	userId := int(middlewares.ExtractToken(e))
	transactions := []models.Transaction{}
	err := config.DB.Debug().Model(&models.Transaction{}).Where("user_id = ?", userId).Find(&transactions).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.TransactionResponse{
			false, "Get transactions failed", nil,
		})
	}
	return e.JSON(http.StatusOK, models.TransactionResponse{
		true, "Get cart success", transactions,
	})
}

func CheckoutController(e echo.Context) error {
	userId := int(middlewares.ExtractToken(e))
	carts := []models.Cart{}
	err := config.DB.Debug().Model(&models.Cart{}).Where("user_id = ?", userId).Find(&carts).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.CartResponse{
			false, "Get cart failed", nil,
		})
	}
	count := config.DB.Debug().Model(&models.Cart{}).Where("user_id = ?", userId).Find(&carts).RowsAffected
	if count == 0 {
		return e.JSON(http.StatusInternalServerError, models.CartResponse{
			false, "Cart is empty", nil,
		})
	} else {
		transaction := models.Transaction{}
		transaction.UserID = userId
		err := config.DB.Create(&transaction).Error
		if err != nil {
			return e.JSON(http.StatusInternalServerError, models.TransactionResponse{
				false, "Create transaction head failed", nil,
			})
		}
		for _, cart := range carts {
			fmt.Println(cart.ProductID)
			detail := models.TransactionDetail{}
			detail.TransactionID = int(transaction.ID)
			detail.ProductID = cart.ProductID
			detail.Qty = cart.Qty
			err := config.DB.Create(&detail).Error
			if err != nil {
				return e.JSON(http.StatusInternalServerError, models.TransactionResponse{
					false, "Create transaction detail failed", nil,
				})
			}
			err = config.DB.Debug().Model(&models.Cart{}).Where("user_id = ?", userId).Delete(&carts).Error
			if err != nil {
				return e.JSON(http.StatusInternalServerError, models.TransactionResponse{
					false, "Clear cart failed", nil,
				})
			}
		}
	}
	return e.JSON(http.StatusOK, models.TransactionResponse{
		true, "Checkout success", nil,
	})
}

func PaymentController(e echo.Context) error {
	userId := int(middlewares.ExtractToken(e))
	input := models.Transaction{}
	e.Bind(&input)
	transaction := models.Transaction{}
	err := config.DB.Debug().Model(&transaction).Where("id = ? AND user_id = ?", input.ID, userId).Update("status", "Paid").Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.CartResponse{
			false, "Payment failed", nil,
		})
	}
	count := config.DB.Debug().Model(&transaction).Where("id = ? AND user_id = ?", input.ID, userId).Update("status", "Paid").RowsAffected
	if count == 0 {
		return e.JSON(http.StatusInternalServerError, models.UserResponse{
			false, "Transaction ID not found", nil,
		})
	}
	return e.JSON(http.StatusOK, models.CartResponse{
		true, "Payment success", nil,
	})
}
