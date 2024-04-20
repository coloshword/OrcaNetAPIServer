 package manageOrcaNet

 import (
     "fmt"
     "os"
     "os/exec"
 )
 //startOrcaNet: starts the OrcaNet full node
func Start() error {
    const path string = "../OrcaNet/OrcaNet"
    // check for the existence of executable

    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        fmt.Println("Cannot find OrcaNet executable")
        return err
    } 

    // we know it exists 
    cmd := exec.Command(path)
    if err := cmd.Run(); err != nil {
        fmt.Println(err)
        fmt.Println("failed to run")
        return nil
    }
    return nil 
} 


