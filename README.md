
**Traveling Salesperson Problem**

This is solved using brute force. Next step is to use goroutines to parallelize
the computations. After that I will employ some algorithms.

The solution is implemented in Go as a web service. The frontend uses React.

**Installation**

  * Install Node.js (I recommend using nvm for this)
  * Install Go
  * Clone this repo
  * `$ cd tsp`
  * `$ npm install` (this takes some time)
  * `$ cd golang`
  * `$ go build tsp_service.go`
  * `$ cd ..`
  * `$ ./golang/tsp_service` (The Go webservice)
  * In other terminal - `$ npm run start` (this will open the app in your default browser)

