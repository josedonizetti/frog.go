package frog

import (
  "testing"
  "fmt"
)

func TestLex(t *testing.T) {
  l := lex("jao", "html { head {} }")

  loop: for {
    item := l.nextItem()
    if item.typ == itemEOF {
      break loop
    }
    fmt.Println(item.typ)
    fmt.Println(item)
  }
}
