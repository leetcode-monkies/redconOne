export default function myAtoi(s: string): number {
  const trimmed = s.trimStart();
  const sign = /-/.test(trimmed[0]) ? -1 : 1;
  let pointer = /[-+]/.test(trimmed[0]) ? 1 : 0;
  let result = 0;

  while (/\d/.test(trimmed[pointer]))
    result = result * 10 + +trimmed[pointer++];

  result *= sign;

  if (result > 2 ** 31 - 1) return 2 ** 31 - 1;
  if (result < -(2 ** 31)) return -(2 ** 31);

  return result;
}
