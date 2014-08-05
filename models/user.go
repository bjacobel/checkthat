package models

type User struct {
    Id          int64
    NfcSerial   int64
    FirstName  string
    LastName   string
    Tel         int
    Devices     []Device
}
