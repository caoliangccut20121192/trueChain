package sm2

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
)
//sm2加密算法
func Encrypt_C1C2C3(src []byte )(error,[]byte){
	//src := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	priv, pub, err := GenerateKey(rand.Reader)
	if err != nil {
		log.Error(err.Error())
		return err,nil
	}

	fmt.Printf("d:%s\n", hex.EncodeToString(priv.D.Bytes()))
	fmt.Printf("x:%s\n", hex.EncodeToString(pub.X.Bytes()))
	fmt.Printf("y:%s\n", hex.EncodeToString(pub.Y.Bytes()))

	cipherText, err := Encrypt(pub, src, C1C2C3)
	if err != nil {
		log.Error(err.Error())
		return err,nil
	}
	fmt.Printf("cipher text:%s\n", hex.EncodeToString(cipherText))
    return nil,cipherText
	//plainText, err := Decrypt(priv, cipherText, C1C2C3)
	//if err != nil {
	//	t.Error(err.Error())
	//	return
	//}
	//fmt.Printf("plain text:%s\n", hex.EncodeToString(plainText))
	//
	//if !bytes.Equal(plainText, src) {
	//	t.Error("decrypt result not equal expected")
	//	return
	//}
}
func Decrypt_C1C2C3(src []byte )(error,[]byte) {

}