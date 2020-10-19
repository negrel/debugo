// +build !assert

package assert

import (
	"net/http"
	"net/url"
)

func // HTTPSuccess asserts that a specified handler returns a success status code.
//
//  assert.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
//
// Returns whether the assertion was successful (true) or not (false).
HTTPSuccess(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) {
}

func // HTTPRedirect asserts that a specified handler returns a redirect status code.
//
//  assert.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
HTTPRedirect(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) {
}

func // HTTPError asserts that a specified handler returns an error status code.
//
//  assert.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
HTTPError(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) {
}

func // HTTPStatusCode asserts that a specified handler returns a specified status code.
//
//  assert.HTTPStatusCode(t, myHandler, "GET", "/notImplemented", nil, 501)
//
// Returns whether the assertion was successful (true) or not (false).
HTTPStatusCode(handler http.HandlerFunc, method, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) {
}

func // HTTPBody is a helper that returns HTTP body of the response. It returns
// empty string if building a new request fails.
HTTPBody(handler http.HandlerFunc, method, url string, values url.Values) {
}

func // HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
//  assert.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
HTTPBodyContains(handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
}

func // HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
//  assert.HTTPBodyNotContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
HTTPBodyNotContains(handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
}