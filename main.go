package main

import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
    "time"
    "math/rand"
)

func quote(w http.ResponseWriter, req *http.Request) {
randsource := rand.NewSource(time.Now().UnixNano())
                randgenerator := rand.New(randsource)
                firstLoc := randgenerator.Intn(100)
                candidate1 := ""

        dat, err := ioutil.ReadFile("quotes.txt")
        if err == nil {
                ascii := string(dat)
                splt := strings.Split(ascii, "\n")
candidate1 = splt[firstLoc]
        }
fmt.Fprintf(w, candidate1)
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/quote", quote)
    http.HandleFunc("/headers", headers)
    http.ListenAndServe(":8080", nil)
}
