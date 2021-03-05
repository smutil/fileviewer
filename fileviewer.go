package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"flag"
	"log"
	"strings"
)

var workingDir string
const VERSION = "v1.0"

var (
	endpointsAccessed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "endpoints_accessed",
			Help: "Total number of accessed to a given endpoint",
		},
		[]string{"accessed_endpoint"},
	)
)

func main() {
	version := flag.Bool("version", false, "returns the fileviewer version")
	port := flag.String("port", "3000", "overwrite default port")
	dest := flag.String("dest", "", "(required) destination/working directory, should not be root /")
	tlsCert := flag.String("tls-crt", "", "certificate path, only needed for ssl service")
	tlsKey := flag.String("tls-key", "", "key path, only needed for ssl service")
	logfile := flag.String("log-file", "", "key path, only needed for ssl service")
	prometheus.MustRegister(endpointsAccessed)
	flag.Parse()
	if *version {
		fmt.Println(VERSION)
		return
	}
	if *dest == "" || *dest == "/" {
		log.Fatal("-dest is required for default destination/working directory and should not be root /, please refer -h")
	} 

	if *logfile != "" {
		file, err := os.OpenFile(*logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)
	}
	

	workingDir = formatDirName(*dest)
	log.Println("working directory is  "+workingDir)
	fs := http.FileServer(http.Dir(workingDir))
	http.Handle("/", fs)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", health)
	
	if *tlsCert != "" && *tlsKey != "" {
		log.Println("application starting on port  "+*port + " (https)")
		err := http.ListenAndServeTLS(":"+*port, *tlsCert, *tlsKey, nil)
		if err  != nil {
			log.Println(err)
            return
		}
		log.Println("application started on port  "+*port + " (https)")
	} else {
		log.Println("application starting on port  "+*port + " (http)")
		err := http.ListenAndServe(":"+*port, nil)
		if err  != nil {
			log.Println(err)
        	return
		}
		log.Println("application started on port  "+*port + " (http)")
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	endpointsAccessed.WithLabelValues("/health").Inc()
	fmt.Fprintln(w, "application is healthy")
}

func formatDirName(s string) string {
	if strings.HasSuffix(s, "/") { 
		s = strings.TrimSuffix(s, "/")
	}
	if !(strings.HasPrefix(s, "/")){
		s = "/" + s
	}
	return s
}

func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}
