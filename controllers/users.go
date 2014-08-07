package controllers

import (
	_ "fmt"
	"github.com/bjacobel/checkthat/models"
	"github.com/jinzhu/gorm"
	"github.com/laurent22/ripple"
	_ "github.com/lib/pq"
	"strconv"
)

type UserController struct {
	db gorm.DB
}

func NewUserController(db gorm.DB) *UserController {
	output := new(UserController)
	output.db = db
	return output
}

func (this *UserController) Get(ctx *ripple.Context) {
	userId, _ := strconv.Atoi(ctx.Params["id"])

	if userId > 0 {
		ctx.Response.Body = this.db.First(&models.User{}, userId)
	} else {
		ctx.Response.Body = this.db.Find(&[]models.User{})
	}
}
