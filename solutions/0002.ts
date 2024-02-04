/**You are given two non-empty linked lists representing two non-negative 
 * integers. The digits are stored in reverse order, and each of their nodes 
 * contains a single digit. Add the two numbers and return the sum as a 
 * linked list.

 * You may assume the two numbers do not contain any leading zero, except the 
 * number 0 itself.
*/

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
