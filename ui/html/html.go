package html

import (
	"fmt"
	"github.com/minhnghia2k3/personal-blog/ui"
	"html/template"
	"io"
	"strings"
)

var (
	home = parse("html/pages/home.html")
)

func parse(file string) *template.Template {
	funcs := template.FuncMap{
		"uppercase": func(v string) string {
			fmt.Println("v:", v)
			return strings.ToUpper(v)
		},
	}
	patterns := []string{
		"html/layout.html",
		file,
	}
	return template.Must(template.New("layout.html").Funcs(funcs).ParseFS(ui.Files, patterns...))
}

type HomeParams struct {
	Title string
}

func HomeShow(w io.Writer, p HomeParams) error {
	return home.Execute(w, p)
}
