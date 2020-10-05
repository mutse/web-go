package main

import(
    "fmt"
    "net/http"
    "text/template"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
    t,_ := template.ParseFiles("./templates/index.html")
    data := map[string]string{
        "name": "mutse",
        "date": "2020-10-05",
    }
    t.Execute(w, data)
}

func main() {
    http.HandleFunc("/", myWeb)

    staticHandle := http.FileServer(http.Dir("./static"))
    http.Handle("/js/", staticHandle)

    fmt.Println("Starting server, please vist http://localhost:8080")

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("server start error:", err)
    }
}

