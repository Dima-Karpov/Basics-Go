package templates

import (
	"html/template"
	"net/http"
	"time"
)

var tp1 = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Date Example</title>
	</head>
	<body>
		<p>{{.Date | dateFormat "Jan 2, 2024"}}</p>
	</body>
</html>
`

var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

func serverTemplate(res http.ResponseWriter, req *http.Request) {
	t := template.New("date")
	t.Funcs(funcMap)
	t.Parse(tp1)
	data := struct{ Date time.Time }{
		Date: time.Now(),
	}

	t.Execute(res, data)
}

func main() {
	http.HandleFunc("/", serverTemplate)
	http.ListenAndServe(":8080", nil)
}
