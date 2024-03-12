export default function isPalindrome(x: number): boolean {
  const reverser = (num: number): number => {
    let result = 0;

    while (num) {
      result = result * 10 + (num % 10);
      num = Math.floor(num / 10);
    }

    return result;
  };
  if (x < 0) return false;

  return reverser(x) === x;
}
