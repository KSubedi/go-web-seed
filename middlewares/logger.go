package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mgutz/ansi"
)

//HTTPLogger logs the http requests to stdout
type HTTPLogger struct {
}

func (h *HTTPLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//Colors to be used on the logger
	red := ansi.ColorCode("red+bh")
	white := ansi.ColorCode("white+bh")
	greenUnderline := ansi.ColorCode("green+buh")
	blackOnWhite := ansi.ColorCode("black+b:white+h")
	//Reset the color
	reset := ansi.ColorCode("reset")

	//Log data
	fmt.Println(red,
		time.Now(),
		reset,
		":",
		white,
		r.RemoteAddr,
		reset, "requested page",
		greenUnderline,
		r.URL.String(),
		reset,
		"with method",
		blackOnWhite,
		r.Method,
		reset)
}
