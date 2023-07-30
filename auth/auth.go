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

	// esta parts2 la agregue yo con la primer parte del token hardcodeado, por que el parts original falla. Tambien comente abajo la verificacion del a hora del token
	//parts2 := "eyJraWQiOiJEMHpcL3VVK0tOdUpCb1JtSlBLd3hCZVhBNThzeHBlVU9DanFGV1dDK0JpTT0iLCJhbGciOiJSUzI1NiJ9"
	//parts2 := "eyJhdF9oYXNoIjoiUlhzRThmVFR0YlpNSWU4dzVRcHRkdyIsInN1YiI6Ijk0MzgyNGY4LTYwZTEtNzBhNi0zODFlLWE4YjhhNjU0OWVlZCIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV9EdmhKNjJJZHoiLCJjb2duaXRvOnVzZXJuYW1lIjoiOTQzODI0ZjgtNjBlMS03MGE2LTM4MWUtYThiOGE2NTQ5ZWVkIiwiYXVkIjoiNTNmYzQ0OTVvdWtqajA0dWRtdmpydmZiMWQiLCJ0b2tlbl91c2UiOiJpZCIsImF1dGhfdGltZSI6MTY4OTIwNDk4NSwiZXhwIjoxNjg5MjkxMzg1LCJpYXQiOjE2ODkyMDQ5ODUsImp0aSI6IjBiZjkyODQ2LTMwZTMtNGQ2Yy05MTI4LTQwMGRmMmY4ZTk1MCIsImVtYWlsIjoibWF0aWFzLm1hcnRpbmV6OTArODdAZ21haWwuY29tIn0"

	if len(parts) != 3 {
		fmt.Println("El token no es válido")
		return false, nil, "El token no es válido"
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	//userInfo, err := base64.StdEncoding.DecodeString(parts2)
	if err != nil {
		fmt.Println("No se puede decodificar la parte del token :", err.Error())
		fmt.Println("Funca TURRITO")
		fmt.Println("Imprimo parts1:")
		fmt.Println(parts[1])
		fmt.Println("Imprimo parts2:")
		fmt.Println("PERRRITO MALVADOOOO:")
		fmt.Println(parts[2])
		fmt.Println("Imprimo parts:")
		fmt.Println(parts)
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

	if tm.Before(ahora) {
		fmt.Println("Fecha expiración token = " + tm.String())
		fmt.Println("Token expirado !")
		return false, err, "Token expirado !!"
	}

	return true, nil, string(tkj.Username)
}
