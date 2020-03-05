package cmd

import (
	"time"

	"github.com/mkan0141/gws/util"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		svg := util.GetGrassSVG(user)
		util.Svg2Png(svg, fileName)
	},
}

var (
	user     string
	fileName string
)

func init() {
	time.Now().Format("20060102")
	RootCmd.Flags().StringVar(&user, "u", "", "GitHubのユーザ名(required)")
	RootCmd.MarkFlagRequired("u")
	RootCmd.PersistentFlags().StringVar(&fileName, "o", time.Now().Format("20060102.png"), "出力ファイル名")
}
