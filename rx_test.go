package rx

import (
	"testing"
)

func TestRouter_Match(t *testing.T) {
	r := NewRouter().NewRoute()
	r = r.Path("/parent/{id}/child/{another_id}")
	r = r.Path("/{foo:[0-9][0-9][0-9]}/{bar:[0-9][0-9][0-9]}/{baz:[0-9][0-9][0-9]}")
	matchInfo := &RouteMatch{}
	match := r.Match("/123/456/789", matchInfo)
	//match = r.Match("/123/456/789", matchInfo)
	print(match)
}
