/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-17 10:18 
# @File : transaction.go
# @Description : 
# @Attention : 
*/
package impl

import (
	"examples/blockchain/config/common/base/fabric"
)

type BaseTransactionImpl struct {
}

func (b *BaseTransactionImpl) GetFromWalletAddress() base.From {
	panic("implement me")
}

func (b *BaseTransactionImpl) GetToWalletAddress() base.To {
	panic("implement me")
}

func (b *BaseTransactionImpl) GetToken() base.Token {
	panic("implement me")
}

func (b *BaseTransactionImpl) SetFromWalletAddress(f base.From) {
	panic("implement me")
}

func (b *BaseTransactionImpl) SetToWalletAddress(t base.To) {
	panic("implement me")
}

func (b *BaseTransactionImpl) SetToken(t base.Token) {
	panic("implement me")
}
