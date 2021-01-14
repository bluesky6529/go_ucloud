package GetBalance

import (
	"encoding/json"
	"log"

	"github.com/tidwall/gjson"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

func GetBalance(AccessKeyID string, AccessKeySecret string) string {
	cfg := ucloud.NewConfig()
	credential := auth.NewCredential()
	credential.PrivateKey = AccessKeyID
	credential.PublicKey = AccessKeySecret
	client := ucloud.NewClient(&cfg, &credential)
	req := client.NewGenericRequest()
	req.SetAction("GetBalance")
	res, err := client.GenericInvoke(req)
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(res.GetPayload())
	log.Printf("%f", gjson.GetBytes(b, "AccountInfo.Amount").Float())
}
