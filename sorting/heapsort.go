package sorting

import (
	"fmt"
	"trab_final/sorting"
)

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

	if esq <= heapSize && nEsq != nil && (nEsq.titulo > nPos.titulo){
		maior = esq
	} else {
		maior = pos
	}

	nMaior = getNodePos (head, maior)

	a = getNodePos(head, pos)
	b = getNodePos(head, maior)

	if dir <= heapSize && nDir != nil && (nDir.titulo > nMaior.titulo){
		maior = dir
	}
	if maior != pos {
		quicksort.Troca(a, b)
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

	for i = heapsize; i >= 0; i-- {
		a := getNodePos(head, i)
		quicksort.Troca(head, a)
		heapSize--
		HeapMaximusPos(head, 0, heapSize)
	}
}
