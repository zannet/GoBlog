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
	// Set the static route to access
	router.Static("/static", "./Static")

	// Create the functions map for template rendering
	var funcMap = template.FuncMap{
		"escape": Helpers.Escape,
	}

	// Enable template loading
	html := template.Must(template.New("").Funcs(funcMap).ParseGlob("Templates/*"))
    router.SetHTMLTemplate(html)

    // Home path
	router.GET( "/",           Controllers.HomeController, Utils.AddTemplatingInfo, Views.HomeView)
	router.GET( "/post/:slug", Controllers.PostController, Utils.AddTemplatingInfo, Views.PostView)

	// Run the web server
    router.Run(":" + Config.Global.Service.Port)
}