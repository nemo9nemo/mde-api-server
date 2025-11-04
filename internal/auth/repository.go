package auth

type User struct {
	ID       string
	Password string
}

// TODO: 실제 DB 연결 로직 필요
var dummyUsers = map[string]string{
	"admin": "1234",
	"user1": "abcd",
}

// 사용자 조회
// TODO: DB연결 후 쿼리로 대체해야 함.
func GetUserByID(id string) (*User, error) {
	pw, ok := dummyUsers[id]
	if !ok {
		return nil, nil
	}

	return &User{ID: id, Password: pw}, nil
}
