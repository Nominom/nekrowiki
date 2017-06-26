package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)


func _start(args []string) {
	fs := http.FileServer(http.Dir("public_html"))
	http.Handle("/", customErrorMW(fs))

	//http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(gfmstyle.Assets)))

	http.Handle("/login", customErrorMW(http.HandlerFunc(loginHandler)))
	http.Handle("/logout", customErrorMW(http.HandlerFunc(logoutHandler)))

	http.Handle(config.homePage, customErrorMW(authMW(http.HandlerFunc(homepageHandler))))

	ufs := http.FileServer(http.Dir("uploads"))
	http.Handle("/uploads/", customErrorMW(authMW(http.StripPrefix("/uploads", ufs))))

	http.Handle("/upload", customErrorMW(authMW(http.HandlerFunc(uploadHandler))))
	http.Handle("/md/", customErrorMW(authMW(http.StripPrefix("/md", http.HandlerFunc(markdownHandler)))))
	http.Handle("/edit/", customErrorMW(authMW(http.StripPrefix("/edit", http.HandlerFunc(editHandler)))))


	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func _stop(args []string) {
	fmt.Print("Not implemented")
}

func _restart(args []string) {
	fmt.Print("Not implemented")
}

func _config(args []string) {
	fmt.Print("Not implemented")
}

func _help() {
	fmt.Print("Usage: \n$./nekrowiki start")
}

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		_help()
		return
	} else {
		function := args[0]

		var additional_args []string

		if len(args) > 1 {
			additional_args = args[1:]
		} else {
			additional_args = make([]string, 0)
		}

		switch function {
		case "start":
			_start(additional_args)
		case "stop":
			_stop(additional_args)
		case "restart":
			_restart(additional_args)
		case "config":
			_config(additional_args)
		default:
			_help()
		}

	}

}