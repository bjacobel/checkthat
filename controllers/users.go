package controllers

import (
    "github.com/bjacobel/checkthat/models"
    "github.com/laurent22/ripple"
    "github.com/jinzhu/gorm"
    _"github.com/lib/pq"
    "strconv"
    "fmt"
    "os"
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
    output.db = db

    return output
}

func (this *UserController) Get(ctx *ripple.Context) {
    user := models.User{}
    users := []models.User{}

    userId, _ := strconv.Atoi(ctx.Params["id"])
    if userId > 0 {
        ctx.Response.Body = this.db.First(&user, userId)
    } else {
        ctx.Response.Body = this.db.Find(&users)
    }
}

// func (this *UserController) Post(ctx *ripple.Context) {
//     body, _ := ioutil.ReadAll(ctx.Request.Body)
//     var user models.UserModel
//     json.Unmarshal(body, &user)
//     ctx.Response.Body = this.userCollection.Add(user)
// }