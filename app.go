package main

import (
  "log"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  "strconv"
)


type StrResponse struct {
  Message string `json:"message"`
}


func HelloHandler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    res := &StrResponse{
      Message: "Hello World!",
    }

    data, err := json.Marshal(res)

    if err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
     }

    w.Header().Set("Content-Type", "application/json")
    w.Write(data)
}

type NumResponse struct {
  Number int `json:"number"`
}


func PlusHandler(w http.ResponseWriter, r *http.Request) {
  defer func() {
    r.Body.Close()
    err := recover()
    if err != nil {
      http.Error(w, "Error", http.StatusInternalServerError)
    }
  }()

  vars := mux.Vars(r)

  num, err := strconv.Atoi(vars["num"])
  if err != nil {
    panic(err)
  }

  res := &NumResponse{
    Number: num + 1,
  }

  data, err := json.Marshal(res)
  if err != nil {
    panic(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(data)
}


func GetHandler(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()

  name, ok := r.URL.Query()["name"]
  print(ok)
  if !ok {
      http.Error(w, "'name' must be required. ", http.StatusBadRequest)
      return
  }

  res := &StrResponse{
    Message: "Hello, " + name[0] + "!",
  }

  data, err := json.Marshal(res)

  if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
   }

  w.Header().Set("Content-Type", "application/json")
  w.Write(data)
}


func PostHandler(w http.ResponseWriter, r *http.Request) {
  defer func() {
    r.Body.Close()
    err := recover()
    if err != nil {
      http.Error(w, "Error", http.StatusInternalServerError)
    }
  }()

  err := r.ParseForm()
  if err != nil {
    panic(err)
  }

  name, ok := r.Form["name"]
  if !ok {
      http.Error(w, "'name' must be required. ", http.StatusBadRequest)
      return
  }

  res := &StrResponse{
    Message: "Hello, " + name[0] + "!",
  }

  data, err := json.Marshal(res)

  if err != nil {
    panic(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(data)
}


func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HelloHandler)
  r.HandleFunc("/plus/{num:[0-9]+}", PlusHandler)
  r.HandleFunc("/get", GetHandler).Methods("GET")
  r.HandleFunc("/post", PostHandler).Methods("POST")
  log.Fatal(http.ListenAndServe(":80", r))
}
