package models

type Device struct {
	Id			int64
	Os			string
	Type		string
	Name		string
	Version		float32
	NfcSerial	int64
    CheckedOut	int64  // use go's time.Now().Unix
	UserId		int64  `sql:"bigint REFERENCES user(id)"`
}
