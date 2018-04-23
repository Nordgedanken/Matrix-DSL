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
	Key    string   `@Ident ":"`
	Arrays []*Array `[ { @@ } ]`
	Value  *Value   `[ | @@ ]`
	Event  *string  `[ | "["@Ident"]" ]`
}

type Property struct {
	Key   string  `@Ident ":"`
	Value *Value  `[ @@ ]`
	Event *string `[ | "["@Ident"]" ]`
}

type Array struct {
	Key        string      ` "-"`
	Properties []*Property `{ @@ }`
}

type Value struct {
	String *string ` @String`
	Bool   *string `| @String`
}
