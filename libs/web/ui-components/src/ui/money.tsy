import { gql, getFragmentData, type FragmentType } from "@wealth-wizard/web/graphql";
import { Decimal } from 'decimal.js';

const MoneyItem_MoneyFragment = gql(/* GraphQL */`
  fragment MoneyItem_MoneyFragment on Money {
    amount
    currency
  }
`);

export function Money(props: { money: FragmentType<typeof MoneyItem_MoneyFragment> }) {
  const { currency, ...money } = getFragmentData(MoneyItem_MoneyFragment, props.money)
  const amount = new Decimal(money.amount).toNumber();

  const text = amount.toLocaleString("nl-US", { style: "currency", currency });

  return (
    <span>{text}</span>
  );
}
