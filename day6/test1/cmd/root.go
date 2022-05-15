package cmd

import "github.com/spf13/cobra"

var cmd = &cobra.Command{} //创建初始化对象Command

func Execute() {
	cmd.Execute()
}
func init() {
	cmd.AddCommand(modelToDataCmd)
}
