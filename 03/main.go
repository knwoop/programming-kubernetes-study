package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var namespace = flag.String("n", "", "namespaces in kubernetes")

func main() {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("failed clientcmd.BuildConfigFromFlags: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("failed kubernetes.NewForConfig: %v", err)
	}

	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods(*namespace).List(ctx,  metav1.ListOptions{})
	if err != nil {
		log.Fatalf("failed Pods.List: %v", err)
	}

	for i, pod := range pods.Items {
		fmt.Printf("[%d] %s\n", i, pod.GetName())
	}
}
