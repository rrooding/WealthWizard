import { gql } from "@wealth-wizard/web/graphql";
import { getClient } from "../../lib/apollo";
import { StocksTableCard } from "@wealth-wizard/web/stocks";

const CurrentlyHeldSecuritiesQueryDocument = gql(/* GraphQL */`
  query CurrentlyHeldSecuritiesQuery {
    ...SecurityItems_QueryFragment
  }
`);

export async function CurrentlyHeldSecuritiesContainer() {
  const { data } = await getClient().query({ query: CurrentlyHeldSecuritiesQueryDocument });

  return (<StocksTableCard securities={data} />);
}
