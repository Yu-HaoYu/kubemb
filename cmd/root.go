package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

const (
	version = "0.0.2\nBy John.Yu"
)

var (
	KubernetesConfigFlags *genericclioptions.ConfigFlags
)

var rootCmd = &cobra.Command{
	Use:     "kubemb",
	Short:   "This is a kubectl magic box plugin",
	Long:    `This is a kubectl magic box plugin`,
	Version: version,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	KubernetesConfigFlags = genericclioptions.NewConfigFlags(true)
	KubernetesConfigFlags.AddFlags(rootCmd.PersistentFlags())
}
