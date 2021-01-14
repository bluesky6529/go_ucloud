package cmd

import (
	"fmt"
	"ucloud/internal/GetBalance"

	_ "ucloud/configs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getbalance = &cobra.Command{
	Use:   "GetBalance",
	Short: "查詢帳戶餘額",
	Long:  "用來查詢帳戶餘額使用方法如下 : \nGetBalance -a <帳戶名稱> ex: GetBalance -a account",
	Run: func(cmd *cobra.Command, args []string) {
		var AccessKeyID = viper.GetString(account + ".AccessKeyID")
		var AccessKeySecret = viper.GetString(account + ".AccessKeySecret")

		request := GetBalance.GetBalance(AccessKeyID, AccessKeySecret)
		fmt.Printf("%s 帳戶餘額: %f", account, request)
	},
}

func init() {
	getbalance.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	getbalance.MarkFlagRequired("account")
}
