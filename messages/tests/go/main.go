package main

import (
  "fmt"
  app "github.com/t1cg/util/messages/go/application"
  user "github.com/t1cg/util/messages/go/user"
)

func main() {
  fmt.Println(app.Messages.SUCCESS.Message)
  fmt.Println(user.Messages.INCORRECT_PASSWORD.Header)
}
