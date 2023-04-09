package user

import (
	"fmt"
	"net/http"

	"github.com/afrizal423/go-gin-secure-newbie/app/business/user"
	"github.com/afrizal423/go-gin-secure-newbie/app/models"
	"github.com/afrizal423/go-gin-secure-newbie/pkg/utils/errors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service user.IUserService
}

func NewUserController(service user.IUserService) *Controller {
	return &Controller{
		service,
	}
}

func (c *Controller) Register(ctx *gin.Context) {
	// alokasikan memori
	// https://dev.to/bhanu011/how-is-new-different-in-go-41lk
	// https://go.dev/doc/effective_go
	user := new(models.User)
	// pastikan data harus berupa json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}

	// proses ke bagian services
	data, err := c.service.Register(*user)
	if err != nil {
		restErr := errors.NewBadRequestError(fmt.Sprintf("%v", err.Error()))
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}
	// tinggal di return aja
	ctx.JSON(http.StatusCreated, gin.H{
		"id":    data.ID,
		"email": data.Email})
}

func (c *Controller) Login(ctx *gin.Context) {
	// user := new(models.User)
	var user models.User
	// pastikan data harus berupa json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}
	token, err := c.service.Login(user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"token": token,
		})
	}
}
