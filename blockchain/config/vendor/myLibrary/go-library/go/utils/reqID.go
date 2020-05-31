package utils

import (
	"time"
	"context"

	"strings"

	"github.com/satori/go.uuid"
)

var reqID struct{}
var reqStartTime struct{}


// SetReqID : 设置一个reqID
func SetReqID(ctx context.Context) context.Context {
	uid:= uuid.NewV4()
	str := strings.Replace(uid.String(), "-", "", -1)
	c := context.WithValue(ctx, reqID, str)
	return c
}

// GetReqID : 获取一个reqID
func GetReqID(ctx context.Context) string {
	id := ctx.Value(reqID)

	switch id.(type) {
	case string:
		return id.(string)
	default:
		return ""
	}
}

// SetReqTime : 设置当前时间(unix nano)到ctx
func SetReqTime(ctx context.Context) context.Context {
	c := context.WithValue(ctx, reqStartTime, time.Now().UnixNano())
	return c
}

// GetReqTime : 获取设置的时间,如果没有返回0
func GetReqTime(ctx context.Context) int64 {
	t := ctx.Value(reqStartTime)

	switch t.(type) {
	case int64:
		return t.(int64)
	default:
		return 0
	}

	return 0
}
