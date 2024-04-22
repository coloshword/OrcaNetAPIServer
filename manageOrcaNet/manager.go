 package manageOrcaNet

 import (
     "fmt"
     "os"
     "os/exec"
     "strings"
     "bufio"
     "time"
     "io"
 )

 const (
     orcaNetPath string = "./OrcaNet/OrcaNet"
     btcctlPath string = "./OrcaNet/cmd/btcctl/btcctl"
     orcaWalletPath string = "./OrcaWallet/btcwallet"
 )

var cmdProcess *exec.Cmd
func Start(params ...string) error {
    _, err := os.Stat(orcaNetPath)
    if os.IsNotExist(err) {
        fmt.Println("Cannot find OrcaNet executable")
        return err
    }

    cmdProcess = exec.Command(orcaNetPath, params...)

    stdout, err := cmdProcess.StdoutPipe()
    if err != nil {
        return fmt.Errorf("failed to create stdout pipe: %w", err)
    }
    stderr, err := cmdProcess.StderrPipe()
    if err != nil {
        return fmt.Errorf("failed to create stderr pipe: %w", err)
    }

    fmt.Println("Start OrcaNet with params: ", params)
    if err := cmdProcess.Start();  err != nil {
        fmt.Println("Failed to start OrcaNet:", err)
        return err
    }
    fmt.Println("OrcaNet started successfully")
    go printOutput(stdout)
    go printOutput(stderr)


    return nil
}


func printOutput(r io.Reader) {
    scanner := bufio.NewScanner(r)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading stream: %v\n", err)
    }
}


// Stop: ends the running OrcaNet instance if its running
func Stop() error {
    if cmdProcess == nil || cmdProcess.Process == nil {
        fmt.Println("OrcaNet process is not currently running.")
        return fmt.Errorf("OrcaNet process is not running")
    }

    fmt.Println("Stopping OrcaNet...")
    // send interrupt sig
    if err := cmdProcess.Process.Signal(os.Interrupt); err != nil {
        fmt.Println("Failed to send interrupt:", err)
        return err
    }
    time.Sleep(5)
    fmt.Println("OrcaNet stopped successfully.")
    return nil
}

//startOrcaWallet: starts the OrcaWallet
func StartOrcaWallet() error {
    // check for the existence of the executable 
    _, err := os.Stat(orcaWalletPath)
    if os.IsNotExist(err) {
        fmt.Println("Cannot find Orcawallet executable")
        return err
    }

    cmd := exec.Command(orcaWalletPath)
    if err := cmd.Start(); err != nil {
        fmt.Println(err)
        fmt.Println("failed to start wallet executable")
        return nil
    }
    fmt.Println("Wallet started successfully")
    return nil
}


// callBtcctlCmd: calls a Btcctl command Exactly as specified in string param, and returns the stdout of btcctl as a string 
// its a singular string, but you can pass as many arguments, we will split the arguments in this fn
func CallBtcctlCmd(cmdStr string) (string, error) {
    params :=  strings.Split(cmdStr, " ") 
    fmt.Println(params)
    cmd := exec.Command(btcctlPath, params...)
    // get the stdout of cmd, CAN HANG but shouldn't be a problem in a btcctl command
    stdout, err := cmd.CombinedOutput() 
    if err != nil {
        return "", fmt.Errorf("failed to execute btcctl commands '%s': %s, error: %v", cmdStr, stdout, err)
    }
    fmt.Println(err);
    return string(stdout), nil
}


