package GetBalance

import (
	"encoding/json"

	"github.com/tidwall/gjson"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

func GetBalance(AccessKeyID string, AccessKeySecret string) float64 {
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
	//response := log.Printf("%f", gjson.GetBytes(b, "AccountInfo.Amount").Float())
	return gjson.GetBytes(b, "AccountInfo.Amount").Float()
}
