package lexer

type Matrix struct {
	Sections   []*Section        `{ @@ }`
	Properties []*PropertyArrays `{ @@ }`
}

type Section struct {
	Identifier string            `"[""[" @Ident "]""]"`
	Properties []*PropertyArrays `{ @@ }`
}

type PropertyArrays struct {
	Key     string   `@Ident ":"`
	Arrays  []*Array `[ { @@ } ]`
	Value   *Value   `[ | @@ ]`
	Special *Special `[ | @@ ]`
}

type Property struct {
	Key     string   `@Ident ":"`
	Value   *Value   `| @@`
	Special *Special `| @@`
}

type Array struct {
	Key        string      ` "-"`
	Properties []*Property `{ @@ }`
}

type Value struct {
	String *string  ` @String`
	Number *float64 `| @Float`
}

type Special struct {
	Key *string `"["@Ident"]"`
}
