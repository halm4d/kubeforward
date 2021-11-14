package forward

import (
	"bytes"
	"fmt"
	"k8s.io/apimachinery/pkg/util/httpstream"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"net/http"
	"net/url"
	"strings"
)

type Conf struct {
	KubeConf  *rest.Config
	Namespace string
	Pod       string
	Ports     Ports
}

type Ports struct {
	Local  string
	Remote string
}

func (f *Conf) Start() {
	dialer := f.createDialer()
	stopChan, readyChan := make(chan struct{}, 1), make(chan struct{}, 1)
	out, errOut := new(bytes.Buffer), new(bytes.Buffer)

	forwarder, err := portforward.New(dialer, []string{fmt.Sprintf("%s:%s", f.Ports.Local, f.Ports.Remote)}, stopChan, readyChan, out, errOut)
	if err != nil {
		panic(err)
	}

	go func() {
		for range readyChan {
		}
		if len(errOut.String()) != 0 {
			panic(errOut.String())
		} else if len(out.String()) != 0 {
			fmt.Printf("Portforwarding %v pod from %v namespace to port %v.\n", f.Pod, f.Namespace, f.Ports.Local)
			fmt.Println(out.String())
		}
	}()

	if err = forwarder.ForwardPorts(); err != nil {
		panic(err)
	}
}

func (f *Conf) createDialer() httpstream.Dialer {
	path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward", f.Namespace, f.Pod)
	hostIP := strings.TrimLeft(f.KubeConf.Host, "htps:/")
	serverURL := url.URL{Scheme: "https", Path: path, Host: hostIP}
	roundTripper, upgrader, err := spdy.RoundTripperFor(f.KubeConf)
	if err != nil {
		panic(err)
	}
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: roundTripper}, http.MethodPost, &serverURL)
	return dialer
}
