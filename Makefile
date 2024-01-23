# Leggi il file di configurazione JSON e assegna i valori alle variabili
CONFIG_FILE := .\configuration\config.json
LOAD_BALANCER := $(shell jq -r .addressLoadBalancer $(CONFIG_FILE)) #inutile
SERVER_ADDRESSES := $(shell jq -r .addressServers[] $(CONFIG_FILE))
NUMBER_OF_SERVERS := $(shell jq -r .numberOfServers $(CONFIG_FILE))

all:
	cd .\main && \
	$(foreach addr, $(SERVER_ADDRESSES), start go run .\server.go $(addr) &&) \
	start go run .\loadBalancer.go && \
	start go run .\client.go && \
	echo "Avviate ulteriori shell"

server:
	cd .\main && \
	$(foreach addr, $(SERVER_ADDRESSES), start go run .\server.go $(addr) &&) \
	echo "Server avviati"

loadBalancer:
	cd .\main && \
	go run .\loadBalancer.go

client:
	cd .\main && \
	start go run .\client.go &
