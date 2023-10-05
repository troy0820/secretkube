package commands

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	toolswatch "k8s.io/client-go/tools/watch"
)

var watcherCmd = &cobra.Command{
	Use:     "watch",
	Short:   "Watches k8s secrets in a namespace",
	Long:    `Watches k8s secrets in a namespace`,
	Aliases: []string{"watch"},
	Run: func(cmd *cobra.Command, args []string) {

		var kubeconfig, ns string
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
		if cmd.Flags().Changed("config") {
			var err error
			kubeconfig, err = cmd.Flags().GetString("config")
			printError(err, cmd, "Error:")
		}
		ns = "default"
		if cmd.Flags().Changed("namespace") {
			var err error
			ns, err = cmd.Flags().GetString("namespace")
			printError(err, cmd, "Error:")
		}
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		printErrorWithExit(err, cmd, "Error:")
		clientset := kubernetes.NewForConfigOrDie(config)
		secretclient := clientset.CoreV1().Secrets(ns)
		watchFunc := func(_ metav1.ListOptions) (watch.Interface, error) {
			return secretclient.Watch(context.Background(), metav1.ListOptions{})
		}
		retryWatcher, err := toolswatch.NewRetryWatcher("1", &cache.ListWatch{WatchFunc: watchFunc})
		if err != nil {
			printError(err, cmd, red("Error: "))
		}
		ch := make(chan struct{})
		go callWatcher(retryWatcher, cobra.Command{}, ch)
		<-ch
	},
}

func callWatcher(watcher watch.Interface, cmd cobra.Command, done chan struct{}) {
	for event := range watcher.ResultChan() {
		sec := event.Object.(*v1.Secret)
		if event.Type == watch.Added || event.Type == watch.Modified {
			cmd.Printf("I watched you add/modify a secret to the %s namespace: %s \n", sec.Namespace, sec.Name)
		}
		if event.Type == watch.Deleted {
			cmd.Printf("I watched you delete the secret to the cluster in the %s namespace.....\n", sec.Namespace)
		}
	}
	done <- struct{}{}
}
