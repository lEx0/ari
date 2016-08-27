package native

import (
	"fmt"

	"github.com/CyCoreSystems/ari"
)

type nativeApplication struct {
	conn *Conn
}

// Data returns the details of a given ARI application
// Equivalent to GET /applications/{applicationName}
func (a *nativeApplication) Data(name string) (d ari.ApplicationData, err error) {
	err = Get(a.conn, "/applications/"+name, &d)
	return
}

// Subscribe subscribes the given application to an event source
// Equivalent to POST /applications/{applicationName}/subscription
func (a *nativeApplication) Subscribe(name string, eventSource string) (err error) {
	var m ari.ApplicationData

	type request struct {
		EventSource string `json:"eventSource"`
	}

	req := request{EventSource: eventSource}
	err = Post(a.conn, "/applications/"+name+"/subscription", &m, &req)
	return err
}

// Unsubscribe unsubscribes (removes a subscription to) a given
// ARI application from the provided event source
// Equivalent to DELETE /applications/{applicationName}/subscription
func (a *nativeApplication) Unsubscribe(name string, eventSource string) (err error) {
	var m ari.ApplicationData

	// TODO: handle Error Responses individually

	// Make the request
	err = Delete(a.conn, "/applications/"+name+"/subscription", &m, fmt.Sprintf("eventSource=%s", eventSource))
	return
}
