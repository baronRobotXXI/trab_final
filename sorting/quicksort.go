package sorting

import (
	"strconv"
	"fmt"
)

func Troca (a, b *Node){
	a.isbn, b.isbn = b.isbn, a.isbn
	a.titulo, b.titulo = b.titulo, a.titulo
	a.autor, b.autor = b.autor, a.autor
	a.genero_area, b.genero_area = a.genero_area, b.genero_area
	a.ano_publi, b.ano_publi = b.ano_publi, a.ano_publi
	a.qtde_disp, b.qtde_disp = b.qtde_disp, a.qtde_disp
}

func Particiona (head *Node, esq int, dir int) *Node {
	i := head
	pivot := head

	for j := head; j != nil; j = j.next{
		if j.isbn < pivot.isbn{
			Troca(i, j)
			i = i.next
		}
	}

	Troca(head, i)
	return i
}

func QuickSort (start, end *Node){
	//sendo start não nulo, diferente de end 
	//e não tendo dado a volta
	if start != nil && start != end && start != end.next{
		pivot := Particiona(start, 0, 0)
		QuickSort(start, pivot)
		QuickSort(pivot.next, end)
	}

}