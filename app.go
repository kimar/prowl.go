package main

import (
  "fmt"
  "net/http"
  "net/url"

  "github.com/zenazn/goji"
  "github.com/zenazn/goji/web"
)

func prowl (c web.C, w http.ResponseWriter, r *http.Request) {
  resp, _ := http.PostForm("https://api.prowlapp.com/publicapi/add", url.Values{"apikey": {c.URLParams["apikey"]}, "application": {"prowl.go"}, "description": {c.URLParams["message"]}});
  fmt.Fprintf(w, resp.Status);
}

func main () {
  goji.Get("/:apikey/:message", prowl);
  goji.Serve();
}
