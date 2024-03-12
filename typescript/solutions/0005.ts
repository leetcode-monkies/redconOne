export default function longestPalindrome(s: string): string {
  const isPalindrome = (left: number, right: number): boolean => {
    while (left < right) {
      if (s[left] !== s[right]) return false;
      left++;
      right--;
    }
    return true;
  };

  for (let len = s.length; len >= 0; len--)
    for (let start = 0; start + len < s.length; start++)
      if (isPalindrome(start, start + len))
        return s.slice(start, start + len + 1);

  return "";
}
