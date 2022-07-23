package models

import "time"

type User struct {
	BaseModel

	Username string    `json:"username"`
	Password string    `json:"password"`
	WorkDate time.Time `json:"work_date"`
}
