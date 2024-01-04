# Sample HTTP Golang API program with Dockerfile

## Command line testing
```
# Access the user1 endpoint
curl http://localhost:8080/user1
# Output: Hello, World!

# Access the user2 endpoint
curl http://localhost:8080/user2

```

## Docker build and run
```
docker build -t sample:test .
docker run -p 8080:8080 sample:test

```


