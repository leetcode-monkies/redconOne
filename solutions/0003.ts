/**
 Given a string s, find the length of the longest substring
 without repeating characters.
 * */

export default function lengthOfLongestSubstring(str: string): number {
  const dict = {};
  let largest = 0;

  for (let p1 = 0, p2 = 0; p2 < str.length; ) {
    if (p1 === p2 || !dict[str[p2]]) dict[str[p2++]] = 1;
    else dict[str[p1++]]--;
    largest = Math.max(largest, p2 - p1);
  }

  return largest;
}
