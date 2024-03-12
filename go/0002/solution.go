package problem0002

type ListNode struct {
	Next *ListNode
	Val  int
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head    *ListNode
		current *ListNode
		carry   = 0
	)

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if head == nil {
			head = &ListNode{Val: sum % 10}
			current = head
		} else {
			current.Next = &ListNode{Val: sum % 10}
			current = current.Next
		}

		carry = sum / 10
	}

	return head
}
