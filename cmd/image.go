package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubemb/pkg/kube"
	"kubemb/pkg/utils"
)

var (
	example = `# Get deployment image versions of all namespaces
%[1]s image

# Get the deployment image version of the specified namespace
%[1]s image -n default`
)

func Image() error {
	clientset := kube.Clientset(KubernetesConfigFlags)

	n, _ := rootCmd.Flags().GetString("namespace")
	deployList, err := clientset.AppsV1().Deployments(n).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return err
	}

	var itemSlice []string
	itemSlice = append(itemSlice, "NAMESPACE\tDEPLOY_NAME\tCONTAINER_NAME\tIMAGE")
	for i := 0; i < len(deployList.Items); i++ {
		d := deployList.Items[i]
		for i := 0; i < len(d.Spec.Template.Spec.Containers); i++ {
			c := d.Spec.Template.Spec.Containers[i]
			tab := "\t"
			item := d.Namespace + tab + d.Name + tab + c.Name + tab + c.Image
			itemSlice = append(itemSlice, item)
		}
	}

	utils.Fprint(itemSlice)

	return nil
}

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "View the currently deployed image version",
	Long: `View the currently deployed image versiondir
`,
	Example: fmt.Sprintf(example, "kubemb"),
	Run: func(cmd *cobra.Command, args []string) {
		if err := Image(); err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
}
