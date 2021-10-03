package emailer

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var emailPassword string
var jwtSecret string
var whiteListerEmails []string

func init() {
	emailPasswordSecret, err := getGIRCAPPPassword()
	if err != nil {
		panic(err)
	}
	emailPassword = emailPasswordSecret

	jwtSecretSecret, err := getJWTSecret()
	if err != nil {
		panic(err)
	}
	jwtSecret = jwtSecretSecret

	whiteListerEmailsSecret, err := getWhiteListerEmails()
	if err != nil {
		panic(err)
	}
	fmt.Println(whiteListerEmailsSecret)
	whiteListerEmails = whiteListerEmailsSecret
}

func getWhiteListerEmails() ([]string, error) {
	name := "projects/gircapp/secrets/WHITELIST_NOTIFICATION_EMAILS/versions/latest"
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return []string{""}, fmt.Errorf("failed to create secretmanager client: %v", err)
	}
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return []string{""}, fmt.Errorf("failed to access secret version: %v", err)
	}

	str := string(result.Payload.Data)

	resultStrSlice := strings.Split(str, ", ")
	return resultStrSlice, nil
}

func getJWTSecret() (string, error) {
	name := "projects/gircapp/secrets/EMAIL_JWT_SECRET/versions/latest"
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secretmanager client: %v", err)
	}
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %v", err)
	}
	return string(result.Payload.Data), nil
}

func getGIRCAPPPassword() (string, error) {
	name := "projects/gircapp/secrets/EMAIL_PASSWORD/versions/latest"
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secretmanager client: %v", err)
	}
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %v", err)
	}
	return string(result.Payload.Data), nil
}

func makeEmailJWT(email string, userID string, specialty string, fullName string, verified string) (jwtValue string, err error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userID,
		"email":     email,
		"specialty": specialty,
		"fullName":  fullName,
		"verified":  verified,
	})

	// Sign and get the complete encoded token as a string using the secret
	hmacSecret := []byte(jwtSecret)
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		panic(err)
	}

	return tokenString, nil
}

func SendConfirmationEmailIfVerified(emailAddress string, specialty string, fullName string, userID string) error {

	jwt, err := makeEmailJWT(emailAddress, userID, specialty, fullName, "true")
	if err != nil {
		return err
	}
	

	fullNameStrSlice := strings.Split(fullName, " ")
	firstName := fullNameStrSlice[0]
	url := fmt.Sprintf("https://girc.app/confirmemail?key=%s", jwt)
	e := email.NewEmail()
	e.From = "admin@globalirc.org"
	e.To = []string{emailAddress}
	e.Subject = fmt.Sprintf("%s, Please Confirm Your GIRC Updated Email Address", firstName)

	html := fmt.Sprintf("<a href=%s>Click here to confirm email.</a>", url)
	e.HTML = []byte(html)
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "admin@globalirc.org", emailPassword, "smtp.gmail.com"))
	return nil

}

func SendConfirmationEmailIfNotVerified(emailAddress string, userID string, fullName string, specialty string) error {

	jwt, err := makeEmailJWT(emailAddress, userID, specialty, fullName, "false")
	if err != nil {
		return err
	}

	fullNameStrSlice := strings.Split(fullName, " ")

	firstName := fullNameStrSlice[0]

	url := fmt.Sprintf("https://girc.app/confirmemail?key=%s", jwt)
	e := email.NewEmail()
	e.From = "admin@globalirc.org"
	e.To = []string{emailAddress}
	e.Subject = fmt.Sprintf("%s, Please Confirm Your GIRC Account", firstName)

	html, err := makeEmailConfirmationEmailTemplate(url)
	if err != nil {
		panic(err)
	}
	e.HTML = []byte(html)
	err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "admin@globalirc.org", emailPassword, "smtp.gmail.com"))
	if err != nil {
		return err
	}
	return nil

}

func DecodeJWTClaims(jwtValue string) (email string, userID string, specialty string, fullName string, verified string, err error) {

	hmacSecret := []byte(jwtSecret)
	token, err := jwt.Parse(jwtValue, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims["foo"], claims["nbf"])
		email := fmt.Sprintf("%s", claims["email"])
		userID := fmt.Sprintf("%s", claims["userID"])
		specialty := fmt.Sprintf("%s", claims["specialty"])
		fullName := fmt.Sprintf("%s", claims["fullName"])
		verified := fmt.Sprintf("%s", claims["verified"])

		return email, userID, specialty, fullName, verified, nil
	} else {
		return "", "", "", "", "", err
	}

}


// TODO: return error ?
func SendDirectorVerificationEmail(fullName string, specialty string, userEmail string, userID string) {
	jwt, err := makeEmailJWT(userEmail, userID, specialty, fullName, "false")
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("https://girc.app/directorverifybyemail?key=%s", jwt)
	e := email.NewEmail()
	e.From = "admin@globalirc.org"
	e.To = whiteListerEmails
	subject := fmt.Sprintf("%s Has Confirmed Their Email ", fullName)
	e.Subject = subject
	//	e.Text = []byte("Text Body is, of course, supported!")
	html := fmt.Sprintf(
		"<a href=%s>Click Here To Confirm User with email address: %s</a><div>User name: %s </div><div>User Specialty: %s</div>",
		url, userEmail, fullName, specialty)
	e.HTML = []byte(html)
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "admin@globalirc.org", emailPassword, "smtp.gmail.com"))
}

func SendUserVerificationWelcomeEmail(userEmail string) {
	// jwt, err := makeEmailJWT(userEmail, userID, "false")
	// if err != nil {
	// 	panic(err)
	// }
	//	url := fmt.Sprintf("https://girc.app/directorverifybyemail?key=%s", jwt)
	e := email.NewEmail()
	e.From = "admin@globalirc.org"
	e.To = []string{userEmail}
	e.Subject = "Welcome to the GIRC App"
	//	e.Text = []byte("Text Body is, of course, supported!")
	// html := fmt.Sprintf("<a href=%s>link text</a>", url)
	html, err := makeWelcomeEmailTemplate()
	if err != nil {
		panic(err)
	}
	e.HTML = []byte(html)
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "admin@globalirc.org", emailPassword, "smtp.gmail.com"))
}
