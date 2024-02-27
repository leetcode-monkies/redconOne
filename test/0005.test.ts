import { expect } from "chai";
import longestPalindrome from "../solutions/0005";

describe("Testing LC 0005", () => {
  it("test case 1", () => {
    const testStr = "babad";
    const result = longestPalindrome(testStr);
    const expected = "bab";
    expect(result).to.equal(expected);
  });

  it("test case 2", () => {
    const testStr = "cbbd";
    const result = longestPalindrome(testStr);
    const expected = "bb";
    expect(result).to.equal(expected);
  });

  it("test case 3", () => {
    const testStr = "tacocat";
    const result = longestPalindrome(testStr);
    const expected = "tacocat";
    expect(result).to.equal(expected);
  });
});
