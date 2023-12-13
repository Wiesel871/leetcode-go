package add_two_numbers

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0

    head := &ListNode{Val: 0, Next: nil}
    tail := head
    l, r := l1, l2
    for l != nil && r != nil {
        tail.Next = &ListNode{Val: (l.Val + r.Val + carry) % 10, Next: nil}
        carry = (l.Val + r.Val + carry) / 10
        tail = tail.Next
        l = l.Next
        r = r.Next
    }

    for l != nil {
        tail.Next = &ListNode{Val: (l.Val + carry) % 10, Next: nil}
        carry = (l.Val + carry) / 10
        tail = tail.Next
        l = l.Next
    }

    for r != nil {
        tail.Next = &ListNode{Val: (r.Val + carry) % 10, Next: nil}
        carry = (r.Val + carry) / 10
        tail = tail.Next
        r = r.Next
    }

    for carry > 0 {
        tail.Next = &ListNode{Val: carry % 10, Next: nil}
        carry /= 10
        tail = tail.Next
    }

    head = head.Next
    return head
}
