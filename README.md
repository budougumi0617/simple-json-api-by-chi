# simple-json-api-by-chi
RESTful JSON API server by Go.
This code is based on http://thenewstack.io/make-a-restful-json-api-go/, but HTTP router is used `go-chi` insteed of `gorilla/mux`.


# Result

`curl` result is below.

```bash
$ curl -D - -X GET http://localhost:8080/todos
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sun, 11 Mar 2018 05:14:15 GMT
Content-Length: 163

[{"id":1,"name":"Write presentation","completed":false,"due":"0001-01-01T00:00:00Z"},{"id":2,"name":"Host meetup","completed":false,"due":"0001-01-01T00:00:00Z"}]

$ curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' -X POST http://localhost:8080/todos
{"id":3,"name":"New Todo","completed":false,"due":"0001-01-01T00:00:00Z"}

$ curl -D - -X GET http://localhost:8080/todos
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sun, 11 Mar 2018 05:14:37 GMT
Content-Length: 237

[{"id":1,"name":"Write presentation","completed":false,"due":"0001-01-01T00:00:00Z"},{"id":2,"name":"Host meetup","completed":false,"due":"0001-01-01T00:00:00Z"},{"id":3,"name":"New Todo","completed":false,"due":"0001-01-01T00:00:00Z"}]
```

# Servr log is below.

```bash
gor *.go
# github.com/budougumi0617/simple-json-api-by-chi

Sample JSON API server by go-chi.

## Routes

<details>
<summary>`/*`</summary>

- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/***
	- _*_
		- [main.Index](/handler.go#L16)

</details>
<details>
<summary>`/todos`</summary>

- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/todos**
	- _POST_
		- [main.TodoCreate](/handler.go#L36)
	- _*_
		- [main.TodoIndex](/handler.go#L21)

</details>
<details>
<summary>`/todos/{todoID}`</summary>

- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/todos/{todoID}**
	- _*_
		- [main.TodoShow](/handler.go#L30)

</details>

Total # of routes: 3

chi-log: logger.go:146: "GET http://localhost:8080/todos HTTP/1.1" from [::1]:57487 - 200 163B in 1.629009ms
chi-log: logger.go:146: "POST http://localhost:8080/todos HTTP/1.1" from [::1]:57504 - 201 74B in 476.071µs
chi-log: logger.go:146: "GET http://localhost:8080/todos HTTP/1.1" from [::1]:57526 - 200 237B in 55.556µs
```

# `docgen` result formated by Markdown

--------------

# github.com/budougumi0617/simple-json-api-by-chi

Sample JSON API server by go-chi.

## Routes

<details>
<summary>`/*`</summary>

- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/***
	- _*_
		- [main.Index](/handler.go#L16)

</details>
<details>
<summary>`/todos`</summary>

- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/todos**
	- _POST_
		- [main.TodoCreate](/handler.go#L36)
	- _*_
		- [main.TodoIndex](/handler.go#L21)

</details>
<details>
<summary>`/todos/{todoID}`</summary>

- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/todos/{todoID}**
	- _*_
		- [main.TodoShow](/handler.go#L30)

</details>

--------------

# `docgen` result formated by JSON

--------------

```json
{
  "router": {
    "middlewares": [
      {
        "pkg": "github.com/budougumi0617/simple-json-api-by-chi/vendor/github.com/go-chi/chi/middleware",
        "func": "Logger",
        "comment": "Logger is a middleware that logs the start and end of each request, along\nwith some useful data about what was requested, what the response status was,\nand how long it took to return. When standard output is a TTY, Logger will\nprint in color, otherwise it will print in black and white. Logger prints a\nrequest ID if one is provided.\n\nAlternatively, look at https://github.com/pressly/lg and the `lg.RequestLogger`\nmiddleware pkg.\n",
        "file": "github.com/budougumi0617/simple-json-api-by-chi/vendor/github.com/go-chi/chi/middleware/logger.go",
        "line": 30
      }
    ],
    "routes": {
      "/*": {
        "handlers": {
          "*": {
            "middlewares": [],
            "method": "*",
            "pkg": "",
            "func": "main.Index",
            "comment": "Index returns simple response.\n",
            "file": "github.com/budougumi0617/simple-json-api-by-chi/handler.go",
            "line": 16
          }
        }
      },
      "/todos": {
        "handlers": {
          "*": {
            "middlewares": [],
            "method": "*",
            "pkg": "",
            "func": "main.TodoIndex",
            "comment": "TodoIndex is not implemented.\n",
            "file": "github.com/budougumi0617/simple-json-api-by-chi/handler.go",
            "line": 21
          },
          "POST": {
            "middlewares": [],
            "method": "POST",
            "pkg": "",
            "func": "main.TodoCreate",
            "comment": "TodoCreate add new Todo to repository\n",
            "file": "github.com/budougumi0617/simple-json-api-by-chi/handler.go",
            "line": 36
          }
        }
      },
      "/todos/{todoID}": {
        "handlers": {
          "*": {
            "middlewares": [],
            "method": "*",
            "pkg": "",
            "func": "main.TodoShow",
            "comment": "TodoShow is not implemented.\n",
            "file": "github.com/budougumi0617/simple-json-api-by-chi/handler.go",
            "line": 30
          }
        }
      }
    }
  }
}
```

--------------

# References
- [Making a RESTful JSON API in Go](http://thenewstack.io/make-a-restful-json-api-go/)
- https://github.com/corylanou/tns-restful-json-api
- [GoでJSON APIを書く](http://sgykfjsm.github.io/blog/2016/03/13/golang-json-api-tutorial/)
- [Go言語で薄く作るAPI(go-chi/chi) #1 最低限構築](http://tikasan.hatenablog.com/entry/2017/11/26/130000)
