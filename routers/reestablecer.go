package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	"github.com/DanMurguia/TT_FreeJobs/bd"
	"github.com/DanMurguia/TT_FreeJobs/jwt"
	"github.com/DanMurguia/TT_FreeJobs/models"
)

/*Reestablecer obtiene un enlace para reestablecer contraseña */
func Reestablecer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario
	var smtpServer string
	from := "your_email@example.com"
	password := "your_password"
	to := t.Email

	if strings.Contains(to, "gmail.com") {
		smtpServer = "smtp.gmail.com"
	} else if strings.Contains(to, "yahoo.com") {
		smtpServer = "smtp.mail.yahoo.com"
	} else if strings.Contains(to, "hotmail.com") || strings.Contains(to, "outlook.com") {
		smtpServer = "smtp-mail.outlook.com"
	} else {
		// Handle error or assign a default SMTP server
	}

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Email Invalido"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}
	documento, existe := bd.IntentoReestablecer(t.Email)
	if !existe {
		http.Error(w, "Usuario inválido", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar general el Token correspondiente "+err.Error(), 400)
		return
	}

	subject := "Reestablecer contraseña"
	url := fmt.Sprintf("http:localhost:8080/update-password/%s", jwtKey)
	body := "Enlace para reestablecer su contraseña de FreeJobs: " + url
	message := []byte("Subject: " + subject + "\r\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpServer)

	err = smtp.SendMail(smtpServer+":587", auth, from, []string{to}, message)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email sent!")

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
