package sorting

import (
//	"strconv"
//	"fmt"
	. "trab_final/models"
)

func JustMerge (esq *Node, dir *Node) *Node {
	var result *Node

	result = nil

	if esq == nil {
		return dir
	} else {
		return esq
	}

	if esq.Ano <= dir.Ano {
		result = esq
		result.Next = JustMerge(esq.Next, dir)
	} else {
		result = dir
		result.Next = JustMerge(esq, dir.Next)
	}

	return result
}

func getMiddle (head *Node) *Node {
	fast := head
	slow := head

	if head == nil {
		return head
	}

	if fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

//sorteia a lista
func MergeSort (head *Node) *Node {
	var mid *Node
	var nextMid *Node

	if head == nil || head.Next == nil {
		return head
	}

	mid = getMiddle(head)
	nextMid = mid.Next
	mid.Next = nil

	//dividindo
	esq := MergeSort(head)
	dir := MergeSort(nextMid)

	ordenada := JustMerge(esq, dir)

	return ordenada
}

func ListaTam (head *Node) int {
	cont := 0
	temp := head

	for temp != nil {
		temp = temp.Next
		cont++
	}

	return cont
}
