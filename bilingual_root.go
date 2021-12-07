package textra

import (
	"encoding/json"
	"fmt"
)

// Bilingual Texts.
const BILINGUALROOT = "bilingual_root"

type BilingualRootRequest struct {
	LangS  string `json:"lang_s"`
	LangT  string `json:"lang_t"`
	Order  string `json:"order"`
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

type BilingualRooListResult struct {
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
			List []BilingualRootList `json:"list"`
		} `json:"result"`
	} `json:"resultset"`
}

type BilingualRootList struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	SourceLanguage string `json:"source_language"`
	TargetLanguage string `json:"target_language"`
	Provide        int    `json:"provide"`
	ProvideTarget  string `json:"provide_target"`
}

func (tra *TexTra) BilingualRootGet(request *BilingualRootRequest) ([]BilingualRootList, error) {
	values := tra.apiValues()
	values.Add("api_name", BILINGUALROOT)
	values.Add("api_param", "get")
	if request != nil {
		values.Add("lang_s", request.LangS)
		values.Add("lang_t", request.LangT)
		values.Add("order", request.Order)
		values.Add("limit", request.Limit)
		values.Add("offset", request.Offset)
	}

	ret, err := tra.apiPost(values)
	if err != nil {
		return nil, err
	}
	result, err := tra.BilingualRootDecode(ret)
	if err != nil {
		return nil, err
	}

	return result.Resultset.Result.List, nil
}

func (tra *TexTra) BilingualRootDecode(ret []byte) (*BilingualRooListResult, error) {
	result := new(BilingualRooListResult)
	if err := json.Unmarshal(ret, result); err != nil {
		return nil, err
	}
	if result.Resultset.Code != 0 {
		return result, fmt.Errorf("%d: %s", result.Resultset.Code, errorText[result.Resultset.Code])
	}
	return result, nil
}
