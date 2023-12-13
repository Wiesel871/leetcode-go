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

func toSlice(l *ListNode) []int {
    var nums []int

    for l != nil {
        nums = append(nums, l.Val)
        l = l.Next
    }

    return nums
}

func TestAddTwoNumbers1(t *testing.T) {
    l1 := fromSlice([]int{2, 4, 3})
    l2 := fromSlice([]int{5, 6, 4})
    l3 := addTwoNumbers(l1, l2)

    if !listEq(l3, fromSlice([]int{7, 0, 8})) {
        t.Error("got ", toSlice(l3), " expected ", []int{7, 0, 8})
    } else {
        t.Log("1. Success")
    }
}

func TestAddTwoNumbers2(t *testing.T) {
    l1 := fromSlice([]int{0})
    l2 := fromSlice([]int{0})
    l3 := addTwoNumbers(l1, l2)

    if !listEq(l3, fromSlice([]int{0})) {
        t.Error("got ", toSlice(l3), " expected ", []int{0})
    } else {
        t.Log("2. Success")
    }
}

func TestAddTwoNumbers3(t *testing.T) {
    l1 := fromSlice([]int{9, 9, 9, 9, 9, 9, 9})
    l2 := fromSlice([]int{9, 9, 9, 9})
    l3 := addTwoNumbers(l1, l2)
    l4 := addTwoNumbers(l2, l1)

    if !listEq(l3, fromSlice([]int{8, 9, 9, 9, 0, 0, 0, 1})) {
        t.Error("got ", toSlice(l3), " expected ", []int{8, 9, 9, 9, 0, 0, 0, 1})
    } else {
        t.Log("3. Success")
    }

    if !listEq(l4, fromSlice([]int{8, 9, 9, 9, 0, 0, 0, 1})) {
        t.Error("got ", toSlice(l4), " expected ", []int{8, 9, 9, 9, 0, 0, 0, 1})
    } else {
        t.Log("4. Success")
    }
}
