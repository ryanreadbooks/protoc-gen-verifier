package verifier

import "fmt"

var (
	ErrCheckType    = fmt.Errorf("type mismatch")
	ErrEqCheckType  = fmt.Errorf("eq only works for integer, boolean and string")
	ErrGtCheckType  = fmt.Errorf("gt only works for integer")
	ErrGteCheckType = fmt.Errorf("gte only works for integer")
	ErrLtCheckType  = fmt.Errorf("lt only works for integer")
	ErrLteCheckType = fmt.Errorf("lte only works for integer")
	ErrNeCheckType  = fmt.Errorf("ne only works for integer, string, bytes, map and repeated type")
	ErrLenCheckType = fmt.Errorf("len only works for string, bytes, map and repeated type")

	ErrAlphaCheckType      = fmt.Errorf("alpha only works for string")
	ErrNumberCheckType     = fmt.Errorf("number only works for string")
	ErrContainsCheckType   = fmt.Errorf("contains only works for string")
	ErrStartsWithCheckType = fmt.Errorf("startswith only works for string")
	ErrEndsWithCheckTag    = fmt.Errorf("endswith only works for string")
)

type TypeChecker interface {
	CheckType(interface{}) error
}

type Skipper struct {
}
