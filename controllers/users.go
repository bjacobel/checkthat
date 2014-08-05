package controllers

import (
    "github.com/laurent22/ripple"
    "github.com/dchest/uniuri"
    "github.com/bjacobel/checkthat/models"
    "encoding/json"
    "io/ioutil"
    "strconv"
)

type UserController struct {
    userCollection models.UserCollection
}

func NewUserController() *UserController {
    output := new(UserController)

    output.userCollection.Users = make(map[int]models.UserModel)

    for i := 0; i < 100; i++ {
        output.userCollection.Add(models.UserModel{0, uniuri.NewLen(10), uniuri.NewLen(10)})
    }

    return output
}

func (this *UserController) Get(ctx *ripple.Context) {
    userId, _ := strconv.Atoi(ctx.Params["id"])
    if userId > 0 {
        ctx.Response.Body = this.userCollection.Get(userId)
    } else {
        ctx.Response.Body = this.userCollection.GetAll()
    }
}

func (this *UserController) Post(ctx *ripple.Context) {
    body, _ := ioutil.ReadAll(ctx.Request.Body)
    var user models.UserModel
    json.Unmarshal(body, &user)
    ctx.Response.Body = this.userCollection.Add(user)
}

func (this *UserController) Put(ctx *ripple.Context) {
    body, _ := ioutil.ReadAll(ctx.Request.Body)
    userId, _ := strconv.Atoi(ctx.Params["id"])
    var user models.UserModel
    json.Unmarshal(body, &user)
    ctx.Response.Body = this.userCollection.Set(userId, user)
}