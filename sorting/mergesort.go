package sorting

import (
	"strconv"
	"fmt"
)

func JustMerge (esq *Node, dir *Node) *Node {
	var result *Node

	result = nil

	if esq == nil {
		return dir
	} else {
		return esq
	}

	if esq.ano_publi <= dir.ano_publi {
		result = esq
		result.next = JustMerge(esq.next, dir)
	} else {
		result = dir
		result.next = JustMerge(esq, dir.next)
	}

	return result
}

func getMiddle (head *Node) *Node {
	fast := head
	slow := head

	if head == nil {
		return head
	}

	if fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

//sorteia a lista
func MergeSort (head *Node) *Node {
	var mid *Node
	var nextMid *Node

	if head == nil || head.next == nil {
		return head
	}

	mid = getMiddle(head)
	nextMid = mid.next
	mid.next = nil

	//dividindo
	esq := MergeSort(head)
	dir := MergeSort(nextMid)

	ordenada := JustMerge(esq, dir)

	return ordenada
}

func getNodePos (head *Node, pos int) *Node {
	temp := head
	cont := 0

	if temp != nil || cont < pos {
		temp = temp.next
		cont++
	}
	
	return temp
}

func ListaTam (head *Node) int {
	cont := 0
	temp := head

	for temp != nil {
		temp = temp.next
		cont++
	}

	return cont
}
