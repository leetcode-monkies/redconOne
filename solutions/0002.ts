import ListNode from "./../utils/ListNode.js";

const addTwoNumbers = (
  l1: ListNode<number> | null,
  l2: ListNode<number> | null,
) => {
  if (!l1) return l2;
  if (!l2) return l1;
  let carry = 0;
  const result: ListNode<number> = new ListNode(0);
  let current = result;

  while (l1 || l2) {
    const val = (l1 ? l1.val : 0) + (l2 ? l2.val : 0) + carry;
    carry = Math.floor(val / 10);
    current.next = new ListNode(val % 10);
    current = current.next;
    if (l1) l1 = l1.next;
    if (l2) l2 = l2.next;
  }
  if (carry) current.next = new ListNode(carry);

  return result.next;
};

export default addTwoNumbers;
