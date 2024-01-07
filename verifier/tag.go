package verifier

import (
	"fmt"
	"strings"
)

const (
	VerifyTagPrefix    = "// @verify:"
	VerifyTagPrefixLen = len(VerifyTagPrefix)
)

var (
	ErrTagWrongPrefix = fmt.Errorf("tag wrong prefix")
	ErrUnsupportedTag = fmt.Errorf("unsupported tag")
)

type Tag struct {
	Rule string
}

func NewTag(r string) *Tag {
	return &Tag{
		Rule: strings.TrimSpace(r),
	}
}

// Valid 判断tag是否可以合法
func (t *Tag) Valid(field interface{}) error {
	hasPrefix := strings.HasPrefix(t.Rule, VerifyTagPrefix)
	if !hasPrefix {
		return ErrTagWrongPrefix
	}
	return nil
}

// parse 解析Tag中的Rule 并且校验field是否适用指定的Rule
func (t *Tag) Parse(fieldName string, fieldValue interface{}, fn func(statement string, deps []string)) error {
	// 检验是否符合tag检验的语法规则
	t.Rule = strings.TrimSpace(t.Rule[VerifyTagPrefixLen:])

	subRules := strings.Split(t.Rule, ",")
	existence := make(map[string]struct{}, len(subRules))
	for _, rule := range subRules {
		params := strings.Split(rule, "=")
		if len(params) == 0 {
			continue
		}

		// tag去重
		ruleTag := params[0]
		if _, ok := existence[ruleTag]; ok {
			continue
		}
		existence[ruleTag] = struct{}{}

		verifier, ok := tagVerifierMapping[ruleTag]
		if !ok {
			return ErrUnsupportedTag
		}
		statement, err := verifier.Generate(fieldName, fieldValue, params...)
		deps := verifier.Dependency()
		if err != nil {
			return err
		}
		fn(statement, deps)
	}
	return nil
}

func OneParamPreCheck(params ...string) (string, error) {
	if len(params) != 1 {
		return "", fmt.Errorf("tag takes no parameter")
	}

	return params[0], nil
}

func TwoParamsPreCheck(params ...string) (string, error) {
	if len(params) != 2 {
		return "", fmt.Errorf("tag takes only one parameter")
	}

	return params[1], nil
}
