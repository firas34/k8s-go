package main

import (
	"flag"
	"fmt"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	configfile := flag.String("kubeconfig", "~/.kube/config-0", "Configfile of current k8s cluster")

	config, err := clientcmd.BuildConfigFromFlags("", *configfile)

	if err != nil {
		// handle error
	}
	fmt.Println(config)

}
