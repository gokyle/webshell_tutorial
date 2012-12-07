package main

import (
        "fmt"
        "github.com/gokyle/webshell"
        "github.com/russross/blackfriday"
        "html/template"
        "io/ioutil"
        "net/http"
        "os"
        "path/filepath"
        "regexp"
)

var (
        server_port = "8080"
        page_tpl = webshell.MustCompileTemplate("templates/page.html")
        htmlToMd = regexp.MustCompile("^(.+)\\.html$")
        slash_replace = regexp.MustCompile("/")
)

type Page struct {
        Title string
        Body template.HTML
}

func init() {
        port := os.Getenv("PORT")
        if port != "" {
                server_port = port
        }
}

func servePage(w http.ResponseWriter, r *http.Request) {
        file := r.URL.Path[1:]
        if file == "" {
                file = "index.html"
        }
        md_file := "pages/" + htmlToMd.ReplaceAllString(file, "$1.md")
        out, err := loadMarkdown(md_file)
        if err != nil {
                webshell.Error404("Page not found.", "text/plain", w, r)
                return
        }
        title := htmlToMd.ReplaceAllString(filepath.Base(file), "$1")
        title = slash_replace.ReplaceAllString(title, " ")
        page := Page{title, template.HTML(string(out))}
        body, err := webshell.BuildTemplate(page_tpl, page)
        if err != nil {
                webshell.Error500(err.Error(), "text/plain", w, r)
                return
        }
        w.Write(body)
}

func loadMarkdown(file string) (out []byte, err error) {
        var in []byte
        in, err = ioutil.ReadFile(file)
        if err != nil {
                return
        }
        out = blackfriday.MarkdownCommon(in)
        return
}

func main() {
        app := webshell.NewApp("webshell tutorial", "", server_port)
        app.AddRoute("/", servePage)
        app.StaticRoute("/assets/", "assets/")
        app.StaticRoute("/examples/", "examples/")
        fmt.Println("[!] ", app.Serve())
}
