## API Server wrapper for OrcaNet(full node) and the OrcaWallet, and Btcctl 

### Goals
Create an HTTP server with a start function that another program can call. This server can be started by calling this function, and then you can just interact with the full node, wallet directly by sending HTTP requests.

### Why
The other blockchain group already went with the executable interact with the cli approach. It doesn't make sense for us to the same work. This method also has the advantage of letting us create an instance running on a server that would allow front end web applications to directly query the server for blockchain information, without having to run a local full node! 
