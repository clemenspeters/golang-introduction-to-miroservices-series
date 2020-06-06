# Building Microservices with Go: 3. RESTful services

[https://www.youtube.com/watch?v=eBeqtmrvVpg](https://www.youtube.com/watch?v=eBeqtmrvVpg)

Run

```bash
go run main.go
```

and from another terminal run

```bash
curl localhost:9090 | jq
curl localhost:9090 -XDELETE -v | jq
```

## Links

Source: [https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_3](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_3)

RESTful best practices: [https://docs.microsoft.com/en-us/azure/architecture/best-practices/api-design](https://docs.microsoft.com/en-us/azure/architecture/best-practices/api-design)

encoding/json: [https://golang.org/pkg/encoding/json/](https://golang.org/pkg/encoding/json/)
