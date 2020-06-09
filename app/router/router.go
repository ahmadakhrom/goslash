package router

import (
	"crud_demo/app/controllers"
	"crud_demo/app/middlewares"
	"html/template"
	"io"

	//modules2 "crud_demo/migrates"
	"github.com/labstack/echo"
)

//for usage templates
// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

var Server = echo.New()

func SetRouter() {

	//html rendere
	//migrates.PathRenderer()
	Map := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"inc": func(i int) int {
			return i + 1
		},
	}
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseFiles(
			"app/views/index.html",
			"app/views/list-user.html",
			"app/views/insert-user.html",
			"app/views/view-user.html",
			"app/views/edit-user.html",
			"app/views/login.html",
			"app/views/logged/dashboard.html",
			"app/views/upload-data.html",

			//running templating
		)).Funcs(Map),
	}
	Server.Renderer = renderer

	// middleware
	Server.Use(middlewares.Cors())
	Server.Use(middlewares.Gzip())
	Server.Use(middlewares.Logger())
	Server.Use(middlewares.Secure())

	//public routes
	Server.GET("/", controllers.Index)
	Server.GET("/login", controllers.Login)
	Server.GET("/login/auth", controllers.LoginAuth)

	//routes users
	route := Server.Group("/user")

	route.GET("/dashboard", controllers.Dashboard)
	route.GET("/upload-data", controllers.UploadData)
	route.GET("/list", controllers.ListUsers)
	route.GET("/insert", controllers.NewUser)
	route.POST("/insert/store", controllers.NewUserStore)
	route.GET("/delete/:id", controllers.DeleteUser)
	route.GET("/:id", controllers.ShowUser)
	route.GET("/edit/:id", controllers.ShowUserByID)
	route.POST("/update/:id", controllers.UpdateUserStore)

}
