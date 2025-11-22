package routes

import (
	controllers "movion/Controllers"
	middleware "movion/MiddleWare"
	constants "movion/const"
	// "time"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

    // public page: no middleware
    r.GET("/home", controllers.GetAllMoviesPage)


    // public routes
    r.POST("/signup", controllers.Signup)
    r.POST("/login", controllers.Login)

    // protected routes user,admin
    api := r.Group("/api")
    api.Use(middleware.RBACmiddleware(constants.User, constants.Admin))
    {
        api.GET("/profile", controllers.Profile)
        api.POST("/update", controllers.UpdateUser)
        api.POST("/logout", controllers.Logout)
    }

    // admin routes
    admin := r.Group("/admin")
    {
        admin.POST("/adminlogin", controllers.AdminLogin)

        protected := admin.Group("")
        protected.Use(middleware.RBACmiddleware(constants.Admin))
        {
            // user CRUD
            protected.GET("/users", controllers.GetAllUsers)
            protected.GET("/user/:id", controllers.GetUser)
            protected.POST("/edit/:id", controllers.EditUser)
            protected.POST("/create", controllers.CreateUser)
            protected.POST("/delete/:id", controllers.DeleteUser)
            protected.GET("/search", controllers.SearchUser)
            protected.POST("/block/:id/block", controllers.BlockUser)

            // movie CRUD (JSON API)
            protected.POST("/movie", controllers.CreateMovie)
            protected.GET("/getmovie/:id", controllers.GetMovie)
            protected.POST("/editmovie/:id", controllers.EditMovie)
            protected.POST("/deletemovie/:id", controllers.DeleteMovie)

            // screen CRUD
            protected.POST("/screen", controllers.CreateScreen)
            protected.GET("/getscreens", controllers.GetAllScreens)
            protected.POST("/editscreen/:id", controllers.EditScreen)
            protected.POST("/deletescreen/:id", controllers.DeleteScreen)
        }
    }
}

