import { expect } from "chai";
import myAtoi from "../solutions/0008";

describe("Testing LC 0008", () => {
  it("test case 1", () => {
    const testStr = "42";
    const result = myAtoi(testStr);
    const expected = 42;
    expect(result).to.equal(expected);
  });

  it("test case 2", () => {
    const testStr = "-42";
    const result = myAtoi(testStr);
    const expected = -42;
    expect(result).to.equal(expected);
  });

  it("test case 3", () => {
    const testStr = "4193 with words";
    const result = myAtoi(testStr);
    const expected = 4193;
    expect(result).to.equal(expected);
  });
});
