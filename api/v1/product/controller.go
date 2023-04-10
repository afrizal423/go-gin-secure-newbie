package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/afrizal423/go-gin-secure-newbie/app/business/product"
	"github.com/afrizal423/go-gin-secure-newbie/app/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Controller struct {
	service product.IProductService
}

func NewProductController(service product.IProductService) *Controller {
	return &Controller{
		service,
	}
}

// Role Admin (1)
// GET, GET ALL, UPDATE, DELETE, POST
// Role User (2)
// GET, GET ALL, POST
func (handler *Controller) GetAllProducts(c *gin.Context) {
	fmt.Println(c.MustGet("userData").(jwt.MapClaims))
	data := c.MustGet("userData").(jwt.MapClaims)
	roles := uint(data["roles"].(float64))

	if roles == 1 {
		// if admin
		dt, err := handler.service.GetAll()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": fmt.Sprintf("Error: %s", err.Error()),
			})
			return
		}
		//fmt.Println(dt)
		c.JSON(http.StatusOK, dt)
		return
	} else {
		// if user
		dt, err := handler.service.GetAllByUserId(uint(data["user_id"].(float64)))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": fmt.Sprintf("Error: %s", err.Error()),
			})
			return
		}
		//fmt.Println(dt)
		c.JSON(http.StatusOK, dt)
		return
	}
}

func (handler *Controller) GetProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	roles := uint(userData["roles"].(float64))
	userID := uint(userData["user_id"].(float64))
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil || uint(productId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid product ID",
		})
		return
	}

	if roles == 1 {
		res, err := handler.service.GetDataById(uint(productId))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintf("Product with id: %d not found\n", productId),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		res, err := handler.service.GetDataByUserId(userID, uint(productId))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintf("Product with id: %d not found\n", productId),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func (handler *Controller) CreateProduct(c *gin.Context) {
	fmt.Println(c.MustGet("userData").(jwt.MapClaims))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))
	product := models.Product{}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if res, err := handler.service.CreateProduct(product, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusCreated, res)
		return
	}
}

func (handler *Controller) UpdateProduct(c *gin.Context) {
	product := models.Product{}
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid product ID",
		})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if res, err := handler.service.UpdateProduct(product, uint(productId)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (handler *Controller) DeleteProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil || uint(productId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid product ID",
		})
		return
	}

	if _, err := handler.service.DeleteProduct(uint(productId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete product",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product deleted successfully",
		})
	}
}
