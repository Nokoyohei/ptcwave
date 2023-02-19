package repository

import (
	"context"

	"github.com/Nokoyohei/ptcwave/v1/entity"
	"github.com/jmoiron/sqlx"
)

type SubPriceRepository struct {
	db *sqlx.DB
}

func NewSubPriceRepository(db *sqlx.DB) (*SubPriceRepository, error) {
	return &SubPriceRepository{db: db}, nil
}

func (cr *SubPriceRepository) FindSubPrice(prev_fetch_date string, now_fetch_date string) (*[]entity.SubPrice, error) {
	ctx := context.Background()
	cards := []entity.SubPrice{}

	err := cr.db.SelectContext(ctx, &cards,
		`WITH n AS(
			SELECT * 
			FROM cards as n
			WHERE n.fetch_date=?
		), p AS(
			SELECT * 
			FROM cards 
			WHERE fetch_date=?
		)
		SELECT 
			name, 
			version, 
			rarity, 
			n.type, 
			status, 
			n.price as now_price, 
			p.price as prev_price, 
			n.price - p.price as sub, 
			n.fetch_date as now_fetch_date, 
			p.fetch_date as prev_fetch_date 
		FROM n JOIN p USING(name, version, rarity, status)
		WHERE p.price - n.price != 0`, prev_fetch_date, now_fetch_date)
	if err != nil {
		return nil, err
	}

	return &cards, nil
}
