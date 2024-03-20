package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"wealth-wizard/api/models"

	"github.com/shopspring/decimal"
)

// Securities is the resolver for the securities field.
func (r *queryResolver) Securities(ctx context.Context) ([]*models.Security, error) {
	var results []struct {
		Isin         string
		Broker       string
		Exchange     string
		Currency     string
		Amount       int
		AveragePrice decimal.Decimal
	}

	result := r.DB.Table("stocks_transactions").
		Select("isin, broker, exchange, currency, SUM(amount) AS amount, COALESCE(SUM(amount * price) / NULLIF(SUM(amount), 0), 0) AS average_price").
		Group("isin, broker, exchange, currency").
		Having("SUM(amount) <> 0").
		Scan(&results)

	if result.Error != nil {
		return nil, result.Error
	}

	var securities []*models.Security

	for _, res := range results {
		security := &models.Security{
			Isin:     res.Isin,
			Broker:   res.Broker,
			Exchange: res.Exchange,
			Amount:   res.Amount,
			AveragePrice: &models.Money{
				Amount:   res.AveragePrice,
				Currency: res.Currency,
			},
		}

		securities = append(securities, security)
	}

	return securities, nil
}
