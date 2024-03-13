export default function convert(s: string, numRows: number): string {
  let pointer = 0;
  let currentRow = 0;
  let dir = 1;
  let result: string[] = [];

  while (pointer < s.length) {
    result[currentRow] = (result[currentRow] || "") + s[pointer];
    currentRow += dir;

    if (currentRow === numRows - 1 || currentRow === 0) dir *= -1;

    pointer++;
  }

  return result.join("");
}
