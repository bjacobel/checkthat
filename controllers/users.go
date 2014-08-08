package controllers

import (
	"github.com/bjacobel/checkthat/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/laurent22/ripple"
	"github.com/agonzalezro/twilio-go"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

type UserController struct {
	db gorm.DB
	twclient *twilio.TwilioRestClient
}

func NewUserController(db gorm.DB, twclient *twilio.TwilioRestClient) *UserController {
	output := new(UserController)
	output.twclient = twclient
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

func (this *UserController) PostPush(ctx *ripple.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)

	requestbody := map[string]int32{}

	json.Unmarshal(body, &requestbody)

	if _, ok := requestbody["device_id"]; !ok {
		ctx.Response.Status = 412
		return
	}

	device := models.Device{}
	this.db.Find(&device, requestbody["device_id"])

	var msg_resp *twilio.MessagesResponse
	var msg_err *twilio.ErrorResponse

	if userId, _ := strconv.Atoi(ctx.Params["id"]) ; userId > 0 {
		// if there is a user id, push to that user

		user := models.User{}
		this.db.Find(&user, userId)

		message := "Hey, "+string(user.FirstName)+"! Somebody needs to use the "+string(device.Model)+" you've checked out. Please return it."
		
		msg_resp, msg_err = this.twclient.Messages.Create("+15074164045", user.Tel, message)
	} else {
		users := []models.User{}
		this.db.Find(&users)

		var message string

		for _, user := range users {
			message = "Hey, "+string(user.FirstName)+"! Somebody needs the "+string(device.Model)+" called "+string(device.Name)+", but CheckThis has lost track of it. If you have it, please check it back in!"

			msg_resp, msg_err = this.twclient.Messages.Create("+15074164045", user.Tel, message)
			
			if msg_err.Code == 0 {
				break
			}
		} 
	}

	if msg_err.Code == 0 {
		ctx.Response.Body = msg_resp
		ctx.Response.Status = 200
	} else {
		ctx.Response.Body = msg_err
		ctx.Response.Status = 500
	}
	
}
