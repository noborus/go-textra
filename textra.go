package textra

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const APIURL = "https://mt-auto-minhon-mlt.ucri.jgn-x.jp/"

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

func New(c Config) TexTra {
	ctx := context.Background()
	conf := &clientcredentials.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		TokenURL:     APIURL + "oauth2/token.php",
	}

	client := conf.Client(ctx)
	token, err := conf.Token(ctx)
	if err != nil {
		log.Fatal(err)
	}
	api := TexTra{}
	api.client = client
	api.token = token
	api.ClientID = c.ClientID
	api.ClientSecret = c.ClientSecret
	api.Name = c.Name

	return api
}

func (tra TexTra) apiPost(values url.Values) ([]byte, error) {
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
