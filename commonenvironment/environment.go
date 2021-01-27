package commonenvironment

import (
	"encoding/base64"
	"encoding/json"
	"html"
	"mime"
	"net/url"
	"strings"

	"github.com/herb-go/herbtext"
)

const DirectiveParserString = "string"

func ParserString(data string) (interface{}, error) {
	return data, nil
}

const DirectiveParserJSON = "json"

func ParserJSON(data string) (interface{}, error) {
	if data == "" {
		return nil, nil
	}
	var result interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

const DirectiveConverterToLower = "lower"

func ConverterToLower(data string) string {
	return strings.ToLower(data)
}

const DirectiveConverterToUpper = "upper"

func ConverterToUpper(data string) string {
	return strings.ToUpper(data)
}

const DirectiveConverterTrim = "trim"

func ConverterTrim(data string) string {
	return strings.TrimSpace(data)
}

const DirectiveConverterBase64 = "base64encode"

func ConverterBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

const DirectiveConverterURLEncode = "urlencode"

func ConverterURLEncode(data string) string {
	return url.PathEscape(data)
}

const DirectiveConverterHTMLEscape = "htmlescape"

func ConverterHTMLEscape(data string) string {
	return html.EscapeString(data)
}

var commaescaper = strings.NewReplacer(
	`,`, "&#44;",
	`&`, "&amp;",
)

const DirectiveConverterCommaEscape = "commaescape"

func ConverterCommaEscape(data string) string {
	return commaescaper.Replace(data)
}

const DirectiveConverterBEncoding = "bencoding"

func ConverterBEncoding(data string) string {
	return mime.BEncoding.Encode("utf-8", data)
}

const DirectiveConverterJSONEscape = "jsonescape"

func ConverterJSONEscape(data string) string {
	escaped, err := json.Marshal([]byte(data))
	if err != nil {
		panic(err)
	}
	return string(escaped[1 : len(escaped)-2])
}

func ApplyDirectives(env *herbtext.Environment) {
	env.SetParser(DirectiveParserString, ParserString)
	env.SetParser(DirectiveParserJSON, ParserJSON)
	env.SetConverter(DirectiveConverterToLower, ConverterToLower)
	env.SetConverter(DirectiveConverterToUpper, ConverterToUpper)
	env.SetConverter(DirectiveConverterTrim, ConverterTrim)
	env.SetConverter(DirectiveConverterBase64, ConverterBase64)
	env.SetConverter(DirectiveConverterURLEncode, ConverterURLEncode)
	env.SetConverter(DirectiveConverterHTMLEscape, ConverterHTMLEscape)
	env.SetConverter(DirectiveConverterCommaEscape, ConverterCommaEscape)
	env.SetConverter(DirectiveConverterBEncoding, ConverterBEncoding)
	env.SetConverter(DirectiveConverterJSONEscape, ConverterJSONEscape)
}

func init() {
	ApplyDirectives(herbtext.DefaultEnvironment())
}
