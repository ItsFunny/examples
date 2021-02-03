package channel

import (
	"bidchain/http_framework/protocol"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type FilterPipelineState int

const (
	RUNNING FilterPipelineState = iota
	CLOSED
)

type FilterPipeline struct {
	head    *FilterContext
	tail    *FilterContext
	state   FilterPipelineState // 0 running 1 close
	channel IHttpChannel
}

func NewFilterPipeline(ch IHttpChannel) *FilterPipeline {
	fp := &FilterPipeline{
		channel: ch,
	}
	return fp
}

func (fp *FilterPipeline) GetChannel() IHttpChannel {
	return fp.channel
}

func (fp *FilterPipeline) Init(head, tail *FilterContext) {
	fp.head = head
	fp.tail = tail
	head.next = tail
	tail.prev = head
	head.SetAdded()
	tail.SetAdded()
}


func (fp *FilterPipeline) isClose() bool {
	return fp.state == CLOSED
}

func (fp *FilterPipeline) checkDuplicateName(name string) {
	if fp.GetContext(name) != nil {
		msg := fmt.Sprintf("duplicate handler name: " + name)
		panic(errors.New(msg))
	}
}

func (fp *FilterPipeline) GetContext(name string) *FilterContext {
	ctx := fp.head.next
	for ctx != fp.tail {
		if ctx.Name == name {
			return ctx
		}
		ctx = ctx.next
	}
	return nil
}

func (fp *FilterPipeline) filterName(name string, filter IFilter) string {
	if name == "" {
		err := errors.New("name is empty")
		panic(err)
	}
	fp.checkDuplicateName(name)
	return name
}


func (fp *FilterPipeline) handleAddContext(ctx *FilterContext) {

	//err := ctx.Filter.FilterAdded(ctx)
	//// 正常情况
	//if err == nil {
	//	ctx.SetAdded()
	//	return
	//}
	//removed := false
	//err = remove(ctx)
	//if err == nil {
	//	ctx.Filter.FilterRemoved(ctx)
	//	ctx.SetRemoved()
	//	removed = true
	//}
	//if removed == true {
	//
	//} else {
	//
	//}

	defer func() {
		if err := recover(); err != nil {
			removed := false
			func() {
				defer func() {
					recover() // 捕获错误
				}()
				remove(ctx)
				func() {
					defer func() {
						// finally
						if err := recover(); err != nil {
							log.Println(err)
						}
						ctx.SetRemoved()
					}()
					ctx.Filter.FilterRemoved(ctx)
				}()
				removed = true
			}()
			if removed {
				err := errors.New(fmt.Sprintf("filter.filterAdded()出错"))
				fp.FireExceptionCaught(&httpFrameworkException{error: err})
			} else {
				err := errors.New(fmt.Sprintf("filter.filterAdded() has thrown an exception; also failed to remove"))
				fp.FireExceptionCaught(&httpFrameworkException{error: err})
			}
		}
	}()

	ctx.Filter.FilterAdded(ctx)
	ctx.SetAdded()
}

func remove(ctx *FilterContext) error {
	if ctx == nil {
		return errors.New("ctx is nil")
	}
	prev := ctx.prev
	next := ctx.next
	if prev == nil {
		msg := fmt.Sprintf("prev is nil, ctx=[%v]", ctx)
		return errors.New(msg)
	}
	if next == nil {
		msg := fmt.Sprintf("next is nil, ctx=[%v]", ctx)
		return errors.New(msg)
	}
	prev.next = next
	next.prev = prev
	return nil
}

func addAfter(ctx, newCtx *FilterContext) {
	newCtx.next = ctx.next
	newCtx.prev = ctx
	ctx.next.prev = newCtx
	ctx.next = newCtx
}

func addBefore(ctx, newCtx *FilterContext) {
	newCtx.prev = ctx.prev
	newCtx.next = ctx
	ctx.prev.next = newCtx
	ctx.prev = newCtx
}

func (fp *FilterPipeline) newContext(name string, filter IFilter) *FilterContext {
	fc := NewFilterContext(fp, filter, name)
	return fc
}

// AddFirst、AddLast可能抛异常
func (fp *FilterPipeline) AddFirst(name string, filter IFilter) {
	name = fp.filterName(name, filter)
	newCtx := fp.newContext(name, filter)
	addAfter(fp.head, newCtx)
	fp.handleAddContext(newCtx)
}

func (fp *FilterPipeline) AddLast(name string, filter IFilter) {
	name = fp.filterName(name, filter)
	newCtx := fp.newContext(name, filter)
	addBefore(fp.tail, newCtx)
	fp.handleAddContext(newCtx)
}

func (fp *FilterPipeline) WriteCommand(cmd protocol.ICommand, ch chan<-FabricResult) {
	invokeWriteCommand(fp.tail, cmd, ch)
}

func (fp *FilterPipeline) FireReadCommand(cmd protocol.ICommand, ch chan<-FabricResult) {
	invokeReadCommand(fp.head, cmd, ch)
}

func (fp *FilterPipeline) FireExceptionCaught(cause error) {
	invokeExceptionCaught(fp.head, cause)
}
