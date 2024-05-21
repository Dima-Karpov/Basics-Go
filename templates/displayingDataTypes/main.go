package main

import (
	"bytes"
	"html/template"
	"net/http"
)

var t *template.Template
var qc template.HTML

func init() {
	t = template.Must(template.ParseFiles("index.html", "quote.html"))
}

type Pega struct {
	Title   string
	Content template.HTML
}
type Quote struct {
	Quote, Persona string
}

func main() {
	q := &Quote{
		Quote: `You keep using that word. I do not think
				in means what you think it means.`,
		Persona: "Inigo Montoya",
	}
	var b bytes.Buffer
	t.ExecuteTemplate(&b, "quote.html", q)
	qc = template.HTML(b.String())
	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8888", nil)
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Pega{
		Title:   "A User",
		Content: qc,
	}
	t.ExecuteTemplate(w, "index.html", p)
}
