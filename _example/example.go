package main

import (
	. "WebGui"
	"net/http"
	"time"
)

func timeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	OnWhile(w, 1)
	t := time.Now().Format(time.RFC3339)
	if er := AddLine(w, t, 1); er != nil {
		P("%v", er)
	}
	F(w, "<a href=\"/js\">js Page</a>\n")
	F(w, "<a href=\"/jsfile\">jsfile Page</a>\n")
	F(w, "<a href=\"/img\">img Page</a>\n")

}

func jsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	RunScriptJS(w, `
	function myAlert() {
    	alert("12345...");
      	return;
  	};
	myAlert();`)
	F(w, "<a href=\"/time\">Time Page</a>\n")
	F(w, "<a href=\"/jsfile\">jsfile Page</a>\n")
	F(w, "<a href=\"/img\">img Page</a>\n")

}
func jsfileHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	RunFileJS(w, "script.js")
	F(w, "<a href=\"/time\">Time Page</a>\n")
	F(w, "<a href=\"/js\">js Page</a>\n")
	F(w, "<a href=\"/img\">img Page</a>\n")
}

func main() {
	PORT := "8001"
	wb := NewRouter(PORT)
	wb.AddHandler("/time", timeHandler)
	wb.AddHandler("/js", jsHandler)
	wb.AddHandler("/jsfile", jsfileHandler)
	wb.AddHandlerHtmlPage("/img", "index_img.html")
	P("%v\n", wb.GetListHandlers())
	wb.Listen(true)
}
