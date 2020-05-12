
package infrastructure


import (
	"fmt"
	"time"
	"html/template"

	"github.com/gin-gonic/gin"

	"github.com/psychedelicnekopunch/gin-clean-architecture/app/interfaces/controllers"
)


type Routing struct {
	DB *DB
	Http *Http
	Gin *gin.Engine
	Port string
	AbsolutePath string
}


func NewRouting(db *DB, http *Http) *Routing {
	c := NewConfig()
	r := &Routing{
		DB: db,
		Http: http,
		Gin: gin.Default(),
		Port: c.Routing.Port,
		AbsolutePath: c.AbsolutePath,
	}
	r.loadTemplates()
	r.setRouting()
	return r
}


func (r *Routing) loadTemplates() {

	r.Gin.Static("/css", r.AbsolutePath + "/assets/css")
	r.Gin.Static("/js", r.AbsolutePath + "/assets/js")

	// 順番大事 フィルター定義 (SetFuncMap()) → テンプレート読込み (LoadHTMLFiles())
	r.Gin.SetFuncMap(template.FuncMap{
		"formatedDate": func (t time.Time) string {
			year, month, day := t.Date()
			return fmt.Sprintf("%d/%02d/%02d", year, month, day)
		},
		"test": func (s string) string {
			return s + " : filter success"
		},
		"sinitize": func (s string) template.HTML {
			return template.HTML(s)
		},
	})
	r.Gin.LoadHTMLFiles(
		r.AbsolutePath + "/app/interfaces/presenters/components/header.tmpl",
		r.AbsolutePath + "/app/interfaces/presenters/components/repos.tmpl",
		r.AbsolutePath + "/app/interfaces/presenters/index.tmpl",
		r.AbsolutePath + "/app/interfaces/presenters/assets/index.tmpl",
		r.AbsolutePath + "/app/interfaces/presenters/cookies/index.tmpl",
		r.AbsolutePath + "/app/interfaces/presenters/forms/index.tmpl",
		r.AbsolutePath + "/app/interfaces/presenters/templates/index.tmpl",
		r.AbsolutePath + "/app/interfaces/presenters/templates/error.tmpl",
	)
}


func (r *Routing) setRouting() {
	indexController := controllers.NewIndexController()
	jsonController := controllers.NewJsonController(r.DB)
	cookiesController := controllers.NewCookiesController()
	templatesController := controllers.NewTemplatesController(r.Http)
	formsController := controllers.NewFormsController()
	assetsController := controllers.NewAssetsController()

	r.Gin.GET("/", func (c *gin.Context) { indexController.Get(c) })

	/**
	 * SSL Sample
	 */
	r.Gin.GET("/.well-known/pki-validation/fileauth.txt", func(c *gin.Context) {
		c.String(200, "YOUR AUTH CODE HERE :)")
	})


	/**
	 * Google Ads Sample
	 */
	r.Gin.GET("/ads.txt", func(c *gin.Context) {
		c.String(200, "google.com, pub-xxxxxxxxxx, DIRECT, xxxxxxxxxx")
	})


	/**
	 * JSON Sample
	 */
	r.Gin.GET("/json/:id", func (c *gin.Context) { jsonController.Get(c) })


	/**
	 * View Sample
	 */
	r.Gin.GET("/templates", func (c *gin.Context) { templatesController.Get(c) })

	r.Gin.GET("/cookies", func (c *gin.Context) { cookiesController.Get(c) })
	r.Gin.POST("/cookies", func (c *gin.Context) { cookiesController.Post(c) })

	r.Gin.GET("/forms", func (c *gin.Context) { formsController.Get(c) })
	r.Gin.POST("/forms", func (c *gin.Context) { formsController.Post(c) })

	r.Gin.GET("/assets", func (c *gin.Context) { assetsController.Get(c) })
}


func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
