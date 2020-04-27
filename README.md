### gosearch

A simple utility written in go that leveraged DuckDuckGo APIs to retrieve links to top web pages for a given keyword

**Usage**

gosearch -keyword cricket -pages 5

**Sample Output**

1. https://duckduckgo.com/Cricket
2. https://duckduckgo.com/Cricket_(insect)
3. https://duckduckgo.com/Cricket_Wireless

**Setup**
1. Clone the repository
2. Run *Make run*
3. You will find all the binaries in */bin* directory

**Test Cases**
1. Internal pkg *scrapper* has a unit test case
2. The test case runs an HTTP Server and mock the response
3. Sample response for test case is stored in */pkg/scrapper/testdata*
4. You can run the test by executing *go test ./...* or *make test* 

 

