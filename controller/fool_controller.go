package controller

import (
	"context"
	"fmt"
	"log"
	"time"
	"web-server/framework"
)


func FooControllerHandler(ctx *framework.Context) error {

	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
	defer cancel()
	panicChan := make(chan interface{}, 1)
	finish := make(chan struct{}, 1)

	go func() {
		defer func() {
			if p := recover() ; p!=nil {
				panicChan <- p
			}
		}()
		time.Sleep(10 * time.Second)
		ctx.Json(200, "ok")

		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		log.Println(p)
		ctx.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.Json(500, "time out")
		ctx.SetHasTimeout()
	}
	return nil
}