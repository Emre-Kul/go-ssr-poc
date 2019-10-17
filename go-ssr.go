package main
// TODO : remove POCGW import 
import (
	"POCGW/utils"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

var con = utils.Connect()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	name, ok := r.URL.Query()["name"]
	var ssrResponse string
	if(!ok || len(name[0]) < 1) {
		ssrResponse = utils.Send(con, "Emre")	
	} else {
		ssrResponse = utils.Send(con, name[0])
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, ssrResponse)
}

func main() {
	fmt.Println("App started!");
    router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
