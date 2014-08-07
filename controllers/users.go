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
		single_user := models.User{}
		this.db.First(&single_user, userId)
		ctx.Response.Body = single_user
	} else {
		user_list := []models.User{}
		this.db.Find(&user_list)
		ctx.Response.Body = user_list
	}
}
