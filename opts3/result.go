package opts3

type Result interface {
	Code() Code
	String() string
}

type resultImpl struct {
	code Code
	string string
}

func (r *resultImpl) Code() Code {
	return r.code
}

func (r *resultImpl) String() string {
	return r.string
}

func NewResult(code Code, str string) Result {
	code.PanicIfInvalid()
	return &resultImpl{
		code:   code,
		string: str,
	}
}
