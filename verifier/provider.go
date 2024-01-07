package verifier

type DependencyProvider interface {
	Dependency() []string
}

type FmtProvider struct{}

func (p *FmtProvider) Dependency() []string {
	return []string{"fmt"}
}

type StrconvProvider struct{}

func (p *StrconvProvider) Dependency() []string {
	return []string{"fmt", "strconv"}
}

type StringsProvider struct{}

func (p *StringsProvider) Dependency() []string {
	return []string{"fmt", "strings"}
}
