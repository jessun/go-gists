package main

import (
	"crypto/md5"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	DMP_USER_HASH_TPL = "%v:%v:v4PL_o~cc!@:%v:%v" //user:password::user:password
	HMAC_SECRET       = "ActionTech-URds Secret"
	jwtSeed           = "gZCBxDRqRmCC2ihn"
)

func main() {
	tk, err := getToken("provision", "passwd")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", tk)

	user, hash, err := parseToken(tk)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user[%s]\nhash[%s]\n", user, hash)

	hash2 := getHash("provision", "passwd")
	fmt.Printf("hash2[%s]\n", hash2)
}

// generate jwt token
func getToken(user, password string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"seed": "gZCBxDRqRmCC2ihn",
		"nbf":  now.Unix(),
		"user": user,
		"hash": fmt.Sprintf("%v", md5.Sum([]byte(fmt.Sprintf(USER_HASH_TPL, user, password, user, password)))),
	})
	return token.SignedString([]byte(_HMAC_SECRET))
}

func parseToken(token string) (user, hash string, err error) {

	jwtToken, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", tk.Header["alg"])
		}
		return []byte(_HMAC_SECRET), nil
	})
	if err != nil {
		return
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		user = claims["user"].(string)
		hash = claims["hash"].(string)
	}

	return user, hash, nil
}

func getHash(a, b string) (hash string) {

	return fmt.Sprintf("%v", md5.Sum([]byte(fmt.Sprintf(USER_HASH_TPL, a, b, a, b))))
}
