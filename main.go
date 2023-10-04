package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	
	r.HandleFunc("/api/hello", func (w http.ResponseWriter, r *http.Request){
		fmt.Println("Hello, world!")
	})

	http.Handle("/", r)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
