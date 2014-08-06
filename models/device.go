package models

type Device struct {
	Id        int64
	Os        string
	Type      string
	Name      string
	Version   float32
	NfcSerial int64
	UserId    int64
}
