package api

import (
	"net"
	"net/http"

	matrixsvc "github.com/KonstantinGasser/davinci/core/domain/matrix/svc"
	"github.com/KonstantinGasser/davinci/core/pkg/asset"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	defaultAddress     = "127.0.0.1:8001"
	defaultStoragePath = "assets"

	// maxUploadSize is set to 1024 byte
	// matching exactly 16x16 pixel á 4 bytes (rgba)
	maxUploadSize = 1 << 10
)

type Api struct {
	addr   string
	router *mux.Router

	// storagePath refers to the host-path
	// under which files (images/gifs) will be stored
	// if not set default will be used
	assets asset.Store

	// dependencies
	matrixSvc matrixsvc.Service
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
func WithStorage(store asset.Store) func(*Api) {
	return func(a *Api) {
		a.assets = store
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
func New(matrixSvc matrixsvc.Service, opts ...func(*Api)) *Api {

	apiSrv := &Api{
		addr:   defaultAddress,
		router: mux.NewRouter(),
		assets: asset.NewStore(defaultStoragePath),

		matrixSvc: matrixSvc,
	}

	for _, opt := range opts {
		opt(apiSrv)
	}
	apiSrv.setup()

	return apiSrv
}

// setup initializes the api routes
func (a *Api) setup() {

	// /api/upload allows to upload either images (16x16)
	// or gifs (16x16)
	a.router.HandleFunc("/api/upload/{format}", a.HandleUpload).Methods("POST")

	// /api/assets/list returns a list of all available assets
	a.router.HandleFunc("/api/assets/list", a.HandleAssetList).Methods("GET")
	// /api/asset returns the requested asset
	a.router.HandleFunc("/api/assets/{asset}", a.HandleAsset).Methods("GET")

	// /api/run allows to render and run a specific image/gif
	// on the LED matrix
	a.router.HandleFunc("/api/run/{asset}", a.HandleUpdates).Methods("GET")

	// /api/draw allows to request a self drawn 16x16 pixel art
	a.router.HandleFunc("/api/draw", a.HandleDraw).Methods("POST")

	// a.router.PathPrefix("").Handler(http.FileServer(http.Dir("./assets")))

	a.router.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

	a.router.Use(
		withLogging,
	)
}

// ListenAndServer start the Api-Server on the given address
func (a *Api) ListenAndServe() error {
	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	logrus.Infof("[API-Server] start listening on: %q\n", a.addr)
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
