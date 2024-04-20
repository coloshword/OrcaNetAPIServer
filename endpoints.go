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
    stdout :=  manageOrcaNet.CallBtcctlCmd(command)
    // return the output of CallBtcctlCommand back to the querier 
    io.WriteString(w, stdout)
}

