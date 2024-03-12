package problem0002

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	testcases := []struct {
		l1, l2, want *ListNode
	}{
		{
			l1:   &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			l2:   &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			l1:   &ListNode{Val: 0},
			l2:   &ListNode{Val: 0},
			want: &ListNode{Val: 0},
		},
		{
			l1: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{
				Val: 9,
				Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{
					Val:  9,
					Next: &ListNode{Val: 9},
				}}},
			}}},
			l2: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{
				Val:  9,
				Next: &ListNode{Val: 9},
			}}},
			want: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{
				Val: 9,
				Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{
					Val:  0,
					Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}},
				}}},
			}}},
		},
	}

	for _, tc := range testcases {
		result := AddTwoNumbers(tc.l1, tc.l2)

		if !compareLists(result, tc.want) {
			t.Errorf("AddTwoNumbers = %v, want: %v\n", result, tc.want)
		}
	}
}

func compareLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
