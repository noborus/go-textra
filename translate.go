package textra

import (
	"encoding/json"
	"net/url"
)

// Machine translation request.
const MT = "mt"

// Machine translation engine.
const GENERAL_EN_JA = "generalNT_en_ja"
const GENERAL_JA_EN = "generalNT_ja_en"

type MTResult struct {
	Resultset struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Request struct {
			URL   string `json:"url"`
			Text  string `json:"text"`
			Split int    `json:"split"`
			Data  string `json:"data"`
		} `json:"request"`
		Result struct {
			Text        string      `json:"text"`
			Information interface{} `json:"information"`
		} `json:"result"`
	} `json:"resultset"`
}

func (tra TexTra) Translate(apiType string, str string) (string, error) {
	values := url.Values{
		"access_token": []string{tra.token.AccessToken},
		"key":          []string{tra.ClientID},
		"api_name":     []string{MT},
		"api_param":    []string{apiType},
		"name":         []string{tra.Name},
		"type":         []string{"json"},
		"text":         []string{str},
	}
	ret, err := tra.apiPost(values)
	if err != nil {
		return "", err
	}

	data := new(MTResult)
	if err := json.Unmarshal(ret, data); err != nil {
		return "", err
	}
	return data.Resultset.Result.Text, nil
}
