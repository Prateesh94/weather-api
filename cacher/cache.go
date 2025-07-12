package cacher

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"
)

var c *cache.Cache
var clients = make(map[string]*Client)
var mu sync.Mutex

type Client struct {
	limiter *rate.Limiter
}

func init() {
	c = cache.New(time.Minute*15, time.Minute*30)
}
func getclient(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	if client, exists := clients[ip]; exists {
		fmt.Println("Client Found")
		return client.limiter
	}
	limit := rate.NewLimiter(rate.Limit(50), 1)
	clients[ip] = &Client{limit}
	fmt.Println("Added client")
	return limit
}
func Addcache(a map[string]interface{}, s string) {
	c.Set(s, a, 15*time.Minute)
	fmt.Println("Added Cache")
}
func Limitmid(a http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		lim := getclient(ip)
		if !lim.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			fmt.Println("error generated")
			return
		}
		a.ServeHTTP(w, r)
	})
}
func Readcache(s string) (any, bool) {
	var f bool
	var d any
	if d, f = c.Get(s); f {
		//fmt.Println(d)
		fmt.Println("Cache found")
		return d, true
	} else {
		return nil, false
	}

}
