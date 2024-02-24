export default function reverse(x: number): number {
  const sign = x >= 0 ? 1 : -1;
  let num = (x *= sign);
  let result = 0;

  while (num) {
    result *= 10;
    result += num % 10;
    num = Math.floor(num / 10);
  }

  return result * sign;
}
