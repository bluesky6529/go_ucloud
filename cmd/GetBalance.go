package cmd

import (
	"log"

	_ "github.com/bluesky6529/go_ucloud/configs"
	_ "github.com/bluesky6529/ucloud/internal/GetBalance"

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
		log.Printf("%s 帳戶餘額: %s", account, request)
	},
}

func init() {
	queryaccountbalance.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	queryaccountbalance.MarkFlagRequired("account")
}