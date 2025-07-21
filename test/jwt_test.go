package test

import (
	"fmt"
	"server/utils"
	"testing"
)

func TestJWT(t *testing.T) {
	// defer cmd.Clean()
	// cmd.Start()

	token, _ := utils.GenerateToken(1, "jimzhang")
	fmt.Println("token: ", token)

	iJwtCustClaims, _ := utils.ParseToken(token)
	fmt.Println("iJwtCustClaims: ", iJwtCustClaims)

	fmt.Println(utils.IsVaildToken(token))
	fmt.Println(utils.IsVaildToken(token + "asdzxqweggcgxcasd"))
	// iJwtCustClaims2, err := utils.ParseToken(token + "asdzxqweggcgxcasd")
	// if err != nil {
	// 	// fmt.Println("err: ", err)
	// 	panic(err.Error())
	// }
	// fmt.Println("iJwtCustClaims2: ", iJwtCustClaims2)
}

// func TestParseToken(t *testing.T){

// }
