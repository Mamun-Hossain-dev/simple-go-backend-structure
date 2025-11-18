package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/**
import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// helper: base64Url encode
func base64UrlEncode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

// helper: base64Url decode
func base64UrlDecode(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}

type Claims struct {
	Sub  string `json:"sub"`
	Role string `json:"role"`
	Name string `json:"name"`
	Exp  int64  `json:"exp"`
	Iat  int64  `json:"iat"`
}

func CreateJWT(claims Claims, secret []byte) (string, error) {
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}

	// Encode herder and payload
	hdrJson, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	payloadJson, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	encodedHeader := base64UrlEncode(hdrJson)
	encodedPayload := base64UrlEncode(payloadJson)
	signingInput := encodedHeader + "." + encodedPayload

	// Signature = HMAC-SHA256(signingInput, secret)
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(signingInput))
	signature := mac.Sum(nil)
	encodedSig := base64UrlEncode(signature)

	// Final JWT: header.payload.signature
	token := signingInput + "." + encodedSig
	return token, nil
}

func VerifyJWT(token string, secret []byte) (Claims, error) {
	var claims Claims
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return claims, errors.New("invalid token format")
	}

	signingInput := parts[0] + "." + parts[1]
	sigBytes, err := base64UrlDecode(parts[2])
	if err != nil {
		return claims, err
	}

	// verify signature
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(signingInput))
	expectedSigBytes := mac.Sum(nil)
	if !hmac.Equal(sigBytes, expectedSigBytes) {
		return claims, errors.New("invalid signature")
	}

	// Decode payload
	payloadJson, err := base64UrlDecode(parts[1])
	if err != nil {
		return claims, err
	}

	if err := json.Unmarshal(payloadJson, &claims); err != nil {
		return claims, err
	}

	// check expiry
	if time.Now().Unix() > claims.Exp {
		return claims, errors.New("token expired, please generate new token")
	}

	return claims, nil
}
 **/

type Claims struct {
	Sub  string `json:"sub"`
	Role string `json:"role"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func CreateToken(secret []byte, userId, role, name string) (string, error) {
	claims := Claims{
		Sub:  userId,
		Role: role,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func VerifyJWT(tokenString string, secret []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	// cast claims
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
