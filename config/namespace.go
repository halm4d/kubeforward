package config

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"regexp"
)

type NameSpaceConfig struct {
	Name     string `yaml:"name"`
	PodRegex string `yaml:"podRegex"`
	Ports    struct {
		Local  string `yaml:"local"`
		Remote string `yaml:"remote"`
	} `yaml:"ports"`
}

func (c *NameSpaceConfig) FindPod(kubeConf *rest.Config) string {
	client, err := kubernetes.NewForConfig(kubeConf)
	if err != nil {
		panic(err)
	}
	list, err := client.CoreV1().Pods(c.Name).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	compile, _ := regexp.Compile(c.PodRegex)
	for _, pod := range list.Items {
		if compile.MatchString(pod.Name) {
			return pod.Name
		}
	}
	panic(fmt.Sprintf("Pod not found with regex: %v", c.PodRegex))
}
