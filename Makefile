.PHONY: all orcanet orcawallet btcctl OrcaNetAPIServer clean

all: orcanet orcawallet btcctl OrcaNetAPIServer

OrcaNetAPIServer:
	@echo "Building OrcaNetAPIServer..."
	@go build -v
	@echo "Built OrcaNetAPIServer successfully."

orcanet:
	@echo "Building OrcaNet..."
	@cd OrcaNet && go build -v
	@echo "Built OrcaNet successfully."

orcawallet:
	@echo "Building OrcaWallet..."
	@cd OrcaWallet && go build -v
	@echo "Built OrcaWallet successfully."

btcctl:
	@echo "Building btcctl..."
	@cd OrcaNet/cmd/btcctl && go build -v
	@echo "Built btcctl successfully."

clean:
	@echo "Cleaning up..."
	@rm -f OrcaNet/OrcaNet OrcaWallet/btcwallet OrcaNet/cmd/btcctl/btcctl OrcaNetAPIServer
	@echo "Clean up done."

