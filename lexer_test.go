package frog

import (
  "testing"
)

func collect(lexer *lexer) (items []item){
  for {
    item := lexer.nextItem()
    items = append(items, item)
    if item.typ == itemEOF {
      break
    }
  }
  return
}

func equal(got, expected []item) bool {
  if len(got) != len(expected) {
    return false
  }

  for i := range got {
    if got[i].typ != expected[i].typ {
      return false
    }

    if got[i].val != expected[i].val {
      return false
    }
  }

  return true
}

var (
    tEOF = item{itemEOF,""}
    tSpace = item{itemSpace," "}
)

type LexTest struct {
  name string
  input string
  items []item
}

var lexTests = []LexTest{
  LexTest{"noSpace","html{}", []item{
    {itemText, "html"},
    {itemText, "{"},
    {itemText, "}"},
    tEOF,
  }},
  LexTest{"spaces", "html { }", []item{
    {itemText, "html"},
    tSpace,
    {itemText, "{"},
    tSpace,
    {itemText, "}"},
    tEOF,
  }},
  //LexTest{"html \"hi there\"", []item{
  //  {itemText, "html"},
  //  {itemString, "\"hi there\""},
  //}},
}

func TestLex(t *testing.T) {
  for _,lexTest := range lexTests {
    lexer := lex(lexTest.input)
    items := collect(lexer)
    if !equal(items, lexTest.items) {
      t.Errorf("test %s items should be equal, got\n\t%+v\nexpected\n\t%+v",lexTest.name, items, lexTest.items)
    }
  }
}
