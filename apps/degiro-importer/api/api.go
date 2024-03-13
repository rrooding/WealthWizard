package api

import (
  "errors"
  "net/http"

  "github.com/Khan/genqlient/graphql"
)

type API struct {
  client graphql.Client
}

func Init(endpoint string) *API {
  return &API{
    client: graphql.NewClient(endpoint, http.DefaultClient),
  }
}

func (a *API) IsOK() (bool, error) {
  _ = `# @genqlient
    query IsOK {
      __typename
    }
  `

  resp, err := IsOK(a.client)
  if err != nil {
    return false, err
  }

  if resp.GetTypename() != "Query" {
    return false, errors.New("Invalid response from server")
  }

  return true, nil
}

func (a *API) CreateTransaction(input NewTransaction) (*CreateTransactionResponse, error) {
  _ = `# @genqlient
    mutation CreateTransaction($input: NewTransaction!) {
      createTransaction(input: $input) {
        isin
        broker
        brokerId
      }
    }
  `
  return CreateTransaction(a.client, input)
}
