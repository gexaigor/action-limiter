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
##### Wait()
This example demonstrates limiting the output rate to 2 times per second.
```go
func main() {
	al := limiter.New(2, time.Second)
	begin := time.Now()
	for i := 0; i < 10; i++ {
		al.Wait()
		fmt.Printf("%d started at %s\n", i, time.Now().Sub(begin))
		//...
	}
}
```

Output:
```sh
0 started at 0s
1 started at 997.8µs
2 started at 1.0021138s
3 started at 1.0023405s
4 started at 2.0026946s
5 started at 2.0026946s
6 started at 3.0173932s
7 started at 3.0183951s
8 started at 4.0252897s
9 started at 4.0252897s
```

##### Try()
This example demonstrates the ability to perform other actions while the limit is over.
```go
func main() {
	al := limiter.New(5, time.Second)
	begin := time.Now()
	for i := 0; i < 10; i++ {
		if al.Try() {
			fmt.Printf("%d started at %s\n", i, time.Now().Sub(begin))
		} else {
			fmt.Println("wait a bit")
		}
	}
}
```

Output
```sh
0 started at 0s
1 started at 0s
2 started at 1.9987ms
3 started at 2.998ms
4 started at 2.998ms
wait a bit
wait a bit
wait a bit
wait a bit
wait a bit
```

### Author
 - [gexaigor](https://github.com/gexaigor "gexaigor")