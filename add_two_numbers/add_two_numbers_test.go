package add_two_numbers

import (
    "testing"
)

func fromSlice(nums []int) *ListNode {
    var head *ListNode
    var tail *ListNode

    for _, num := range nums {
        node := &ListNode{Val: num}
        if head == nil {
            head = node
            tail = node
        } else {
            tail.Next = node
            tail = node
        }
    }
    return head
}

func listEq(l1 *ListNode, l2 *ListNode) bool {
    for l1 != nil && l2 != nil {
        if l1.Val != l2.Val {
            return false
        }

        l1 = l1.Next
        l2 = l2.Next
    }

    return l1 == nil && l2 == nil
}

func testerIter(l []int, r []int, exp []int, t *testing.T) {
    res := addTwoNumbers(fromSlice(l), fromSlice(r))
    if !listEq(res, fromSlice(exp)) {
        t.Error("expected", exp, "got", res)
    }
}

func testerRec(l []int, r []int, exp []int, t *testing.T) {
    res := addTwoNumbers2(fromSlice(l), fromSlice(r))
    if !listEq(res, fromSlice(exp)) {
        t.Error("expected", exp, "got", res)
    }
}

func TestAddTwoNumbers1(t *testing.T) {
    testerIter([]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}, t)
    testerIter([]int{0}, []int{0}, []int{0}, t)
    testerIter([]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}, t)
    testerIter([]int{9, 9, 9, 9}, []int{9, 9, 9, 9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}, t)

    testerRec([]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}, t)
    testerRec([]int{0}, []int{0}, []int{0}, t)
    testerRec([]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}, t)
    testerRec([]int{9, 9, 9, 9}, []int{9, 9, 9, 9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}, t)
}

