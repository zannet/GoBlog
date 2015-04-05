package Router

import (
	"html/template"
	"github.com/gin-gonic/gin"
	"../Views"
	"../Controllers"
	"../Middlewares"
	"../Helpers"
	"../Config"
)

func Init() {
	// Init a new Router
	router := gin.Default()
	// Set the static route to access
	router.Static("/static", "./Static")
	router.Static("/admin", "./Admin")

	// Create the functions map for template rendering
	var funcMap = template.FuncMap{
		"escape": Helpers.Escape,
	}

	// Enable template loading
	html := template.Must(template.New("").Funcs(funcMap).ParseGlob("Templates/*"))
    router.SetHTMLTemplate(html)

    // Home path
	router.GET( "/",           Controllers.HomeController,  Middlewares.AddTemplatingInfo, Views.HomeView)
	router.GET( "/post/:slug", Controllers.PostController,  Middlewares.AddTemplatingInfo, Views.PostView)

	// API Paths
	router.GET ("/api/posts",       Controllers.PostListController,   Views.PostListAPIView)
	router.GET ("/api/post/:slug",  Controllers.PostController,       Views.PostAPIView)
	router.PUT ("/api/post/:slug",  Controllers.PostUpdateController, Views.PostAPIView)

	// Run the web server
    router.Run(":" + Config.Global.Service.Port)
}