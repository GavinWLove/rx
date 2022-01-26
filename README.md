rx
====
Inspired from mux,thanks mux!

## Install
```sh
go get github.com/GavinWLove/rx
```
## Examples
```go
	r := rx.NewRouter().NewRoute()
	r = r.Path("/foo/{id}/aa")
	r = r.Path( "/proxy/{id}/*")
	r = r.Path("/parent/{id}/child/{another_id}")
	r = r.Path("/{foo:[0-9][0-9][0-9]}/{bar:[0-9][0-9][0-9]}/{id:[0-9]+}")
	matchInfo := &rx.RouteMatch{}
	var match bool
	match = r.Match("/parent/123/child/456", matchInfo)
	match = r.Match("/123/456/789", matchInfo)
	match = r.Match("/foo/-/aa", matchInfo)
	match = r.Match("/proxy/myid/1", matchInfo)
```
