package main
// TODO : remove POCGW import 
import (
	"POCGW/utils"
	"POCGW/services"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"net/http/pprof"
)

var con = utils.Connect()

func newsHandler(w http.ResponseWriter, r *http.Request) {
	news := services.FetchNews2()
	ssrResponse := utils.Send(con, "NEWS-CONTENT", news)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, ssrResponse)
}
/*
func footerHandler(w http.ResponseWriter, r *http.Request) {
	var ssrResponse string
	ssrResponse = utils.Send(con, "FOOTER-CONTENT", "Emre")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, ssrResponse)
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	var ssrResponse string
	ssrResponse = utils.Send(con, "HEADER-CONTENT", "")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, ssrResponse)
}
*/
func main() {
	fmt.Println("App started!");
    router := mux.NewRouter()
	router.HandleFunc("/news", newsHandler).Methods("GET")
	// router.HandleFunc("/footer", footerHandler).Methods("GET")
	// router.HandleFunc("/header", headerHandler).Methods("GET")



	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))


	http.ListenAndServe(":8080", router)
}
