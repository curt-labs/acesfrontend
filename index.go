package main

import (
	"flag"
	"github.com/curt-labs/acesfrontend/app/controllers/part_controller"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	"html/template"
	"log"
	"net/http"
)

var (
	listenAddr = flag.String("por", "3001", "http listen address")
)

func main() {
	flag.Parse()
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{ //rendering middleware
		Directory:  "app/templates",
		Layout:     "main",
		Extensions: []string{".tmpl", ".html"},
		Funcs: []template.FuncMap{
			{},
		},
		Delims:          render.Delims{"{{", "}}"},
		Charset:         "UTF-8",
		IndentJSON:      true,
		HTMLContentType: "text/html",
	}))
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"https://*.foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	m.Group("(?i)/getPart", func(r martini.Router) {
		r.Get("", part_controller.Search)
		r.Post("", part_controller.Get)
	})

	m.Group("(?i)/findVehicle", func(r martini.Router) {
		r.Get("", part_controller.Search2)
		r.Post("", part_controller.Get)
	})

	log.Printf("Starting server on 127.0.0.1:%s\n", *listenAddr)
	log.Fatal(http.ListenAndServe(":"+*listenAddr, m))

}
