>>> ESPECIFICAÇÃO DO PROJETO <<<

Este projeto consiste em escrever um programa para o gerenciamento de uma biblioteca universitária utilizando estruturas de dados avançadas para otimizar as operações de busca, inservão e organização dos dados.

Para todas as operações pedidas, crie um sistema de menus para facilitar o acesso dentro de seu sistema. O menu pode ser baseado em texto.

PARTE 1: ÁRVORE AVL

1. Implementar uma árvore AVL para armazenar livros com os seguintes campos:
    * Código ISBN (chave)
    * Título
    * Autor
    * Ano de Publicação
    * Quantidade disponível

2. Implementar as seguintes operações:
    * Inserção de novo livro
    * Remoção de livro 
    * Busca por ISBN
    * Balanceamento automático
    * Exibição da árvore em pré-ordem, em ordem e pós-ordem
    * Carregar dados de arquivo texto
    * Salvar dados para arquivo texto

PARTE 2: GRAFO DE RELACIONAMENTOS 

1. Implementar um grafo utilizando lista de adjacências para representar as relações entre os livros: 
    * Vértices são livros 
    * Arestas representam relações (mesmo autor, mesmo gênero, etc) 
    * Peso das arestas indicam força da relação (1-5) 

2. Implementar os seguintes algoritmos: 
    * Busca em profundidade (DFS) 
    * Busca em largura (BFS) 
    * Dijkstra para encontrar o caminho mais forte entre dois livros. Note que é um problema de encontrar o caminho de custo máximo.  
    * Função para sugerir livros relacionados 
    * Carregar dados de arquivo texto 
    * Salvar dados para arquivo texto. 

PARTE 3: ALGORITMOS DE ORDENAÇÃO

1. Implementar três algoritmos de ordenação:
    * QuickSort
    * HeapSort
    * MergeSort

2. Criar funções para ordenar os livros por:
    * ISBN (com QuickSort)
    * Titulo (com HeapSort)
    * Ano de publicação (com MergeSort)

BONUS:
    * Implementar interface gráfica para visualização das estruturas
    * Implementar árvore Rubro-Negra além da AVL
