#!/usr/bin/env bash

# This script is used to start and stop docker-compose services
# It can also be used to generate a new docker-compose file

if [ "$1" == "start" ]; then
    docker compose -f docker-compose/docker-compose.yaml --env-file docker-compose/.env up -d
elif [ "$1" == "stop" ]; then
    docker compose -f docker-compose/docker-compose.yaml --env-file docker-compose/.env down
elif [ "$1" == "generate" ]; then
    echo "Generating docker-compose file"
    echo "What is the name of the service?"
    read -r serviceName
    mkdir -p docker-compose/"$serviceName"
    cp docker-compose/compose-template.yaml docker-compose/"$serviceName"/docker-compose.yaml
    sed -i "s/service_name/$serviceName/g" docker-compose/"$serviceName"/docker-compose.yaml
    {
        echo "  - path: ./$serviceName/docker-compose.yaml"
        echo "    project_directory: .."
        echo "    env_file: docker-compose/.env"
        echo ""
    } >>docker-compose/docker-compose.yaml
    echo "Docker-compose file for $serviceName generated"
else
    echo "Please specify start or stop as argument"
fi
