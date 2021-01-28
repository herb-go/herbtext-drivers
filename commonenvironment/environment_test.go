package commonenvironment

import (
	"testing"

	"github.com/herb-go/herbtext"
)

func TestDefaultParser(t *testing.T) {
	p := herbtext.DefaultEnvironment().GetParser(DirectiveDefaultParser)
	if p == nil {
		t.Fatal(p)
	}
	v, err := p("test")
	if err != nil {
		panic(err)
	}
	str, ok := v.(string)
	if !ok || str != "test" {
		t.Fatal(str, ok)
	}
}

func TestParserString(t *testing.T) {
	p := herbtext.DefaultEnvironment().GetParser(DirectiveParserString)
	if p == nil {
		t.Fatal(p)
	}
	v, err := p("test")
	if err != nil {
		panic(err)
	}
	str, ok := v.(string)
	if !ok || str != "test" {
		t.Fatal(str, ok)
	}
}

func TestParserJSON(t *testing.T) {
	p := herbtext.DefaultEnvironment().GetParser(DirectiveParserJSON)
	if p == nil {
		t.Fatal(p)
	}
	v, err := p("\"test\"")
	if err != nil {
		panic(err)
	}
	str, ok := v.(string)
	if !ok || str != "test" {
		t.Fatal(str, ok)
	}
	v, err = p("")
	if err != nil || v != nil {
		panic(err)
	}

}

type convertertestset struct {
	Directive string
	Given     string
	Wanted    string
}

var testset = []*convertertestset{
	{DirectiveConverterToLower, "ABCD", "abcd"},
	{DirectiveConverterToUpper, "abcd", "ABCD"},
	{DirectiveConverterTrim, " abcd ", "abcd"},
	{DirectiveConverterBase64Encode, "abcd", "YWJjZA=="},
	{DirectiveConverterURLEncode, "abcd?", "abcd%3F"},
	{DirectiveConverterHTMLEscape, "<ab,cd>", "&lt;ab,cd&gt;"},
	{DirectiveConverterCommaEscape, "<ab,cd>", "<ab&#44;cd>"},
	{DirectiveConverterJSONEscape, "ab\"cd", `ab\"cd`},
}

func TestConverters(t *testing.T) {
	for _, v := range testset {
		c := herbtext.DefaultEnvironment().GetConverter(v.Directive)
		if c == nil {
			t.Fatal()
		}
		result := c(v.Given)
		if result != v.Wanted {
			t.Fatal(v.Directive, v.Given, v.Wanted, result)
		}
	}
}
