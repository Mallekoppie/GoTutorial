package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	addr := "0.0.0.0:11000"

	e := fasthttp.ListenAndServeTLS(addr, "./loadtest.cer", "./loadtest.pkcs8", func(c *fasthttp.RequestCtx) {
		_, err := c.WriteString("Hello World")
		if err != nil {
			log.Println("Error writing response: ", err.Error())
		}
	})

	if e != nil {
		log.Println("HTTP server error: ", e.Error())
	}
}
