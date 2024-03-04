import { expect } from "chai";
import maxArea from "../solutions/0011";

describe("Testing LC 0011", () => {
  it("test case 1", () => {
    const testHeights = [1, 8, 6, 2, 5, 4, 8, 3, 7];
    const result = maxArea(testHeights);
    const expected = 49;
    expect(result).to.equal(expected);
  });

  it("test case 2", () => {
    const testHeights = [1, 1];
    const result = maxArea(testHeights);
    const expected = 1;
    expect(result).to.equal(expected);
  });
});
