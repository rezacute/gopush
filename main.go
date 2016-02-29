package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
    "github.com/pusher/pusher-http-go"
    "github.com/gorilla/mux"
)
var appId string
var appKey string
var secret string
func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/reply/{answer}", InvitationReply)
    appId = os.Getenv("APP_ID")
    appKey = os.Getenv("APP_KEY")
    secret = os.Getenv("SECRET")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo Index!")
    client := pusher.Client{
    AppId: appId,
    Key: appKey,
    Secret: secret,
    }

    data := map[string]string{"message": "hello world"}

    client.Trigger("test_channel", "my_event", data)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}
func InvitationReply(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    answer := vars["answer"]

    if answer == "" {
      w.WriteHeader(http.StatusBadRequest)
      return
    }else{
      fmt.Fprintln(w, "Reply :", answer)
    }
    client := pusher.Client{
    AppId: appId,
    Key: appKey,
    Secret: secret,
    }

  data := map[string]string{"message": answer}

  client.Trigger("test_channel", "my_event", data)
}
