package gsp

import (
	"fmt"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	type user struct {
		Name string
		Age  int
	}
	go func() {
		for {

			time.Sleep(10 * time.Millisecond)
			GetEvent("hello").Publish(&user{Name: "tom", Age: 18})
		}
	}()
	go func() {
		for {
			GetEvent("hello").Publish(time.Now().UnixNano() / 1e6)
			time.Sleep(15 * time.Millisecond)
		}
	}()

	//time.Sleep(1 * time.Second)

	go func() {
		GetEvent("hello").Subscribe(helloworld)
	}()

	go func() {
		GetEvent("hello").Subscribe(helloworld1)
	}()
	time.Sleep(10 * time.Second)
	GetEvent("hello").Publish("hello world")
}

func helloworld(i interface{}) {
	fmt.Printf("receive : %+v\n", i)
}
func helloworld1(i interface{}) {
	fmt.Printf("receive1 : %+v\n", i)
}
