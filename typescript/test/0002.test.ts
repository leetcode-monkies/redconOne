import { expect } from "chai";
import ListNode from "@utils/ListNode";
import addTwoNumbers from "@solutions/0002";

describe("Testing LC 0002", () => {
  it("test case 1", () => {
    const testL1 = new ListNode(2, new ListNode(4, new ListNode(3)));
    const testL2 = new ListNode(5, new ListNode(6, new ListNode(4)));
    const result = addTwoNumbers(testL1, testL2);

    expect(result).to.deep.equal(
      new ListNode(7, new ListNode(0, new ListNode(8))),
    );
  });

  it("test case 2", () => {
    const testL1 = new ListNode(0);
    const testL2 = new ListNode(0);
    const result = addTwoNumbers(testL1, testL2);
    expect(result).to.deep.equal(new ListNode(0));
  });

  it("test case 3", () => {
    const testL1 = new ListNode(
      9,
      new ListNode(
        9,
        new ListNode(
          9,
          new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9)))),
        ),
      ),
    );
    const testL2 = new ListNode(
      9,
      new ListNode(9, new ListNode(9, new ListNode(9))),
    );
    const result = addTwoNumbers(testL1, testL2);

    expect(result).to.deep.equal(
      new ListNode(
        8,
        new ListNode(
          9,
          new ListNode(
            9,
            new ListNode(
              9,
              new ListNode(
                0,
                new ListNode(0, new ListNode(0, new ListNode(1))),
              ),
            ),
          ),
        ),
      ),
    );
  });
});
