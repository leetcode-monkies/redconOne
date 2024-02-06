/** Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order. */

export default function twoSum(nums: number[], target: number): number[] {
  const dict = {};

  for (const [idx, num] of nums.entries())
    if (dict[target - num] !== undefined) return [dict[target - num], idx];
    else dict[num] = idx;

  return [];
}
