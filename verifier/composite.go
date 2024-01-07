package verifier

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

type LenVerifier struct {
	FmtProvider
	name string
}

func (v *LenVerifier) Generate(field string, value interface{}, params ...string) (judgement string, err error) {
	param, err := TwoParamsPreCheck(params...)
	if err != nil {
		return "", errors.WithMessage(err, "param of len invalid")
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.String, reflect.Slice, reflect.Map:
		judgement, err = v.lenNotEqualJudgement(field, param, constructErrMsg(field, v.name))
	default:
		return "", ErrLenCheckType
	}

	return judgement, err
}

func (v *LenVerifier) lenNotEqualJudgement(field, param, errMsg string) (string, error) {
	return v.lenJudgement(field, param, "!=", errMsg)
}

func (v *LenVerifier) lenEqualJudgement(field, param, errMsg string) (string, error) {
	return v.lenJudgement(field, param, "==", errMsg)
}

func (v *LenVerifier) lenGreaterJudgement(field, param, errMsg string) (string, error) {
	return v.lenJudgement(field, param, ">", errMsg)
}

func (v *LenVerifier) lenGreaterEqualJudgement(field, param, errMsg string) (string, error) {
	return v.lenJudgement(field, param, ">=", errMsg)
}

func (v *LenVerifier) lenLessJudgement(field, param, errMsg string) (string, error) {
	return v.lenJudgement(field, param, "<", errMsg)
}

func (v *LenVerifier) lenLessEqualJudgement(field, param, errMsg string) (string, error) {
	return v.lenJudgement(field, param, "<=", errMsg)
}

func (v *LenVerifier) lenJudgement(field, param string, symbol, info string) (string, error) {
	wantLen, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		return "", errors.WithMessage(err, "expected len must be integer")
	}

	const template = `if len(%s) %s %d {
		return fmt.Errorf("%s")
	}`

	return fmt.Sprintf(template, field, symbol, wantLen, info), nil
}
