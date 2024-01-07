package verifier

var (
	tagVerifierMapping = map[string]Verifier{
		"eq":  &EqVerifier{name: "eq"},
		"gt":  &GtVerifier{name: "gt"},
		"gte": &GteVerifier{name: "gte"},
		"lt":  &LtVerifier{name: "lt"},
		"lte": &LteVerifier{name: "lte"},
		"ne":  &NeVerifier{name: "ne"},

		"len":        &LenVerifier{name: "len"},
		"alpha":      &AlphaVerifier{name: "alpha"},
		"number":     &NumberVerifier{name: "number"},
		"contains":   &ContainsVerifier{name: "contains"},
		"startswith": &StartsWithVerifier{name: "startswith"},
		"endswith":   &EndsWithVerifier{name: "endswith"},
	}
)
