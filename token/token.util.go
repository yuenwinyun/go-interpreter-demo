package token

func LookupIdentifier(ident string) Type {
	if tok, ok := builtInKeyWords[ident]; ok {
		return tok
	}
	return IDENT
}
