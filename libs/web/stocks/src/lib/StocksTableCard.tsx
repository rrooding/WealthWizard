import { Card, Table, TableHeader, TableHead, TableBody, TableRow, TableCell } from "@wealth-wizard/web/ui-components";

export function StocksTableCard() {
  return (
    <Card>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">Stock Symbol</TableHead>
            <TableHead>Quantity</TableHead>
            <TableHead>Price</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow>
            <TableCell className="font-medium">AAPL</TableCell>
            <TableCell>10</TableCell>
            <TableCell>$150.00</TableCell>
          </TableRow>
          <TableRow>
            <TableCell className="font-medium">GOOG</TableCell>
            <TableCell>5</TableCell>
            <TableCell>$2000.00</TableCell>
          </TableRow>
          <TableRow>
            <TableCell className="font-medium">MSFT</TableCell>
            <TableCell>15</TableCell>
            <TableCell>$250.00</TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </Card>
  )
}
