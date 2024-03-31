import { Card, CardHeader, CardTitle, CardContent } from "@wealth-wizard/web/ui-components";

export function PortfolioValueCard() {
  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between pb-2 space-y-0">
        <CardTitle className="text-sm font-medium">Total Value</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="text-2xl font-bold">$123,456.78</div>
        <p className="text-xs text-gray-500 dark:text-gray-400">+20.1% from last month</p>
      </CardContent>
    </Card>
  );
}
