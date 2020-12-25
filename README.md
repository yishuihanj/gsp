# gsp
golang 发布订阅 模块

订阅

```go
GetEvent("hello").Subscribe(helloworld) //event中带的是标识， interface{}

func helloworld(i interface{}) {
	fmt.Printf("receive : %+v\n", i)
}
```

发布

```go
type user struct{
    Name string
    Age int
}
GetEvent("hello").Publish(&user{Name:"tom",Age:18})  //publish中是 interface{}
```

