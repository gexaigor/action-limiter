# Action limiter
[![Build Status](https://travis-ci.org/gexaigor/action-limiter.svg?branch=main)](https://travis-ci.org/gexaigor/action-limiter) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/gexaigor/action-limiter)](https://goreportcard.com/report/github.com/gexaigor/action-limiter)

The Action limiter is very simple package allows you to limit the number of actions at a time. Wait method blocks the thread if more limit actions have already been performed.

### Features
 - small and simple API
 - thread safe
 - low memory usage

### Install
```sh
go get github.com/gexaigor/action-limiter
```

### Usage
In this example, the action limiter controls the launch of no more than 5 goroutines every 2 seconds.
```go
func main() {
	var wg sync.WaitGroup
	al := limiter.New(5, time.Second*2)
	for i := 0; i < 20; i++ {
		al.Wait()
		wg.Add(1)
		go do(i, &wg)
	}
	wg.Wait()
}

func do(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}
```
