package trab_final

/*const ISBN_TAM = 15
const TITULO_TAM = 100
const AUTOR_TAM = 70
const TAM_GENERO = 20*/

type Node struct {
	Isbn		 string
	Titulo		 string
	Autor 		 string
	Genero 		 string
	Ano 		 int16
	Qtde_disp 	 int16
	Next 		 *Node			//para lista
}

type Tree struct {
	Dado		*Node
	Left		*Tree
	Right		*Tree
	Altura		int16
}