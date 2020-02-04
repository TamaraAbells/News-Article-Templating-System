package controllers

import (
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

// Controller Instance
type App struct {
	*revel.Controller
}

// home page renderer
func (c App) Index() revel.Result {
	return c.RenderTemplate("App/Index.html")
}

// Article Template renderer
func (c App) ArticleTemplate() revel.Result {
	return c.RenderTemplate("App/Article.html")
}

// Article Template data receiver; going to implement model interaction soon
func (c App) SubmitArticle() revel.Result {
	text := c.Params.Get("text")
	print(text)
	return c.RenderText("test")
}
