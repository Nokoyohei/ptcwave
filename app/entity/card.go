package entity

import "fmt"

type Card struct {
	name string `db:"name" json:"name"`
	version string `db:"version" json:"version"`
	rarity string `db:"rarity" json:"rarity"`
	_type string `db:"type" json:"type"`
	price int `db:"price" json:"price"`
	status string `db:"status" json:"status"`
}

func NewCard(
	name string, 
	version string,
	rarity string,
	_type string,
	price int,
	status string,
) Card {
	return Card{
		name: name,
		version: version,
		rarity: rarity,
		_type: _type,
		price: price,
		status: status,
	}
}