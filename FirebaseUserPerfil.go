package clases

type FirebaseUserPerfil struct {
	Kind          string     `json:"kind"`
	FederatedId   string     `json:"federatedId"`
	ProviderId    string     `json:"providerId"`
	LocalId       string     `json:"localId"`
	EmailVerified string     `json:"emailVerified"`
	Email         string     `json:"email"`
	OauthIdToken  string     `json:"oauthIdToken"`
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	FullName      string     `json:"fullName"`
	DisplayName   string     `json:"displayName"`
	IdToken        string     `json:"IdToken"`
	PhotoUrl      string     `json:"photoUrl"`
	RefreshToken  string     `json:"refreshToken"`
	OauthExpireIn string     `json:"oauthExpireIn"`
	ExpiresIn     string     `json:"expiresIn"`
	RawUserInfo   string     `json:"rawUserInfo"`
	Errores       ClsErrores `json:"Errores"`
}

type FirebaseVerifyAssertion struct {
	PostBody            string `json:"postBody"`
	RequestUri          string `json:"requestUri"`
	ReturnIdpCredential bool   `json:"returnIdpCredential"`
	ReturnSecureToken   bool   `json:"returnSecureToken"`
}
