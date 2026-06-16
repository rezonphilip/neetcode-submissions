/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head1, head2 := list1, list2

    if head1 == nil && head2 == nil {
        return nil
    }

    if head1 == nil {
        return head2
    }

    if head2 == nil {
        return head1
    }



    var head *ListNode
    if head1.Val < head2.Val {
        head = head1
        head1 = head1.Next
    }else {
        head = head2
        head2 = head2.Next
    }
    curr := head
    for true {
        if head1 == nil && head2 == nil {
            return head
        }
        if head1 == nil {
            curr.Next = head2
            return head
        }

        if head2 == nil {
            curr.Next = head1
            return head
        }

        if head1.Val < head2.Val {
            curr.Next = head1
            curr = curr.Next
            head1 = head1.Next
        }else {
            curr.Next = head2
            curr = curr.Next
            head2 = head2.Next
        } 
    }
    return head
}