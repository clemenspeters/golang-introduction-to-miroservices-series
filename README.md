# Building Microservices with Go: 6 JSON Validation

[https://www.youtube.com/watch?v=gE8\_-8KoOLc](https://www.youtube.com/watch?v=gE8_-8KoOLc)

Run

```bash
go run main.go
```

and from another terminal run

```bash
curl localhost:9090 | jq
curl localhost:9090 -v -d '{"id": 1, "name": "tea", "description": "A nice cup of tea."}'
curl localhost:9090/1  -v -X PUT -d '{"id": 1, "name": "tea", "description": "A nice cup of tea5."}'
```

## Links

Source: [https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_6](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_6)
