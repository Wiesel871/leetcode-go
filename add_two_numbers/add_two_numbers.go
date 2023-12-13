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
    lv, rv := 0, 0
    for l != nil || r != nil || carry > 0 {
        if (l != nil) {
            lv = l.Val
            l = l.Next
        } else {
            lv = 0
        }

        if (r != nil) {
            rv = r.Val
            r = r.Next
        } else {
            rv = 0
        }

        tail.Next = &ListNode{Val: (lv + rv + carry) % 10, Next: nil}
        carry = (lv + rv + carry) / 10
        tail = tail.Next
    }


    head = head.Next
    return head
}
