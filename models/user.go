package models

type User struct {
    Id          int64
    NFCSerial   int64
    First_name  string
    Last_name   string
    Tel         int
    Devices     []Device
}
