package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidoToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")

	parts2 := "eyJraWQiOiJEMHpcL3VVK0tOdUpCb1JtSlBLd3hCZVhBNThzeHBlVU9DanFGV1dDK0JpTT0iLCJhbGciOiJSUzI1NiJ9"

	if len(parts) != 3 {
		fmt.Println("El token no es válido")
		return false, nil, "El token no es válido"
	}

	//userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	userInfo, err := base64.StdEncoding.DecodeString(parts2)
	if err != nil {
		fmt.Println("No se puede decodificar la parte del token :", err.Error())
		fmt.Println("Funca TURRITO")
		fmt.Println(parts[2])
		fmt.Println("Imprimo parts 1:")
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("No se puede decodificar la estructura JSON ", err.Error())
		return false, err, err.Error()
	}

	ahora := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	//if tm.Before(ahora) {
	//	fmt.Println("Fecha expiración token = " + tm.String())
	//	fmt.Println("Token expirado !")
	//	return false, err, "Token expirado !!"
	//}

	return true, nil, string(tkj.Username)
}
