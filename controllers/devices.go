package controllers

import (
	"fmt"
	"github.com/bjacobel/checkthat/models"
	"github.com/jinzhu/gorm"
	"github.com/laurent22/ripple"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

type DeviceController struct {
	db gorm.DB
}

type JoinedResult struct {
	DeviceId 			int64
	DeviceOs 			string
	DeviceType 			string
	DeviceName 			string
	DeviceVersion		float32
	DeviceNfcSerial		int64
	DeviceCheckedOut 	int64
	UserId 				int64
	UserNfcSerial 		int64
	UserFirstName		string
	UserLastName 		string
	UserTel				string
}

func NewDeviceController() *DeviceController {
	output := new(DeviceController)

	db, dberr := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@ec2-54-197-241-67.compute-1.amazonaws.com:5432/%s", os.Getenv("PGUSER"), os.Getenv("PGPW"), os.Getenv("PGDB")))
	if dberr != nil {
		panic(dberr)
	}

	db.AutoMigrate(models.Device{})

	output.db = db

	return output
}

func (this *DeviceController) Get(ctx *ripple.Context) {
	deviceId, _ := strconv.Atoi(ctx.Params["id"])

	scan := JoinedResult{}

	listResponse := map[string]*gorm.DB{}
	listResponse["checked_out"] = this.db.Table("devices").Joins("left join users on users.id = devices.user_id").Select(`
		devices.id as device_id,
		devices.os as device_os,
		devices.type as device_type,
		devices.name as device_name,
		devices.version as device_version,
		devices.nfc_serial as device_nfc_serial,
		devices.checked_out as device_checked_out,
		users.id as user_id,
		users.nfc_serial as user_nfc_serial,
		users.first_name as user_first_name,
		users.last_name as user_last_name,
		users.tel as user_tel
	`).Where("user_id >= ?", 2).Find(&scan)
	listResponse["checked_in"] = this.db.Where("user_id = ?", 1).Find(&[]models.Device{})

	if deviceId > 0 {
		ctx.Response.Body = this.db.First(&models.Device{}, deviceId)
	} else {
		ctx.Response.Body = listResponse
	}
}

func(this *DeviceController) PostCheckout(ctx *ripple.Context) {
	//deviceId, _ := strconv.Atoi(ctx.Params["id"])
	body, _ := ioutil.ReadAll(ctx.Request.Body)

	pc := map[string]int64{}

	json.Unmarshal(body, &pc)

	if _, ok := pc["device_uid"]; ! ok {
		ctx.Response.Status = 422
		return
	}
	if _, ok := pc["user_uid"]; ! ok {
		ctx.Response.Status = 422
		return
	}

	ctx.Response.Status = 200
}
