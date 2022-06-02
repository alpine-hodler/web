package morningstar

// * This is a generated file, do not edit

// Token is an OAuth bearer token authorized using a base64-encoded user id and password.
type Token struct {
	AccessToken string `json:"access_token" bson:"access_token"`
	ExpiresIn   int    `json:"expires_in" bson:"expires_in"`
	TokenType   string `json:"token_type" bson:"token_type"`
}
