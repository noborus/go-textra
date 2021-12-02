package textra

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Machine translation request.
const TERM = "term"

type TermSetResult struct {
	Resultset struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Request struct {
			URL string `json:"url"`
		} `json:"request"`
		Result struct {
			Pid int `json:"pid"`
			Cid int `json:"cid"`
		} `json:"result"`
	} `json:"resultset"`
}

type TermSearchRequest struct {
	Text      string
	MatchType string
	Pid       string
	Order     string
	Limit     int
	Offset    int
}

type TermSearchResult struct {
	Resultset struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Request struct {
			URL    string `json:"url"`
			Limit  int    `json:"limit"`
			Offset int    `json:"offset"`
		} `json:"request"`
		Result struct {
			Term []TermSearchTerm `json:"term"`
		} `json:"result"`
	} `json:"resultset"`
}

type TermSearchTerm struct {
	Pid    int    `json:"pid"`
	Cid    int    `json:"cid"`
	Source string `json:"source"`
	Target string `json:"target"`
}

func (tra TexTra) TermSearch(request *TermSearchRequest) ([]TermSearchTerm, error) {
	values := url.Values{
		"access_token": []string{tra.token.AccessToken},
		"key":          []string{tra.ClientID},
		"name":         []string{tra.Name},
		"api_name":     []string{TERM},
		"api_param":    []string{"search"},
		"type":         []string{"json"},
		"text":         []string{request.Text},
		"match_type":   []string{request.MatchType},
		"pid":          []string{fmt.Sprintf("%d", request.Pid)},
	}
	ret, err := tra.apiPost(values)
	if err != nil {
		return nil, err
	}
	result, err := tra.TermSearchDecode(ret)
	if err != nil {
		return nil, err
	}

	return result.Resultset.Result.Term, nil
}

func (tra TexTra) TermSearchDecode(ret []byte) (*TermSearchResult, error) {
	result := new(TermSearchResult)
	if err := json.Unmarshal(ret, result); err != nil {
		return nil, err
	}
	if result.Resultset.Code != 0 {
		return result, fmt.Errorf("%d: %s", result.Resultset.Code, errorText[result.Resultset.Code])
	}
	return result, nil

}

func (tra TexTra) TermSet(pid string, text_s string, text_t string) (int, int, error) {
	values := url.Values{
		"access_token": []string{tra.token.AccessToken},
		"key":          []string{tra.ClientID},
		"name":         []string{tra.Name},
		"api_name":     []string{TERM},
		"api_param":    []string{"set"},
		"type":         []string{"json"},
		"pid":          []string{pid},
		"text_s":       []string{text_s},
		"text_t":       []string{text_t},
	}
	ret, err := tra.apiPost(values)
	if err != nil {
		return 0, 0, err
	}

	data := new(TermSetResult)
	if err := json.Unmarshal(ret, data); err != nil {
		return 0, 0, err
	}
	return data.Resultset.Result.Pid, data.Resultset.Result.Cid, nil
}
