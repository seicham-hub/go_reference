//トークンの生成

app.Post("/signin", func(ctx iris.Context) {
	claims := UserClaims{
		Username: "kataras",
	}

	token, err := signer.Sign(claims)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}

	//最後にクライアントに送る
	ctx.Write(token)
})

//トークンの認証

verifier := jwt.NewVerifier(jwt.HS256, []byte("secret"))
verifyMiddleware := verifier.Verify(func() interface{} {
    return new(UserClaims) // must return a pointer to the claims struct type.
})

//ミドルウェアをルートパーティに登録する,これはトークンがこちらで発行されたものかどうかチェック？
protected := app.Party("/protected")
protected.Use(verifyMiddleware)

//ミドルウェアを単独ルートに登録する。
app.Get("/todos", verifyMiddleware, getTodos)

// 暗号化した内容(claims)を取得する。
protected.Get("/todos", func(ctx iris.Context) {
    claims := jwt.Get(ctx).(*UserClaims)
    ctx.WriteString(claims.Username)
})

// 7.1 Get the Context User, if and when the claims implements one or more Context.User method
user := ctx.User()
username, _ := user.GetUsername()

//  Get the Verified Token information:
verifiedToken := jwt.GetVerifiedToken(ctx)
verifiedToken.Token // the original request token.

// The VerifiedToken looks like this:
type VerifiedToken struct {
    Token          []byte // The original token.
    Header         []byte // The header (decoded) part.
    Payload        []byte // The payload (decoded) part.
    Signature      []byte // The signature (decoded) part.
    StandardClaims Claims // Any standard claims extracted from the payload.
}


func (u *Clients) GenerateTokenPair(ctx iris.Context) *jwt.TokenPair {
	if u.ID == 0 {
		return nil
	} else {
		tempID := strconv.Itoa(int(u.ID))
		refreshClaims := jwt.Claims{Subject: tempID}

		accessClaims := UserClaims{
			ID:   tempID,
			Mail: u.Mail,
			exp:  time.Now().UTC().Add(time.Hour * time.Duration(1)).Unix(),
			iat:  time.Now().UTC().Unix(),
		}
		tokenPair, err := signer.NewTokenPair(accessClaims, refreshClaims, refreshTokenMaxAge)
		if err != nil {
			ctx.Application().Logger().Errorf("token pair: %v", err)
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return nil
		}
		u.AccessToken, _ = strconv.Unquote(string(tokenPair.AccessToken))
		u.RefreshToken, _ = strconv.Unquote(string(tokenPair.RefreshToken))

		ups := map[string]interface{}{"access_token": u.AccessToken, "refresh_token": u.RefreshToken}
		if err := Update(&Clients{}, ups, u.ID); err != nil {
			log.Printf("Token保存失敗：%+v\n", u.ID)
			log.Printf("Token保存失敗：%+v\n", err)
		}
		return &tokenPair
	}
}