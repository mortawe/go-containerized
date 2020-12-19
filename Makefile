.PHONY: netsetgo
netsetgo:
	@ #wget 'https://github.com/teddyking/netsetgo/releases/download/0.0.1/netsetgo'
	@ sudo cp ./assets/netsetgo /usr/local/bin/
	@ sudo chown root:root /usr/local/bin/netsetgo
	@ sudo chmod 4755 /usr/local/bin/netsetgo
	@ echo 'netsetgo: finished'


.PHONY: genrootfs
genrootfs:
	@ mkdir -p /tmp/go-containerized/rootfs
	@ tar -C /tmp/go-containerized/rootfs -xf assets/alpine-minirootfs-3.12.3-x86_64.tar.gz
	@ echo 'genrootfs: finished'

.PHONY: build
build:
	@ go build -o ./bin/go-conteinerized ./src/cmd/
	@ echo 'build: finished'

.PHONY: install
install: netsetgo genrootfs build
	@ sudo cp ./bin/go-conteinerized /usr/local/bin/go-conteinerized
	@ sudo chown root:root /usr/local/bin/go-conteinerized
