package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/Nokoyohei/ptcwave/v1/database"
	"github.com/Nokoyohei/ptcwave/v1/entity"
	"github.com/Nokoyohei/ptcwave/v1/repository"
)

type CardRequest struct {
	FetchDate string `query:"fetch_date"`
}

type CardResponse struct {
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	Rarity    string    `json:"rarity"`
	Type      string    `json:"type"`
	Price     int       `json:"price"`
	Status    string    `json:"status"`
	FetchDate time.Time `json:"fetch_date"`
}

type CardsResponse struct {
	Cards []CardResponse `json:"cards"`
}

func Cardlist(c echo.Context) error {
	repository, err := repository.NewCardlistRepository(database.DB)
	var req CardRequest
	if err := c.Bind(&req); err != nil {
		fmt.Println(err.Error())
		return err
	}

	var cards *[]entity.Card
	if req.FetchDate == "" {
		cards, err = repository.FindAll()
	} else {
		cards, err = repository.FindByFetchDate(req.FetchDate)
	}

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	rs := make([]CardResponse, 0, 500)
	for _, v := range *cards {
		rs = append(rs, CardResponse{
			v.Name,
			v.Version,
			v.Rarity,
			v.Type,
			v.Price,
			v.Status,
			v.FetchDate,
		})
	}

	return c.JSON(http.StatusOK, CardsResponse{rs})
}
