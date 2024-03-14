package tests

import (
	"bytes"
	"htmx/templates"
	"testing"
)

func TestAddSource(testing *testing.T) {
	t := templates.NewTemplates()
	t.AddSource("data")
	err := t.Parse()
	if err != nil {
		testing.Error("Error parsing templates ", err.Error())
	}
	buffer := bytes.NewBuffer([]byte{})
	data := make(map[string]interface{})
	data["Message"] = "Hello, World!"
	data["Title"] = "Htmx test"
	err = t.Execute(buffer, "index", data)
	if err != nil {
		testing.Error("Error executing template", err.Error())

	}
	// Print the buffer to the console
	println("Template : ", buffer.String())
}
