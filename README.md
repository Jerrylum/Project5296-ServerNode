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

# Generate Blob Files

This repository contains a script to generate blob files. It can be used to generate any size of file for testing purposes. The script is located at `./generate.go`. You can use the following command to generate blob files in any size:

```bash
go run generate.go -m=100
```

This script is written in Go, so you need to have Go installed on your machine to run this script. This will generate a file named `100.out` with 100MB of data located at the ./public/download directory. The file name is the size of the file in MB.
