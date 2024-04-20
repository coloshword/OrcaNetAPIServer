package main

import (
    "fmt"
    "io"
    "net/http"
    "errors"
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
    const command string = "getblockchaininfo 2"
    manageOrcaNet.callBtcctlCmd(command)
   
}

// startOrcaNet: starts an OrcaNet full node instance for the server to communicate with
func startOrcaNet() (error) {
   return manageOrcaNet.Start()
}

func main() {
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/hello", getHello)
    http.HandleFunc("/getBlockchainInfo", getBlockchainInfo)
    fmt.Println("starting orcanet")
    startOrcaNet()    
    err := http.ListenAndServe(":3333", nil)
    if errors.Is(err, http.ErrServerClosed) {
        fmt.Println("server is closed")
    } else if err != nil {
        fmt.Printf("error starting the server %s\n ", err)
    }
}

