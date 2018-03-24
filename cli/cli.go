package cli

import (
        "fmt"
        "strconv"
        "time"
        "encoding/json"
        "github.com/anvie/port-scanner"
        "github.com/parnurzeal/gorequest"
        
)

const minPort int=1
const maxPort int=65535
// ScanOption IP,Start and End port
type ScanOption struct {
        IP        string
        StartPort int
        StopPort  int
}

// Walk each node
func walk(items Nodes) {

        if len(items.Nodes) > 0 {
                for i := 0; i < len(items.Nodes); i++ {
                        fmt.Println(fmt.Sprintf("key is %s Value is %s", items.Nodes[i].Key, items.Nodes[i].Value))
                        subnode := items.Nodes[i].Nodes
                        if len(subnode.Nodes) > 0 {
                                walk(subnode)
                        }
                }
        }

}

// isValidPort check
func isValidPort (port int) bool {

	return port > minPort && port < maxPort
}
// Scan with option
func Scan(opt *ScanOption) {

        if (opt.IP == "" ||  !isValidPort(opt.StartPort)  || !isValidPort(opt.StopPort)) {
                panic("Please,check IP address or port numbers!")
        }
        // scan opt.IP with a 2 second timeout per port in 5 concurrent threads
        ps := portscanner.NewPortScanner(opt.IP, 2*time.Second, 5)

        // get opened port
        fmt.Printf("scanning port %d-%d...\n", opt.StartPort, opt.StopPort)

        openedPorts := ps.GetOpenedPort(opt.StartPort, opt.StopPort)
        var app EtcdResult
        for i := 0; i < len(openedPorts); i++ {
                port := openedPorts[i]
                resp, body, errs := gorequest.New().SetCurlCommand(false).SetDebug(false).Get("http://" + opt.IP + ":"+ strconv.Itoa(opt.StartPort) + "/v2/keys/?recursive=true").End()

                fmt.Println("body sample " + fmt.Sprintf("%d", body[0]))
                fmt.Println(resp.Status)
                fmt.Println("****************")

                if errs != nil {
                        fmt.Println(errs)
                }

                err := json.Unmarshal([]byte(body), &app)
                if err != nil {
                        fmt.Println(err)
                }

                walk(app.Node.Nodes)

                fmt.Print(" ", port, " [open]")
                fmt.Println("  -->  ", ps.DescribePort(port))
        }

}