windows:
	CONFIG_FILE := .\configuration\config.json
	LOAD_BALANCER := $(shell jq -r .addressLoadBalancer $(CONFIG_FILE)) #inutile
	SERVER_ADDRESSES := $(shell jq -r .addressServers[] $(CONFIG_FILE))
	NUMBER_OF_SERVERS := $(shell jq -r .numberOfServers $(CONFIG_FILE))

	cd .\main && \
	$(foreach addr, $(SERVER_ADDRESSES), start go run .\server.go $(addr) &&) \
	start go run .\loadBalancer.go && \
	start go run .\client.go && \
	echo "Avviate ulteriori shell"

unix:
	# !!!---IMPORTANTE---!!!
	# Per eseguire correttamente in Unix bisogna modificare il path nel file `config.go`:
	#Sostituire `const filename = "..\\configuration\\config.json"` con `const filename = "../configuration/config.json"`.

	CONFIG_FILE := ./configuration/config.json
	LOAD_BALANCER := $(shell jq -r .addressLoadBalancer $(CONFIG_FILE))
	SERVER_ADDRESSES := $(shell jq -r .addressServers[] $(CONFIG_FILE))
	NUMBER_OF_SERVERS := $(shell jq -r .numberOfServers $(CONFIG_FILE))

	cd ./main && \
	$(foreach addr, $(SERVER_ADDRESSES), go run ./server.go $(addr) &) \
	go run ./loadBalancer.go & \
	go run ./client.go & \
	echo "Avviate ulteriori shell"