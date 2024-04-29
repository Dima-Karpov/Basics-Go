package codeTiming

import (
	"bytes"
	"testing"
	"text/template"
)

func BenchmarkTemplates(b *testing.B) {
	b.Logf("b.N is %d\n", b.N)
	tp1 := "Hello {{.Name}}"
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t, _ := template.New("test").Parse(tp1)
		t.Execute(&buf, data)
		buf.Reset()
	}
}

func BenchmarkCompiledTemplates(b *testing.B) {
	b.Logf("b.N is %d\n", b.N)
	tp1 := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tp1)
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t.Execute(&buf, data)
		buf.Reset()
	}
}
