package repository

import (
	"context"
	"fmt"

	"github.com/Nokoyohei/ptcwave/v1/entity"
	"github.com/jmoiron/sqlx"
)

type SubPriceRepository struct {
	db *sqlx.DB
}

func NewSubPriceRepository(db *sqlx.DB) (*SubPriceRepository, error) {
	return &SubPriceRepository{db: db}, nil
}

func (cr *SubPriceRepository) FindSubPrice(base_fetch_date string, at_fetch_date string, shop string) (*[]entity.SubPrice, error) {
	ctx := context.Background()
	cards := []entity.SubPrice{}

	table_name := fmt.Sprintf("%s_cards", shop)

	err := cr.db.SelectContext(ctx, &cards,
		fmt.Sprintf(`WITH n AS(
			SELECT * 
			FROM %s as n
			WHERE n.fetch_date=?
		), p AS(
			SELECT * 
			FROM %s
			WHERE fetch_date=?
		), sub AS(
		SELECT 
			name, 
			version, 
			rarity, 
			n.type, 
			status, 
			n.price as now_price, 
			CASE WHEN p.price IS NULL THEN 0 ELSE p.price END as prev_price,
			CASE WHEN p.price IS NULL THEN n.price ELSE n.price - p.price END as sub
		FROM n LEFT JOIN p USING(name, version, rarity, status)
		)
		SELECT * FROM sub
		WHERE sub != 0`, table_name, table_name), at_fetch_date, base_fetch_date)
	if err != nil {
		return nil, err
	}

	return &cards, nil
}
