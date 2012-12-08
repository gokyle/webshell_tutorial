## Handling Forms

There are two types of forms to deal with: urlencoded and multipart.

### urlencoded forms
For forms using the default `application/x-www-form-urlencoded`
mime type, the `ParseForm` method on the [`Request`
object](http://golang.org/pkg/net/http/#Request) should be used.

The `http.Request` value (we'll use the standard name of `r`) has a `Form`
field that is initialised when `ParseForm` is called. This field is a
[`net/url.Values`](http://golang.org/pkg/net/url/#Values) type. We can
retrieve the values, which will all be strings, using the `Get`
accessor:

```
        r.ParseForm()
        username := r.Form.Get("user")
```

Alternatively, the `FormValue` method can be used, which calls
`ParseForm` automatically:

```
        username := r.FormValue("user)

### multipart forms 
Still figuring this out.


## Examples
I've written a an example for us to look through; because they
contain several files, I've supplied them as compressed archives.

* urlencoded: [tarball](/examples/forms/urlencoded.tgz),
[zip](/examples/forms/urlencoded.zip)

### urlencoded

[`server.go`](/examples/forms/urlencoded/server.go):

```
package main

import (
	"github.com/gokyle/webshell"
	"net/http"
)

type Page struct {
	Processed bool
	User      *FormData
}

type FormData struct {
	Name  string
	Email string
}

var home_template = webshell.MustCompileTemplate("templates/index.html")

func main() {
	app := webshell.NewApp("minimal app", "127.0.0.1", "8080")
	app.AddRoute("/", home)
	app.Serve()
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		getForm(w, r)
	} else {
		showForm(w, r)
	}
}

func showPage(page Page, w http.ResponseWriter, r *http.Request) {
	out, err := webshell.BuildTemplate(home_template, page)
	if err != nil {
		webshell.Error500(err.Error(), "text/plain", w, r)
	} else {
		w.Write(out)
	}
}

func showForm(w http.ResponseWriter, r *http.Request) {
	page := Page{false, nil}
	showPage(page, w, r)
}

func getForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := FormData{r.Form.Get("name"), r.Form.Get("email")}
	page := Page{true, &user}
	showPage(page, w, r)
}
```

This expects
[`templates/index.html`](/examples/forms/urlencoded/templates/index.html):

```
<!doctype html>
<html>
  <head>
    <meta charset="UTF-8" />
    <title>urlencoded form example</title>
  </head>

  <body>
    <h1>It's A Form!</h1>
{{if .Processed}}
    {{with .User}}
    <p>Hello, {{.Name}}!</p>
    {{end}}
{{else}}
    <form action="/" method="post" name="userdata">
      Name: <input type="text" name="name"><br /> 
      Email: <input type="text" name="email"><br />
      <input type="submit">
    </form>
{{end}}
  </body>
</html>
```

