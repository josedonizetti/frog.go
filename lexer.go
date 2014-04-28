package frog

import (
  "unicode/utf8"
  "fmt"
)

type item struct {
    typ itemType
    val string
}

func (i item) String() string {
    switch i.typ {
    case itemEOF:
        return "EOF"
    case itemError:
        return i.val
    }
    if len(i.val) > 10 {
        return fmt.Sprintf("%.10q...", i.val)
    }
    return fmt.Sprintf("%q", i.val)
}

type itemType int

const (
    itemError itemType = iota
    itemEOF
    itemString
    itemText
    itemId
    itemClass
    itemExpression
    itemStatement
)

type lexer struct {
  name string
  input string
  start int
  pos int
  width int
  items chan item
}

const eof = -1

type stateFn func(*lexer) stateFn

func lex(name, input string) (*lexer) {
  l := &lexer{
    name: name,
    input: input,
    items: make(chan item),
  }

  go l.run()

  return l
}

func (l *lexer) run() {
  for state := lexText; state != nil; {
    state = state(l)
  }
  close(l.items)
}

func (l *lexer) emit(t itemType) {
  fmt.Println(l.start, l.pos)
  l.items <- item{t, l.input[l.start:l.pos]}
  l.start = l.pos
}

func (l *lexer) nextItem() item {
  item := <-l.items
  return item
}

func lexText(l *lexer) stateFn {
  for {
    r := l.next()

    if isSpace(r) {
      l.backup()
      l.emit(itemText)
      l.ignoreSpace()
    }

    if r == eof { break }
  }

  if l.pos > l.start {
    l.emit(itemText)
  }

  l.emit(itemEOF)
  return nil
}

func (l *lexer) next() (rune) {
  if l.pos >= len(l.input) {
    l.width = 0
    return eof
  }

  r, w := utf8.DecodeRuneInString(l.input[l.pos:])
  l.width = w
  l.pos += l.width
  return r
}

func (l *lexer) backup() {
  l.pos -= 1
}

func (l *lexer) ignoreSpace() {
  l.pos += 1
  l.start += 1
}

func isSpace(r rune) bool {
  return r == ' ' || r == '\t'
}
