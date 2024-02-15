import { StocksTableCard, PortfolioValueCard } from "@wealth-wizard/web/stocks";

export default function StockPage() {
  return (
    <div>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <PortfolioValueCard />
      </div>
      <div>
        <StocksTableCard />
        </div>
    </div>
  );
}
