package forms

import (
	"github.com/kirves/revel-forms/fields"
	"testing"
)

var (
	txt, psw, btn fields.FieldInterface
)

func TestFieldRender(t *testing.T) {
	field := fields.TextField("test")
	field.AddClass("test")
	field.AddClass("class")
	field.SetId("testId")
	field.SetParam("param1", "val1")
	field.AddCss("css1", "val1")
	field.SetStyle(fields.BASE)
	t.Log("Rendered field:", field.Render())
	txt = field
}

func TestPasswordRender(t *testing.T) {
	field := fields.PasswordField("test")
	field.AddClass("test")
	field.AddClass("class")
	field.SetId("testId")
	field.SetStyle(fields.BASE)
	t.Log("Rendered field:", field.Render())
	psw = field
}

func TestButtonRender(t *testing.T) {
	field := fields.SubmitButton("btn", "Click me!")
	field.SetStyle(fields.BASE)
	t.Log("Rendered button:", field.Render())
	btn = field
}

func TestFormRender(t *testing.T) {
	form := BaseForm(POST, "")
	form.AddField(txt)
	form.AddField(psw)
	form.AddField(btn)
	t.Log("Rendered form:", form.Render())
}
