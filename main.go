package main

import (
"fmt"
"bytes"
"strconv"
"os"
"log"
"encoding/base64"
"regexp"
"net/http"
"io/ioutil"
)

func print_image(body []byte) {

  body64 := base64.StdEncoding.EncodeToString(body)
  term := os.Getenv("TERM")
  matched, err := regexp.MatchString("screen-\\w+",term)

  if err != nil {
    log.Fatal(err)
  }

  length := len(body)
  var buf bytes.Buffer

  if matched {
    buf.WriteString("\033Ptmux;\033\033]")
  } else{
    buf.WriteString("\033]")
  }

  buf.WriteString("1337;File=")
  buf.WriteString("size=")
  buf.WriteString(strconv.Itoa(length))
  buf.WriteString(";inline=1")
  buf.WriteString(":")
  buf.WriteString(body64)
  buf.WriteString("")
  buf.WriteString("\n")

  fmt.Println(buf.String())
}

func main() {

  url := "https://api.thecatapi.com/v1/images/search?format=src&mime_types=image/gif"
  req, _ := http.NewRequest("GET", url, nil)
  res, _ := http.DefaultClient.Do(req)
  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  print_image(body)
}
