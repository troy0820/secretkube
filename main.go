package main

import (
	"encoding/base64"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	fmt.Println("Using kubeconfig: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	secrets, err := clientset.CoreV1().Secrets("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	if len(secrets.Items) > 0 {
		fmt.Printf("There are %d secrets in the cluster\n", len(secrets.Items))
		for _, secret := range secrets.Items {
			fmt.Printf("secret %s \t %s \n", secret.GetName(), base64.StdEncoding.EncodeToString([]byte(secret.GetNamespace())))
		}
	} else {
		fmt.Println("No secrets found!")
	}
}
