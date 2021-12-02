package textra

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Machine translation request.
const MT = "mt"

// Machine translation engine.
const (
	// General - NT 【English - Japanese】 汎用NT 【英語 - 日本語】
	GENERAL_EN_JA = "generalNT_en_ja"
	// Patents - NT 特許NT 【英語 - 日本語】
	PATENTNT_EN_JA = "patentNT_en_ja"
	// 対話NT(音声翻訳エンジン専用) 【英語 - 日本語】
	VOICETRANT_EN_JA = "voicetraNT_en_ja"
	// Finance - NT 金融NT 【英語 - 日本語】
	FSANT_EN_JA = "fsaNT_ja_en"
	// 法令契約NT 【英語 - 日本語】
	LAWSNT_EN_JA = "lawsNT_ja_en"
	// General - NT+ (For Experimental Use Only / Technology Transfer Not Allowed) 【English - Japanese】
	// 汎用NT＋(実験用・技術移転不可) 【英語 - 日本語】
	MINNSNONT_EN_JA = "minnaNT_en_ja"

	// General - NT 【Japanese - English】
	GENERAL_JA_EN = "generalNT_ja_en"
	// Patents - NT 特許NT 【英語 - 日本語】
	PATENTNT_JA_EN = "patentNT_ja_en"
	// 対話NT(音声翻訳エンジン専用) 【英語 - 日本語】
	VOICETRANT_JA_EN = "voicetraNT_ja_en"
	// Finance - NT 金融NT 【英語 - 日本語】
	FSANT_JA_EN = "fsaNT_ja_en"
	// 法令契約NT 【英語 - 日本語】
	LAWSNT_JA_EN = "lawsNT_ja_en"
	// General - NT+ (For Experimental Use Only / Technology Transfer Not Allowed)
	// 汎用NT＋(実験用・技術移転不可) 【英語 - 日本語】
	MINNSNONT_JA_EN = "minnaNT_ja_en"
)

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
			Text        string `json:"text"`
			Information json.RawMessage
		} `json:"result"`
	} `json:"resultset"`
}

type MTInformation struct {
	TextS    string `json:"text-s"`
	TextT    string `json:"text-t"`
	Sentence []struct {
		TextS string `json:"text-s"`
		TextT string `json:"text-t"`
		Split []struct {
			TextS   string `json:"text-s"`
			TextT   string `json:"text-t"`
			Process struct {
				Regex []struct {
					Text    string `json:"text"`
					Result  string `json:"result"`
					Pattern string `json:"pattern"`
					Replace string `json:"replace"`
				} `json:"regex"`
				ReplaceBefore []struct {
					TextS string `json:"text-s"`
					TextT string `json:"text-t"`
					TermS string `json:"term-s"`
					TermT string `json:"term-t"`
				} `json:"replace-before"`
				Preprocess []interface{} `json:"preprocess"`
				Translate  struct {
					Reverse []struct {
						Selected int    `json:"selected"`
						IdN      int    `json:"id-n"`
						IdR      int    `json:"id-r"`
						NameN    string `json:"name-n"`
						NameR    string `json:"name-r"`
						TextS    string `json:"text-s"`
						TextT    string `json:"test-t"`
						TextR    string `json:"test-r"`
						Score    int    `json:"score"`
					} `json:"reverse"`
					Specification []interface{}     `json:"specification"`
					TextS         string            `json:"text-s"`
					TextT         string            `json:"text-t"`
					Associate     [][]interface{}   `json:"associate"`
					Oov           interface{}       `json:"oov"`
					Exception     string            `json:"exception"`
					Associates    [][][]interface{} `json:"associates"`
				} `json:"translate"`
				ReplaceAfter []struct {
					TextS string `json:"text-s"`
					TextT string `json:"text-t"`
					TermS string `json:"term-s"`
					TermT string `json:"term-t"`
				} `json:"replace-after"`
			} `json:"process"`
		} `json:"split"`
	} `json:"sentence"`
}

func (tra *TexTra) Translate(apiType string, original string) (string, error) {
	data, err := tra.TranslateRaw(apiType, original)
	if err != nil {
		return "", err
	}
	return data.Resultset.Result.Text, nil
}

func (tra *TexTra) TranslateRaw(apiType string, original string) (*MTResult, error) {
	values := url.Values{
		"access_token": []string{tra.token.AccessToken},
		"key":          []string{tra.ClientID},
		"api_name":     []string{MT},
		"api_param":    []string{apiType},
		"name":         []string{tra.Name},
		"type":         []string{"json"},
		"text":         []string{original},
	}
	ret, err := tra.apiPost(values)
	if err != nil {
		return nil, err
	}
	data, err := mtDecode(ret)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func mtDecode(jsonStr []byte) (*MTResult, error) {
	result := new(MTResult)
	if err := json.Unmarshal(jsonStr, result); err != nil {
		return nil, err
	}
	if result.Resultset.Code != 0 {
		return result, fmt.Errorf("%d: %s", result.Resultset.Code, errorText[result.Resultset.Code])
	}
	return result, nil
}

func MTInfoDecode(result *MTResult) (*MTInformation, error) {
	info := new(MTInformation)
	if err := json.Unmarshal(result.Resultset.Result.Information, info); err != nil {
		return nil, err
	}
	return info, nil
}
