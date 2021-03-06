package command

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/luopengift/autossh/modules"
	"github.com/luopengift/autossh/pkg/endpoint"
	"github.com/luopengift/golibs/channel"
	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
)

// Batch Batch
type Batch struct {
	fail    int
	succ    int
	timeout int
	mutex   *sync.Mutex
	workers *channel.Channel
	results *channel.Channel
	quit    chan bool
}

// NewBatch NewBatch
func NewBatch(fork, timeout int) *Batch {
	batch := new(Batch)
	batch.timeout = timeout
	batch.mutex = new(sync.Mutex)
	batch.workers = channel.NewChannel(fork)
	batch.results = channel.NewChannel(fork)
	batch.quit = make(chan bool)
	return batch
}

// Result Result
type Result struct {
	Addr string
	Out  []byte
	Err  error
}

// Execute Execute
func (b *Batch) Execute(endpoints endpoint.Endpoints, mod, args string) error {
	startTime := time.Now()
	if len(endpoints) == 0 {
		return fmt.Errorf(`主机数量为0, 请使用"-i/-files"指定`)
	}
	module, ok := modules.Modules[mod]
	if !ok {
		return fmt.Errorf("module missing: %v", mod)
	}
	if err := module.Parse(args); err != nil {
		return fmt.Errorf("module init error: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(b.timeout)*time.Second)
	defer cancel()
	go func() {
		for {
			res, _ := b.results.Get()
			result := res.(Result)
			b.displayResult(b.results.Total(), result)
			if b.results.Total() == int64(len(endpoints)) {
				b.quit <- true
			}
		}
	}()

	for _, endpoint := range endpoints {
		b.workers.Add()
		go func(ctx context.Context, endpoint *ssh.Endpoint) {
			result := Result{Addr: endpoint.IP}
			result.Out, result.Err = module.Run(ctx, endpoint)
			b.results.Put(result)
			b.workers.Done()
		}(ctx, endpoint)
	}
	format := "[主机数量]:%d, [成功]:%d, [失败]:%d, [超时]:%d|[执行时间]:%s"
	select {
	case <-ctx.Done():
		log.Warn(format, len(endpoints), b.succ, b.fail, len(endpoints)-b.succ-b.fail, time.Since(startTime).String())
	case <-b.quit:
		log.Info(format, len(endpoints), b.succ, b.fail, len(endpoints)-b.succ-b.fail, time.Since(startTime).String())

	}
	return nil
}

func (b *Batch) displayResult(no int64, result Result) {
	if result.Err != nil {
		b.fail++
		log.Error("[%d] %s | %s =>\n%s", no, result.Addr, "FAIL", string(result.Out)+result.Err.Error())
	} else {
		b.succ++
		log.Info("[%d] %s | %s =>\n%s", no, result.Addr, "SUCC", string(result.Out))
	}
}
