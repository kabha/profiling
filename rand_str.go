package main

import (
    "fmt"
    "math/rand"
    "time"
    "log"
    "github.com/gorilla/mux"
    "net/http"
    "net/http/pprof"

)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init(){
rand.Seed(time.Now().Unix())
}


func getRandomString(n  int) string {
b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}


//http request handler 
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf( w, "%s",getRandomString(10))
}





func main(){
fmt.Println(getRandomString(10))
// Create a new HTTP multiplexer
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// Register our handler for the / route
	router.HandleFunc("/randomstring",handler)

	// Add the pprof routes
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	router.Handle("/debug/pprof/block", pprof.Handler("block"))
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	// Start listening on port 8080
	if err := http.ListenAndServe(":8080",router); err != nil {
		log.Fatal(fmt.Sprintf("Error when starting or running http server: %v", err))
	}

}







