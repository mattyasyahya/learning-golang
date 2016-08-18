package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res,
		`
		<html>
			<head>
				<title>Sample Go Server</title>
			</head>
			<body>
				<h1>Go Http Server</h1>
			</body>
		</html>
		`,
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("input"))))

	http.ListenAndServe(":9000", nil)
}
