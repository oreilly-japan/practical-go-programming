package tips

import (
	"net/http"
)

// hello-start
type HelloHandler struct{}

func (HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

// hello-end
