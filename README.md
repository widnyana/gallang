Gallang
-------
http service that provide oEmbed metadata, using [gin](https://github.com/gin-gonic/gin). 

need demo? [click here](http://gallang-xyonite.rhcloud.com/embed?url=https://www.youtube.com/watch?v=KbPWi1gshzI)

this app rely on [golang oembed's](https://github.com/widnyana/oembed) shoulder.


Instalation
-----------
since this is a "functionaly work", after defining your `$GOPATH`, execute this command:

```go get github.com/widnyana/gallang```

this app require you to define `GALLANG_IP` and `GALLANG_PORT` as your env variable, 
`gallang` will be run and listening on those env variable.
 
```
export GALLANG_IP=127.0.0.1
export GALLANG_PORT=8080
```

then copy the `main.go.example` to `main.go`, and you can run your app with go `run main.go`

Endpoint
--------

- `/` as index
- `/embed?url=<oembed provider url>` as oembed info parser.


License
--------
[MIT License](http://widnyana.mit-license.org)