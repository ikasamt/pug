package pug

import (
	"fmt"
	"strings"
)

const SEP = "\n"
const TAB = "	"
const TAB_CODE = 9
const SPACE_CODE = 32

type Token struct{
	Indent int
	Name   string
	Value  string
}

func (t *Token) String() string{
	return fmt.Sprintf("%d, %s, %s", t.Indent, t.Name, t.Value)
}

func (t *Token) NameLower() string{
	return strings.ToLower(t.Name)
}


func (t *Token) OpenTag() string{
	switch t.NameLower() {
	case `br`, `hr`, `i`, `meta`:
		return fmt.Sprintf("<%s/>", t.Name)
	default:
		return fmt.Sprintf("<%s>", t.Name)
	}
}

func (t *Token) CloseTag() string{
	switch t.NameLower() {
	case `br`, `hr`, `i`, `meta`:
		return ""
	default:
		return fmt.Sprintf("</%s>", t.Name)
	}
}

func CountIndent(s string) (count int){
	for _, char := range s{
		switch char {
		case TAB_CODE:
			count++
		case SPACE_CODE:
			count++
		default:
			return
		}
	}
	return
}

func CutTag(s string) (tag string){
	for _, char := range s{
		switch char {
		case TAB_CODE:
			return
		case SPACE_CODE:
			return
		default:
			tag = tag + string(char)
		}
	}
	return
}


func NewToken(s string) Token {
	t := Token{}
	t.Indent = CountIndent(s)
	t.Name = CutTag(s[t.Indent:len(s)])
	t.Value = strings.TrimSpace(s[t.Indent+len(t.Name):len(s)])
	return t
}

func Parse(src string) (tokens []Token) {
	lines := strings.Split(src, SEP)
	for _, line := range lines {
		if line == `` {
			continue
		}

		token := NewToken(line)
		if token.Name == `` {
			continue
		}
		tokens = append(tokens, token)
	}
	return
}

func Render(tokens []Token) string {
	htmlLines := []string{}

	deepestIndent := 0
	closeTags := map[int][]Token{}
	for _, token := range tokens{
		if len(closeTags[token.Indent]) == 0{
			closeTags[token.Indent] = []Token{}
		}

		// Pop CloseTags
		currentCloseTags := closeTags[token.Indent]
		for _, token:= range currentCloseTags{
			htmlLines = append(htmlLines, token.CloseTag())
		}

		// Push CloseTags
		closeTags[token.Indent] = []Token{}
		closeTags[token.Indent] = append(closeTags[token.Indent], token)

		// Pop OpenTag, Value
		htmlLines = append(htmlLines, token.OpenTag())
		htmlLines = append(htmlLines, token.Value)

		if deepestIndent < token.Indent{
			deepestIndent = token.Indent
		}
	}
	for indent:=deepestIndent; indent>=0; indent-- {
		for _, token:= range closeTags[indent]{
			htmlLines = append(htmlLines, token.CloseTag())
		}
	}
	return strings.Join(htmlLines, SEP)
}

func Do(src string) string {
	tokens := Parse(src)
	return Render(tokens)
}