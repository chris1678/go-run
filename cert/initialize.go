/*
@Date : 2022/6/3 18:45
@Author : cirss
*/
package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/chris1678/go-run/config"
	"github.com/chris1678/go-run/logger"
	"github.com/chris1678/go-run/utils"
	"io/ioutil"
	"os"
)

var privateCert []byte
var publicCert []byte

/**
 * @Description Initialize token加密证书初始化
 **/
func Initialize() {

	c := config.ApplicationConfig
	if !utils.PathExist(c.PrivateCert) &&
		!utils.PathExist(c.PublicCert) &&
		c.PrivateCert != "" &&
		c.PublicCert != "" {
		loadingPublicCert(c)
		loadingPrivateCert(c)
	}
}

/**
 * @Description loadingPublicCert
 * @Param c
 **/
func loadingPublicCert(c *config.Application) {
	file, err := os.Open(c.PublicCert)
	if err != nil {
		logger.LogHelper.Fatalf("证书文件%s打开失败", c.PublicCert)
	}
	publicCert, err = ioutil.ReadAll(file)
	if err != nil {
		logger.LogHelper.Fatalf("证书%s读取失败", c.PublicCert)
	}
}

/**
 * @Description loadingPrivateCert
 * @Param c
 **/
func loadingPrivateCert(c *config.Application) {
	file, err := os.Open(c.PrivateCert)
	if err != nil {
		logger.LogHelper.Fatalf("证书文件%s打开失败", c.PrivateCert)
	}
	privateCert, err = ioutil.ReadAll(file)
	if err != nil {
		logger.LogHelper.Fatalf("证书%s读取失败", c.PrivateCert)
	}
}

/**
 * @Description RsaEncrypt token加密
 * @Param origData
 * @Param publicKey
 * @return []byte
 * @return error
 **/
func Encode(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicCert)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

/**
 * @Description rsaDecrypt
 * @Param ciphertext
 * @Param privateKey
 * @return []byte
 * @return error
 **/
func rsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

/**
 * @Description RsaDecode token解密
 * @Param str
 * @return []byte
 * @return error
 **/
func Decode(str string) ([]byte, error) {

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(str)))
	n, err := base64.StdEncoding.Decode(dst, []byte(str))
	if err != nil {
		return []byte{}, err
	}
	dst = dst[:n]

	decrypt, err := rsaDecrypt(dst, privateCert)
	if err != nil {
		return []byte{}, err
	}

	return decrypt, nil
}
