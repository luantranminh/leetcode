// // https://leetcode.com/problems/sort-list/
// package main

// import (
// 	"fmt"
// )

// // ListNode Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// // space: O(n) time: O(nlogn)
// // func sortList(head *ListNode) *ListNode {
// // 	var arr []int

// // 	for head != nil {
// // 		arr = append(arr, head.Val)
// // 		head = head.Next
// // 	}

// // 	sort.Slice(arr, func(i, j int) bool {
// // 		return arr[i] < arr[j]
// // 	})

// // 	tmp := initListNode(arr)
// // 	return tmp
// // }

// // space: O(1) time: O(n^2)
// func sortList(head *ListNode) *ListNode {
// 	result := head

// 	for head != nil {
// 		tmpHead := head
// 		for tmpHead != nil {
// 			if head.Val > tmpHead.Val {
// 				tmpVal := head.Val
// 				head.Val = tmpHead.Val
// 				tmpHead.Val = tmpVal
// 			}

// 			tmpHead = tmpHead.Next
// 		}

// 		head = head.Next
// 	}

// 	return result
// }

// func initListNode(arr []int) *ListNode {
// 	var prvNode *ListNode
// 	var head *ListNode

// 	for _, v := range arr {
// 		newNode := ListNode{Val: v, Next: nil}

// 		if prvNode == nil {
// 			head = &newNode
// 		}

// 		if prvNode != nil {
// 			prvNode.Next = &newNode
// 		}

// 		prvNode = &newNode
// 	}

// 	return head
// }

// func printListNode(l *ListNode) {
// 	for l != nil {
// 		fmt.Println(l.Val)

// 		l = l.Next
// 	}
// }

// func main() {
// 	l := initListNode([]int{3, 2, 1, 9})

// 	z := sortList(l)

// 	printListNode(z)
// }
