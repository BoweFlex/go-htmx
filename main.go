package main

import (
    "fmt"
    "io" 
    "flag" // Allows handling named CLI parameters, easier to work with than os.Args
    "net/http"
    "log"
    "html/template"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello From BoweFlex's http server\n")
    io.WriteString(w, r.Method)
}

func main() {
    var port int
    var host string

    flag.IntVar(&port, "port", 1234, "Provide a port number")
    flag.StringVar(&host, "host", "", "Provide a host")

    flag.Parse()

    fullHost := fmt.Sprintf("%v:%v", host, port)

    http.Handle("/style/",
        http.StripPrefix("/style/", http.FileServer(http.Dir("./web/style"))))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("./web/templates/index.html"))
        tmpl.Execute(w, nil)
    })

    http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("./web/templates/fragments/results.html"))
        data := map[string][]Stock {
            "Results": SearchTicker(r.URL.Query().Get("key")),
        }
        tmpl.Execute(w, data)
    })

    http.HandleFunc("/stock/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			ticker := r.PostFormValue("ticker")
			stk := SearchTicker(ticker)[0]
			val := GetDailyValues(ticker)
			tmpl := template.Must(template.ParseFiles("./templates/index.html"))
			tmpl.ExecuteTemplate(w, "stock-element",
				Stock{Ticker: stk.Ticker, Name: stk.Name, Price: val.Open})
		}
	})

    log.Println("Listening on ", fullHost)
    log.Fatal(http.ListenAndServe(fullHost, nil))

}
