package tlgstat

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io"
    "io/ioutil"
    "github.com/shkatovdm/domain"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "<html><body>Hi, I'm a tlgstat and I'm good!</body></html>")
}

/*
    Test with this curl command:
    curl -H "Content-Type: application/json" -d '{"name":"name_value","link":"link_value", "сategory":"сategory_value","subscribers":"subscribers_value"}' http://localhost:8181/channels
*/
func channelAdd(w http.ResponseWriter, r *http.Request) {
    var channel domain.Channel
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &channel); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    
    fmt.Printf("%v", channel)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
}
