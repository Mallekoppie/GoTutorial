/*
 * LoadTest.V1
 *
 * No description
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "Tutorial/LoadTest/swagger"
	//	"context"
	//"time"
	//"github.com/pkg/profile"
)

func main() {
	/*
		// Tutorial on profiling
		// command go tool pprof --pdf ./LoadTest.exe C:\Users\Heinrich\AppData\Local\Temp\profile196261455\cpu.pprof > result.pdf
		defer profile.Start().Stop()
		log.Printf("Server started")

		router := sw.NewRouter()

		srv := &http.Server{Addr: ":11000"}

		go func() {
			//log.Fatal(http.ListenAndServeTLS(":11000", "./loadtest.cer", "./loadtest.pkcs8", router))
			//log.Fatal(http.ListenAndServe(":11000", router))
			srv.Handler = router
			srv.ListenAndServeTLS("./loadtest.cer", "./loadtest.pkcs8")
		}()

		time.Sleep(time.Minute * 1)
		srv.Shutdown(context.TODO())
	*/

	// no perf logging

	router := sw.NewRouter()
	log.Printf("Server started")

	log.Fatal(http.ListenAndServeTLS(":11000", "./loadtest.cer", "./loadtest.pkcs8", router))
	//log.Fatal(http.ListenAndServe(":11000", router))

}
