package pg

import (
	"fmt"
	"testing"
)

const testuserID = "1234567890"

var testencryptedUserID string

func TestEncryption(t *testing.T) {
	encrypted, err := encryptUserID(testUserID)
	if err != nil {
		panic(err)
	}
	fmt.Println(encrypted)

	// for _, v := range encrypted {
	// 	fmt.Println(v)
	// }

	decrypted, err := decryptUserID(encrypted)
	if err != nil {
		panic(err)
	}

	if testUserID != decrypted {
		t.Errorf("keys do not match, \n %v \n %v", testUserID, decrypted)
	}
}
