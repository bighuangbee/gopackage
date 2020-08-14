package jwtService

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

/**
* @Author: bigHuangBee
* @Date: 2020/4/24 22:36
 */

//const JWT_Encrtpy = "1be9261239972389b9c3e9e653664e93a"

const JWT_Encrtpy = "jlkwdaa"

const JWT_KEY_SYS = "sysUser"
const JWT_KEY_DUTY = "dutyUser"
const JWT_KEY_DRONE = "droneUser"

const USER_TYPE_MOINTOR = 1
const USER_TYPE_DUTY = 2
const USER_TYPE_DRONE = 3

var USER_TYPE  = map[int]string{
	USER_TYPE_MOINTOR: JWT_KEY_SYS,
	USER_TYPE_DUTY: JWT_KEY_DUTY,
	USER_TYPE_DRONE: JWT_KEY_DRONE,
}

type UserClaims struct {
	UserName string `json:"username"`
	UserType int8 `json:"userType"`
	CompanyId int32 `json:"company_id"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

type UserJwt struct{
	Type 		string	//用户类型
	Encrtpy 	string	//密钥
	TokenKey 	string	//token索引
	InValidTokenKey string	//失效token索引
}

func NewUserJwt(userType string) *UserJwt{

	return &UserJwt{
		//Type:            userType,
		Encrtpy:         userType + JWT_Encrtpy,
		TokenKey:        userType + ":token_%s",
		InValidTokenKey: userType + "TokenInvalid:%s",
	}
}

func (user *UserJwt) EncodeToken(claims *UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(user.Encrtpy))
}

func (user *UserJwt) DecodeToken(tokenStr string) (*UserClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(user.Encrtpy), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// token索引键
func (user *UserJwt)CreateTokenKey(userName string) (string){
	return fmt.Sprintf(user.TokenKey, userName)
}

