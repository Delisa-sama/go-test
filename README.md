# News

### Setting up Docker Swarm cluster
    docker-machine create --driver virtualbox --virtualbox-cpu-count <CPU_COUNT> --virtualbox-memory <MEM> --virtualbox-disk-size <DISK_SIZE> swarm-manager-0
    eval "$(docker-machine env swarm-manager-0)"
    docker network create --driver overlay <NETWORK_NAME>
    docker swarm init --advertise-addr <ADDR>

### Deploy rabbitmq
    ./support.sh

### Deploy microservices
    ./services.sh
    
### Dev run (without docker)
    ./dev.sh