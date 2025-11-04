package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT 키
var jwtSecret = []byte("super-secret-key")

// 토큰 구조체
type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

// 로그인
func LoginService(id, pw string) (*TokenPair, error) {
	// TODO: DB 연결 후 실제 사용자 인증 로직 추가
	user, err := GetUserByID(id)
	if err != nil {
		return nil, errors.New("서버 오류: 사용자 조회 실패")
	}
	if user == nil {
		return nil, errors.New("존재하지 않는 사용자입니다")
	}
	if user.Password != pw {
		return nil, errors.New("비밀번호가 올바르지 않습니다")
	}

	// AccessToken (15분)
	accessToken, err := createToken(id, 15*time.Minute)
	if err != nil {
		return nil, err
	}

	// RefreshToken (24시간)
	refreshToken, err := createToken(id, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// JWT 생성
func createToken(id string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(duration).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWT 검증
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("잘못된 서명 방식")
		}
		return jwtSecret, nil
	})
}
