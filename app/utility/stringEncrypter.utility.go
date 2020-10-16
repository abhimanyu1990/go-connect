package utility

import(
	"crypto/aes"
    "crypto/cipher"
	"crypto/rand"
	"io"
	"encoding/hex"
)

var logger = Logger()

func Encrypt(key, text []byte) (string, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    ciphertext := make([]byte, aes.BlockSize + len(text))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)
	return hex.EncodeToString(ciphertext), nil
}
