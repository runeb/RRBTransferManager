package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

var fixtures map[string]string = make(map[string]string)

func generateFixtures() {
	fmt.Println("Generating fixtures...")

	// dd if=/dev/urandom bs=1000 count=1 of=fixtures/small
	app := "dd"
	arg0 := "if=/dev/urandom"
	arg1 := "bs=1000"
	arg2 := "count=1"
	arg3 := "of=fixtures/small"
	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))

	// md5 fixtures/small

	app = "md5"
	arg0 = "-q"
	arg1 = "fixtures/small"
	cmd = exec.Command(app, arg0, arg1)
	stdout, err = cmd.Output()
	md5 := string(stdout)
	strings.TrimSpace(md5)
	fixtures["small"] = strings.ToUpper(md5)

	fmt.Println(fixtures)
}

func manifest(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(fixtures)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func cleanUp() {
	fmt.Println("Deleting fixtures...")
	app := "rm"
	arg0 := "fixtures/small"
	cmd := exec.Command(app, arg0)
	stdout, _ := cmd.Output()
	print(string(stdout))
}

func main() {
	// Trap ctrl-c for exiting
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("captured %v, stopping server and cleaning up..\n", sig)
			cleanUp()
			os.Exit(0)
		}
	}()

	generateFixtures()
	http.HandleFunc("/manifest", manifest)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		http.ServeFile(w, r, "fixtures/"+r.URL.Path[1:])
	})
	fmt.Println("Ready for testing")
	fmt.Println("Ctrl-C to quit")
	panic(http.ListenAndServe(":2222", nil))
}
