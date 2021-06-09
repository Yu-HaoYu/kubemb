package kube

import (
	"fmt"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func Clientset(configFlags *genericclioptions.ConfigFlags) *kubernetes.Clientset {
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		fmt.Println(err, "failed to read kubeconfig")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err, "failed to create clientset")
	}

	return clientset
}
