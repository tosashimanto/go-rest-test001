

-------------------------

$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
{"v":"HELLO, WORLD","err":null}

$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
{"v":12}

