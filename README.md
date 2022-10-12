# GoNet - Easy HTTP Libary for Go
- GET
- POST
- Only Syntax is Changed! Still in contact with the original net/http (Made for beginners)

## How to use
```go
1. Use the 'go get' command to download the package.
	Command: go get github.com/9dl/GoNet/Libary
2. Import in your Project
	Example: import GoNet "github.com/9dl/GoNet/Libary"
```

# Get Request Example
```go
Client := GoNet.NewClient(false)
Request, _ := GoNet.GetRequest("Url")
Resp, _ := GoNet.NewResponse(Client, Request)
Response, _ := GoNet.GetResponse(Resp)

fmt.Printf(Response)
```
# Post Requests Example
```go
Client := GoNet.NewClient(false)
Request, _ := GoNet.PostRequest("Url", "Content Type", Body)
Resp, _ := GoNet.NewResponse(Client, Request)
Response, _ := GoNet.GetResponse(Resp)

fmt.Printf(Response)
```
