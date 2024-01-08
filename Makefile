# Leggi il file di configurazione JSON e assegna i valori alle variabili
CONFIG_FILE := config.json
#LOAD_BALANCER := $(shell jq -r .addressLoadBalancer $(CONFIG_FILE))
#SERVER_1 := $(shell jq -r .addressServers[0] $(CONFIG_FILE))
#SERVER_2 := $(shell jq -r .addressServers[1] $(CONFIG_FILE))
#SERVER_3 := $(shell jq -r .addressServers[2] $(CONFIG_FILE))

SERVER_1=localhost:8085
SERVER_2=localhost:8086
SERVER_3=localhost:8087

server:
	cd .\main && \
	go run .\server.go $(SERVER_1) & \
	go run .\server.go $(SERVER_2) & \
	go run .\server.go $(SERVER_3)

loadBalancer:
	cd .\main && \
	go run .\loadBalancer.go

client:
	cd .\main && \
	go run .\client.go 5 5
