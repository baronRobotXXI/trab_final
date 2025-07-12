package avl

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"github.com/baronRobotXXI/trab_final"
)

type Tree struct {
	Dado 		*Node
	Left		*Tree
	Right		*Tree
	Altura		int16
}

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
	var t *Tree

	t.Dado = d
	t.Left = nil
	t.Right = nil
	t.Altura = 1

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
	pont_a.Altura = int16(Max(Alt(pont_a.Left), Alt(pont_a.Right)) + 1)

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

