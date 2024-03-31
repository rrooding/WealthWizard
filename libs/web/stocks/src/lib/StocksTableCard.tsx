import { gql, getFragmentData, type FragmentType } from "@wealth-wizard/web/graphql";
import { Card, Table, TableHeader, TableHead, TableBody, TableRow } from "@wealth-wizard/web/ui-components";
import { StocksTableRow } from "./StocksTableRow";

const SecurityItems_QueryFragment = gql(/* GraphQL */`
  fragment SecurityItems_QueryFragment on Query {
    securities {
      isin
      ...SecurityItem_SecurityFragment
    }
  }
`);

export function StocksTableCard(props: { readonly securities: FragmentType<typeof SecurityItems_QueryFragment>}) {
  const query = getFragmentData(SecurityItems_QueryFragment, props.securities)

  return (
    <Card>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">Stock Symbol</TableHead>
            <TableHead>Quantity</TableHead>
            <TableHead>Average Price</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {query.securities.map(security => <StocksTableRow security={security} key={security.isin} />)}
        </TableBody>
      </Table>
    </Card>
  );
}
