package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	path, _ := filepath.Abs("./gowiki/data/" + filename)
	return os.WriteFile(path, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	path, _ := filepath.Abs("./gowiki/data/" + filename)
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

var validPath = regexp.MustCompile("^/(view|edit|save)/([a-zA-z\\d]+)$")

// getTitle
func _(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid page title")
	}
	return m[2], nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title, err := getTitle(w, r)
	//if err != nil {
	//	return
	//}
	p, err := loadPage(title)
	if err != nil {
		//p = &Page{Title: title, Body: []byte("No content")}
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, _ *http.Request, title string) {
	//title, err := getTitle(w, r)
	//if err != nil {
	//	return
	//}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{
			Title: title,
		}
	}
	//fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	//	"<form action=\"/save/%s\" method=\"POST\">"+
	//	"<textarea name=\"body\">%s</textarea><br>"+
	//	"<input type=\"submit\" value=\"Save\">"+
	//	"</form>",
	//	p.Title, p.Title, p.Body)

	//path, _ := filepath.Rel("", "gowiki/edit.html")
	//t, _ := template.ParseFiles(path)
	//t.Execute(w, p)

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title, err := getTitle(w, r)
	//if err != nil {
	//	return
	//}
	body := r.FormValue("body")
	p := &Page{
		Title: title,
		Body:  []byte(body),
	}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Here we will extract the page title from the Request,
		// and call the provided handler 'fn'
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

var templates = template.Must(template.ParseFiles("./gowiki/tmpl/edit.html", "./gowiki/tmpl/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	//path, _ := filepath.Abs("gowiki/" + tmpl)
	//t, err := template.ParseFiles(path + ".html")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	//err = t.Execute(w, p)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//files, _ := template.ParseFiles("./gowiki/page/index.html")
	//files.Execute(w, nil)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	file, _ := os.ReadFile("./gowiki/page/index.html")
	_, err := w.Write(file)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
	}
}

func main() {
	//p1 := Page{"TestPage", []byte("This is a sample Page.")}
	//p1.save()
	//p2, _ := loadPage("TestPage")
	//fmt.Println(string(p2.Body))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
