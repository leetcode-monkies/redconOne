import { expect } from "chai";
import convert from "../solutions/0006";

describe("Testing LC 0006", () => {
  it("test case 1", () => {
    const testStr = "PAYPALISHIRING";
    const testRows = 3;
    const result = convert(testStr, testRows);
    const expected = "PAHNAPLSIIGYIR";
    expect(result).to.equal(expected);
  });

  it("test case 2", () => {
    const testStr = "PAYPALISHIRING";
    const testRows = 4;
    const result = convert(testStr, testRows);
    const expected = "PINALSIGYAHRPI";
    expect(result).to.equal(expected);
  });

  it("test case 3", () => {
    const testStr = "A";
    const testRows = 1;
    const result = convert(testStr, testRows);
    const expected = "A";
    expect(result).to.equal(expected);
  });
});
