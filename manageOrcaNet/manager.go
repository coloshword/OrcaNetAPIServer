 package manageOrcaNet

 import (
     "fmt"
     "os"
     "os/exec"
     "strings"
 )

 const (
     orcaNetPath string = "../OrcaNet/OrcaNet"
     btcctlPath string = "../OrcaNet/cmd/btcctl/btcctl"
 )
 //startOrcaNet: starts the OrcaNet full node
func Start() error {
    // check for the existence of executable

    _, err := os.Stat(orcaNetPath)
    if os.IsNotExist(err) {
        fmt.Println("Cannot find OrcaNet executable")
        return err
    } 

    // we know it exists 
    cmd := exec.Command(orcaNetPath)
    if err := cmd.Start();  err != nil {
        fmt.Println(err)
        fmt.Println("failed to run")
        return nil
    }
    fmt.Println("OrcaNet started successfully")
    return nil
} 

// callBtcctlCmd: calls a Btcctl command Exactly as specified in string param, and returns the stdout of btcctl as a string 
// its a singular string, but you can pass as many arguments, we will split the arguments in this fn
func CallBtcctlCmd(cmdStr string) string {
    params :=  strings.Split(cmdStr, " ") 
    fmt.Println(params)
    cmd := exec.Command(btcctlPath, params...)
    // get the stdout of cmd, CAN HANG but shouldn't be a problem in a btcctl command
    stdout, err := cmd.CombinedOutput() 
    if err != nil {
        fmt.Println("failed to execute Btcctl command")
        return ""
    }
    fmt.Println(err);
    return string(stdout)
}


