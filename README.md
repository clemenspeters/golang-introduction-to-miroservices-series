# Building Microservices with Go: 7 Documenting RESTful APIs with Swagger

[https://www.youtube.com/watch?v=07XhTqE-j8k](https://www.youtube.com/watch?v=07XhTqE-j8k)

Run

```bash
go run main.go
```

and from another terminal run

```bash
curl localhost:9090 | jq
curl localhost:9090 -v -d '{"name": "Tea", "sku": "abc-ab-tea", "price": 1.0}'
curl localhost:9090/1  -v -X PUT -d '{"name": "Water", "sku": "abc-ab-water", "price": 2.0}'
```

## Links

Source: [https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_7](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_7)
