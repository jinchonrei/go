package myTemplate

import (
	"MyJson"
	"bytes"
	"fmt"
	"html/template"
	"strconv"
)

type Person struct {
	UserName string
	Email    string
}

const HTML = `
<!DOCTYPE html>
<html lang="en">
     <head>
        <meta charset="utf-8">
        <title>selected attribute</title>
    </head>
    <body>
        <form method="GET">
            <div>
                <label>Places:</label>
                <select id="places" name="places">
                    {{range .}}
                    <option value="{{.Value}}" id="{{.Id}}" {{if . }}selected{{end}}>{{.Text}}</option>
                    {{end}}
                </select>
            </div>
            <input type="submit" value="submit">
        </form>
    </body>
</html>
`

type Option struct {
	Value, Id, Text string
	Selected        bool
}

type Body struct {
	Path    string
	Options []Option
}

var placesPageTmpl *template.Template = template.Must(template.ParseFiles("../myTemplate/a.html"))

func GetTemplate() (buff string) {
	var test MyJson.Tests = MyJson.ReadFile()
	fmt.Println(test.List)

	options := []Option{
		/*	Option{"Value1", "Id1", "Text1", false},
			Option{"Value2", "Id2", "Text2", true},
			Option{"Value3", "Id3", "Text3", false},*/
	}

	for i := 0; i < len(test.List); i++ {
		options = append(options, Option{test.List[i], strconv.Itoa(i), test.List[i], false})
	}

	body := Body{
		test.Field,
		options,
	}

	buf := bytes.Buffer{}
	if err := placesPageTmpl.Execute(&buf, body); err != nil {
		fmt.Println("Failed to build page", err)
		buff = ""
	} else {
		fmt.Println(buf.String())
		buff = buf.String()
	}
	return
}
