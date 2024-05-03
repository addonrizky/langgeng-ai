package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("okee")

	r := mux.NewRouter()

	r.HandleFunc("/getAnswer", getAnswer).Methods("GET")

	_ = http.ListenAndServe(":4004", r)

	fmt.Println("WOOWWIIWIWI")
}

func getAnswer(res http.ResponseWriter, req *http.Request) {
	// echo := exec.Command("echo", os.Getenv("PATH"))
	// fmt.Println("ECHO", echo)

	q := req.URL.Query().Get("question")
	output, err := exec.Command("/root/.venvs/MyEnv/bin/python", "minimal.py", q).Output()
	if err != nil {
		fmt.Println("error nih : ", err)
	}

	fmt.Println(string(output))
	res.Write([]byte(output))
}
