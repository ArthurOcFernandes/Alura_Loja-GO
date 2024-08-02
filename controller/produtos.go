package controller

import (
	"arthur_loja/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosProdutos := models.BuscaTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	idDoProduto := r.URL.Query().Get("id")

	models.Delete(idDoProduto)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func Edit(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")

	produto := models.BuscaProdutoPorId(idProduto)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Erro na conversão do id para int:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preco para float64:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade para int:", err)
		}

		models.Update(idConvertido, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}
