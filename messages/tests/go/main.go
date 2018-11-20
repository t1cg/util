package main

import (
  "fmt"
  app "github.com/t1cg/util/messages/application"
  user "github.com/t1cg/util/messages/user"
)

func main() {
  fmt.Println(user.Messages.CONNECTION_REFUSED.Header)
  fmt.Println(app.Messages.SUCCESS.Message)
}
