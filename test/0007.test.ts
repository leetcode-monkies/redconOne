import { expect } from "chai";
import reverse from "../solutions/0007";

describe("Testing LC 0003", () => {
  it("test case 1", () => {
    const testNum = 123;
    const result = reverse(testNum);
    const expected = 321;
    expect(result).to.equal(expected);
  });

  it("test case 2", () => {
    const testNum = -123;
    const result = reverse(testNum);
    const expected = -321;
    expect(result).to.equal(expected);
  });

  it("test case 3", () => {
    const testNum = 120;
    const result = reverse(testNum);
    const expected = 21;
    expect(result).to.equal(expected);
  });
});
