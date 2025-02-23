package crypto

type AESSealer struct{}

func (*AESSealer) Seal(message, key string) (encoded string, err error) {
	enckey := [32]byte{}
	copy(enckey[:], key)
	return AESEncrypt(message, &enckey)
}

func (*AESSealer) Unseal(message, key string) (decoded string, err error) {
	enckey := [32]byte{}
	copy(enckey[:], key)
	return AESDecrypt(message, &enckey)
}
