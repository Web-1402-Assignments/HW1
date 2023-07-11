# HW1

docker nginx command : $ docker run --name nginx-test -v ~/docker-nginx/default.conf:/etc/nginx/conf.d/default.conf -v ~/docker-nginx/example.com+5-key.pem:/etc/nginx/ssl/example.com+5-key.pem -v ~/docker-nginx/example.com+5.pem:/etc/nginx/ssl/example.com+5.pem --network=host -d nginx
