package main

import (
	"context"
	"fmt"
	"os"

	"github.com/troy0820/secretkube/internal/commands"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		os.Exit(1)
	}
	clientset := kubernetes.NewForConfigOrDie(config)
	secretClientSet := clientset.CoreV1().Secrets("default")
	watcher, err := secretClientSet.Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		os.Exit(1)
	}
	for event := range watcher.ResultChan() {
		switch event.Type {
		case watch.Added:
			fmt.Fprintf(os.Stdout, "\n I've watched you add a secret to the cluster in the default namespace.....\n")
			watcher.Stop()
		}
	}
	//Executes secretkube cli
	commands.Execute()
}
