package cmd

import (
	"fmt"
	"time"

	"github.com/mkan0141/gws/util"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(util.GetGrassSVG(user))
	},
}

var (
	user string
	file string
)

func init() {
	time.Now().Format("20060102")
	RootCmd.Flags().StringVar(&user, "u", "", "GitHubのユーザ名(required)")
	RootCmd.MarkFlagRequired("u")
	RootCmd.PersistentFlags().StringVar(&file, "o", time.Now().Format("20060102"), "出力ファイル名")
}
