package controllers

import (
    "github.com/laurent22/ripple"
    "github.com/bjacobel/checkthat/models"
    "encoding/json"
    "io/ioutil"
    "strconv"
)

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