import { expect } from "chai";
import lengthOfLongestSubstring from "../solutions/0003";

describe("Testing LC 0003", () => {
  it("test case 1", () => {
    const testStr = "abcabcbb";
    const result = lengthOfLongestSubstring(testStr);
    const expected = 3;
    expect(result).to.equal(expected);
  });

  it("test case 2", () => {
    const testStr = "bbbbb";
    const result = lengthOfLongestSubstring(testStr);
    const expected = 1;
    expect(result).to.equal(expected);
  });

  it("test case 3", () => {
    const testStr = "pwwkew";
    const result = lengthOfLongestSubstring(testStr);
    const expected = 3;
    expect(result).to.equal(expected);
  });
});
