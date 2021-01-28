package commonenvironment

import (
	"encoding/base64"
	"encoding/json"
	"html"
	"net/url"
	"strings"

	"github.com/herb-go/herbtext"
)

//DirectiveParserString directive for parser string
const DirectiveParserString = "string"

//ParserString parser string.
//Return given data as parsed result.
func ParserString(data string) (interface{}, error) {
	return data, nil
}

//DirectiveParserJSON directive for parser json
const DirectiveParserJSON = "json"

//ParserJSON parser json.
//Parse give data in json format.
//Nil will be returned if empty data given
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

//DirectiveConverterToLower directive for converter tolower.
const DirectiveConverterToLower = "lower"

//ConverterToLower converter tolower.
//Convert give data to lower.
func ConverterToLower(data string) string {
	return strings.ToLower(data)
}

//DirectiveConverterToUpper directive for converter touppwer.
const DirectiveConverterToUpper = "upper"

//ConverterToUpper converter toupper.
//Convert give data to lower.
func ConverterToUpper(data string) string {
	return strings.ToUpper(data)
}

//DirectiveConverterTrim directive for converter trim.
const DirectiveConverterTrim = "trim"

//ConverterTrim converter trim.
//Trim  data spaces.
func ConverterTrim(data string) string {
	return strings.TrimSpace(data)
}

//DirectiveConverterBase64Encode directive for converter base64encode.
const DirectiveConverterBase64Encode = "base64encode"

//ConverterBase64Encode converter base64encode.
//Base64 encode data.
func ConverterBase64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

//DirectiveConverterURLEncode directive for converter urlencode.
const DirectiveConverterURLEncode = "urlencode"

//ConverterURLEncode converter urlencode.
//Base64 encode data.
func ConverterURLEncode(data string) string {
	return url.PathEscape(data)
}

//DirectiveConverterHTMLEscape directive for converter html escape.
const DirectiveConverterHTMLEscape = "htmlescape"

//ConverterHTMLEscape converter html escape.
//Escape data as html.
func ConverterHTMLEscape(data string) string {
	return html.EscapeString(data)
}

var commaescaper = strings.NewReplacer(
	`,`, "&#44;",
	`&`, "&amp;",
)
var commaunescaper = strings.NewReplacer(
	"&#44;", `,`,
	"&amp;", `&`,
)

//DirectiveConverterCommaEscape directive for converter comma escape.
const DirectiveConverterCommaEscape = "commaescape"

//ConverterCommaEscape converter comma escape.
//Escape ',' amd '&' in data as html.
func ConverterCommaEscape(data string) string {
	return commaescaper.Replace(data)
}

//ConverterCommaUnescape converter comma escape.
//Escape ',' amd '&' in data as html.
func ConverterCommaUnescape(data string) string {
	return commaunescaper.Replace(data)
}

//DirectiveConverterJSONEscape directive for converter json escape
const DirectiveConverterJSONEscape = "jsonescape"

//ConverterJSONEscape converter json escape
//Escape data for safely used in json string.
func ConverterJSONEscape(data string) string {
	escaped, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(escaped[1 : len(escaped)-1])
}

//ApplyDirectives apply common directives to given plain enviroment.
func ApplyDirectives(env *herbtext.PlainEnvironment) {
	env.SetParser(DirectiveParserString, ParserString)
	env.SetParser(DirectiveParserJSON, ParserJSON)
	env.SetConverter(DirectiveConverterToLower, ConverterToLower)
	env.SetConverter(DirectiveConverterToUpper, ConverterToUpper)
	env.SetConverter(DirectiveConverterTrim, ConverterTrim)
	env.SetConverter(DirectiveConverterBase64Encode, ConverterBase64Encode)
	env.SetConverter(DirectiveConverterURLEncode, ConverterURLEncode)
	env.SetConverter(DirectiveConverterHTMLEscape, ConverterHTMLEscape)
	env.SetConverter(DirectiveConverterCommaEscape, ConverterCommaEscape)
	env.SetConverter(DirectiveConverterJSONEscape, ConverterJSONEscape)
}

func init() {
	env := herbtext.Clone(herbtext.DefaultEnvironment())
	ApplyDirectives(env)
	herbtext.SetDefaultEnvironment(env)
}
