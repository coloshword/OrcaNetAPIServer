package main

import (
    "fmt"
    "io"
    "net/http"
    "github.com/coloshword/OrcaNetAPIServer/manageOrcaNet"
    "strings"
	"strconv"
	"encoding/json"
)
// some basic endpoints

// getRoot: the root endpoint ('/')
func getRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Println("got / request")
    io.WriteString(w, "This is the root of the API server")
}

// getHello: the hello endpoint
func getHello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("got hello request")
    io.WriteString(w, "Hello, HTTP!\n")
}

// getBlockchainInfo: endpoint to get the blockchain info
func getBlockchainInfo(w http.ResponseWriter, r *http.Request) {
    fmt.Println("getBlockchainInfo request") 
    const command string = "getblockchaininfo"
    stdout, err :=  manageOrcaNet.CallBtcctlCmd(command)
    fmt.Println(err)
    // return the output of CallBtcctlCommand back to the querier 
    io.WriteString(w, stdout)
}
// getNewAddress: endpoint to get a new wallet address
// this wallet address can be used for mining rewards / sending / receiving transactions
// For security purposes, it is recommended to create a new address everytime 
func getNewAddress(w http.ResponseWriter, r *http.Request) {
    fmt.Println("getNewAddress request")
    const command string = "getnewaddress --wallet"
    stdout, err := manageOrcaNet.CallBtcctlCmd(command)
    fmt.Println(err)
    io.WriteString(w, stdout)
}

// getBalance: gets the balance of the wallet 
func getBalance(w http.ResponseWriter, r *http.Request) {
    fmt.Println("getBalance endpoint")
    const command string = "getbalance --wallet"
    stdout, err := manageOrcaNet.CallBtcctlCmd(command)
    fmt.Println(err)
    io.WriteString(w, stdout)
}

// mine: endpoint to start mining, mining rewards go to the associated wallet on this node 
func mine(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Mine endpoint")
    // to mine, we need to restart the OrcaNet node with --generate and --mingaddr=<newAddress>
    const getAddressCmd string = "getnewaddress --wallet"
    stdout, err := manageOrcaNet.CallBtcctlCmd(getAddressCmd)
    if err != nil {
        fmt.Println("error getting a new address for mining")
        io.WriteString(w, "error getting a new address for mining")
    }
    //std out is the address
    address := strings.TrimSpace(stdout)
    orcaNetParams := []string{"--generate", "--miningaddr=" + address}
    fmt.Println("orcaParams[0] " + orcaNetParams[0])
    fmt.Println("orcaParams[1] " + orcaNetParams[1])
    // we need to kill the OrcaNet process first
    if err := manageOrcaNet.Stop(); err != nil {
        fmt.Println("failed to end OrcaNet:", err)
        http.Error(w, "failed to kill original OrcaNet instance to start mining", http.StatusInternalServerError)
        return 
    }
    if err := manageOrcaNet.Start(orcaNetParams...); err != nil {
        fmt.Println("failed to start mining:", err)
        http.Error(w, "failed to start mining", http.StatusInternalServerError)
        return
    }
    io.WriteString(w, "Mining successfully started")
}

// sendToAddress: endpoint to send n coins to an address
// if you want to send coins to a specific wallet, ask the recepient to getNewAddress and pass that address to the query string 
// Usage: make a JSON request with 2 fields "coins" and "address"
func sendToAddress(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Coins	string `json:"coins"`
		Address string `json:"address"`
        SenderWalletPass string `json:"senderwalletpass"`
	}
    fmt.Println("send to address endpoint")
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Error reading request body", http.StatusBadRequest)
        return
    }
    if request.Coins == "" || request.Address == "" {
        http.Error(w, "Missing coins or address in request", http.StatusBadRequest)
        return
    }

    if _, err := strconv.ParseFloat(request.Coins, 64); err != nil {
        http.Error(w, "Invalid number format for coins", http.StatusBadRequest)
        return
    }   
    // unlock the wallet first 
    err := unlockWallet(request.SenderWalletPass, w)
    if err != nil {
        fmt.Println("failed to unlock wallet, wallet sender pass likely wrong ")
        http.Error(w, "error unlocking wallet", http.StatusInternalServerError)
    }
	err2 := sendCoins(request.Coins, request.Address, w)
    if err2 != nil {
		fmt.Println("error sending coins")
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    fmt.Println("coins: " + request.Coins + "address " + request.Address)

    fmt.Fprintf(w, "Successfully sent %s coins to %s\n", request.Coins, request.Address)

	
}

func unlockWallet(walletPass string, w http.ResponseWriter) error {
    command := fmt.Sprintf("--wallet walletpassphrase \"%s\" 100", walletPass)
    stdout, err := manageOrcaNet.CallBtcctlCmd(command)
    if err != nil {
        fmt.Println(err)
        io.WriteString(w, "error unlocking wallet")
        return err
    }
    io.WriteString(w, stdout)
    return nil
}
func sendCoins(numCoins, address string, w http.ResponseWriter) error {
    command := fmt.Sprintf("--wallet sendtoaddress %s %s", address, numCoins)
    stdout, err := manageOrcaNet.CallBtcctlCmd(command)
    if err != nil {  
        fmt.Println(err)
        io.WriteString(w, "error sending coins")
        return err
    }
    io.WriteString(w, stdout)
    return nil
}



// stopMine: endpoint to stop mining
func stopMine(w http.ResponseWriter, r *http.Request) {
    fmt.Println("stop mine endpoint")
}





