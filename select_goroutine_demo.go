package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 *  @Description: 简单实例
 *
 *  @param
 *  @return
 */
func testSimpleDemo() {
	//  创建2个通道
	ch1 := make(chan string)
	ch2 := make(chan string)

	//  启动Goroutine向ch1写入数据
	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "from ch1"
	}()

	//  启动Goroutine向ch2写入数据
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "from ch2"
	}()

	// select 监听2个通道,那个先就绪执行那个
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	time.Sleep(time.Second * 2)
}

/**
 *  @Description: 模拟超时请求
 *
 *  @param
 *  @return
 */
func request(name string, delay time.Duration, resChan chan string) {
	time.Sleep(delay)
	resChan <- fmt.Sprintf("%s res", name)
}

func testMultipleGoroutineDemo() {
	reschan := make(chan string)
	go request("req1", time.Second*2, reschan)
	go request("req2", time.Second*4, reschan)
	go request("req3", time.Second*3, reschan)

	select {
	case res := <-reschan:
		fmt.Println(res)
	case <-time.After(time.Second * 2): // 超时操作
		fmt.Println("timeout")
	}

	// 关闭通道
	close(reschan)
}

func worker(taskId int, taskChan chan string, quitChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d start\n", taskId)
	for {
		select {
		case task := <-taskChan:
			fmt.Printf("worker:%d,process task:%s\n", taskId, task)
		case <-quitChan:
			fmt.Printf("worker:%d,quit\n", taskId)
			return
		default: // default 分支会在所有 case 都未就绪时立即执行，实现 “非阻塞” 的通道操作。
			fmt.Printf("worker:%d,无数据\n", taskId)
		}
	}
}

func testStopGoroutineDemo() {
	taskChan := make(chan string)
	quitChan := make(chan bool)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, taskChan, quitChan, &wg)
	}

	for i := 0; i < 5; i++ {
		taskChan <- fmt.Sprintf("msg:%d", i)
	}

	time.Sleep(time.Second * 2)
	fmt.Println("stop goroutine")
	close(quitChan)

	wg.Wait()
	fmt.Printf("all goroutine done")

}

func testSelectGoroutineDemo() {
	// testSimpleDemo()

	// testMultipleGoroutineDemo()

	testStopGoroutineDemo()
}
