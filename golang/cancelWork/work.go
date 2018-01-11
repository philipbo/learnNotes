package main

import (
	"context"
	"log"
	"time"
)

func work(ctx context.Context, w int) {
	time.Sleep(time.Duration(w) * time.Second)
	select {
	case <-ctx.Done():
		log.Printf("work %d 取消干活", w)
		return
	default:
		//ignore
	}

	log.Printf("work %d 正常干完", w)

}
