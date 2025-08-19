 package main

import (
    "fmt"
    "strings"
    )

type Formatter interface {
    Format(text string) string
}

type UpperCaseFormatter struct{}
func (u *UpperCaseFormatter) Format(text string) string {
    return strings.ToUpper(text)
}

type LowerCaseFormatter struct{}
func (l *LowerCaseFormatter) Format(text string) string {
    return strings.ToLower(text)
}

type TextProcessor struct {
    formatter Formatter
}

func (t *TextProcessor) SetFormatter(f Formatter) {
    t.formatter = f
}

func (t *TextProcessor) Process(text string) string {
    return t.formatter.Format(text)
}

func main() {
    tp := &TextProcessor{}

    // Use upper case strategy
    tp.SetFormatter(&UpperCaseFormatter{})
    fmt.Println(tp.Process("Hello Strategy!")) // Output: HELLO STRATEGY!

    // Use lower case strategy
    tp.SetFormatter(&LowerCaseFormatter{})
    fmt.Println(tp.Process("Hello Strategy!")) // Output: hello strategy!
}

