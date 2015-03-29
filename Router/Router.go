package Router

import (
	"html/template"
	"github.com/gin-gonic/gin"
	"../Views"
	"../Controllers"
	"../Utils"
	"../Helpers"
	"../Config"
)

func Init() {
	// Init a new Router
	router := gin.Default()

	router.Static("/static", "./Static")

	// Create the functions map for template rendering
	var funcMap = template.FuncMap{
		"escape": Helpers.Escape,
	}

	// Enable template loading
	html := template.Must(template.New("").Funcs(funcMap).ParseGlob("Templates/*"))
    router.SetHTMLTemplate(html)

	router.GET( "/", Controllers.HomeController, Utils.AddTemplatingInfo, Views.HomeView)

	// Run in the port 8080
    router.Run(":" + Config.Global.Service.Port)
}