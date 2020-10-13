package handler

import (
	"log"
	"net/http"
	"strings"
)

type DomainHandler struct {
	domainMap map[string] http.Handler
}

func (d DomainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	log.Print("Remote address: " + r.RemoteAddr)
	log.Println("Host: " + host)
	if strings.Contains(host, "localhost") {
		d.domainMap["localhost"].ServeHTTP(w, r)
	} else {
		d.domainMap[host].ServeHTTP(w,r)
	}

}

