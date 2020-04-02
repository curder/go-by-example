package channels

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
// 1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
// 2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
// 3. 主goroutine从resultChan取出结果并打印到终端输出
func TestCalculatorSum(t *testing.T) {
	var wg sync.WaitGroup
	jobChan := make(chan int64, 100)
	resultChan := make(chan int64, 100)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			rand.Seed(time.Now().Unix()) // 防止生成的随机数一模一样
			randomInt := rand.Int63()
			jobChan <- randomInt
			time.Sleep(time.Millisecond * 500) // 休眠 1 S
		}
	}()

	wg.Add(24)
	for i := 0; i < 24; i++ {
		go func() {
			defer wg.Done()
			result, ok := <-jobChan
			if !ok {
				return
			}
			var sum int64
			var n = result
			for n > 0 {
				sum += n % 10 // 取模
				n = n / 10
			}
			t.Log(result, sum)
			resultChan <- sum
		}()
	}

	for result := range resultChan {
		t.Log(result)
	}

	wg.Wait()
}

// 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
// 1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
// 2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
// 3. 主goroutine从resultChan取出结果并打印到终端输出
func TestCalculatorInt64Sum(t *testing.T) {
	type job struct {
		value int64
	}
	type result struct {
		job *job
		sum int64
	}

	var wg sync.WaitGroup
	var jobChan = make(chan *job)
	var resultChan = make(chan *result)

	wg.Add(1)
	generateRand := func(j chan<- *job) {
		defer wg.Done()
		for {
			jobChan <- &job{
				value: rand.Int63(),
			}
			time.Sleep(time.Millisecond * 1000)
		}
	}

	fetchRand := func(j <-chan *job, r chan<- *result) {
		defer wg.Done()
		res := <-j

		var sum int64
		var temp = res.value
		// 求和
		for temp > 0 {
			sum += temp % 10
			temp = temp / 10
		}
		r <- &result{
			job: &job{value: res.value},
			sum: sum,
		}
	}

	go generateRand(jobChan) // 启动一个 gorountine 向jobChan通道中传值

	wg.Add(24)
	for i := 0; i < 24; i++ { // 启动 24 个 gorountine 取jobChan中的值并计算传递给resultChan通道
		go fetchRand(jobChan, resultChan)
	}

	for res := range resultChan {
		t.Log(res.job.value, res.sum)
	}

	wg.Wait()
}
