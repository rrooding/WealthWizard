import { gql, getFragmentData, type FragmentType } from "@wealth-wizard/web/graphql";
import { Decimal } from 'decimal.js';

export const MoneyItem_MoneyFragment = gql(/* GraphQL */`
  fragment MoneyItem_MoneyFragment on Money {
    amount
    currency
  }
`);

export function Money(props: { readonly money: FragmentType<typeof MoneyItem_MoneyFragment> }) {
  const { currency, ...money } = getFragmentData(MoneyItem_MoneyFragment, props.money)
  const amount = new Decimal(money.amount).toNumber();

  const text = amount.toLocaleString('nl-NL', { style: "currency", currency });

  return (
    <span>{text}</span>
  );
}
