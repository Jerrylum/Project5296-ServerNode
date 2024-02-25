# Introduction

This is a Nginx Docker compose file that uses the official Nginx image from Docker Hub. It is used to serve static files for the project Project5296 experiments.

To simulate different server constraints, you can modify the `default.conf` file. The `nginx.conf` file is the main configuration file for the server, and the `default.conf` file is the configuration file for the default server block.

Inside the `default.conf` file, `limit_conn` and `limit_rate` are used to simulate the server's connection and bandwidth constraints. The `limit_conn` directive is used to limit the number of connections per IP address, and the `limit_rate` directive is used to limit the bandwidth per connection.

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
