package router

// import (
// 	"net/http"
// 	"fmt"
//

// 	"github.com/gorilla/mux"
// )

// type muxRouter struct {}

// var (
// 	muxDispatcher = mux.NewRouter()
// )

// func NewMuxRouter() Router {
// 	return &muxRouter{}
// }

// func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)){
// 	muxDispatcher.HandleFunc(uri, f).Methods("GET")
// }

// func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)){
// 	muxDispatcher.HandleFunc(uri, f).Methods("POST")
// }

// func (*muxRouter) SERVE(port string){
// 	fmt.Printf("Mux HTTP server running on port %v", port)
// 	http.ListenAndServe(port, muxDispatcher)
// }
