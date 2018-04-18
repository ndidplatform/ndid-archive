package idp

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func CreateSignature(reqMsg string, accessorPrivk []byte) ([]byte, error) {

	// string to byte array
	message := []byte(reqMsg)

	block, _ := pem.Decode(accessorPrivk)
	if block == nil {
		fmt.Println("Error cannot create block")
		return nil, nil
	}

	// Create private key
	privk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Prepare required parametered Message - Signature
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	PSSmessage := message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	// Signed message
	signature, err := rsa.SignPSS(rand.Reader, privk, newhash, hashed, &opts)
	fmt.Println("Signature :" + base64.StdEncoding.EncodeToString(signature))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return signature, nil
}

func TestVerifySignature(signMsg []byte, reqMsg string) error {

	//string to byte array
	message := []byte(reqMsg)

	//2.Prepare public key
	accessorPubk := []byte("-----BEGIN PUBLIC KEY-----\n" +
		"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyzm4qZY1I9jFpYa/82pg" +
		"jO9lkKVhvXelPPyRfiEj4kgzezaOAlZWbv11F2JmV0cuB8RtnEgjoepbwuBxn97m" +
		"XZ45HmWLp4IES/UPNeUHvxoTtaEU78WXS+NGl/yx8ai0fhGrpnT3wEc5Wpr1R2vA" +
		"IUyCnRpIN4uYEFqCEbpxk6HB3z46bGp5FzImnF2EcMRUJr6bAS0ncq6Tqd0VNIuD" +
		"36/xcaSz4LL2QYjwohKMhK3hIRY4Nrlb9Q2Y8lQrLHuDGQ/7FGf2T76A/boVMxrh" +
		"4sKVwBCtXnCNZe91jHsNQkKacoET+PooQjWCmRfxcz6fzPhbzHYdAUspJdlate0y" +
		"3QIDAQAB" +
		"\n-----END PUBLIC KEY-----")

	block, _ := pem.Decode([]byte(accessorPubk))
	if block == nil {
		fmt.Println("Error cannot create block")
		return nil
	}

	fmt.Println("Create public key object")
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	pubk := pub.(*rsa.PublicKey)

	fmt.Println("Prepare parameter ")
	//Prepare required parameter
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	PSSmessage := message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	//Verift signature against message
	fmt.Println("Verify")
	err = rsa.VerifyPSS(pubk, newhash, hashed, signMsg, &opts)

	//Return result
	if err != nil {
		fmt.Println("Verify Signature failed !!!")
		return err
	} else {
		fmt.Println("Verify Signature successful")
		return nil
	}

}

/*
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyzm4qZY1I9jFpYa/82pg
jO9lkKVhvXelPPyRfiEj4kgzezaOAlZWbv11F2JmV0cuB8RtnEgjoepbwuBxn97m
XZ45HmWLp4IES/UPNeUHvxoTtaEU78WXS+NGl/yx8ai0fhGrpnT3wEc5Wpr1R2vA
IUyCnRpIN4uYEFqCEbpxk6HB3z46bGp5FzImnF2EcMRUJr6bAS0ncq6Tqd0VNIuD
36/xcaSz4LL2QYjwohKMhK3hIRY4Nrlb9Q2Y8lQrLHuDGQ/7FGf2T76A/boVMxrh
4sKVwBCtXnCNZe91jHsNQkKacoET+PooQjWCmRfxcz6fzPhbzHYdAUspJdlate0y
3QIDAQAB
-----END PUBLIC KEY-----*/

/*
privk := "-----BEGIN RSA PRIVATE KEY-----\n" +
	"MIIEpAIBAAKCAQEAyzm4qZY1I9jFpYa/82pgjO9lkKVhvXelPPyRfiEj4kgzezaO" +
	"AlZWbv11F2JmV0cuB8RtnEgjoepbwuBxn97mXZ45HmWLp4IES/UPNeUHvxoTtaEU" +
	"78WXS+NGl/yx8ai0fhGrpnT3wEc5Wpr1R2vAIUyCnRpIN4uYEFqCEbpxk6HB3z46" +
	"bGp5FzImnF2EcMRUJr6bAS0ncq6Tqd0VNIuD36/xcaSz4LL2QYjwohKMhK3hIRY4" +
	"Nrlb9Q2Y8lQrLHuDGQ/7FGf2T76A/boVMxrh4sKVwBCtXnCNZe91jHsNQkKacoET" +
	"+PooQjWCmRfxcz6fzPhbzHYdAUspJdlate0y3QIDAQABAoIBAFvlXwZ4oNDz3fQK" +
	"qdPlX4F7Y37z+e2WI8cfIp9ZVwOkyHrH4ZFW/0CzJfaMDWEcE8l5XmNUD6DQ++t0" +
	"WI0FW0AsIwIGww7c8Rpz1wv04/rbTNxN2czOmaq8PBjQMJrpQazcSrU7oSh2TI4X" +
	"EiDYrYmhMlLeJpfPbNTASJ7Bo0gxB5MugZVEJzmWnDrv4d5wt1HxE3hZag3k7tn7" +
	"rpMEUaj9JvySuQNSndQC/fNeCAI/LxYERH1nzv9IRJvYT3KoTSwBNV0WpoY4iyTl" +
	"IeP5SZ+hEhBlgpYQ+L+vaz4tF3czZEqGT5tOK98mq6E7OA3iYEH1uKD56vWXq8WS" +
	"0msGirUCgYEA77ahFJ9vyxL/fyGGG8iw+H3d6YTPGz10owhBOGtIS2AAEs3/FdX8" +
	"I3HwH8s/LsQACo4IRTf7uLnWZLiq8sZT5MZby9KiqTfY+SyyGB20aEk6j80yldqR" +
	"1n18czVpmR0trcPWCbbt/jtQccDXSw+wJpThzEaT3FH00a+aHnKraJ8CgYEA2Qhz" +
	"x70fD4MSbQBmOJkgRkqIsKTqzIFkOjAy6Mdn/09HYSAK+/wDxfo4dFGEeg7zQyzS" +
	"/wOTYjEWz5G4z2Yuzp4u5iW0TdZdZyacGuTXlN0EdmiPrlHmhD8eAfAGqtl6uWfA" +
	"ASFnoFp5wJ7VjFWgrp52X4EpmRqAD+hbCLFNZwMCgYEAwfDxbDz+dr99JT9bQFlk" +
	"DPfQteD0qyZSmqDQG7R4vjsCdDRkACerooXJb//Bs7VyDxgQyufbaf443i6maFtb" +
	"fsmZdVOKtPvxONAXmvlVf9ZXYgbzuVgoC/Bk4tVBBVcdIOxD3II7FSiqEg6iYSFl" +
	"NCHaeapUcLmCKLA5Mg82bSMCgYA/U/MMvkd47EMNIUeyjiPXvtnhyU1l/P905ymD" +
	"uOPEoKpwUbxCyMeFFj0w27RvNTuQR10N4ko4JaDFUnz9r5BK4+dFao0RBVLdzWtR" +
	"gaLSIieyMVJziBxeTFiCMjqP0cO4o+hnrAqjxKKwLOic+UNYkI5z0amErjHd8mvV" +
	"vn6aZwKBgQCaU3B1KY1Jgoo3NiLjz3uu769LgIbF0lwgBuP/jRScByaemPA8ts6M" +
	"xabA9d+cg2ICfQJQdkkS101Zm/Zdcop4whzQPL1imT4Y9o83t6Uc8qcwqMglyiFY" +
	"JDFYEgkt/x/+/oAyd9PTU9B2ubBU39JY+Hv7PnToEydNmhGQGd2F1w==\n" +
	"-----END RSA PRIVATE KEY-----"*/
