echo "removing all containers"
# shellcheck disable=SC2046
docker rm -vf $(docker ps -a -q)
#echo "removing all images"
## shellcheck disable=SC2046
#docker rmi -f $(docker images -a -q)