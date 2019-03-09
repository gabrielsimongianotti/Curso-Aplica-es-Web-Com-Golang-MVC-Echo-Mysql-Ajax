package routers

import (
	"github.com/Curso-Aplica-es-Web-Com-Golang-MVC-Echo-Mysql-Ajax/controllers"
	"github.com/labstack/echo"
)

// App é uma instancia de echo
var App *echo.Echo

func init() {
	App = echo.New()
	//pagina inicial
	App.GET("/", controllers.Home)

	api := App.Group("/v1")

	api.POST("/insert", controllers.Inserir)
}
