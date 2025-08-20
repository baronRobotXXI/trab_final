package sorting

import (
	//"strconv"
	//"fmt"
	. "trab_final/models"
)

func Troca (a, b *Node){
	a.Isbn, b.Isbn = b.Isbn, a.Isbn
	a.Titulo, b.Titulo = b.Titulo, a.Titulo
	a.Autor, b.Autor = b.Autor, a.Autor
	a.Genero, b.Genero = a.Genero, b.Genero
	a.Ano, b.Ano = b.Ano, a.Ano
	a.Qtde_disp, b.Qtde_disp = b.Qtde_disp, a.Qtde_disp
}

func Particiona (head *Node, esq int, dir int) *Node {
	i := head
	pivot := head

	for j := head; j != nil; j = j.Next{
		if j.Isbn < pivot.Isbn{
			Troca(i, j)
			i = i.Next
		}
	}

	Troca(head, i)
	return i
}

func QuickSort (start, end *Node){
	//sendo start não nulo, diferente de end 
	//e não tendo dado a volta
	if start != nil && start != end && start != end.Next{
		pivot := Particiona(start, 0, 0)
		QuickSort(start, pivot)
		QuickSort(pivot.Next, end)
	}

}
