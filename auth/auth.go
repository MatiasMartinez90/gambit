package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Estructura que viaja en el token
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
	fmt.Println("Ingresando a la funcion ValidoToken")
	parts := strings.Split(token, ".")
	parts2 := "eyJzdWIiOiI5NDM4MjRmOC02MGUxLTcwYTYtMzgxZS1hOGI4YTY1NDllZWQiLCJ0b2tlbl91c2UiOiJhY2Nlc3MiLCJzY29wZSI6ImF3cy5jb2duaXRvLnNpZ25pbi51c2VyLmFkbWluIG9wZW5pZCBlbWFpbCIsImF1dGhfdGltZSI6MTY5MjMxNTg0NSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLnVzLWVhc3QtMS5hbWF6b25hd3MuY29tXC91cy1lYXN0LTFfRHZoSjYySWR6IiwiZXhwIjoxNjkyNDAyMjQ1LCJpYXQiOjE2OTIzMTU4NDUsInZlcnNpb24iOjIsImp0aSI6IjNlZmE5MzRkLTY5NjMtNDBmNy04NDM0LTQzZjZmM2FhOTU2MCIsImNsaWVudF9pZCI6IjUzZmM0NDk1b3VramowNHVkbXZqcnZmYjFkIiwidXNlcm5hbWUiOiI5NDM4MjRmOC02MGUxLTcwYTYtMzgxZS1hOGI4YTY1NDllZWQifQ"

	if len(parts) != 3 {
		fmt.Println("El token no es valido")
		return false, nil, "El token no es valido"
	}

	//userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	userInfo, err := base64.StdEncoding.DecodeString(parts2)
	if err != nil {
		fmt.Println("No se puede decodificar la parte del token : ", err.Error())
		fmt.Println(parts2)
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("No se puede decodificar en la estructura JSON", err.Error())
		return false, err, err.Error()
	}

	ahora := time.Now()
	//creo una variable para poder comparar las fechas
	tm := time.Unix(int64(tkj.Exp), 0)

	//Funcion para comparar fechas
	if tm.Before(ahora) {
		fmt.Println("Fecha de expiracion token" + tm.String())
		fmt.Println("Token expirado !")
		return false, err, "Token expierado !!"
	}

	return true, nil, string(tkj.Username)

}
