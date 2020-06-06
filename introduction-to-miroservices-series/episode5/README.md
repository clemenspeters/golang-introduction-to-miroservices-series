# Building Microservices with Go: 5. Gorilla framework for RESTful services

[https://www.youtube.com/watch?v=DD3JlT_u0DM](https://www.youtube.com/watch?v=DD3JlT_u0DM)

Run

```bash
go run main.go
```

and from another terminal run

```bash
curl localhost:9090 | jq
curl localhost:9090 -v -d '{"id": 1, "name": "tea", "description": "A nice cup of tea."}'
curl localhost:9090/1  -v -XPUT -d '{"id": 1, "name": "tea", "description": "A nice cup of tea5."}'
```

## Links

Source: [https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_5](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_5)
