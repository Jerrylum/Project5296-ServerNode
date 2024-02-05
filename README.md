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
