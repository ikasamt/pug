package pug_test

import (
	"github.com/Joker/hpp"
	"testing"
)
import "github.com/ikasamt/pug"


func TestPug_render(t *testing.T) {
	src := `
html
	body
		h1 aa
	`
	actual := hpp.PrPrint(pug.Do(src))
	expected := hpp.PrPrint(`
<html>
	<body>
		<h1>aa </h1>
	</body>
</html>
	`)
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestPug_render2(t *testing.T) {
	src := `
html
	body
		h1 aa
		div
			div
			span aaa
			br
			hr
		div
			span footer
	`
	actual := hpp.PrPrint(pug.Do(src))
	expected := hpp.PrPrint(`
<html>
	<body>
		<h1>aa </h1>
		<div>
			<div></div>
			<span>aaa </span>
			<br/>
			<hr/>
		</div>
		<div>
			<span>footer </span>
		</div>
	</body>
</html>
`)
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
