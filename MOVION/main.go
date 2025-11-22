package main

import (
	config "movion/Config"
	routes "movion/Routes"

	"github.com/gin-gonic/gin"
)

func main(){
	r:= gin.Default()

	config.InitDB()

	// load templates (all subfolders)
	 r.LoadHTMLGlob("templates/*.html")
// r.LoadHTMLGlob("templates/**/*.html")




	routes.RegisterRoutes(r)

	r.Run()
}