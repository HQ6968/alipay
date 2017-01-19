package alipay

import (
	"github.com/smartwalle/going/http"
	"io/ioutil"
)

// TradeWapPay https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.stK0ff&docType=4&apiId=1046
func (this *AliPay) TradeWapPay(param AliPayTradeWapPay) (html string, err error) {
	req, err := http.NewRequest("POST", this.apiDomain, this.URLValues(param))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return "", err
	}

	rep, err := this.client.Do(req)
	if err != nil {
		return "", err
	}
	defer rep.Body.Close()

	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return "", err
	}
	html = string(data)
	return html, err
}