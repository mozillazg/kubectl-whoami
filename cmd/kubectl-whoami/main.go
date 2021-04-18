package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/mozillazg/kubectl-whoami/pkg/cert"
	"github.com/mozillazg/kubectl-whoami/pkg/version"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func showVersion() {
	fmt.Printf(`kubectl-whoami:
 Version:    %s
 Go version: %s
 Git commit: %s
 Built:      %s
 OS/Arch:    %s/%s
`,
		version.Version, runtime.Version(), version.Commit, version.Date, runtime.GOOS, runtime.GOARCH)
}

func main() {
	raw := flag.Bool("raw", false, "output raw json result")
	ver := flag.Bool("version", false, "show version and exit")
	flag.Parse()

	if *ver {
		showVersion()
		return
	}

	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Printf("parse kubeconfig failed: %+v", err)
		os.Exit(1)
	}
	c, err := cert.GetCertInfo(cfg)
	if err != nil {
		fmt.Printf("parse certinfo failed: %+v", err)
		os.Exit(1)
	}

	d := ""
	if *raw {
		d = cert.ToJSON(*c)
	} else {
		d = cert.Summary(*c)
	}
	fmt.Println(d)
}