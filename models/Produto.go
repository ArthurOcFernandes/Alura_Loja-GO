package models

import (
	base "arthur_loja/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := base.ConectaComBancoDeDados()

	selectTodosProdutos, err := db.Query("Select * from produtos order by id")

	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}

	for selectTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p := Produto{Id: id, Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade}
		produtos = append(produtos, p)

	}

	defer db.Close()

	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := base.ConectaComBancoDeDados()

	query := "INSERT INTO PRODUTOS (NOME, DESCRICAO, PRECO, QUANTIDADE) VALUES ($1, $2, $3, $4)"

	insercao, err := db.Prepare(query)

	if err != nil {
		panic(err.Error())
	}

	insercao.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func Delete(id string) {
	db := base.ConectaComBancoDeDados()

	deletarProduto, err := db.Prepare("Delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)

	defer db.Close()

}

func BuscaProdutoPorId(id string) Produto {

	db := base.ConectaComBancoDeDados()

	produtoQuery, err := db.Query("Select id, nome, descricao, preco, quantidade from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	var idProduto, quantidade int
	var nome, descricao string
	var preco float64

	if produtoQuery.Next() {
		err = produtoQuery.Scan(&idProduto, &nome, &descricao, &preco, &quantidade)
	}

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{Id: idProduto, Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade}

	defer db.Close()

	return produto

}

func Update(id int, nome, descricao string, preco float64, quantidade int) {

	db := base.ConectaComBancoDeDados()

	updateQuery, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateQuery.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
