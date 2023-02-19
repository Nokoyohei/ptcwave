package entity

import "time"

type SubPrice struct {
	Name          string    `db:"name"`
	SubPrice      int       `db:"sub"`
	Version       string    `db:"version" json:"version"`
	Rarity        string    `db:"rarity" json:"rarity"`
	Type          string    `db:"type" json:"type"`
	PrevPrice     int       `db:"prev_price"`
	NowPrice      int       `db:"now_price"`
	Status        string    `db:"status" json:"status"`
	PrevFetchDate time.Time `db:"prev_fetch_date"`
	NowFetchDate  time.Time `db:"now_fetch_date"`
}
