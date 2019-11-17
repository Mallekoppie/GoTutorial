package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "changed"

	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	data := AgentData{}
	data.Agents = []string{"One", "Two"}
	c.RenderTemplate("App/Hello.html")
	c.ViewArgs["data"] = data
	return c.Render()
}

type Bla struct {
	Name   string ` json:"name" `
	Second string ` json:"second" `
}

type Stuff struct {
	Foo string ` json:"foo" xml:"foo" `
	Bar int    ` json:"bar" xml:"bar" `
}

func (c App) Test() revel.Result {
	//	data := make(map[string]interface{})
	//	data["error"] = nil
	//	stuff := Stuff{Foo: "xyz", Bar: 999}
	//	data["stuff"] = stuff
	//	return c.RenderJSON(data)

	test := Bla{Name: "bla", Second: "another"}

	return c.RenderJSON(test)
}

type AgentData struct {
	Agents []string `json:"agents"`
}

func (c App) Settings() revel.Result {
	c.RenderTemplate("App/Settings.html")

	return c.Render()
}

func (c App) UpdateSettings() revel.Result {
	return c.Render()
}
