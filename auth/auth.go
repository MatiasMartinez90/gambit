package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJson struct {
	Sub       string
	Event_id  string
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
	fmt.Println("Imprimo el token")
	fmt.Println(parts)

	if len(parts) != 3 {
		fmt.Println("El token no es valido")
		return false, nil, "El token no es valido"
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])

	if err != nil {
		fmt.Println("No se puede decodificar la parte del token :", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJson
	err = json.Unmarshal(userInfo, &tkj)

	if err != nil {
		fmt.Println("No se puede decodificar en la estructura JSON ", err.Error())
		return false, err, err.Error()
	}

	ahora := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(ahora) {
		fmt.Println("Fecha de exiracion token = " + tm.String())
		fmt.Println("Tokenexpirado !")
		return false, err, "Token expirado :) :)"
	}

	return true, nil, string(tkj.Username)
}
