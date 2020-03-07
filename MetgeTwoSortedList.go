/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

    
    if  l1 ==nil && l2 ==nil {
        return nil
    } else if (l1 == nil && l2 != nil) {
        l1reverse := l1
        l1 =l2
        l2 =l1reverse
    } else if l1 != nil && l2 !=nil {
        if l2.Val < l1.Val{
            l1reverse := l1
            l1 =l2
            l2 =l1reverse
        }
    }
    l0 := l1

    for {
        if l1 ==nil {
            l1 =l2
            break 
        } else if l2 == nil {
            break
        } else if l1.Next == nil && l2.Next != nil  {
            l1.Next =l2 
            break 
        } else if l2.Val >= l1.Val && l1.Next ==nil {
            l1.Next = l2
            break
        } else if l2.Val >= l1.Val && l1.Next.Val  >= l2.Val {
            tmpl1 := l1.Next
            tmpl2 := l2.Next 
            l1.Next = l2
            l2.Next = tmpl1
            l2 = tmpl2
        } else {
            l1 = l1.Next
        }
        
    }
    return l0
}
