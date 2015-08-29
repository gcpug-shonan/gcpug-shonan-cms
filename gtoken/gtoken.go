package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

var (
	secretJSON = flag.String("json", "", "Service coount json key.")
)

type secret struct {
	Type         string `json:"type"`
	PrivateKey   string `json:"private_key"`
	PrivateKeyID string `json:"private_key_id"`
	ClientID     string `json:"client_id"`
	ClientEmail  string `json:"client_email"`
}

func usage() {
	fmt.Fprint(os.Stderr, "Usage: gtoken --json [ServiceAccount JSON key PATH]")
	os.Exit(2)
}

func main() {
	flag.Parse()
	if *secretJSON == "" {
		usage()
	}
	jf, e := ioutil.ReadFile(*secretJSON)
	if e != nil {
		fmt.Fprintf(os.Stderr, "File error: %v\n", e)
		os.Exit(1)
	}

	var s secret
	json.Unmarshal(jf, &s)

	conf := &jwt.Config{
		Email:      s.ClientEmail,
		PrivateKey: []byte(s.PrivateKey),
		TokenURL:   google.JWTTokenURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/appengine.admin",
			"https://www.googleapis.com/auth/cloud-platform",
		},
	}

	ts := conf.TokenSource(oauth2.NoContext)
	t, e := ts.Token()
	if e != nil {
		fmt.Fprint(os.Stderr, "Token error: %v\n", e)
	}
	fmt.Fprintf(os.Stdout, "%v\n", t.AccessToken)
}
