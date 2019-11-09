package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	maxTar     int    = 1000000
	message    string = "<p>Thou darest request php from my site? Thine ass is banished to the tarpit!"
	messageEnd string = "<p>I release thee from the tarpit. Do not do it again."
)

func phptarpit(w http.ResponseWriter, r *http.Request) {
	ipSet, _ := r.Header["X-Forwarded-For"]
	var ip string
	if len(ipSet) >= 1 {
		ip = ipSet[0]
	} else {
		ip = "Unknown"
	}

	log.Printf("%s - Caught in tarpit\n", ip)

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<!DOCTYPE html><html><head><title>No PHP exploits allowed!</title><body>")

	for i := 0; i < maxTar; i++ {
		if _, err := fmt.Fprintln(w, message); err != nil {
			if strings.Contains(err.Error(), "broken pipe") {
				log.Printf("%s - Tarred for %d seconds\n", ip, i)
				return
			}

			log.Println(err)
			return
		}

		w.(http.Flusher).Flush()
		time.Sleep(time.Second * 1)
	}

	log.Printf("%s - Tarred for %d (max) seconds\n", ip, maxTar)

	fmt.Fprintln(w, messageEnd)
}

func main() {
	http.HandleFunc("/", phptarpit)
	http.ListenAndServe(":8081", nil)
}
