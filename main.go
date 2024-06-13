package main

import (
	"fmt"
	"golangNative/helpers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Memulai Server . . . ")
	routers := httprouter.New()
	routers.GET("/" , func(w http.ResponseWriter , r *http.Request , p httprouter.Params){
		fmt.Fprint(w , "helllow welcome to golang ade ..")
	})
	server := http.Server{Addr: "localhost:8888" , Handler : routers}
	err := server.ListenAndServe()
	helpers.PanicError((err))
}