#!/bin/zsh
#echo "creating docker group"
#sudo groupadd docker
#echo "adding user to a docker group"
#sudo usermod -aG docker $USER
#echo "relogging"
#newgrp docker
# shellcheck disable=SC2046
echo "removing all containers"
# shellcheck disable=SC2046
docker rm -vf $(docker ps -a -q)
echo "removing all images"
# shellcheck disable=SC2046
docker rmi -f $(docker images -a -q)
echo "running prometheus on port 9090"
docker run -p 9090:9090 -d --name prometheus --net=host -v $PWD/prometheus:/etc/config prom/prometheus --config.file=/etc/config/prometheus.yml
echo "running node exporter on port 9100"
docker run -p 9100:9100 -d --name node_exporter --net=host -v $PWD/:/etc/config prom/node-exporter --path.rootfs=/etc/config
echo "running grafana on port 3000"
docker run -d -p 3000:3000 --name grafana --net=host grafana/grafana
