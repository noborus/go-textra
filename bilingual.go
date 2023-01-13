package textra

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// A library for registering, updating, deleting, and searching bilingual collections.
// Operate the id (pid) managed by Billingual_root.

// Bilingual entries.
const BILINGUAL = "bilingual"

type BilingualSearchRequest struct {
	Text      string `json:"text"`
	MatchType string `json:"match_type"`
	Pid       string `json:"pid"`
	Order     string `json:"order"`
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
}

type BilingualSearchResult struct {
	Resultset struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Request struct {
			URL    string `json:"url"`
			Limit  int    `json:"limit"`
			Offset int    `json:"offset"`
		} `json:"request"`
		Result struct {
			Bilingual []BilingualSearchItem `json:"bilingual"`
		} `json:"result"`
	} `json:"resultset"`
}

type BilingualSearchItem struct {
	Pid    int    `json:"pid"`
	Cid    int    `json:"cid"`
	Source string `json:"source"`
	Target string `json:"target"`
}

func (tra *TexTra) BilingualSearch(request *BilingualSearchRequest) ([]BilingualSearchItem, error) {
	values := tra.apiValues()
	values.Add("api_name", BILINGUAL)
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

	result, err := tra.BilingualSearchDecode(ret)
	if err != nil {
		return nil, err
	}

	return result.Resultset.Result.Bilingual, nil
}

func (tra *TexTra) BilingualSearchDecode(ret []byte) (*BilingualSearchResult, error) {
	result := new(BilingualSearchResult)
	if err := json.Unmarshal(ret, result); err != nil {
		return nil, err
	}

	if result.Resultset.Code != 0 {
		return result, fmt.Errorf("%d: %s", result.Resultset.Code, errorText[result.Resultset.Code])
	}
	return result, nil
}

type BilingualSetResult struct {
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

func (tra *TexTra) BilingualSet(pid string, textS string, textT string) (int, int, error) {
	values := tra.apiValues()
	values.Add("api_name", BILINGUAL)
	values.Add("api_param", "set")
	values.Add("text_s", textS)
	values.Add("text_t", textT)
	return tra.bilingualSet(values)
}

func (tra *TexTra) BilingualUpdate(pid string, cid string, textS string, textT string) (int, int, error) {
	values := tra.apiValues()
	values.Add("api_name", BILINGUAL)
	values.Add("api_param", "update")
	values.Add("pid", pid)
	values.Add("cid", cid)
	values.Add("text_s", textS)
	values.Add("text_t", textT)
	return tra.bilingualSet(values)
}

func (tra *TexTra) BilingualDelete(pid string, cid string) (int, int, error) {
	values := tra.apiValues()
	values.Add("api_name", BILINGUAL)
	values.Add("api_param", "delete")
	values.Add("pid", pid)
	values.Add("cid", cid)
	return tra.bilingualSet(values)
}

func (tra *TexTra) bilingualSet(values url.Values) (int, int, error) {
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

func BilingualSetDecode(jsonStr []byte) (*TermSetResult, error) {
	result := new(TermSetResult)
	if err := json.Unmarshal(jsonStr, result); err != nil {
		return nil, err
	}
	return result, nil
}
