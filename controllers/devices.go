package controllers

import (
	
	"github.com/bjacobel/checkthat/models"
	"github.com/jinzhu/gorm"
	"github.com/laurent22/ripple"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"time"
	_ "github.com/lib/pq"
)

type DeviceController struct {
	db gorm.DB
}

type JoinedResult struct {
	Id            int64
	Os            string
	Type          string
	Name          string
	Manufacturer  string
	Model         string
	Version       float32
	NfcSerial     int64
	CheckedOut    int64
	UserId        int64
	UserNfcSerial int64
	UserFirstName string
	UserLastName  string
	UserTel       string
}

type ResponseStruct struct {
	CheckedOut   []JoinedResult
	CheckedIn    []models.Device
}

func NewDeviceController(db gorm.DB) *DeviceController {
	output := new(DeviceController)
	output.db = db
	return output
}

func (this *DeviceController) Get(ctx *ripple.Context) {
	deviceId, _ := strconv.Atoi(ctx.Params["id"])

	if deviceId > 0 {
		singleDevice := models.Device{}
		this.db.First(&singleDevice, deviceId)
		ctx.Response.Body = singleDevice
	} else {
		checkedout := []JoinedResult{}
		this.db.Table("devices").Joins("left join users on users.id = devices.user_id").Select(`
			devices.id as id,
			devices.os as os,
			devices.type as type,
			devices.name as name,
			devices.manufacturer as manufacturer,
			devices.model as model,
			devices.version as version,
			devices.nfc_serial as nfc_serial,
			devices.checked_out as checked_out,
			users.id as user_id,
			users.nfc_serial as user_nfc_serial,
			users.first_name as user_first_name,
			users.last_name as user_last_name,
			users.tel as user_tel
		`).Where("user_id >= ?", 1).Find(&checkedout)

		checkedin := []models.Device{}
		this.db.Where("user_id = ?", 0).Find(&checkedin)

		listResponse := ResponseStruct{}
		listResponse.CheckedOut = checkedout
		listResponse.CheckedIn = checkedin

		ctx.Response.Body = listResponse
	}
}

func (this *DeviceController) PostCheckout(ctx *ripple.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)

	pc := map[string]int64{}

	json.Unmarshal(body, &pc)

	if _, ok := pc["device_uid"]; !ok {
		ctx.Response.Status = 412
		return
	}
	if _, ok := pc["user_uid"]; !ok {
		ctx.Response.Status = 412
		return
	}

	device := models.Device{}
	this.db.Where("nfc_serial = ?", pc["device_uid"]).First(&device)
	
	user := models.User{}
	this.db.Where("nfc_serial = ?", pc["user_uid"]).First(&user)

	if device.Id == 0 {
		ctx.Response.Status = 412
		return
	}
	
	if device.Id == 0 {
		ctx.Response.Status = 412
		return
	}

	device.CheckedOut = time.Now().Unix()
	device.UserId = user.Id

	this.db.Save(&device)

	ctx.Response.Status = 200
	ctx.Response.Body = device
}
