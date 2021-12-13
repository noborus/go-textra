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

// API Process error code.
const (
	APIKeyError                     = 500 // API key error
	NameError                       = 501 // Name error
	MaximumRequestDayError          = 502 // Maximum request error(day)
	MaximumRequestMinError          = 504 // Maximum request error(min)
	MaximumRequestSimultaneousError = 505 // Maximum request error (simultaneous requests)
	OAuthAuthenticationError        = 510 // OAuth authentication error
	OAuthHeaderError                = 511 // OAuth header error
	AccessURLError                  = 520 // Access URL error
	AccessURL2Error                 = 521 // Access URL error
	RequestKeyError                 = 522 // Request key error
	RequestNameError                = 523 // Request name error
	RequestParameterError           = 524 // Request parameter error
	RequestParameterSizeError       = 525 // Request parameter error (for data size limit)
	AuthorizationError              = 530 // Authorization error
	ExecutionError                  = 531 // Execution error
	Nodata                          = 532 // No data
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

	token, err := conf.Token(ctx)
	if err != nil {
		return nil, err
	}

	api := TexTra{}
	api.client = conf.Client(ctx)
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
