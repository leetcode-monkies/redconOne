import { expect } from "chai";
import twoSum from "./../solutions/0001.js";

describe("Testing LC 0001", () => {
  it("test case 1", () => {
    const testNums = [2, 7, 11, 15];
    const testTarget = 9;
    const result = twoSum(testNums, testTarget);
    const expected = [0, 1];
    expect(result).to.have.members(expected);
  });

  it("test case 2", () => {
    const testNums = [3, 2, 4];
    const testTarget = 6;
    const result = twoSum(testNums, testTarget);
    const expected = [1, 2];
    expect(result).to.have.members(expected);
  });

  it("test case 3", () => {
    const testNums = [3, 3];
    const testTarget = 6;
    const result = twoSum(testNums, testTarget);
    const expected = [0, 1];
    expect(result).to.have.members(expected);
  });
});
