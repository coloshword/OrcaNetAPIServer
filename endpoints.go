package main

import (
    "fmt"
    "io"
    "net/http"
    "github.com/coloshword/OrcaNetAPIServer/manageOrcaNet"
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

