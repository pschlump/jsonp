# jsonp A small library for allowing JSONp requests

Go (golang) routines for supporting JSONp.

The primary interface is:

``` golang
	func JsonP(s string, res http.ResponseWriter, req *http.Request) string 
```

This allows converting a return string in JSON to a JSONp callback.   If the parameter on the url "callback" is
supplied then the returned string is wrapped in the callback name.

An example handler using this is:

``` golang

	func handleVersion(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		io.WriteString(res, jsonp.JsonP(fmt.Sprintf(`{"status":"success", "version":"1.0.0"}`+"\n"), res, req))
	}

```

