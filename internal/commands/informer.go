package commands

import (
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1ac "k8s.io/client-go/applyconfigurations/core/v1"
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
		ctx := cmd.Context()

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
			AddFunc: func(obj any) {
				secretObj := obj.(*v1.Secret)
				cmd.Println("I see that you added a secret", secretObj.Name, secretObj.Namespace)
			},
			DeleteFunc: func(obj any) {
				secretObj := obj.(*v1.Secret)
				cmd.Println("I see that you deleted a secret", secretObj.Name, secretObj.Namespace)
				cmd.Println("I'm going to create it again", secretObj.Name, secretObj.Namespace)
				timeString := strconv.Itoa(int(time.Now().UnixNano()))
				secApply, err := v1ac.ExtractSecret(secretObj, FieldManagerSecretKube)
				printError(err, cmd, "Error:")
				secApply.WithLabels(map[string]string{"newly-created": timeString})
				_, err = clientset.CoreV1().Secrets(secretObj.Namespace).Apply(ctx, secApply, metav1.ApplyOptions{FieldManager: FieldManagerSecretKube})
				if err != nil {
					cmd.PrintErrf("server side apply failed %v", err)
				}
			},
		})
		ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
		defer cancel()
		secretInformer.Run(ctx.Done())
	},
}
