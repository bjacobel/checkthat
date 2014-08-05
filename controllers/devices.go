package controllers

import (
    "github.com/bjacobel/checkthat/models"
    "github.com/laurent22/ripple"
    "github.com/jinzhu/gorm"
    _"github.com/lib/pq"
    "strconv"
    "fmt"
    "os"
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
    output.db = db

    return output
}

func (this *DeviceController) Get(ctx *ripple.Context) {
    device := models.Device{}
    devices := []models.Device{}

    deviceId, _ := strconv.Atoi(ctx.Params["id"])
    if deviceId > 0 {
        ctx.Response.Body = this.db.First(&device, deviceId)
    } else {
        ctx.Response.Body = this.db.Find(&devices)
    }
}