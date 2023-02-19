package entity

import "time"

type Card struct {
	Name      string    `db:"name" json:"name"`
	Version   string    `db:"version" json:"version"`
	Rarity    string    `db:"rarity" json:"rarity"`
	Type      string    `db:"type" json:"type"`
	Price     int       `db:"price" json:"price"`
	Status    string    `db:"status" json:"status"`
	FetchDate time.Time `db:"fetch_date"`
}

func NewCard(
	name string,
	version string,
	rarity string,
	_type string,
	price int,
	status string,
	fetch_date time.Time,
) *Card {
	return &Card{
		Name:      name,
		Version:   version,
		Rarity:    rarity,
		Type:      _type,
		Price:     price,
		Status:    status,
		FetchDate: fetch_date,
	}
}
