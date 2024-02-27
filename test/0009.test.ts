import { expect } from "chai";
import isPalindrome from "../solutions/0009";

describe("Testing LC 0009", () => {
  it("test case 1", () => {
    const testArg1 = 121;
    const result = isPalindrome(testArg1);
    const expected = true;
    expect(result).to.equal(expected);
  });
  it("test case 2", () => {
    const testArg1 = -121;
    const result = isPalindrome(testArg1);
    const expected = false;
    expect(result).to.equal(expected);
  });
  it("test case 3", () => {
    const testArg1 = 10;
    const result = isPalindrome(testArg1);
    const expected = false;
    expect(result).to.equal(expected);
  });
});
