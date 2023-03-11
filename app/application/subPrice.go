package application

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Nokoyohei/ptcwave/v1/database"
	"github.com/Nokoyohei/ptcwave/v1/entity"
	"github.com/Nokoyohei/ptcwave/v1/repository"
)

type SubPriceRequest struct {
	Base string `query:"base" validate:"required"`
	At   string `query:"at" validate:"required"`
}

type SubPriceResponse struct {
	Name          string `json:"name"`
	SubPrice      int    `json:"sub"`
	Version       string `json:"version"`
	Rarity        string `json:"rarity"`
	Type          string `json:"type"`
	PrevPrice     int    `json:"base_price"`
	NowPrice      int    `json:"at_price"`
	Status        string `json:"status"`
	PrevFetchDate string `json:"base_fetch_date"`
	NowFetchDate  string `json:"at_fetch_date"`
}

type SubPricesResponse struct {
	Cards []SubPriceResponse `json:"cards"`
}

func SubPrice(c echo.Context) error {
	repository, err := repository.NewSubPriceRepository(database.DB)
	var req SubPriceRequest
	if err := c.Bind(&req); err != nil {
		fmt.Println(err.Error())
		return err
	}

	var cards *[]entity.SubPrice
	cards, err = repository.FindSubPrice(req.Base, req.At, "rush")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	rs := make([]SubPriceResponse, 0, 500)
	for _, v := range *cards {
		rs = append(rs, SubPriceResponse{
			v.Name,
			v.SubPrice,
			v.Version,
			v.Rarity,
			v.Type,
			v.PrevPrice,
			v.NowPrice,
			v.Status,
			req.Base,
			req.At,
		})
	}

	return c.JSON(http.StatusOK, SubPricesResponse{rs})
}
