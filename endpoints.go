package main

import (
    "fmt"
    "io"
    "net/http"
    "github.com/coloshword/OrcaNetAPIServer/manageOrcaNet"
    "strings"
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

// stopMine: endpoint to stop mining
func stopMine(w http.ResponseWriter, r *http.Request) {
    fmt.Println("stop mine endpoint")
}





