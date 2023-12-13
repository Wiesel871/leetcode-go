package add_two_numbers

func addTwoNumbersRec(l *ListNode, r *ListNode, carry int) *ListNode {
    if (l == nil && r == nil && carry == 0) {
        return nil
    }

    val := 0
    var ln (*ListNode) = nil
    var rn (*ListNode) = nil

    if (l != nil) {
        val += l.Val
        ln = l.Next
    }
    if (r != nil) {
        val += r.Val
        rn = r.Next
    }
    val += carry
    res := &ListNode{Val: val % 10}
    res.Next = addTwoNumbersRec(ln, rn, val / 10)
    return res
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
    return addTwoNumbersRec(l1, l2, 0)
}
