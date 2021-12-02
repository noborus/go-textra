package textra

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Machine translation request.
const TERMROOT = "term_root"

type TermRootRequest struct {
	LangS  string
	LangT  string
	Order  string
	Limit  int
	Offset int
}

type TermRootSetResult struct {
	Resultset struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Total   int    `json:"total"`
		Request struct {
			URL    string `json:"url"`
			Limit  int    `json:"limit"`
			Offset int    `json:"offset"`
		} `json:"request"`
		Result struct {
			List []TermRootList `json:"list"`
		} `json:"result"`
	} `json:"resultset"`
}

type TermRootList struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	SourceLanguage string `json:"source_language"`
	TargetLanguage string `json:"target_language"`
	Provide        int    `json:"provide"`
}

func (tra *TexTra) TermRootGet(request *TermRootRequest) ([]TermRootList, error) {
	values := url.Values{
		"access_token": []string{tra.token.AccessToken},
		"key":          []string{tra.ClientID},
		"name":         []string{tra.Name},
		"api_name":     []string{TERMROOT},
		"api_param":    []string{"get"},
		"type":         []string{"json"},
	}
	ret, err := tra.apiPost(values)
	if err != nil {
		return nil, err
	}
	result, err := tra.TermRootDecode(ret)
	if err != nil {
		return nil, err
	}

	return result.Resultset.Result.List, nil
}

func (tra *TexTra) TermRootDecode(ret []byte) (*TermRootSetResult, error) {
	result := new(TermRootSetResult)
	if err := json.Unmarshal(ret, result); err != nil {
		return nil, err
	}
	if result.Resultset.Code != 0 {
		return result, fmt.Errorf("%d: %s", result.Resultset.Code, errorText[result.Resultset.Code])
	}
	return result, nil
}
