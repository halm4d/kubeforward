package main

import (
	"github.com/halm4d/kubeforward/config"
	"github.com/halm4d/kubeforward/forward"
	"github.com/halm4d/kubeforward/util"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	c := config.New()
	c.Load()

	kubeConf, err := clientcmd.BuildConfigFromFlags("", c.KubeConfig)
	if err != nil {
		panic(err)
	}

	ns := c.FindNamespace(util.ReadNamespaceArg())

	f := forward.Conf{
		KubeConf:  kubeConf,
		Namespace: ns.Name,
		Pod:       ns.FindPod(kubeConf),
		Ports: forward.Ports{
			Local:  ns.Ports.Local,
			Remote: ns.Ports.Remote,
		},
	}
	f.Start()
}
