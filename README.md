# profiling



Profiling web applications in Golang
this project includes the implementation of the web profiling process infrastructure in AVRNG
What is a profiler?
Profiler is a dynamic performance analysis tool that provides critical execution insights in various dimensions which enable resolving
performance issues, locating memory leaks, goroutines contention and more.
**kinds of profiles : **


Goroutine: stack traces of all current Goroutines

CPU: stack traces of CPU returned by runtime

Heap: a sampling of memory allocations of live objects

Allocation: a sampling of all past memory allocations

Thread: stack traces that led to the creation of new OS threads

Block: stack traces that led to blocking on synchronization primitives

Mutex: stack traces of holders of contended mutexes

**Profile implementation : **
there are three way to achieve and create profiling which are described below :


use "go test" to generate the profile -
Support for profiling built into the standard testing package.
As an example, the following command runs all benchmarks and writes CPU and memory profile to cpu.prof and mem.prof:

go test -cpuprofile cpu.prof -memprofile mem.prof -bench .



use "code profiling" / "profiling in code"-
to use the profile directly within the code ,
For example you can start a CPU profile using pprof.StartCPUProfile(io.Writer) and then stop it by pprof.StopCPUProfile().
Now you have the CPU profile written to the given io.Writer.


use Http server (profiling web applciation) -
this option is used to get a live profile for long running service , in this project we are going to cover and show how it should be implemented


How to use profile ?
pprof is a tool for visualization and analysis of profiling data.
It reads a collection of profiling samples in profile.proto format and generates reports to visualize and help analyze the data.
It can generate both text and graphical reports.
All of the ways that mentioned above are using runtime/pprof under the hood and this package writes runtime profiling data in the format expected by the pprof visualization tool.
**profile Web applciation **
in this section we are describing the steps that needed to be followed to bring the profile web application up and ready for use
0. create your go file (for example random.go)


import the net/http/pprof & net/http packages


register the reqiured routes


create the http handler which hanlding the http request and integrated to go , from handler function we are calling the code that we want to profile
import "net/http/pprof"
// HTTP request handler func handler(w http.ResponseWriter, r *http.Request) { // return the random string fmt.Fprintf(w, "%s", randomString( )) }


in your main() go function
4.1. create a new gorilla HTTP server
router := mux.NewRouter()
4.2. register our handler for the / route
router.HandleFunc("randomstring", handler)
4.3. add the pprof routes (the profile urls)
router.HandleFunc("/debug/pprof/", pprof.Index) router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline) router.HandleFunc("/debug/pprof/profile", pprof.Profile) router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)`` router.HandleFunc("/debug/pprof/trace", pprof.Trace) router.Handle("/debug/pprof/block", pprof.Handler("block")) router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine")) router.Handle("/debug/pprof/heap", pprof.Handler("heap")) router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
4.4. start listening on port 8080
	// Start listening on port 8080 if err := http.ListenAndServe(":8080",router); err != nil { log.Fatal(fmt.Sprintf("Error when starting or running http server: %v", err)) } 


raise up the http server




open terminal #1 , and hit the following to build the application :
go build -o myserver random.go .


run the binary
./myserver


open terminal #3 and hit the profile urls (cpu profile):
go tool pprof -seconds 30 myserver http://localhost:8080/debug/pprof/profile


simulate http client server to hit an http request to your machine , open terminal #4 and hit the following lot of times :
ab -k -c 8 -n 100000 "http://localhost:8080/randomstring"`


for web visualization and graph presentation hit the following :
_ (pprof) web_
