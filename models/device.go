package models

type Device struct {
    Id          int64
    OS          string
    Type        string
    Version     float32
    NFCSerial   int64
    UserId      int64
}