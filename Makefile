all: orcanet orcawallet btcctl 

orcanet:
	@echo "Building OrcaNet..."
	@go build
	@cd OrcaNet && go build

orcawallet:
	@cd OrcaWallet && go build

btcctl:
	@cd OrcaNet && cd cmd && cd btcctl && go build

clean:
	rm -f OrcaNet/OrcaNet OrcaWallet/btcdwallet OrcaNet/cmd/btcctl/btcctl OrcaNetAPIServer
