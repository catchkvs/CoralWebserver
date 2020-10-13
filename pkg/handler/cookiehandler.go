package handler

import (
	"github.com/speps/go-hashids"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type CookieHandler struct {
	domainMap map[string] http.Handler
}

func (d CookieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	addCookie(w, "session-id", newHashId(), 30 * time.Minute)
}


// addCookie will apply a new cookie to the response of a http request
// with the key/value specified.
func addCookie(w http.ResponseWriter, name, value string, ttl time.Duration) {
	expire := time.Now().Add(ttl)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

func newHashId() string {
	var hd = hashids.NewData()
	hd.Salt = "Parrot Salt"
	h, err := hashids.NewWithData(hd)
	handleError(err)
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	randomness := rand.Int()
	rand.Seed(time.Now().UnixNano())
	a := []int {year, month, day, hour, minute, second, randomness}
	for i := len(a) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	id, _ := h.Encode(a)
	return id
}

func handleError(err error) {
	if err != nil {
		log.Println("handling error::::", err)

	}
}