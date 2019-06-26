package main

import(
	"context"
	"encoding/json"

	//"encoding/json"
	"log"
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"

	"github.com/yalp/jsonpath"
)

func main(){
	configURL := "http://localhost:8180/auth/realms/golang"
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, configURL)
	if err != nil {
		panic(err)
	}

	clientID := "gotutorial"
	clientSecret := "ee297078-16dc-4908-8e3e-b4718506037f"

	redirectURL := "http://localhost:10000"
	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),
		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}
	state := "somestate"

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rawAccessToken := r.Header.Get("Authorization")
		if rawAccessToken == "" {
			w.WriteHeader(401)
			w.Write([]byte("fokof"))
			return
		}

		parts := strings.Split(rawAccessToken, " ")
		if len(parts) != 2 {
			w.WriteHeader(400)
			return
		}
		idToken, err := verifier.Verify(ctx, parts[1])

		if err != nil {
			http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
			return
		}

		var claimsResult *json.RawMessage

		claimError := idToken.Claims(&claimsResult)

		if claimError != nil {
			log.Println("Unable to get claims:", claimError.Error())
		}

		bytes, errMarshal := claimsResult.MarshalJSON()

		if errMarshal != nil {
			log.Println("Marshal to string failed: ", errMarshal)
		}

		//log.Println("Payload: ", string(bytes))

		var payload interface{}

		jsonPathMarshallErr := json.Unmarshal(bytes, &payload)

		if jsonPathMarshallErr != nil {
			log.Println("Error marshalling to interface", jsonPathMarshallErr.Error())
		}

		test, _ := jsonpath.Read(payload, "$.test")

		log.Println("get claim result: ", test)

		roles , _ := jsonpath.Read(payload, "$.roles")

		log.Println("get roles claim result: ", roles)

		result := roles.([]interface{})

		log.Print("first role: ", result[0])

		rolesAsString := result[0].(string)

		log.Println("Role as string: ", rolesAsString)

		w.Write([]byte("hello world"))
	})

	log.Fatal(http.ListenAndServe("localhost:10000", nil))
}