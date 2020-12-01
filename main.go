package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/execute", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		commandLine := query.Get("cmd")

		arguments := strings.Split(commandLine, " ")

		cmd := exec.Command(arguments[0], arguments[1:]...)

		out, e := cmd.Output()

		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error: " + e.Error()))

		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(out)
		}
	})

	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		fmt.Println(e)
	}

}
