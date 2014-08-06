package controllers

import (
	"fmt"
	"github.com/bjacobel/checkthat/models"
	"github.com/jinzhu/gorm"
	"github.com/laurent22/ripple"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	// "encoding/json"
	// "io/ioutil"
)

type UserController struct {
	db gorm.DB
}

func NewUserController() *UserController {
	output := new(UserController)

	db, dberr := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@ec2-54-197-241-67.compute-1.amazonaws.com:5432/%s", os.Getenv("PGUSER"), os.Getenv("PGPW"), os.Getenv("PGDB")))
	if dberr != nil {
		panic(dberr)
	}

    db.AutoMigrate(models.User{})
	
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

// func (this *UserController) Post(ctx *ripple.Context) {
//     body, _ := ioutil.ReadAll(ctx.Request.Body)
//     var user models.UserModel
//     json.Unmarshal(body, &user)
//     ctx.Response.Body = this.userCollection.Add(user)
// }
