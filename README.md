# tcp-rst-srv

I created this "shitty" code just for the test purpouses
e.g. I tested how reverse proxy server behave when it receives
TCP reset from a backend.

There is required one argument **-port** to run the server:

```sh
$ go run main.go -port 9004
2023/02/28 16:36:30 Starting TPC server on '127.0.0.1:9004'
2023/02/28 16:36:53 TCP connection received from: 127.0.0.1:51028
```

Once the server is running you can test RESET behaviour:

```sh
$ curl -v http://127.0.0.1:9004
*   Trying 127.0.0.1:9004...
* Connected to 127.0.0.1 (127.0.0.1) port 9004 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:9004
> User-Agent: curl/7.85.0
> Accept: */*
>
* Recv failure: Connection reset by peer
* Closing connection 0
curl: (56) Recv failure: Connection reset by peer
```
