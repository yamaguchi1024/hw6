package app

import (
  "fmt"
  "os"
  "net/http"
)

func init() {
  http.HandleFunc("/", main)
  http.HandleFunc("/pata", handlePata)
  http.HandleFunc("/train", handleTrain)
}

func main(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  fmt.Fprintf(w, "Nyaaaaaaaaaaaaaaaaaaann\n")
}

func join(a, b string) string {
  res := ""
  i := 0
  for i < len(a) {
    res += a[i:i+1]
    if i < len(b) {
      res += b[i:i+1]
    }
    i++
  }
  for j := i; j < len(b); j++ {
    res += b[j:j+1]
  }
  return res
}

func handlePata(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    r.ParseForm()
    a := r.FormValue("a")
    b := r.FormValue("b")
    fmt.Fprintf(w, "パタトクカシーー: %s\n", join(a,b))
  } else {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "<html> <head> <title>パタトクカシーー</title> <link rel=\"stylesheet\" href=\"pure/pure-min.css\"> </head> <body> <hr> <form method=\"get\"> <input type=text name=a><br> <input type=text name=b><br> <input type=submit> </form> </body> </html> \n")
  }
}

func handleTrain(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    r.ParseForm()
    a := r.FormValue("a")
    b := r.FormValue("b")
    fmt.Fprintf(w, ": %s\n", train(a,b))
  } else {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "<html> <head> <title>パタトクカシーー</title> <link rel=\"stylesheet\" href=\"pure/pure-min.css\"> </head> <body> <hr> <form method=\"POST\"> <input type=text name=a><br> <input type=text name=b><br> <input type=submit> </form> </body> </html> \n")
  }
}
