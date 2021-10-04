package api

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	defaultAddress     = "127.0.0.1:8080"
	defaultStoragePath = "./assets"
)

type Api struct {
	addr       string
	router     *mux.Router
	middleware []Middleware

	// storagePath refers to the host-path
	// under which files (images/gifs) will be stored
	// if not set default will be used
	storagePath string
}

// WithHostAddr applies a IP:Port pair to the API on which
// it will start to listen
func WithHostAddr(addr string) func(*Api) {
	return func(a *Api) {
		a.addr = addr
	}
}

// WithStorage applies a custom storage path to the API
// under which uploaded image and gifs will be stored
func WithStorage(path string) func(*Api) {
	return func(a *Api) {
		a.storagePath = path
	}
}

// Middleware
func WithMiddleware(m ...mux.MiddlewareFunc) func(*Api) {
	return func(a *Api) {
		a.router.Use(m...)
	}
}

// New returns a new Api instance. If no address is provided
// the API will listen on its default address "127.0.0.1:8080"
func New(opts ...func(*Api)) *Api {

	apiSrv := &Api{
		addr:        defaultAddress,
		router:      mux.NewRouter(),
		storagePath: defaultStoragePath,
	}

	for _, opt := range opts {
		opt(apiSrv)
	}
	apiSrv.setup()

	return apiSrv
}

// setup initializes the api routes
func (a *Api) setup() {

	// /upload allows to upload either images (16x16)
	// or gifs (16x16)
	a.router.HandleFunc("/upload", nil)

	// /run allows to render and run a specific image/gif
	// on the LED matrix
	a.router.HandleFunc("/run/{asset}", nil)

	// /draw allows to request a self drawn 16x16 pixel art
	a.router.HandleFunc("/draw", nil)

}

// ListenAndServer start the Api-Server on the given address
func (a *Api) ListenAndServe() error {
	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		return err
	}

	return http.Serve(listener, a.router)
}

type Middleware = http.Handler

// withLogging adds logging to the API route
func withLogging(next http.Handler) Middleware {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("[%s] incoming request from: %v\n", r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
