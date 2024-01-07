package verifier

import (
	stdError "errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

// EqVerifier 生成eq表达式的判断代码
type EqVerifier struct {
	FmtProvider
	name string
}

// Generate judgement statement for tag like 'eq=100'
func (v *EqVerifier) Generate(field string, value interface{}, params ...string) (judgement string, err error) {
	param, err := TwoParamsPreCheck(params...)
	if err != nil {
		return "", errors.WithMessage(err, "param of eq invalid")
	}

	var number bool = false

	switch value.(type) {
	case int32, int64, uint32, uint64:
		number = true
	case bool:
		judgement, err = v.booleanNotEqualJudgement(field, param)
	case string:
		judgement, err = v.stringNotEqualJudgement(field, param)
	default:
		return "", ErrEqCheckType
	}

	if number {
		judgement, err = integerNotEqualJudgement(field, value, param, constructErrMsg(field, v.name))
		if stdError.Is(err, ErrCheckType) {
			return judgement, ErrEqCheckType
		}
	}

	return judgement, err
}

func (v *EqVerifier) booleanNotEqualJudgement(field, param string) (string, error) {
	want, err := strconv.ParseBool(param)
	if err != nil {
		return "", errors.WithMessage(err, "param of eq for bool must be 'true' or 'false'")
	}

	const template = `if %s != %v {
		return fmt.Errorf("%s")
	}`
	return fmt.Sprintf(template, field, want, constructErrMsg(field, v.name)), nil
}

func (v *EqVerifier) stringNotEqualJudgement(field, param string) (string, error) {
	want := param
	const template = `if %s != "%s" {
		return fmt.Errorf("%s")
	}`
	return fmt.Sprintf(template, field, want, constructErrMsg(field, v.name)), nil
}

// GtVerifier 生成gt表达式的判断代码
type GtVerifier struct {
	FmtProvider
	name string
}

func (v *GtVerifier) Generate(field string, value interface{}, params ...string) (judgement string, err error) {
	param, err := TwoParamsPreCheck(params...)
	if err != nil {
		return "", errors.WithMessage(err, "param of gt invalid")
	}

	judgement, err = integerLessEqualJudgement(field, value, param, constructErrMsg(field, v.name))
	if stdError.Is(err, ErrCheckType) {
		return "", ErrGtCheckType
	}

	return judgement, err
}

// GteVerifier 生成gte表达式的判断代码
type GteVerifier struct {
	FmtProvider
	name string
}

func (v *GteVerifier) Generate(field string, value interface{}, params ...string) (judgement string, err error) {
	param, err := TwoParamsPreCheck(params...)
	if err != nil {
		return "", errors.WithMessage(err, "param of gte invalid")
	}

	judgement, err = integerLessJudgement(field, value, param, constructErrMsg(field, v.name))
	if stdError.Is(err, ErrCheckType) {
		return "", ErrGteCheckType
	}

	return judgement, err
}

// LtVerifier 生成lt表达式的判断代码
type LtVerifier struct {
	FmtProvider
	name string
}

func (v *LtVerifier) Generate(field string, value interface{}, params ...string) (judgement string, err error) {
	param, err := TwoParamsPreCheck(params...)
	if err != nil {
		return "", errors.WithMessage(err, "param of lt invalid")
	}

	judgement, err = integerGreaterEqualJudgement(field, value, param, constructErrMsg(field, v.name))
	if stdError.Is(err, ErrCheckType) {
		return "", ErrLtCheckType
	}

	return judgement, err
}

// LteVerifier 生成lte表达式的判断代码
type LteVerifier struct {
	FmtProvider
	name string
}

func (v *LteVerifier) Generate(field string, value interface{}, params ...string) (judgement string, err error) {
	param, err := TwoParamsPreCheck(params...)
	if err != nil {
		return "", errors.WithMessage(err, "param of lte invalid")
	}

	judgement, err = integerGreaterJudgement(field, value, param, constructErrMsg(field, v.name))
	if stdError.Is(err, ErrCheckType) {
		return "", ErrLteCheckType
	}

	return judgement, err
}

// NeVerifier 生成ne表达式的判断代码
type NeVerifier struct {
	FmtProvider
	name string
}

func (v *NeVerifier) Generate(field string, value interface{}, params ...string) (judgement string, err error) {
	param, err := TwoParamsPreCheck(params...)
	if err != nil {
		return "", errors.WithMessage(err, "param of ne invalid")
	}

	var number bool = false

	switch value.(type) {
	case int32, int64, uint32, uint64:
		number = true
	case string:
		judgement, err = v.stringEqualJudgement(field, param)
	default:
		// 可还需要处理repeated和map的情况
		kind := reflect.TypeOf(value).Kind()
		if kind == reflect.Map || kind == reflect.Slice {
			lv := LenVerifier{}
			return lv.lenEqualJudgement(field, param, constructErrMsg(field, v.name))
		}
		return "", ErrNeCheckType
	}

	if number {
		judgement, err = integerEqualJudgement(field, value, param, constructErrMsg(field, v.name))
		if stdError.Is(err, ErrCheckType) {
			return "", ErrNeCheckType
		}
	}

	return judgement, err
}

func (v *NeVerifier) stringEqualJudgement(field, param string) (string, error) {
	want := param
	const template = `if %s == "%s" {
		return fmt.Errorf("%s")
	}`
	return fmt.Sprintf(template, field, want, constructErrMsg(field, v.name)), nil
}
