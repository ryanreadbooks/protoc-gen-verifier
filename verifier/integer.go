package verifier

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

type makeForIntegerReq struct {
	bitSize  int
	unsigned bool
	symbol   string
	errInfo  string
}

func integerEqualJudgement(field string, value interface{}, param, errMsg string) (string, error) {
	return integerCmpJudgement(field, value, param, "==", errMsg)
}

func integerGreaterJudgement(field string, value interface{}, param, errMsg string) (string, error) {
	return integerCmpJudgement(field, value, param, ">", errMsg)
}

func integerGreaterEqualJudgement(field string, value interface{}, param, errMsg string) (string, error) {
	return integerCmpJudgement(field, value, param, ">=", errMsg)
}

func integerLessJudgement(field string, value interface{}, param, errMsg string) (string, error) {
	return integerCmpJudgement(field, value, param, "<", errMsg)
}

func integerLessEqualJudgement(field string, value interface{}, param, errMsg string) (string, error) {
	return integerCmpJudgement(field, value, param, "<=", errMsg)
}

func integerNotEqualJudgement(field string, value interface{}, param, errMsg string) (string, error) {
	return integerCmpJudgement(field, value, param, "!=", errMsg)
}

func integerCmpJudgement(fieldName string, fieldValue interface{}, param string, symbol, errInfo string) (string, error) {
	var judgement string
	var err error

	switch fieldValue.(type) {
	case int32:
		judgement, err = makeForInteger(fieldName, param, &makeForIntegerReq{
			bitSize:  32,
			unsigned: false,
			symbol:   symbol,
			errInfo:  errInfo,
		})
	case int64:
		judgement, err = makeForInteger(fieldName, param, &makeForIntegerReq{
			bitSize:  64,
			unsigned: false,
			symbol:   symbol,
			errInfo:  errInfo,
		})
	case uint32:
		judgement, err = makeForInteger(fieldName, param, &makeForIntegerReq{
			bitSize:  32,
			unsigned: true,
			symbol:   symbol,
			errInfo:  errInfo,
		})
	case uint64:
		judgement, err = makeForInteger(fieldName, param, &makeForIntegerReq{
			bitSize:  64,
			unsigned: true,
			symbol:   symbol,
			errInfo:  errInfo,
		})
	default:
		return "", ErrCheckType
	}

	return judgement, err
}

func makeForInteger(field, param string, req *makeForIntegerReq) (string, error) {
	want, err := strconv.ParseInt(param, 10, req.bitSize)
	if err != nil {
		return "", errors.WithMessage(err, "param must be integer")
	}

	var intType string
	const template = `if %s %s %s(%d) {
		return fmt.Errorf("%s")
	}`
	if req.bitSize == 32 {
		if req.unsigned {
			intType = "uint32"
		} else {
			intType = "int32"
		}
	} else {
		if req.unsigned {
			intType = "uint64"
		} else {
			intType = "int64"
		}
	}
	return fmt.Sprintf(template, field, req.symbol, intType, want, req.errInfo), nil
}
