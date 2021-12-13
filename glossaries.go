package textra

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Glossary entries.
const TERM = "term"

type TermSearchRequest struct {
	Text      string `json:"text"`
	MatchType string `json:"match_type"`
	Pid       string `json:"pid"`
	Order     string `json:"order"`
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
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

func (tra *TexTra) TermSearch(request *TermSearchRequest) ([]TermSearchTerm, error) {
	values := tra.apiValues()
	values.Add("api_name", TERM)
	values.Add("api_param", "search")
	values.Add("pid", request.Pid)
	values.Add("text", request.Text)
	values.Add("match_type", request.MatchType)
	values.Add("order", request.Order)
	values.Add("limit", request.Limit)
	values.Add("offset", request.Offset)

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

func (tra *TexTra) TermSearchDecode(ret []byte) (*TermSearchResult, error) {
	result := new(TermSearchResult)
	if err := json.Unmarshal(ret, result); err != nil {
		return nil, err
	}
	if result.Resultset.Code != 0 {
		return result, fmt.Errorf("%d: %s", result.Resultset.Code, errorText[result.Resultset.Code])
	}
	return result, nil
}

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

func (tra *TexTra) TermSet(pid string, textS string, textT string) (int, int, error) {
	values := tra.apiValues()
	values.Add("api_name", TERM)
	values.Add("api_param", "set")
	values.Add("pid", pid)
	values.Add("text_s", textS)
	values.Add("text_t", textT)
	return tra.termSet(values)
}

func (tra *TexTra) TermUpdate(pid string, cid string, textS string, textT string) (int, int, error) {
	values := tra.apiValues()
	values.Add("api_name", TERM)
	values.Add("api_param", "update")
	values.Add("pid", pid)
	values.Add("cid", cid)
	values.Add("text_s", textS)
	values.Add("text_t", textT)
	return tra.termSet(values)
}

func (tra *TexTra) TermDel(pid string, cid string) (int, int, error) {
	values := tra.apiValues()
	values.Add("api_name", TERM)
	values.Add("api_param", "delete")
	values.Add("pid", pid)
	values.Add("cid", cid)
	return tra.termSet(values)
}

func (tra *TexTra) termSet(values url.Values) (int, int, error) {
	ret, err := tra.apiPost(values)
	if err != nil {
		return 0, 0, err
	}

	data, err := termSetDecode(ret)
	if err != nil {
		return 0, 0, err
	}
	return data.Resultset.Result.Pid, data.Resultset.Result.Cid, nil
}

func termSetDecode(jsonStr []byte) (*TermSetResult, error) {
	result := new(TermSetResult)
	if err := json.Unmarshal(jsonStr, result); err != nil {
		return nil, err
	}
	return result, nil
}
