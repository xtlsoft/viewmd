package main

import (
	"fmt"
	"io/ioutil"

	"github.com/andlabs/ui"
	"gitlab.com/golang-commonmark/markdown"
)

func realMain() {
	file := "README.md"
	parsed, _ := readFile(file)
	fmt.Println(parsed[1].Tag())
	restoreText()
	appendWithAttributes("This is a title", ui.TextSize(35), ui.TextWeightSemiBold)
	appendWithAttributes("\r\nHello, everyone! A demo from me.")
	appendWithAttributes("\r\nNow let's have a look at it.", ui.TextItalicItalic, ui.UnderlineDouble)
	appendWithAttributes("\r\nSecond Paragraph", ui.TextSize(24), ui.TextFamily("Ubuntu Mono"))
	appendWithAttributes("\r\nImportant Things Should Have Important Colours, like me.", ui.TextColor{
		R: 255,
		G: 0,
		B: 0,
		A: 50,
	})
	appendWithAttributes("But I don't have any colour, in fact, I'm important.")
}

func readFile(file string) ([]markdown.Token, error) {
	md := markdown.New()
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return md.Parse(f), nil
}
