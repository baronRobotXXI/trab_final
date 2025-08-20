package sorting

import (
//	"fmt"
	. "trab_final/models"
	//"trab_final/sorting"
)

func getNodePos (head *Node, pos int) *Node {
	temp := head
	cont := 0

	if temp != nil || cont < pos {
		temp = temp.Next
		cont++
	}
	
	return temp
}

func HeapMaximusPos (head *Node, pos int, heapSize int) {
	esq := (pos * 2) + 1
	dir := (pos * 2) + 2
	maior := pos
	nPos := getNodePos (head, pos)
	nEsq := getNodePos (head, esq)
	nDir := getNodePos (head, dir)

	var a, b *Node

	var nMaior *Node

	if nPos == nil {
		return
	}

	if esq <= heapSize && nEsq != nil && (nEsq.Titulo > nPos.Titulo){
		maior = esq
	} else {
		maior = pos
	}

	nMaior = getNodePos (head, maior)

	a = getNodePos(head, pos)
	b = getNodePos(head, maior)

	if dir <= heapSize && nDir != nil && (nDir.Titulo > nMaior.Titulo){
		maior = dir
	}
	if maior != pos {
		Troca(a, b)
		HeapMaximusPos(head, maior, heapSize)
	}
}

func BuildHeapMaximus (head *Node, fim int){
	var i int

	for i = (fim - 1) / 2; i >= 0; i-- {
		HeapMaximusPos(head, i, fim)
	}
}

func HeapSort (head *Node) {
	var i int
	heapSize := ListaTam(head) - 1

	BuildHeapMaximus(head, heapSize)

	for i = heapSize; i >= 0; i-- {
		a := getNodePos(head, i)
		Troca(head, a)
		heapSize--
		HeapMaximusPos(head, 0, heapSize)
	}
}
