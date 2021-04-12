package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

 var Template = `
<!DOCTYPE html>
<title> Choose your own adventrue </title>
<h1>{{.Title}}</h1>
{{range .Story}}
<p>
{{.}}
</p>
{{end}}
{{range $Option := .Options}}
<li>
  <a href="/{{.Arc}}">{{.Text}}</a> 
</li>
{{end}}
`

func ParsingJSON(fileByte []byte) (s Story, err error) {
	s = Story{}
	if err := json.Unmarshal(fileByte, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func CoyaHandler(s Story) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request) {
		rPath := r.URL.Path
		log.Println("Request path:", rPath)
		path := strings.Split(rPath, "/")
		if content, ok := s[path[1]]; ok {
			t, err := template.New("default").Parse(Template)
			if err !=nil {
				log.Println("template error:", err)
			}
			t.Execute(w, content)
		}
	}
}

// 模板渲染
func main() {
	f, err := os.ReadFile("gopher.json")
	if err != nil {
		panic(err) }
	s ,_ := ParsingJSON(f)
	if err != nil {
		panic(err)
	}
	muxHandler := CoyaHandler(s)
	log.Println("Start server and listen 8080")
	http.ListenAndServe(":8080", muxHandler)
}
