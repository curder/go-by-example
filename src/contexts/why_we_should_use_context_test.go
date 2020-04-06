package contexts

import (
	"context"
	"sync"
	"testing"
	"time"
)

// 修改全局变量，通知go routine退出
func TestQuitGoRoutineByGlobalVariable(t *testing.T) {
	var wg = sync.WaitGroup{}
	var quit bool

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			t.Logf("go routine")
			time.Sleep(time.Millisecond * 500)
			if quit {
				break
			}
		}
	}()

	time.Sleep(time.Second * 2)
	quit = true // 修改全局变量的值

	wg.Wait()
}

// 通过channel给go routine发送信号
func TestQuitGoRoutineByChannel(t *testing.T) {
	var wg = sync.WaitGroup{}
	var quitCh = make(chan struct{}, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
	FOR_LOOP:
		for {
			t.Log("go routine")
			time.Sleep(time.Millisecond * 500)
			select {
			case <-quitCh:
				t.Log("break")
				break FOR_LOOP // 此处使用break LABEL的方式退出。使用break无法退出for循环
			default:
			}
		}
	}()

	time.Sleep(time.Second * 2)
	quitCh <- struct{}{} // 往channel中添加一个信号
	close(quitCh)        // 关闭channel

	wg.Wait()
}

// 通过context控制go routine
func TestQuitGoRoutineByContext(t *testing.T) {
	var wg = sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
	LOOP:
		for {
			t.Log("go routine")
			time.Sleep(time.Millisecond * 500)
			select {
			case <-ctx.Done(): // 等待上下文发送通知
				break LOOP
			default:
			}
		}
	}(ctx)

	time.Sleep(time.Second * 2)

	cancel() // 给ctx发送通知，通知子go routine结束

	wg.Wait()
}

// 通过 context 控制多级的go routine
func TestQuitGoRoutinesByContext(t *testing.T) {
	var wg = sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		// 在 go routine 中定义子 go routine
		go func(subCtx context.Context) {
		SUB_LOOP:
			for {
				t.Log("go sub routine")
				time.Sleep(time.Millisecond * 500)
				select {
				case <-subCtx.Done(): // 等待上下文发送通知
					break SUB_LOOP
				default:
				}
			}
		}(ctx)

	LOOP:
		for {
			t.Log("go routine")
			time.Sleep(time.Millisecond * 500)
			select {
			case <-ctx.Done(): // 等待上下文发送通知
				break LOOP
			default:
			}
		}
	}(ctx)

	time.Sleep(time.Second * 2)

	cancel() // 给ctx发送通知，通知子go routine和孙子go routine结束

	wg.Wait()

}
