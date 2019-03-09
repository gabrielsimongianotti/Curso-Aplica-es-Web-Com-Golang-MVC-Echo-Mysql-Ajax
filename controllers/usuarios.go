package controllers

import (
	"fmt"
	"net/http"

	"github.com/Curso-Aplica-es-Web-Com-Golang-MVC-Echo-Mysql-Ajax/models"
	"github.com/labstack/echo"
)

//Home é a página inicial da minha aplicação
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world")
}
func Inserir2(c echo.Context) error {
	fmt.Println("ok")
	return c.String(http.StatusOK, "Hello, world")
}

func Inserir(c echo.Context) error {
	fmt.Println("inserir ok")
	nome := c.FormValue("nome")
	email := c.FormValue("email")
	fmt.Print(" nome: ")
	fmt.Println(nome)
	var usuario models.Usuarios
	usuario.Nome = nome
	usuario.Email = email
	fmt.Print(" usuario.nome: ")
	fmt.Println(nome)
	if nome != "" && email != "" {
		if _, err := models.UsuarioModel.Insert(usuario); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"menssagem": "Não possivel adicionar o registro no banco de dados",
			})
		}
		return c.JSON(http.StatusCreated, map[string]string{
			"menssagem": "O registro foi armazenado com sucesso!",
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"menssagem": "Os campos precisam ser passados!",
	})
}
