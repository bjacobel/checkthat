package controllers

import (
	"encoding/json"
	_ "fmt"
	"github.com/bjacobel/checkthat/models"
	"github.com/jinzhu/gorm"
	"github.com/laurent22/ripple"
	_ "github.com/lib/pq"
	"io/ioutil"
	"strconv"
)

type DeviceController struct {
	db gorm.DB
}

type JoinedResult struct {
	Id            int64
	Os            string
	Type          string
	Name          string
	Version       float32
	NfcSerial     int64
	CheckedOut    int64
	UserId        int64
	UserNfcSerial int64
	UserFirstName string
	UserLastName  string
	UserTel       string
}

func NewDeviceController(db gorm.DB) *DeviceController {
	output := new(DeviceController)
	output.db = db
	return output
}

func (this *DeviceController) Get(ctx *ripple.Context) {
	deviceId, _ := strconv.Atoi(ctx.Params["id"])

	scan := []JoinedResult{}

	listResponse := map[string]*gorm.DB{}

	listResponse["checked_out"] = this.db.Table("devices").Joins("left join users on users.id = devices.user_id").Select(`
		devices.id as id,
		devices.os as os,
		devices.type as type,
		devices.name as name,
		devices.version as version,
		devices.nfc_serial as nfc_serial,
		devices.checked_out as checked_out,
		users.id as user_id,
		users.nfc_serial as user_nfc_serial,
		users.first_name as user_first_name,
		users.last_name as user_last_name,
		users.tel as user_tel
	`).Where("user_id >= ?", 1).Find(&scan)

	listResponse["checked_in"] = this.db.Where("user_id = ?", 0).Find(&[]models.Device{})

	if deviceId > 0 {
		ctx.Response.Body = this.db.First(&models.Device{}, deviceId)
	} else {
		ctx.Response.Body = listResponse
	}
}

func (this *DeviceController) PostCheckout(ctx *ripple.Context) {
	deviceId, _ := strconv.Atoi(ctx.Params["id"])
	body, _ := ioutil.ReadAll(ctx.Request.Body)

	pc := map[string]int64{}

	json.Unmarshal(body, &pc)

	if _, ok := pc["device_uid"]; !ok {
		ctx.Response.Status = 422
		return
	}
	if _, ok := pc["user_uid"]; !ok {
		ctx.Response.Status = 422
		return
	}

	ctx.Response.Status = 200
	ctx.Response.Body = this.db.Find(&[]models.Device{}, deviceId)
}
