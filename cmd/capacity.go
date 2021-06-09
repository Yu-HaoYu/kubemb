package cmd

import (
	"context"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubemb/pkg/kube"
	"kubemb/pkg/utils"
)

func Capacity() error {
	clientset := kube.Clientset(KubernetesConfigFlags)

	n, _ := rootCmd.Flags().GetString("namespace")
	deployList, err := clientset.AppsV1().Deployments(n).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return err
	}

	var itemSlice []string
	itemSlice = append(itemSlice, "NAMESPACE\tCONTAINER NAME\tCPU REQUESTS\tMEMORY REQUESTS\tCPU LIMITS\tMEMORY LIMITS")
	for i := 0; i < len(deployList.Items); i++ {
		d := deployList.Items[i]
		n := d.Namespace
		name := d.Spec.Template.Spec.Containers[0].Name
		rCPU := d.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu().String()
		rMem := d.Spec.Template.Spec.Containers[0].Resources.Requests.Memory().String()
		lCPU := d.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu().String()
		lMem := d.Spec.Template.Spec.Containers[0].Resources.Limits.Memory().String()

		itemSlice = append(itemSlice, n+"\t"+name+"\t"+rCPU+"\t"+rMem+"\t"+lCPU+"\t"+lMem)
	}

	utils.Fprint(itemSlice)

	return nil
}

var capacityCmd = &cobra.Command{
	Use:     "capacity",
	Aliases: []string{"cap"},
	Short:   "Alias \"cap\", view the current resource usage of Node and Pod",
	Long:    "View the current resource usage of Node and Pod",
	Run: func(cmd *cobra.Command, args []string) {
		if err := Capacity(); err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(capacityCmd)
}
