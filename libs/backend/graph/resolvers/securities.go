package resolvers

import (
	"context"
	"wealth-wizard/backend/graph/model"
)

func Securities(ctx context.Context) ([]*model.Security, error) {
  var securities []*model.Security
  security := model.Security{
    ID: "1",
    Name: "Vanguard FTSE All-World UCITS ETF",
    Symbol: "VWRL",
  }

  securities = append(securities, &security)
  return securities, nil
}
