import { gql, getFragmentData, type FragmentType } from "@wealth-wizard/web/graphql";
import { TableRow, TableCell, Money } from "@wealth-wizard/web/ui-components";

const SecurityItem_SecurityFragment = gql(/* GraphQL */`
  fragment SecurityItem_SecurityFragment on Security {
    isin
    amount
    averagePrice {
      ...MoneyItem_MoneyFragment
    }
  }
`);

export function StocksTableRow(props: { security: FragmentType<typeof SecurityItem_SecurityFragment> }) {
  const security = getFragmentData(SecurityItem_SecurityFragment, props.security)

  return (
    <TableRow>
      <TableCell className="font-medium">{security.isin}</TableCell>
      <TableCell>{security.amount}</TableCell>
      <TableCell><Money money={security.averagePrice} /></TableCell>
    </TableRow>
  );
}
