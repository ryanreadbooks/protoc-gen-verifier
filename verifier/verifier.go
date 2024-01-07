package verifier

import "fmt"

const (
	errTemplate = "verifier.Verify: field '%s' failed on '%s' tag"
)

func constructErrMsg(field, tag string) string {
	return fmt.Sprintf(errTemplate, field, tag)
}



type Verifier interface {
	DependencyProvider
	// Generate 生成对应的校验语句
	Generate(field string, value interface{}, params ...string) (string, error)
}

type emptyVerifier struct {
	name string
}

func (v *emptyVerifier) Generate(field string, value interface{}, params ...string) (string, error) {
	return "", nil
}
