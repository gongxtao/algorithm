package array

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil { return nil }
	if l1 == nil && l2 != nil { return l2 }
	if l1 != nil && l2 == nil { return l1 }

	sum := l1.Val + l2.Val
	carry := sum / 10
	rl := &ListNode{Val: sum % 10}
	header := rl
	for {
		sum = carry
		lval, ok1 := getNextVal(l1)
		if ok1 {
			sum += lval
		}
		lval, ok2 := getNextVal(l2)
		if ok2 {
			sum += lval
		}
		if !ok2 && !ok1 && sum <= 0 {
			return header
		}

		carry = sum / 10
		rl.Next = &ListNode{Val: sum % 10}
		rl = rl.Next
	}
}

func getNextVal(l *ListNode) (val int, b bool) {
	if l.Next != nil {
		val = l.Next.Val
		l.Next = l.Next.Next
		b = true
	}
	return
}

