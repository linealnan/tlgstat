package tlgstat

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io"
    "io/ioutil"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "<html><body>Hi, I'm a <strike>taxman</strike> statman and I'm fine!</body></html>")
}

/*
    Test with this curl command:
    curl -H "Content-Type: application/json" -d '{"event_action":"event_action_test_value"}' http://localhost:8181/utm
*/
func utmMarkupAdd(w http.ResponseWriter, r *http.Request) {
    var channel Channel
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
