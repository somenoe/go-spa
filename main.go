package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Hello World!"),
		app.A().Href("/").Text("back to Home"),
	)
}

func (h *hello) OnPreRender(ctx app.Context) {
	ctx.Page().SetTitle("A Hello")
}

type home struct {
	app.Compo
}

func (h *home) Render() app.UI {
	return app.Div().Body(
		app.H1().Class("text-2xl font-bold text-red-500").Text("Home Page"),
		app.P().Class().Text("Quo usque tandem abutere, Catilina, patientia nostra!"),
		app.P().Class("text-sm").Text("small mak mak"),
		app.A().Href("/hello").Text("go to Hello"),

		app.Link().Href("output.css").Rel("stylesheet"),
	)
}

func (h *home) OnPreRender(ctx app.Context) {
	ctx.Page().SetTitle("Home")
}

func main() {
	app.Route("/", func() app.Composer { return &home{} })
	app.Route("/hello", func() app.Composer { return &hello{} })

	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite("dist", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err != nil {
		log.Fatal(err)
	}
}
