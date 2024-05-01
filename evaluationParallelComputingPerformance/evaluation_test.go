package evaluationParallelComputingPerformance

import (
	"bytes"
	"testing"
	"text/template"
)

func BenchmarkParallelTemplates(b *testing.B) {
	tp1 := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tp1)
	data := &map[string]string{
		"Name": "World",
	}

	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}
