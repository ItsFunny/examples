package chaincode

import (
	"bidchain/protocol/store/store_base"
	"bidchain/protocol/transport/base"
)

func Transport2StoreIntegerWrapper(src *base.IntegerWrapper) *store_base.StoreIntegerWrapper {
	if src == nil {
		return nil
	}
	var dest store_base.StoreIntegerWrapper
	dest.Value = src.Value
	return &dest
}

func Store2TransportIntegerWrapper(src *store_base.StoreIntegerWrapper) *base.IntegerWrapper {
	if src == nil {
		return nil
	}
	var dest base.IntegerWrapper
	dest.Value = src.Value
	return &dest
}

func Transport2StoreLongWrapper(src *base.LongWrapper) *store_base.StoreLongWrapper {
	if src == nil {
		return nil
	}
	var dest store_base.StoreLongWrapper
	dest.Value = src.Value
	return &dest
}

func Store2TransportLongWrapper(src *store_base.StoreLongWrapper) *base.LongWrapper {
	if src == nil {
		return nil
	}
	var dest base.LongWrapper
	dest.Value = src.Value
	return &dest
}

func Transport2StoreStringWrapper(src *base.StringWrapper) *store_base.StoreStringWrapper {
	if src == nil {
		return nil
	}
	var dest store_base.StoreStringWrapper
	dest.Value = src.Value
	return &dest
}

func Store2TransportStringWrapper(src *store_base.StoreStringWrapper) *base.StringWrapper {
	if src == nil {
		return nil
	}
	var dest base.StringWrapper
	dest.Value = src.Value
	return &dest
}
