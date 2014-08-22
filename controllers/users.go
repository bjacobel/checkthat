package controllers

import (
	"encoding/json"
	"github.com/bjacobel/checkthat/models"
	"github.com/bjacobel/ripple"
	"github.com/bjacobel/twilio-go"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"io/ioutil"
	"strconv"
)

type UserController struct {
	db       gorm.DB
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

func (this *UserController) Post(ctx *ripple.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	requestbody := map[string]string{}
	json.Unmarshal(body, &requestbody)

	newUser := models.User{}

	if _, ok := requestbody["FirstName"]; !ok {
		ctx.Response.Status = 412
		return
	} else {
		newUser.FirstName = requestbody["FirstName"]
	}

	if _, ok := requestbody["LastName"]; !ok {
		ctx.Response.Status = 412
		return
	} else {
		newUser.LastName = requestbody["LastName"]
	}

	if _, ok := requestbody["Tel"]; !ok {
		ctx.Response.Status = 412
		return
	} else {
		newUser.Tel = requestbody["Tel"]
	}

	if _, ok := requestbody["NfcSerial"]; !ok {
		ctx.Response.Status = 412
		return
	} else {
		derp, _ := strconv.Atoi(requestbody["NfcSerial"])
		newUser.NfcSerial = int64(derp)
	}

	this.db.Save(&newUser)

	ctx.Response.Status = 200
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

	if device.UserId > 0 {
		// if checked out to a user. push to them only

		user := models.User{}
		this.db.Find(&user, device.UserId)

		message := "Hey, " + string(user.FirstName) + "! Somebody needs to use the " + string(device.Model) + " you've checked out. Please return it as soon as you're done with it!"

		msg_resp, msg_err = this.twclient.Messages.Create("+1 617-860-2277", "+15072103812", message)

	} else {
		// we've lost it? Ask everyone.

		users := []models.User{}
		this.db.Find(&users)

		var message string

		for _, user := range users {
			message = "Hey, " + string(user.FirstName) + "! Somebody needs the " + string(device.Model) + " called " + string(device.Name) + ", but CheckThis has lost track of it. If you have it, please check it back in!"

			msg_resp, msg_err = this.twclient.Messages.Create("+1 617-860-2277", user.Tel, message)

			if msg_err.Code == 0 {
				break
			}
		}
	}

	if msg_err.Code == 0 {
		ctx.Response.Body = msg_resp
		ctx.Response.Status = 200
	} else {
		ctx.Response.Body = msg_err.Message
		ctx.Response.Status = msg_err.Code
	}

}
