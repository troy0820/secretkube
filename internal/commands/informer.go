package commands

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

var informerCmd = &cobra.Command{
	Use:     "inform",
	Short:   "informs about secrets",
	Long:    `Informs about secrets `,
	Aliases: []string{"inform"},
	Run: func(cmd *cobra.Command, args []string) {

		var kubeconfig string
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
		if cmd.Flags().Changed("config") {
			var err error
			kubeconfig, err = cmd.Flags().GetString("config")
			printError(err, cmd, "Error:")
		}
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		printErrorWithExit(err, cmd, "Error:")
		clientset := kubernetes.NewForConfigOrDie(config)
		informerFactory := informers.NewSharedInformerFactory(clientset, time.Hour*24)
		secretInformer := informerFactory.Core().V1().Secrets().Informer()
		secretInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				secretObj := obj.(*v1.Secret)
				cmd.Println("I see that you added a secret", secretObj.Name, secretObj.Namespace)
			},
			DeleteFunc: func(obj interface{}) {
				secretObj := obj.(*v1.Secret)
				cmd.Println("I see that you deleted a secret", secretObj.Name, secretObj.Namespace)
				cmd.Println("I'm going to create it again", secretObj.Name, secretObj.Namespace)
				sec := &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      secretObj.Name,
						Namespace: secretObj.Namespace,
					},
					Data: secretObj.Data,
				}
				_, err := clientset.CoreV1().Secrets(secretObj.Namespace).Create(context.Background(), sec, metav1.CreateOptions{})
				if err != nil {
					cmd.Println("can't create secret", err)
				}
			},
		})
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()
		secretInformer.Run(ctx.Done())
	},
}
