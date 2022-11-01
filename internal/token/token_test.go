package token

import (
	"fmt"
	"testing"
)

func TestGenTOken(t *testing.T) {
	s, err := GenToken("admin")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Printf("token: %s\n", s)
}

func TestParseToken(t *testing.T) {
	cc, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiaXNzIjoibXktcHJvamVjdCIsImV4cCI6MTY2NzkwMTEyM30.9vYa15tjVU6oEirXiTwqnEInRPRSDqXoVkEgitEG3Sk")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Printf("cc: %s", cc)
}
