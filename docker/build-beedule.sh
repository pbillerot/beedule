

docker ps -a | grep beedule | awk '{print $1}' | xargs docker rm -f
docker-compose up -d --build
docker image prune -f