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
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
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
