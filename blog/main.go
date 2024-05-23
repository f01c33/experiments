package main

import (
	db "blog/data"
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"

	"github.com/spf13/viper"
)

// htmx
// posts (markdown)
// home (several posts) truncated posts
// categories
// search
// comments
// archive (?)
// blog map
// social buttons
// rss feed
// header & footer

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

// type Category struct {
// 	Url  string
// 	Name string
// }

// func newCategory(url, name string) Category {
// 	return Category{
// 		Url:  url,
// 		Name: name,
// 	}
// }

type Comment struct {
	Name string
	Date string
	Body string
}

func newComment(name, date, body string) Comment {
	return Comment{
		Name: name,
		Date: date,
		Body: body,
	}
}

type Post struct {
	Comments []Comment
	Body     string
}

func newPost(comments []Comment, body string) Post {
	return Post{
		Comments: comments,
		Body:     body,
	}
}

type Index struct {
	Posts []Post
}

func newIndex(posts []Post) Index {
	return Index{
		Posts: posts,
	}
}

//go:generate sqlc generate

//go:embed schema.sql
var ddl string

func main() {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	ctx := context.Background()

	rawDB, err := sql.Open("sqlite3", viper.GetString("db"))
	if err != nil {
		panic(err)
	}
	if _, err := rawDB.ExecContext(ctx, ddl); err != nil {
		panic(err)
	}
	DB := db.New(rawDB)
	_ = DB
	e := echo.New()

	e.Renderer = newTemplate()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		comments := []Comment{newComment("caue", "01/02/2024", "asiodjhkljfdshalksfjhasfkjdhasdfklj")}
		body := "askdçjfhaçskjdlfhçasdkfjhsdfalkçjhasdflkjdasfhlkjsdfahkldasfjhsdflkjhasfdlkjsdfhlkjasdfhkljsdfahsdfakljhsf"
		posts := []Post{newPost(comments, body), newPost(comments, body)}
		index := newIndex(posts)
		return c.Render(http.StatusOK, "index", index)
		// return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/post/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, c.Param("id"))
	})
	e.GET("/search/:q", func(c echo.Context) error {
		return c.String(http.StatusOK, c.Param("q"))
	})
	e.Logger.Fatal(e.Start(":8080"))
}
