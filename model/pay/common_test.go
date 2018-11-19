package pay

import (
	"net/url"
	"testing"
)

var priKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDfjnxqHfLmryXXnaDOQen4n+5Jf6jnbzncmqZGQDts/CgHUfca
aRCTAm+FIyN5FRokuXZWK6jkH/jLnegn/Hr5OoJSD9q9y0BHxYrZaJt04IpgbuqO
r5DtsndPnRjks3Y/QwyCj3BY6xj0Fo3DISGzzQgfohL9WPnoDggbQmqXvwIDAQAB
AoGBAJIBCvx7RtKmfT6OsiFDJz27pfLWr0dHezC6x/GFrcoF/Vaaj5nuGGcK5i67
vkUsJQDrJ4Arz4f94Y2KOb8zxFOc8C7tTM04Q9BMFQTgDpwckS0lDrlfd6kHStDI
2wPy66DTo9nlzluNRgPWvBT8Tn+LbwmBz7D8eEMObKEUieZRAkEA9UB5huAcCldC
wibTzfn2GuMtYVNPVu4a1QL/tUY1nGBO3kyWR16zxiUWt+VdCzbT0mBp63VtDCWt
Z6dXqny7YwJBAOlam6za2HYtVN9ghfNSE/3hCaaoeGFThjTcrj2FNn01VDinS72k
eKK/iTlibzLI9p+rEMaBggQ2PLD6bn0IVvUCQQCvMb+eebmOKYem6dWj7kvAKUjh
nYGvt6ezQtEnzV++tY2hf1Ra52vEv/napB4zRJdMUVNYwCmF4+Rbh084mqHBAkEA
n0VruchJNCfupOQhqRjdckv1pV2ZhHxYvp3dAzp4HW+Xw29UP+URPavTgmpQEW6e
/g3pTkO4tR07wWO8o/RcPQJAavtpJsZJfxuc3+AI74dcwZhbN+ImqYsbyG2D3HxI
jvBQ59pr2yOzs2CGRy8VvG/G89p/kKMD+ySjouZj/488nw==
-----END RSA PRIVATE KEY-----
`

var pubKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDfjnxqHfLmryXXnaDOQen4n+5J
f6jnbzncmqZGQDts/CgHUfcaaRCTAm+FIyN5FRokuXZWK6jkH/jLnegn/Hr5OoJS
D9q9y0BHxYrZaJt04IpgbuqOr5DtsndPnRjks3Y/QwyCj3BY6xj0Fo3DISGzzQgf
ohL9WPnoDggbQmqXvwIDAQAB
-----END PUBLIC KEY-----`

var a = "count=1&dealId=2500207776&giftCardMoney=0&hbBalanceMoney=0&hbMoney=0&orderId=2645985458&partnerId=0&payMoney=1&payTime=1542124311&payType=1117&promoDetail=&promoMoney=0&returnData=&status=2&totalMoney=1&tpOrderId=18607895241&unitPrice=1&userId=3269234686470"
var b = "oOFlK3oOlVSbfLfgL%2FbgqO2rDIxt%2FVFQdxYeoZlj9WHXMrAW4hzyGwlplNJ%2FZzip9h0MRUSDwLQxDDhkFF1s59%2FzY%2BInhB6e%2FNyR%2Bpaurl2Sj1%2Fdj4Q0rbkuqMdKhU1A4A4jmG3fxTTx7FqExOcBjSxz6peLtmXyHDKhshoYYKc%3D"
var pubKey1 = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCtzQERUhDNLkR/yTGZxHU6cy5a
bC7kxeXqdL4Jq7v6A0MqhdK8T/cDr05DHKLmm0QlEdN0Jz2mg8QMkDW9pfPVWxQh
TXikYhLBvdM4qas45LCKuswvMMjDw+lI4zD5M2Zh2T7S+mqkdqg8cOmvLlCvVEg4
5Su/Bj9HrMlGaBHPpQIDAQAB
-----END PUBLIC KEY-----`

func TestRsaVerySignWithSha1Base64(t *testing.T) {
	b, _ := url.QueryUnescape(b)
	if err := CheckSign(a, b, pubKey1); err != nil {
		t.Error(err)
	}
}
