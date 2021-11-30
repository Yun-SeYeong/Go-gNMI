package main

import (
  "context"
  "fmt"

  "github.com/aristanetworks/glog"
  "github.com/aristanetworks/goarista/gnmi"
  pb "github.com/openconfig/gnmi/proto/gnmi"
)

var cfg = &gnmi.Config{
    Addr:     "192.168.100.172:5901",
    Username: "admin",
    Password: "admin",
}

func main() {
    paths := []string{"/interfaces/interface[name=Ethernet2]/state"}
    var origin = "openconfig"
    //var origin = "eos_native"
    ctx := gnmi.NewContext(context.Background(), cfg)
    client, err := gnmi.Dial(cfg)
    if err != nil {
        glog.Fatal(err)
    }

    req, err := gnmi.NewGetRequest(gnmi.SplitPaths(paths), origin)
    if err != nil {
        glog.Fatal(err)
    }
    if cfg.Addr != "" {
        if req.Prefix == nil {
            req.Prefix = &pb.Path{}
        }
        req.Prefix.Target = cfg.Addr
    }
    fmt.Println(req)

    err = gnmi.GetWithRequest(ctx, client, req)
    if err != nil {
        glog.Fatal(err)
        fmt.Println(err)
    }   
}