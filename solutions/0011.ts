export default function maxArea(height: number[]): number {
  const getArea = (l: number, r: number): number => {
    const currentHeight = Math.min(height[l], height[r]);
    const width = r - l;
    return width * currentHeight;
  };
  let left = 0;
  let right = height.length - 1;
  let max = 0;

  while (left < right) {
    const currentArea = getArea(left, right);
    max = Math.max(currentArea, max);

    height[left] < height[right] ? left++ : right--;
  }

  return max;
}
