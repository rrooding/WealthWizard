import { render, screen } from '@testing-library/react';
import { Money, MoneyItem_MoneyFragment } from './Money';
import { makeFragmentData } from "@wealth-wizard/web/graphql";

describe('Money', () => {
  it('renders the amount in the specified currency format', () => {
    const moneyProps = makeFragmentData({
      amount: 1000.23,
      currency: 'USD',
    }, MoneyItem_MoneyFragment);

    render(<Money money={moneyProps} />);

    expect(screen.getByText('US$ 1.000,23')).toBeInTheDocument();
  });
});
