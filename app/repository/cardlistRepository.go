package repository

import (
	"context"

	"github.com/Nokoyohei/ptcwave/v1/entity"
	"github.com/jmoiron/sqlx"
)

type CardlistRepository struct {
	db *sqlx.DB
}

func NewCardlistRepository(db *sqlx.DB) (*CardlistRepository, error) {
	return &CardlistRepository{db: db}, nil
}

func (cr *CardlistRepository) FindAll() (*[]entity.Card, error) {
	ctx := context.Background()
	cards := []entity.Card{}

	err := cr.db.SelectContext(ctx, &cards,
		"SELECT name, version, rarity, type, price, status, fetch_date  FROM cards")
	if err != nil {
		return nil, err
	}

	return &cards, nil
}
