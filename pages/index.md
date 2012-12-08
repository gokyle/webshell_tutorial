# Let's Build a Webapp!

When I started writing some webapps in [Go](http://golang.org), I found
myself writing the same functions and patterns over and over. To alleviate
this, I've been coalescing these commonly-used pieces into a framework
called [webshell](http://gokyle.github.com/webshell/), so-called because
it started as a shell for new webapps. I think it provides a fairly useful
set of tools; hopefully you'll find it useful.

This site was built using the framework; the [source](http://github.com/gokyle/webshell_tutorial)
is available on [Github](https://www.github.com). I have written this
anticipating that you know some Go already. If you don't, I'd recommend taking
a look at the [Go language specification](http://golang.org/ref/spec) and
[Effective Go](http://golang.org/doc/effective_go/html). They are both pretty
short reads and should get you started.

The package documentation can be found in the
[webshell godocs](http://gopkgdoc.appspot.com/github.com/gokyle/webshell). I've
made it a priority to make sure the framework is well-documented.

***Note***: *Several sections link to example source files. The linked examples
can be downloaded and run; to keep them from being compiled into this app,
I've named them as text files. Just remove the `.txt` ending in the file name.*

Let's get started!

## Table of Contents

### The Webshell Core
* [Introduction: a minimal app](/intro)
* [Adding routes](/routes)
* [Writing basic request handlers](/basic_handlers) **(WIP)**
* Templating pages
* Handling forms
* [Deploying your app](/deployment)

### Webshell Subpackages
* Add an asset cache to speed up serving static files with `webshell/assetcache`
* Add authentication and sessions with `webshell/auth`

## Resources
* [`net/http`](http://golang.org/pkg/net/http)
* [Go language spec](http://golang.org/ref/spec)

