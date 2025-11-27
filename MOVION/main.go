package main

import (
	config "movion/Config"
	routes "movion/Routes"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}


func main(){
	r:= gin.Default()
	r.Use(CORSMiddleware())

	config.InitDB()

	// load templates (all subfolders)
	//  r.LoadHTMLGlob("templates/*.html")
// r.LoadHTMLGlob("templates/**/*")


routes.RegisterRoutes(r)

	r.Run()
}