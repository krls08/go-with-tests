// quiz... please consider the following code and it is a piece of a webserver handler:
//
//	   switch r.Method {
//		case http.MethodPost:
//			addNewItem()
//	       fallthrough
//	   case http.MethodGet:
//	       renderHTMLPage()
//		default:
//			answerMethodNotAllowed()
//		}
//
// now answer:
//
//	what's going to happen given curl -I or HEAD request?
//	what's going to happen if I send a GET request?
//	what's going to happen if I send a POST request?
//

package main

import (
	"log"
	"net/http"
)

func handleFoo(w http.ResponseWriter, r *http.Request) {
	log.Println("handleFoo called")
	switch r.Method {
	case http.MethodPost:
		//addNewItem()
		log.Println("add new item")
		fallthrough
	case http.MethodGet:
		//renderHTMLPage()
		log.Println("render page")
	default:
		//answerMethodNotAllowed()
		log.Println("now allowed")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/foo", handleFoo)
	http.ListenAndServe(":3333", mux)
	//http.ListenAndServe(":3333", nil)
}
