package core

import (
	"context"
	"fmt"
	"github.com/luopengift/autossh/modules"
	"github.com/luopengift/golibs/channel"
	"github.com/luopengift/golibs/logger"
	"github.com/luopengift/golibs/ssh"
	"sync"
	"time"
)

var (
	startTime = time.Now()
)

type Batch struct {
	fail    int
	succ    int
	timeout int
	mutex   *sync.Mutex
	workers *channel.Channel
	results *channel.Channel
	quit	chan bool
}

func NewBatch(fork, timeout int) *Batch {
	logger.SetTimeFormat("")
	logger.SetLevel(logger.NULL)
	batch := new(Batch)
	batch.timeout = timeout
	batch.mutex = new(sync.Mutex)
	batch.workers = channel.NewChannel(fork)
	batch.results = channel.NewChannel(fork)
	batch.quit	= make(chan bool)
	return batch
}

type Result struct {
	Addr string
	Out  []byte
	Err  error
}

func (b *Batch) Execute(servers []*ssh.Endpoint, mod, args string) error {
	module, ok := modules.Modules[mod]
	if !ok {
		return fmt.Errorf("module missing: %v", mod)
	}
	if err := module.Init(args); err != nil {
		return fmt.Errorf("module init error: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(b.timeout)*time.Second)
	defer cancel()
	go func() {
		for {
			res, _ := b.results.Get()
			result := res.(Result)
			b.displayResult(b.results.Total(), result)
			if b.results.Total() == int64(len(servers)) {
				b.quit <- true
			}
		}
	}()

	for _, endpoint := range servers {
		b.workers.Add()
		go func(ctx context.Context,endpoint *ssh.Endpoint) {
			result := Result{Addr: endpoint.Ip}
			result.Out, result.Err = module.Run(ctx, endpoint)
			b.results.Put(result)
			b.workers.Done()
		}(ctx, endpoint)
	}
	format := "[主机数量]:%d, [成功]:%d, [失败]:%d, [超时]:%d|[执行时间]:%s"
	select {
	case <-ctx.Done():
		logger.Warn(format, len(servers), b.succ, b.fail, len(servers) - b.succ - b.fail, time.Since(startTime).String())
	case <- b.quit:
		logger.Info(format, len(servers), b.succ, b.fail, len(servers) - b.succ - b.fail, time.Since(startTime).String())

	}
	return nil
}

func (b *Batch) displayResult(no int64, result Result) {
	if result.Err != nil {
		b.fail += 1
		logger.Error("[%d] %s | %s =>\n%s", no, result.Addr, "FAIL", string(result.Out)+result.Err.Error())
	} else {
		b.succ += 1
		logger.Info("[%d] %s | %s =>\n%s", no, result.Addr, "SUCC", string(result.Out))
	}
}
