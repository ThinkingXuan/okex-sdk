package rest

import (
	"context"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api"
	requstAccount "github.com/amir-the-h/okex/requests/rest/account"
	"log"
	"testing"
)

const apiKey = ""
const secretKey = ""
const passphrase = ""

// 查看账户余额
// https://www.okx.com/docs-v5/zh/#trading-account-rest-api-get-balance
func TestAccountBalance(t *testing.T) {
	client := initClient()

	req := requstAccount.GetBalance{
		Ccy: []string{"USDT"},
	}
	res, err := client.Rest.Account.GetBalance(req)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v", res.Balances[0])

}

// 查看持仓信息
// https://www.okx.com/docs-v5/zh/#trading-account-rest-api-get-positions
func TestAccountPosition(t *testing.T) {
	client := initClient()

	// InstType: 产品类型
	// PosID: 持仓ID
	// InstID: 交易产品ID
	// 都是非必填项
	req := requstAccount.GetPositions{
		InstID:   []string{"BTC-USDT-SWAP"},
		PosID:    []string{""},
		InstType: okex.SwapInstrument,
	}
	res, err := client.Rest.Account.GetPositions(req)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v", res)
}

// 查看账户持仓风险
// https://www.okx.com/docs-v5/zh/#trading-account-rest-api-get-account-and-position-risk
func TestAccountAndPositionRisk(t *testing.T) {
	client := initClient()

	// InstType: 产品类型
	req := requstAccount.GetAccountAndPositionRisk{
		InstType: okex.SwapInstrument,
	}
	res, err := client.Rest.Account.GetAccountAndPositionRisk(req)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v", res)
}

// 查看账单流水
// https://www.okx.com/docs-v5/zh/#trading-account-rest-api-get-bills-details-last-7-days
// https://www.okx.com/docs-v5/zh/#trading-account-rest-api-get-bills-details-last-3-months
func TestBills(t *testing.T) {
	client := initClient()

	// InstType: 产品类型
	req := requstAccount.GetBills{
		InstType: okex.SwapInstrument,
	}
	res7days, err := client.Rest.Account.GetBills(req, false)
	if err != nil {
		log.Fatalln(err)
	}
	//res3months, err := client.Rest.Account.GetBills(req, true)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	log.Printf("%#v", res7days)
	//log.Printf("%#v", res3months)
}

func initClient() *api.Client {
	dest := okex.NormalServer // The main API server
	ctx := context.Background()
	client, err := api.NewClient(ctx, apiKey, secretKey, passphrase, dest)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
