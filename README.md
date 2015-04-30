Simple BasicAuth server written in GoLang using routes and standard library.

Test using:

```
$ curl -u foo:bar http://localhost/api/json

$ curl -X POST -d Hello=Foo http://localhost:1380/api/submit
Authorisation Required

$ curl -u foo:bar -X POST -d Hello=Foo http://localhost:1380/api/submit
post request received

```
