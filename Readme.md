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

# Datadog run time metric

DD_AGENT_MAJOR_VERSION=7 DD_API_KEY=XXxxxxxxxxxx DD_SITE="us5.datadoghq.com" bash -c "$(curl -L https://s3.amazonaws.com/dd-agent/scripts/install_mac_os.sh)"


api_key: xxxxxxx
site: us5.datadoghq.com





