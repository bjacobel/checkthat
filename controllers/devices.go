package controllers

import (
	"fmt"
	"github.com/bjacobel/checkthat/models"
	"github.com/jinzhu/gorm"
	"github.com/laurent22/ripple"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

type DeviceController struct {
	db gorm.DB
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

	listResponse := map[string]*gorm.DB{}
	listResponse["checked_out"] = this.db.Where("user_id >= ?", 2).Find(&[]models.Device{})
	listResponse["checked_in"] = this.db.Where("user_id = ?", 1).Find(&[]models.Device{})

	if deviceId > 0 {
		ctx.Response.Body = this.db.First(&models.Device{}, deviceId)
	} else {
		ctx.Response.Body = listResponse
	}
}
