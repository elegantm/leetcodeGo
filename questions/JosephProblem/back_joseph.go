package leetcode

import (
	"fmt"
)

// linkList
type linkList struct {
	val  int
	next *linkList
}

func josephProblem(n, m int) {
	head := createLinkList(n)
	printLinkList(head)
	fmt.Println("the result is:")
	result := joseph(head, m)
	printLinkList(result)
}

func joseph(head *linkList, m int) *linkList {
	cur := head
	for cur.next != cur {
		for i := 1; i < m-1; i++ {
			cur = cur.next
		}
		cur.next = cur.next.next
		cur = cur.next
	}
	return cur
}

func createLinkList(n int) *linkList {
	head := &linkList{}
	cur := head

	for i := 1; i <= n; i++ {
		tmp := &linkList{
			val: i,
		}
		cur.next = tmp
		cur = cur.next
	}
	cur.next = head.next
	return head.next

}

func printLinkList(head *linkList) {
	cur := head
	for cur.next != head {
		fmt.Printf("%d -> ", cur.val)
		cur = cur.next
	}
	fmt.Printf("%d ", cur.val)
	fmt.Println()

}
