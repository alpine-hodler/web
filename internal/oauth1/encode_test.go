package oauth1

// func TestOauth1(t *testing.T) {
// 	t.Run("twitter", func(t *testing.T) {
// 		// These subtests comes form developer.twitter.com:
// 		// https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature
// 		consumerSecret := "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw"
// 		consumerTokenSecret := "LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE"
// 		params := map[string]string{
// 			"oauth_consumer_key":     "xvz1evFS4wEEPTGEFPHBog",
// 			"oauth_signature_method": "HMAC-SHA1",
// 			"oauth_timestamp":        "1318622958",
// 			"oauth_version":          "1.0",
// 			"oauth_token":            "370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb",
// 			"oauth_nonce":            "kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg",
// 			"include_entities":       "true",
// 			"status":                 "Hello Ladies + Gentlemen, a signed OAuth request!",
// 		}

// 		req, err := http.NewRequest("POST", "https://api.twitter.com/1.1/statuses/update.json?include_entities=true", nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		t.Run("normalizeParameterString", func(t *testing.T) {
// 			// This test comes form developer.twitter.com:
// 			// https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature
// 			norm := normalizedParameterString(params)
// 			expected := "include_entities=true&oauth_consumer_key=xvz1evFS4wEEPTGEFPHBog&oauth_nonce=kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1318622958&oauth_token=370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb&oauth_version=1.0&status=Hello%20Ladies%20%2B%20Gentlemen%2C%20a%20signed%20OAuth%20request%21"
// 			require.Equal(t, expected, norm)
// 		})
// 		t.Run("signatureBase", func(t *testing.T) {
// 			// This test comes form developer.twitter.com:
// 			// https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature
// 			expected := "POST&https%3A%2F%2Fapi.twitter.com%2F1.1%2Fstatuses%2Fupdate.json&include_entities%3Dtrue%26oauth_consumer_key%3Dxvz1evFS4wEEPTGEFPHBog%26oauth_nonce%3DkYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1318622958%26oauth_token%3D370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb%26oauth_version%3D1.0%26status%3DHello%2520Ladies%2520%252B%2520Gentlemen%252C%2520a%2520signed%2520OAuth%2520request%2521"
// 			require.Equal(t, expected, signatureBase(req, params))
// 		})
// 		t.Run("signer", func(t *testing.T) {
// 			signer := new(HMACSigner)
// 			signer.ConsumerSecret = consumerSecret
// 			str, err := signer.Sign(consumerTokenSecret, signatureBase(req, params))
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			require.Equal(t, "hCtSmYh+iHYCEqBWrE7C7hYmtUk=", str)
// 		})
// 		// t.Run("SetRequestTokenAuthHeadeHMAC", func(t *testing.T) {
// 		// 	if err := SetRequestTokenAuthHeadeHMAC(req, consumerSecret, consumerTokenSecret, params); err != nil {
// 		// 		t.Fatal(err)
// 		// 	}
// 		// 	fmt.Println(req.Header)
// 		// 	// require.Equal(t, req.Header
// 		// })
// 	})
// }
