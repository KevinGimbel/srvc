package srvc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Server represents a webserver
type Server struct {
	Port string
}

var registered = make(map[string]bool)

// New creates a webserver
func New(p string) Server {
	return Server{Port: p}
}

// Run starts a webserver
func (s *Server) Run() error {
	return http.ListenAndServe(s.Port, nil)
}

// AddHandler adds a new handler
func (s *Server) AddHandler(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	registered[pattern] = true

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		setConfiguredHeaders(w, pattern)
		handler(w, r)
	})
}

func setConfiguredHeaders(w http.ResponseWriter, p string) {
	c := GetConfig()
	headers := c.Headers

	for _, v := range headers {
		w.Header().Set(v.Key, v.Value)
	}

	for _, v := range c.Routes[p].Headers {
		w.Header().Set(v.Key, v.Value)
	}
}

// CreateConfiguredHandlers takes the configuration file and assigns handlers
func (s *Server) CreateConfiguredHandlers() {
	c := GetConfig()
	for k, v := range c.Routes {
		if !registered[k] {
			// This anonymous function is used to scope the k and v variables
			// See https://stackoverflow.com/questions/44044245/golang-register-multiple-routes-using-range-for-loop-slices-map
			func(p string, v RouteConfig) {
				s.AddHandler(p, func(w http.ResponseWriter, r *http.Request) {
					setConfiguredHeaders(w, p)
					if len(v.Content) <= 0 {
						v.Content = ""
					}

					if v.File != "" {
						fb, err := ioutil.ReadFile(v.File)
						if err != nil {
							log.Fatal(err)
						}
						v.Content = string(fb)
					}

					w.Write([]byte(v.Content))
				})
				fmt.Println(fmt.Sprintf("Registered handler for http://localhost%s%s", s.Port, p))
			}(k, v)
		} else {
			fmt.Printf("[INFO] Route %s already registered. Skipping.\n", k)
		}
	}
}
