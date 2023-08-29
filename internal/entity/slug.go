package entity

import "time"

type Slug struct {
	Id      int       `db:"id"`
	Name    string    `db:"name"`
	Chance  float64   `db:"chance"`
	Created time.Time `db:"created_at"`
	Updated time.Time `db:"updated_at"`
}
