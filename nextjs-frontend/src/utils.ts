export function dollarStringFormatter(price: number): string  {
  return new Intl.NumberFormat("en-US", {
    style: "currency",
    currency: "USD",
  }).format(price)
}

export function localeDateFormatter(dateString: string) {
  return new Date(dateString).toLocaleDateString("en-US", {
    weekday: "long",
    day: "2-digit",
    month: "2-digit",
    year: "numeric"
  });
}
