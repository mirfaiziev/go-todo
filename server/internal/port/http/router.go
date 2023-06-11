package http

import (
	"context"
	"database/sql"
	"net/http"
	"regexp"
	"strings"
	"todo-service/internal/port/http/handler"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

func router(db *sql.DB) func(http.ResponseWriter, *http.Request) {

	hwh := handler.NewHttpHandler(db)
	th := handler.NewTodoHandler(db)

	var routes = []route{
		newRoute(http.MethodGet, "/hello", hwh.ShowHelloWorld),
		newRoute(http.MethodGet, "/todo", th.ShowList),
		newRoute(http.MethodPost, "/todo", th.AddTodo),

	}

	return func(w http.ResponseWriter, r *http.Request) {
		var allow []string
		for _, route := range routes {

			matches := route.regex.FindStringSubmatch(r.URL.Path)
			if len(matches) > 0 {
				if r.Method != route.method {
					allow = append(allow, route.method)
					continue
				}
				ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
				route.handler(w, r.WithContext(ctx))
				return
			}
		}
		if len(allow) > 0 {
			w.Header().Set("Allow", strings.Join(allow, ", "))
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.NotFound(w, r)
	}
}

type ctxKey struct{}
