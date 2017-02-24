
**Traveling Salesperson Problem**

This is solved using brute force. Next step is to use goroutines to parallelize
the computations. After that I will employ some algorithms.

The solution is implemented in Go as a web service. The frontend uses React.

**Installation**

  * Install Node.js (I recommend using nvm for this)
  * Install Go
  * Install the tsp package
    `$ go get github.com/arcofdescent/golang/tsp`
  * Clone this repo
  * `$ cd tsp`
  * `$ npm install` (this takes some time)
  * `$ go build golang/tsp_service.go`
  * `$ ./golang/tsp_service` (The Go webservice)
  * In another terminal - `$ npm run start` (this will open the app in your default browser)

