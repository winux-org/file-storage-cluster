// https://techviewleo.com/how-to-install-and-use-docker-in-linux-mint/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// create response binary data
	data := []byte("Hello World!") // slice of bytes

	// write `data` to response
	res.Write(data)
}

// SetupHTTPRouting configures and setup the endpoints
func SetupHTTPRouting() {
	http.HandleFunc("/init", initFileUploadingHandler)

	router := mux.NewRouter()
	router.HandleFunc("/file/{hash}", uploadFileHandler)
	http.Handle("/", router)
}

type initState struct {
	Hash string `json:"hash"`
}

func initFileUploadingHandler(w http.ResponseWriter, r *http.Request) {
	fileUUID, _ := uuid.NewRandom()
	newFile := initState{Hash: fileUUID.String()}
	newFileJSON, _ := json.Marshal(newFile)
	// w := io.PipeWriter{}
	// defer w.Close()
	// w.Write(newFileJSON)
	w.Write(newFileJSON)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	b := make([]byte, 8)
	fmt.Println("started processing response")
	for {
		//n, err := r.Read(b)
		n, err := r.Body.Read(b)

		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	io.WriteString(w, "Hello world!")
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Print(vars["hash"])
	io.WriteString(w, "Hello world!")
}
