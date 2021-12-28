package main

import (
	"bytes"
	"net/http"
	"text/template"
)

type Page struct {
	Title           string
	Author          string
	Header          string
	PageDescription string
	Content         string
	URI             string
}

/*func loadfile(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}*/

func handler(w http.ResponseWriter, r *http.Request) {
	var builder bytes.Buffer
	builder.WriteString("sayfa sayısı bilmem kaç\n")
	builder.WriteString("kapak kalınlığı\n")
	builder.WriteString("rastgele bilgi\n")
	builder.WriteString("specs\n")
	builder.WriteString("specs\n")
	builder.WriteString("bilgi yaz\n")

	uri := "www.google.com"

	page := Page{
		Title:           "title",
		Author:          "author",
		Header:          "header",
		PageDescription: "pagedescrition",
		Content:         builder.String(),
		URI:             "http://" + uri}

	t, _ := template.ParseFiles("page.html")
	t.Execute(w, page)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
