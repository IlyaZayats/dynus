package entity

import "time"

type User struct {
	Id      int       `db:"id"`
	Created time.Time `db:"created_at"`
	Updated time.Time `db:"updated_at"`
}
