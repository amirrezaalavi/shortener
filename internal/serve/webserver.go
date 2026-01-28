package server

import (
	"fmt"
	"net/http"
)

func server() {
	start_http_server()
}
func start_http_server() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/", link_extractor)
	http.ListenAndServe("127.0.0.2:8001", mux)

}
func link_extractor(resp http.ResponseWriter, requ *http.Request) {
	fmt.Fprintln(resp, requ.URL.Path)
	var long_url string = requ.URL.Path[5:]
	var gen_url, err string
	gen_url, err = generate_url(long_url)
	if err != "" {
		fmt.Fprintln(resp, "You've got an error: "+err)
		return
	}
	fmt.Fprintln(resp, "Success, Generated URL is: \n"+gen_url)
}
