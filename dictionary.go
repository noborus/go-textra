package textra

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Machine translation request.
const LOOKUP = "lookup"

type LookUpResult struct {
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
			Lookup []LookupItem `json:"lookup"`
		} `json:"result"`
	} `json:"resultset"`
}

type LookupItem struct {
	Position string `json:"position"`
	Length   int    `json:"length"`
	Hit      string `json:"hit"`
	Term     []struct {
		ID     interface{} `json:"id"`
		Pid    string      `json:"pid"`
		Source string      `json:"source"`
		Target string      `json:"target"`
	} `json:"term"`
}

func (tra TexTra) Lookup(text string, pid string, lang_s string) ([]LookupItem, error) {
	values := url.Values{
		"access_token": []string{tra.token.AccessToken},
		"key":          []string{tra.ClientID},
		"name":         []string{tra.Name},
		"api_name":     []string{LOOKUP},
		"type":         []string{"json"},
		"text":         []string{text},
		"pid":          []string{pid},
		"lang_s":       []string{lang_s},
	}
	ret, err := tra.apiPost(values)
	if err != nil {
		return nil, err
	}

	data, err := lookupDecodee(ret)
	if err != nil {
		return nil, err
	}
	return data.Resultset.Result.Lookup, nil
}

func lookupDecodee(jsonStr []byte) (*LookUpResult, error) {
	result := new(LookUpResult)
	if err := json.Unmarshal(jsonStr, result); err != nil {
		return nil, err
	}
	if result.Resultset.Code != 0 {
		return result, fmt.Errorf("%d: %s", result.Resultset.Code, errorText[result.Resultset.Code])
	}
	return result, nil
}
