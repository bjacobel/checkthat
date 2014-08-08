package models

type Device struct {
	Id              int64
	Os              string
	Type            string
	Manufacturer    string
	Model           string
	Name            string
	Version         float32
	NfcSerial       int64
	LastActivity    int64
	UserId          int64 `sql:"bigint REFERENCES user(id)"`
}
