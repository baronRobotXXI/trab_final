package main

import (
	"fmt"
	"os"
//	"strings"
	 "trab_final/avl"
	. "trab_final/models"
	
)

/*type Node struct {
	Isbn		string
	Titulo		string
	Autor		string
	Genero		string
	Ano			int16
	Qtde_disp	int16
	Next		*Node			//para lista
}*/

func main() {
	var t *Tree
	var opcao int16
	var isbn string
	var novo Node = Node{}
	var err error
	//var temp int16

	fmt.Println("Querido usuario, digite a opcao desejada:")
	fmt.Println("1. Inserir com arquivo")
	fmt.Println("2. Imprimir pre ordem")
	fmt.Println("3. Imprimir em ordem")
	fmt.Println("4. Imprimir pos ordem")
	fmt.Println("5. Inserir Elemento")
	fmt.Println("6. Apagar Elemento")
	fmt.Println("7. Salvar no Arquivo")
	fmt.Println("8. Encerrar")
	fmt.Scanf("%d", &opcao)

	for  {
		switch opcao {
		case 1:
			t, err = avl.InsereFromArquivo("records.txt")
			if err == nil {
				fmt.Println("File opened successfully.")
			}
			break

		case 2:
			fmt.Println("A arvore em pré ordem é:")
			avl.PreOrdem(t)
			break

		case 3:
			fmt.Println("A arvore em pós ordem é:")
			avl.PosOrdem(t)
			break

		case 4:
			fmt.Println("A arvore em ordem é:")
			avl.EmOrdem(t)
			break

		case 5:
			fmt.Println("Digite o ISBN")
			fmt.Scanf("%s", novo.Isbn)
			fmt.Println("Digite o titulo")
			fmt.Scanf("%s", novo.Titulo)
			fmt.Println("Digite o autor")
			fmt.Scanf("%s", novo.Autor)
			fmt.Println("Digite o genero")
			fmt.Scanf("%s", novo.Genero)
			fmt.Println("Digite o ano da publicacão")
			fmt.Scanf("%d", novo.Ano)
			fmt.Println("Digite a quantidade")
			fmt.Scanf("%d", novo.Qtde_disp)

			t = avl.Inserir(t, &novo)

			break

		case 6:
			fmt.Println("Digitar o ISBN a ser apagado")
			fmt.Scanf("%s",isbn)
			t = avl.Remover(t, isbn)

			break

		case 7:
			if err := avl.SalvaArquivo(t, "records.txt"); err != nil {
				fmt.Println("Erro ao salvar arquivo!", err)
			}
			break

		case 8:
			os.Exit(0)

		default:
			fmt.Println("Opcao invalida")
		}

	fmt.Println("Querido usuario, digite a opcao desejada:")
	fmt.Println("1. Inserir com arquivo")
	fmt.Println("2. Imprimir pre ordem")
	fmt.Println("3. Imprimir em ordem")
	fmt.Println("4. Imprimir pos ordem")
	fmt.Println("5. Inserir Elemento")
	fmt.Println("6. Apagar Elemento")
	fmt.Println("7. Salvar no Arquivo")
	fmt.Println("8. Encerrar")
	fmt.Scanf("%d", &opcao)		
	}
}
