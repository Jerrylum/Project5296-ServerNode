# Start the Server

Please read https://hub.docker.com/_/nginx for more information.

```bash
sudo docker compose up
```

# Retrive Config Files

You can retrieve the unmodified config files from the container using the following commands:

```bash
sudo docker run --rm --entrypoint=cat nginx /etc/nginx/nginx.conf > nginx.conf
sudo docker run --rm --entrypoint=cat nginx /etc/nginx/conf.d/default.conf > default.conf
```
