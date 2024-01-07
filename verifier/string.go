package verifier

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// AlphaVerifier 用于生成字符串是否全部是字母的判断代码
type AlphaVerifier struct {
	FmtProvider
	name string
}

func (v *AlphaVerifier) Generate(field string, value interface{}, params ...string) (judgement string, err error) {
	_, err = OneParamPreCheck(params...)
	if err != nil {
		return "", errors.WithMessage(err, "param of alpha invalid")
	}

	if _, ok := value.(string); !ok {
		return "", ErrAlphaCheckType
	}

	const template = `for _, ch := range %s {
		if !(ch >= 'a' && ch <= 'z') && !(ch >= 'A' && ch <= 'Z') {
			return fmt.Errorf("%s")
		}
	}`

	return fmt.Sprintf(template, field, constructErrMsg(field, v.name)), nil
}

// NumberVerifier 用于生成字符串是否构成数字的判断代码
type NumberVerifier struct {
	StrconvProvider
	name string
}

func (v *NumberVerifier) Generate(field string, value interface{}, params ...string) (judgement string, err error) {
	_, err = OneParamPreCheck(params...)
	if err != nil {
		return "", errors.WithMessage(err, "param of number invalid")
	}

	if _, ok := value.(string); !ok {
		return "", ErrNumberCheckType
	}

	const template = `if _, err := strconv.ParseFloat(%s, 64); err != nil {
		return fmt.Errorf("%s")
	}`

	return fmt.Sprintf(template, field, constructErrMsg(field, v.name)), nil
}

// ContainsVerifier 用于生成字符串是否包含某子字符串的判断代码
type ContainsVerifier struct {
	StringsProvider
	name string
}

func (v *ContainsVerifier) Generate(field string, value interface{}, params ...string) (judgement string, err error) {
	param, err := TwoParamsPreCheck(params...)
	if err != nil {
		return "", errors.WithMessage(err, "param of contains invalid")
	}

	// param中可以通过 | 划分多个候选词
	candidates := strings.Split(param, "|")

	const template = `containsCandidates := %s
		hasContains := false
		for _, c := range containsCandidates {
			if strings.Contains(%s, c) {
				hasContains = true
				break
			}
		}
		if !hasContains {
			return fmt.Errorf("%s")
		}
	`

	return fmt.Sprintf(template, makeStringSliceStr(candidates), field, constructErrMsg(field, v.name)), nil
}

// StartsWithVerifier 用于生成字符串是否以特定子串开头的判断代码
type StartsWithVerifier struct {
	StringsProvider
	name string
}

func (v *StartsWithVerifier) Generate(field string, value interface{}, params ...string) (judgement string, err error) {
	param, err := TwoParamsPreCheck(params...)
	if err != nil {
		return "", errors.WithMessage(err, "param of startswith invalid")
	}

	// param中可以通过 | 划分多个候选词
	candidates := strings.Split(param, "|")

	const template = `startsWithCandidates := %s
	hasStartsWith := false
		for _, c := range startsWithCandidates {
			if strings.Contains(%s, c) {
				hasStartsWith = true
				break
			}
		}
		if !hasStartsWith {
			return fmt.Errorf("%s")
		}
	`

	return fmt.Sprintf(template, makeStringSliceStr(candidates), field, constructErrMsg(field, v.name)), nil
}

// EndsWithVerifier 用于生成字符串是否以特定子串结尾的判断代码
type EndsWithVerifier struct {
	StringsProvider
	name string
}

func (v *EndsWithVerifier) Generate(field string, value interface{}, params ...string) (judgement string, err error) {
	param, err := TwoParamsPreCheck(params...)
	if err != nil {
		return "", errors.WithMessage(err, "param of endswith invalid")
	}

	// param中可以通过 | 划分多个候选词
	candidates := strings.Split(param, "|")

	const template = `endsWithCandidates := %s
		hasEndsWith := false
		for _, c := range endsWithCandidates {
			if strings.Contains(%s, c) {
				hasEndsWith = true
				break
			}
		}
		if !hasEndsWith {
			return fmt.Errorf("%s")
		}
	`

	return fmt.Sprintf(template, makeStringSliceStr(candidates), field, constructErrMsg(field, v.name)), nil
}

func makeStringSliceStr(strs []string) string {
	builder := strings.Builder{}
	builder.WriteString("[]string{")
	for i := range strs {
		builder.WriteByte('"')
		builder.WriteString(strs[i])
		builder.WriteByte('"')
		if i != len(strs)-1 {
			builder.WriteByte(',')
		}
	}
	builder.WriteString("}")

	return builder.String()
}
