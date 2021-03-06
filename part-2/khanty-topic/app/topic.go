package main

import (
    "fmt"
    "strings"
    "log"
    "encoding/json"
    "net/http"
    "github.com/samarkanov/khanty-app/utils"
)

type HantyTopicContext map[string] []string

func get_context() (HantyTopicContext) {
    // let's make the service retain the context information
    // we will rework it later on

    var ctx HantyTopicContext = make(HantyTopicContext)

    ctx["ChangeColor"] = []string{"#5e72e4", "#f3a4b5", "#ffd600",
                                  "#2bffc6", "#fd5d93", "#ffffff"}

    ctx["SendMessage"] = []string{"SayHi", "SayStopIt"}

    return ctx
}

func handle_request(w http.ResponseWriter, r * http.Request) {
    // allow CORS: so the request can be done from anywhere
    // It is ok for test environment
    w.Header().Set("Access-Control-Allow-Origin", "*")

    var res string
    var ctx HantyTopicContext
    ctx = get_context()

    if r.URL.Path == "/" {
        // return all supported topics
        supported_topics := make([]string, 0, len(ctx))
        for k := range ctx {
           supported_topics = append(supported_topics, k)
        }
        var reply = make(map[string] []string)
        reply["topics"] = supported_topics
        res_, _ := json.Marshal(reply)
        res = string(res_)
    } else {
        // return list of states for a given topic
        topic := strings.Split(r.URL.Path, "/")[1]
        topic = strings.Split(topic, "/")[0]
        var reply = make(map[string]interface{})
        reply["topic"] = topic
        reply["states"] = ctx[topic]
        res_, _ := json.Marshal(reply)
        res = string(res_)
    }

    fmt.Fprintf(w, res)
}

func main() {
    http.HandleFunc("/", handle_request)
    log.Fatal(http.ListenAndServe(":"+utils.Portno("topic"), nil))
}
