package idp

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func RequestUserAccept(c echo.Context) error {
	fmt.Println("Received User Accept Request")

	// Binding request
	ReqToUser := new(Request)
	if err := c.Bind(ReqToUser); err != nil {
		return err
	}

	// Request to IDP API

	fmt.Println("NameSpace :" + ReqToUser.NameSpace)
	fmt.Println("Identifier :" + ReqToUser.Identifier)
	fmt.Println("RequestMessage :" + ReqToUser.RequestMessage)
	fmt.Println("RequestId :" + ReqToUser.RequestId)
	fmt.Println("End Request User Accept")
	return c.JSON(http.StatusCreated, nil)
}

func ResponseUserAccept(c echo.Context) error {

	fmt.Println("Received User Accept Response From IDP API")
	reqMsg := "Request bookbank"
	//Create signature
	//Get accessor private key from local key store
	privk := []byte("-----BEGIN RSA PRIVATE KEY-----\n" +
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
		"JDFYEgkt/x/+/oAyd9PTU9B2ubBU39JY+Hv7PnToEydNmhGQGd2F1w==" +
		"\n-----END RSA PRIVATE KEY-----")

	signature, err := CreateSignature(reqMsg, privk)
	if err != nil {
		fmt.Println(err)
	}

	//Verify signature
	TestVerifySignature(signature, reqMsg)

	//Mock up reponse message to IDP API
	var body Reponse
	body.RequestId = "ef6f4c9c-818b-42b8-8904-3d97c4c520f6"
	body.Namespace = "citizenid"
	body.Identifier = "01234567890123"
	body.Secret = "MAGIC"
	body.Loa = 3
	body.Approval = "CONFIRM"
	body.Signature = signature
	body.AccessorId = "12a8f328-53da-4d51-a927-3cc6d3ed3feb"

	fmt.Println("End Request User Accept")
	return c.JSON(http.StatusCreated, body)
}
