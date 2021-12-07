package textra

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const APIURL = "https://mt-auto-minhon-mlt.ucri.jgn-x.jp/"

const (
	APIKeyError                     = 500
	NameError                       = 501
	MaximumRequestDayError          = 502
	MaximumRequestMinError          = 504
	MaximumRequestSimultaneousError = 505
	OAuthAuthenticationError        = 510
	OAuthHeaderError                = 511
	AccessURLError                  = 520
	AccessURL2Error                 = 521
	RequestKeyError                 = 522
	RequestNameError                = 523
	RequestParameterError           = 524
	RequestParameterSizeError       = 525
	AuthorizationError              = 530
	ExecutionError                  = 531
	Nodata                          = 532
)

var errorText = map[int]string{
	APIKeyError:                     "API key error",
	NameError:                       "Name error",
	MaximumRequestDayError:          "Maximum request error(day)",
	MaximumRequestMinError:          "Maximum request error(min)",
	MaximumRequestSimultaneousError: "Maximum request error (simultaneous requests)",
	OAuthAuthenticationError:        "OAuth authentication error",
	OAuthHeaderError:                "OAuth header error",
	AccessURLError:                  "Access URL error",
	AccessURL2Error:                 "Access URL error",
	RequestKeyError:                 "Request key error",
	RequestNameError:                "Request name error",
	RequestParameterError:           "Request parameter error",
	RequestParameterSizeError:       "Request parameter error (for data size limit)",
	AuthorizationError:              "Authorization error",
	ExecutionError:                  "Execution error",
	Nodata:                          "No data",
}

type Config struct {
	ClientID     string
	ClientSecret string
	Name         string
}

type TexTra struct {
	client       *http.Client
	token        *oauth2.Token
	ClientID     string
	ClientSecret string
	Name         string
	APIName      string
	APIParam     string
}

func New(c Config) (*TexTra, error) {
	ctx := context.Background()
	conf := &clientcredentials.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		TokenURL:     APIURL + "oauth2/token.php",
	}

	client := conf.Client(ctx)
	token, err := conf.Token(ctx)
	if err != nil {
		return nil, err
	}
	api := TexTra{}
	api.client = client
	api.token = token
	api.ClientID = c.ClientID
	api.ClientSecret = c.ClientSecret
	api.Name = c.Name

	return &api, nil
}

func (tra *TexTra) apiValues() url.Values {
	return url.Values{
		"access_token": []string{tra.token.AccessToken},
		"key":          []string{tra.ClientID},
		"name":         []string{tra.Name},
		"type":         []string{"json"},
	}
}

func (tra *TexTra) apiPost(values url.Values) ([]byte, error) {
	resp, err := tra.client.PostForm(APIURL+"api/", values)
	if err != nil {
		return nil, err
	}
	ret, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
