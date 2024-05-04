package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("okee")

	r := mux.NewRouter()

	r.HandleFunc("/getAnswer", getAnswer).Methods("GET")
	r.HandleFunc("/enrich", enrichKnowledge).Methods("POST")

	_ = http.ListenAndServe(":4004", r)

	fmt.Println("WOOWWIIWIWI")
}

func getAnswer(res http.ResponseWriter, req *http.Request) {
	q := req.URL.Query().Get("question")
	output, err := exec.Command("/root/.venvs/MyEnv/bin/python", "minimal.py", q).Output()
	if err != nil {
		fmt.Println("error nih : ", err)
	}

	fmt.Println(string(output))
	res.Write([]byte(output))
}

func enrichKnowledge(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var i interface{}
	var body []byte
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&i)
	if err != nil {
		//fmt.Println(err)
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("204 - send post with empty/broken body"))
		//no body or empty body maybe
		return
	} else {
		body, _ = json.Marshal(i)
		// // And now set a new body, which will simulate the same data we read:
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}

	fmt.Println(string(body))

	data := map[string]interface{}{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	appendDocument(data)

	w.Write([]byte("OK"))
}

func appendDocument(data map[string]interface{}) {
	knowledge := data["knowledge"].(string)

	file, err := os.OpenFile("kamus.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Could not open kamus.txt")
		return
	}

	defer file.Close()

	_, err2 := file.WriteString("\n" + knowledge)

	if err2 != nil {
		fmt.Println("Could not write text to kamus.txt")

	} else {
		fmt.Println("Operation successful! Text has been appended to kamus.txt")
	}

	return
}
