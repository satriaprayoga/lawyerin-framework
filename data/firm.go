package data

import "time"

type Firm struct {
	FirmID      int
	FirmName    string
	FirmAddress string
	FirmSince   time.Time
}
