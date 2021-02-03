package types

type Head struct {
	ErrCode int64
	ErrDesc string
}

func NewHead(errCode int64, errDesc string) *Head {
	h := &Head{
		ErrCode: errCode,
		ErrDesc: errDesc,
	}
	return h
}

func (h *Head) SetErrCode(errCode int64) {
	h.ErrCode = errCode
}

func (h *Head) GetErrCode() int64 {
	return h.ErrCode
}

func (h *Head) SetErrDesc(errDesc string) {
	h.ErrDesc = errDesc
}

func (h *Head) GetErrDesc() string {
	return h.ErrDesc
}
