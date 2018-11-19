package auth

import (
	"testing"
)

func TestDecrypt(t *testing.T) {
	session_key := "1df09d0a1677dd72b8325aec59576e0c"
	iv := "1df09d0a1677dd72b8325Q=="
	ciphertext := "OpCoJgs7RrVgaMNDixIvaCIyV2SFDBNLivgkVqtzq2GC10egsn+PKmQ/+5q+chT8xzldLUog2haTItyIkKyvzvmXonBQLIMeq54axAu9c3KG8IhpFD6+ymHocmx07ZKi7eED3t0KyIxJgRNSDkFk5RV1ZP2mSWa7ZgCXXcAbP0RsiUcvhcJfrSwlpsm0E1YJzKpYy429xrEEGvK+gfL+Cw=="
	app_key := "y2dTfnWfkx2OXttMEMWlGHoB1KzMogm7"

	_, err := Decrypt(ciphertext, session_key, iv, app_key)
	if err != nil {
		t.Error(err)
	}

}
