package weapp

import (
	"github.com/medivhzhan/weapp/util"
)

const (
	verifySignatureAPI = "/cgi-bin/soter/verify_signature"
)

// VerifySignatureResponse 生物认证秘钥签名验证请求返回数据
type VerifySignatureResponse struct {
	Response
	IsOk bool `json:"is_ok"`
}

// VerifySignature 生物认证秘钥签名验证
// @accessToken 接口调用凭证
// @openID 用户 openid
// @data 通过 wx.startSoterAuthentication 成功回调获得的 resultJSON 字段
// @signature 通过 wx.startSoterAuthentication 成功回调获得的 resultJSONSignature 字段
func VerifySignature(accessToken, openID, data, signature string) (*VerifySignatureResponse, error) {
	api, err := util.TokenAPI(BaseURL+verifySignatureAPI, accessToken)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"openid":         openID,
		"json_string":    data,
		"json_signature": signature,
	}

	res := new(VerifySignatureResponse)
	if err := util.PostJSON(api, params, res); err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, res.ErrorWithInfo("failed to verify signature")
	}

	return res, nil
}
