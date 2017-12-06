package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"./util"
	"github.com/PuloV/ics-golang"
)

var (
	logPath = flag.String("logpath", "tmp/webcal-service.log", "path to log file")
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	parser := ics.New()
	parserChan := parser.GetInputChan()
	parserChan <- "http://www.facebook.com/ical/u.php?uid=100003108879840&key=AQBV-iXb90SgTeLj"
	outputChan := parser.GetOutputChan()
	for event := range outputChan {
		util.Log.Println(event.GetImportedID())
		fmt.Fprintf(w, "Event: %s \n", event.String())
	}
	parser.Wait()
}

func listenAndLog(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	util.Log.Println(path)
	fi, err := os.Open("2eventsCal.ics")
	util.Check(err)
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	util.Log.Println("copying response")
	io.Copy(w, fi)
	util.Log.Println("end")

}

func main() {
	flag.Parse()
	util.LogInit(*logPath)
	util.Log.Println("Initialized")
	http.HandleFunc("/webcal/validate", validateHandler)
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/listen/cal.ics", listenAndLog)
	http.ListenAndServe(":8080", nil)
}
