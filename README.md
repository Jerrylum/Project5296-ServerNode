Read https://hub.docker.com/_/nginx for more information

# Start the server

```bash
sudo docker compose up
```

# Retrive config files

```bash
sudo docker run --rm --entrypoint=cat nginx /etc/nginx/nginx.conf > nginx.conf
sudo docker run --rm --entrypoint=cat nginx /etc/nginx/conf.d/default.conf > default.conf
```

# Curl command

```bash
# Include header
curl -i http://127.0.0.1
# With range and header
curl -r 0-1 -i http://127.0.0.1
```

# More curl commands

```bash
curl -i http://16.163.217.155 http://127.0.0.1:8080
curl -H "Connection: keep-alive" -H "Keep-Alive: timeout=5, max=100" -i http://16.163.217.155 http://127.0.0.1:8080

nc 16.163.217.155 80
```

# Download cat image
```bash
curl http://127.0.0.1/download/200.jpg > output.jpg

curl -r 0-20000 http://127.0.0.1/download/200.jpg > part1
curl -r 20000 http://127.0.0.1/download/200.jpg > part2

cat part1 part2 > output.jpg

```
