## Routes

From the [`net/http` documentation](http://golang.org/pkg/net/http/#ServeMux):

> ServeMux is an HTTP request multiplexer. It matches the URL of each incoming
> request against a list of registered patterns and calls the handler for the
> pattern that most closely matches the URL.
> 
> Patterns name fixed, rooted paths, like "/favicon.ico", or rooted subtrees,
> like "/images/" (note the trailing slash). Longer patterns take precedence
> over shorter ones, so that if there are handlers registered for both
> "/images/" and "/images/thumbnails/", the latter handler will be called for
> paths beginning "/images/thumbnails/" and the former will receive requests
> for any other paths in the "/images/" subtree.
> 
> Patterns may optionally begin with a host name, restricting matches to URLs
> on that host only. Host-specific patterns take precedence over general
> patterns, so that a handler might register for the two patterns "/codesearch"
> and "codesearch.google.com/" without also taking over requests for
> "http://www.google.com/".
> 
> ServeMux also takes care of sanitizing the URL request path, redirecting any
> request containing . or .. elements to an equivalent .- and ..-free URL.

This means the route we setup will handle any incoming paths. To prove this,
let's create a new [`routes.go`](/examples/intro/routes.gp) app:

```
package main

import (
        "fmt"
        "github.com/gokyle/webshell"
        "net/http"
)

var app *webshell.WebApp

func main() {
        app = webshell.NewApp("minimal app", "127.0.0.1", "8080")
        app.AddRoute("/", itWorks)
        app.Serve()
}

func itWorks(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("It worked!"))

        route := fmt.Sprintf("\nhandling route for %s%s", app.Address(),
                r.URL.Path)
        w.Write([]byte(route))
}
```

There are two major things to note about this revision:
0. The `app` variable is now global to the app; any function can now access
it to get the app's name or address.
0. We get the path from the request by accessing the request's `URL` field,
which is an instance of [`net/url`](http://golang.org/pkg/net/url/#URL).
