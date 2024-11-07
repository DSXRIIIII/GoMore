package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left int, right int) *ListNode {
	if left > right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	var mid int = (left + right) >> 1
	return mergeTwo(merge(lists, left, mid), merge(lists, mid+1, right))
}

func mergeTwo(node1 *ListNode, node2 *ListNode) *ListNode {
	dummy := ListNode{}
	cur := &dummy
	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			cur.Next = node1
			node1 = node1.Next
		} else {
			cur.Next = node2
			node2 = node2.Next
		}
		cur = cur.Next
	}
	if node1 == nil {
		cur.Next = node2
	}
	if node2 == nil {
		cur.Next = node1
	}
	return dummy.Next
}
