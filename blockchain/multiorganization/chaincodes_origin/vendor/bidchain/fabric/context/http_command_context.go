package context

type HttpCommandContext struct {
	IBidchainContext // 上下文
	//Args []byte
	ChannelName string
	ChaincodeName string
	//FuncName string  // 最后通过反射调用的方法需要首字母大写
	Contract interface{} // 合约
}

func NewHttpCommandContext(ctx IBidchainContext, cc interface{}, chaincodeName, funcName string) *HttpCommandContext {
	hcc := &HttpCommandContext{}
	hcc.IBidchainContext = ctx
	hcc.ChannelName = ctx.GetChannelID()
	hcc.ChaincodeName = chaincodeName
	//hcc.FuncName = funcName
	hcc.Contract = cc
	return hcc
}


func (hcc *HttpCommandContext) GetContract() interface{} {
	return hcc.Contract
}
