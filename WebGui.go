package WebGui

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

/*
decorators funcs
*/
var S = fmt.Sprintf
var F = fmt.Fprintf
var P = fmt.Printf

type Router struct {
	PORT         string // ":8001"
	Mux          *http.ServeMux
	CntrHandlers int // count additions handlers
}

// use port default: "8001"
func NewRouter(port string) *Router {
	router := Router{}
	router.PORT = ":" + port
	router.CntrHandlers = 0
	router.Mux = http.NewServeMux()
	return &router
}

// get list handler addresses -> ["/home", "/about" ...]
func (router *Router) GetListHandlers() []string {
	list := make([]string, 0)
	res := S("%#v", router.Mux)
	r := strings.Split(res, "m:map[string]")
	c := 0
	c2 := 0
	for _, v := range r {
		for _, val := range strings.Split(v, ":(http.HandlerFunc)") {
			c++
			if c < 3 {
				continue
			}
			for _, value := range strings.Split(val, "pattern") {
				c2++
				if c2%2 != 0 {
					continue
				}
				path := strings.Split(value, "}")[0]
				path = path[2 : len(path)-1]
				list = append(list, path)
			}
		}
	}
	return list
}

// ("/home", "index_home.html")
func (router *Router) AddHandlerHtmlPage(path, filename string) {
	router.CntrHandlers++
	f := func(w http.ResponseWriter, r *http.Request) {
		var temp = template.Must(template.ParseFiles(filename))
		temp.Execute(w, nil)
	}
	router.Mux.HandleFunc(path, f)
}

// ("/home", homeHandlerFunc)
func (router *Router) AddHandler(path string, function func(w http.ResponseWriter, r *http.Request)) {
	router.CntrHandlers++
	router.Mux.HandleFunc(path, function)
}

// running server; information in console output - true/false
func (router *Router) Listen(info bool) {
	if info {
		fmt.Printf("Server running port%v time:%v\n", router.PORT, time.Now().Format(time.RFC1123))
	}
	http.ListenAndServe(router.PORT, router.Mux)
}

// enable loop refresh page
func OnWhile(w http.ResponseWriter, sec int) {
	F(w, "<head><meta http-equiv=\"refresh\" content=%v></head>", sec)
}

// add header in page; (w, "Hello", 3) -> <h3>Hello</h3>
func AddLine(w http.ResponseWriter, line string, size int) error {
	if 0 < size && size < 7 {
		F(w, "<div><h%v>%v</h%v></div>", size, line, size)
		return nil
	}
	lineerror := S("<h%v>%v</h%v>", size, line, size)
	return errors.New(S("error line compile: %v\n", lineerror))
}

// run javascript string
func RunScriptJS(w http.ResponseWriter, codejs string) {
	F(w, "<script type=\"text/javascript\">%v</script>", codejs)
}

// run javavscript file
func RunFileJS(w http.ResponseWriter, filename string) error {
	data, er := os.ReadFile(filename)
	if er != nil {
		return errors.New(S("error get file: %v\n", filename))
	}
	codejs := string(data)
	F(w, "<script type=\"text/javascript\">%v</script>", codejs)
	return nil
}
