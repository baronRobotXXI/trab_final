package avl

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
  	. "trab_final/models"
	//dot import used here because all structures are
	//declared exclusively there
)


func Max (a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func Alt (t *Tree) int16 {
	if t == nil {
		return 0
	}

	return t.Altura
}

func NovoNo (d *Node) *Tree {
	t := &Tree {
		Dado:		d,
		Left:		nil,
		Right:		nil,
		Altura:		1,
	}

	return t
}

func RotacaoDireita (t *Tree) *Tree {
	var pont_a *Tree
	var pont_b *Tree

	pont_a = t.Left
	pont_b = nil

	if pont_a == nil {
		return t
	}
	if pont_a.Right != nil {
		pont_b = pont_a.Right
	}

	pont_a.Right = t
	t.Left = pont_b

	//altera balanceamento
	t.Altura = int16(Max(Alt(t.Left), Alt(t.Right)) + 1)
	pont_a.Altura = int16(Max (Alt(pont_a.Left), Alt(pont_a.Right)) + 1)

	return pont_a
}

func RotacaoEsquerda (t *Tree) *Tree {
	var pont_a *Tree
	var pont_b *Tree

	pont_a = t.Right
	pont_b = nil

	if pont_a == nil {
		return t
	}
	if pont_a.Left != nil {
		pont_b = pont_a.Left
	}

	pont_a.Left = t
	t.Right = pont_b

	//altera balanceamento
	t.Altura = int16(Max(Alt(t.Left), Alt(t.Right)) + 1)
	pont_a.Altura = int16(Max(Alt(pont_a.Left), Alt(pont_a.Right)) + 1)

	return pont_a
}

func getBalance (t *Tree) int16 {
	if t == nil {
		return 0
	}

	return Alt(t.Left) - Alt(t.Right)
}

func Inserir (t *Tree, d *Node) *Tree {
	var balance int16

	if t == nil {
		return (NovoNo(d))
	}

	if d.Isbn < t.Dado.Isbn {
		t.Left = Inserir(t.Left, d)
	} else if d.Isbn > t.Dado.Isbn {
		t.Right = Inserir(t.Right, d)
	} else {
		return t 						
		//se isbn igual, não insere
	}

	//atualizando o balnceamento de cada no

	t.Altura = Max(Alt(t.Left), Alt(t.Right)) + 1
	balance = getBalance(t)


	if balance > 1 && (d.Isbn < t.Left.Dado.Isbn) {
		return RotacaoDireita(t)
	}
	if balance < -1 && (d.Isbn > t.Right.Dado.Isbn) {
		return RotacaoEsquerda(t)
	}
	if balance > 1 && (d.Isbn > t.Left.Dado.Isbn) {
		t.Left = RotacaoDireita(t.Left)
		return RotacaoDireita(t)
	}
	if balance < -1 && (d.Isbn < t.Right.Dado.Isbn) {
		t.Right = RotacaoDireita(t.Right)
		return RotacaoEsquerda(t)
	}

	return t
}

func InsereFromArquivo (nomeArq string) (*Tree, error) {
	var t *Tree
	fmt.Println("*")

	arquivo, err := os.Open(nomeArq)
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir arquivo: %v", err)
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")

		if len(fields) < 6 {
			continue			//skip lines without proper fields
		}

		//parsing values
		ano, _ := strconv.ParseInt(fields[4], 10, 16)
		qtde, _ := strconv.ParseInt(fields[5], 10, 16)

		valor := &Node {
			Isbn:		fields[0],
			Titulo:		fields[1],
			Autor:		fields[2],
			Genero:		fields[3],
			Ano:		int16(ano),
			Qtde_disp:	int16(qtde),
		}

		t = Inserir(t, valor)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("erro na leitura do arquivo: %v", err)
	}

	return t, nil
}

func menorNo (t *Tree) *Tree {
	temp := t

	for (temp.Left != nil) {
		temp = temp.Left
	}

	return temp
}

func Remover (t *Tree, isbn string) *Tree {
	var temp *Tree
	var balance int16

	if t == nil { return t }

	if isbn < t.Dado.Isbn {
		t.Left = Remover(t.Left, isbn)
	} else if isbn > t.Dado.Isbn {
		t.Right = Remover(t.Right, isbn)
	} else {
		if t.Left == nil || t.Right == nil {
			if t.Left != nil {
				temp = t.Left
			} else {
				temp = t.Right
			}

			if temp == nil {
				temp = t
				t = nil
			} else {
				*t = *temp
			}

			temp = nil		//garbage collector does free(temp)
		} else {
			temp = menorNo(t.Right)
			t.Dado.Isbn = temp.Dado.Isbn
			t.Right = Remover(t.Right, temp.Dado.Isbn)
		}
	}

	if t == nil {
		return t
	}

	//atualiza balanceamento
	t.Altura = 1 + Max(Alt(t.Left), Alt(t.Right))
	balance = getBalance(t)

	if balance > 1 && getBalance(t.Left) >= 0 {
		return RotacaoDireita(t)
	}
	if balance > 1 && getBalance(t.Left) < 0 {
		t.Left = RotacaoEsquerda(t.Left)
		return RotacaoDireita(t)
	}
	if balance < -1 && getBalance(t.Right) <= 0 {
		return RotacaoEsquerda(t)
	}
	if balance < -1 && getBalance(t.Right) > 0 {
		t.Right = RotacaoDireita(t.Right)
		return RotacaoEsquerda(t)
	}

	return t
}

func SalvaArquivo (t *Tree, arq string) error {
	var save func(t *Tree) error

	f, err := os.Create(arq)
	if err != nil {
		return fmt.Errorf("error opening file %v", err)
	}
	defer f.Close()

	save = func(t *Tree) error {
		if t == nil {
			return nil
		}
	

		if err := save(t.Left); err != nil {
			return err
		}

		_, err := fmt.Fprintf(f, "%s,%s,%s,%s,%d,%d\n",
				t.Dado.Isbn, t.Dado.Titulo,
				 t.Dado.Autor, t.Dado.Genero,
				t.Dado.Ano, t.Dado.Qtde_disp)
		if err != nil {
			return err
		}

		return save(t.Right)
	}

	return save(t)
}

func PreOrdem (t *Tree) {
	if t == nil {
		return
	}

	fmt.Printf("Book: %s, %s\n", t.Dado.Isbn, t.Dado.Titulo)
	PreOrdem(t.Left)
	PreOrdem(t.Right)
}

func PosOrdem (t *Tree) {
	if t == nil	{
		return
	}

	PosOrdem(t.Left)
	PosOrdem(t.Right)
	fmt.Printf("Book: %s, %s\n", t.Dado.Isbn, t.Dado.Titulo)
}

func EmOrdem (t *Tree) {
	if t == nil {
		return
	}

	EmOrdem(t.Left)
	fmt.Printf("Book: %s, %s\n", t.Dado.Isbn, t.Dado.Titulo)
	EmOrdem(t.Right)
}
