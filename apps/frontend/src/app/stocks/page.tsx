import { PortfolioValueCard } from "@wealth-wizard/web/stocks";
import { CurrentlyHeldSecuritiesContainer } from "./CurrentlyHeldSecuritiesContainer";

export default async function StockPage() {
  return (
    <div>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <PortfolioValueCard />
      </div>
      <div>
        <CurrentlyHeldSecuritiesContainer />
      </div>
    </div>
  );
}
