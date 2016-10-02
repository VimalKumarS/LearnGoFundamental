package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func main1() {

}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}
func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func pipeFileter() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

func (this *Promise) Then(sucess func(interface{}) error, failure func(error)) *Promise {
	result := new(Promise)
	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		select {
		case obj := <-this.successChannel:
			newErr := sucess(obj)
			if newErr == nil {
				result.successChannel <- obj

			} else {
				result.failureChannel <- newErr
			}
		case err := <-this.failureChannel:
			failure(err)
			result.failureChannel <- err
		}
	}()
	return result
}

func SavePOpromise(po *PurchaseOrder, shouldFail bool) *Promise {
	result := new(Promise)
	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		if shouldFail {
			result.failureChannel <- errors.New("failsed to save")
		} else {
			po.Number = 1234
			result.successChannel <- po
		}
	}()

	return result
}
func PromiseEx() {
	po := new(PurchaseOrder)
	po.Valuse = 42.27
	SavePOpromise(po, false).Then(func(obj interface{}) error {
		po := obj.(*PurchaseOrder)
		fmt.Printf("%d", po.Number)
		return nil
	}, func(err error) {

	})
	fmt.Scanln()
}

type PurchaseOrder struct {
	Number int
	Valuse float64
}

func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.Number = 1234
	callback <- po
}

func CallBack() {
	po := new(PurchaseOrder)
	po.Valuse = 42.27
	ch := make(chan *PurchaseOrder)
	go SavePO(po, ch)
	newPo := <-ch
	fmt.Printf("%v", newPo)
}

type Button struct {
	eventListener map[string][]chan string
}

func MakeButton() *Button {
	result := new(Button)
	result.eventListener = make(map[string][]chan string)
	return result
}
func (this *Button) AddEventListener(event string, responseChannel chan string) {
	if _, present := this.eventListener[event]; present {
		this.eventListener[event] = append(this.eventListener[event], responseChannel)
	} else {
		this.eventListener[event] = []chan string{responseChannel}
	}
}
func (this *Button) RemoveEventListern(event string, listernChannel chan string) {
	if _, present := this.eventListener[event]; present {
		for idx, _ := range this.eventListener[event] {
			if this.eventListener[event][idx] == listernChannel {
				this.eventListener[event] = append(this.eventListener[event][:idx], this.eventListener[event][idx+1:]...)
				break
			}
		}
		//this.eventListener[event]=append(this.eventListener[event],responseChannel)
	}
}

func (this *Button) triggerEvent(event string, response string) {
	if _, present := this.eventListener[event]; present {
		for _, handler := range this.eventListener[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}
func EventBasedCon() {
	btn := MakeButton()
	handlerOne := make(chan string)
	handlerTwo := make(chan string)
	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)
	go func() {
		for {
			msg := <-handlerOne
			fmt.Println(msg)
		}
	}()
	go func() {
		for {
			msg := <-handlerTwo
			fmt.Println(msg)
		}
	}()
	btn.triggerEvent("click", "Button Clicked")
	btn.RemoveEventListern("click", handlerTwo)
	btn.triggerEvent("click", "Button Clicked")

	fmt.Scanln()

}

func createLog() {
	runtime.GOMAXPROCS(4)
	f, _ := os.Create("./log.txt")
	f.Close()
	logch := make(chan string, 50)
	go func() {
		for {
			msg, ok := <-logch
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + msg)
			} else {
				break
			}

		}
	}()
	mutex := make(chan bool, 1)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d + %d =%d \n", i, j, i+j)
				logch <- msg
				fmt.Print(msg)
				<-mutex

			}()
		}
	}
	fmt.Scanln()
}

func ChannelSync() {
	runtime.GOMAXPROCS(4)

	mutex := make(chan bool, 1)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			go func() {
				fmt.Printf("%d + %d =%d \n", i, j, i+j)
				<-mutex

			}()
		}
	}
	fmt.Scanln()
}

func mutex1() {
	mutex := new(sync.Mutex)
	runtime.GOMAXPROCS(4)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex.Lock()
			go func() {
				fmt.Printf("%d + %d =%d \n", i, j, i+j)
				mutex.Unlock()
			}()
		}
	}
	fmt.Scanln()
}
