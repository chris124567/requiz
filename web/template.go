package web

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func (h handler) writeTemplateData(writer http.ResponseWriter, path string, data interface{}) error {
	name := filepath.Base(path)
	funcs := template.FuncMap{
		"json": func(x interface{}) template.JS {
			data, _ := json.Marshal(x)
			return template.JS(string(data))
		},
		"timestamp": func(x int64) string {
			return time.Unix(x, 0).Format("January 2, 2006")
		},
		"replaceNewline": func(s string) template.HTML {
			return template.HTML(strings.ReplaceAll(strings.ReplaceAll(template.HTMLEscapeString(s), "\n", "<br>"), "\t", "    "))
		},
	}

	template, err := template.New(name).Funcs(funcs).ParseFiles(path, filepath.Join(h.template, "base.html.tmpl"))
	if err != nil {
		return err
	}
	return template.ExecuteTemplate(writer, name, &data)
}
