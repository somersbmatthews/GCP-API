package emailer

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"github.com/gircapp/api/pg"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
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
	whiteListerEmails = whiteListerEmailsSecret
}

func getWhiteListerEmails() ([]string, error) {
	name := "projects/gircapp/secrets/EMAIL_JWT_SECRET/versions/latest"
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
	name := "projects/gircapp/secrets/GIRCAPP3_EMAILPASSWORD/versions/latest"
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

func makeEmailJWT(email string, userID string, verified string) (jwtValue string, err error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"email":  email,
		"verified": verified,
	})

	// Sign and get the complete encoded token as a string using the secret
	hmacSecret := []byte(jwtSecret)
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		panic(err)
	}

	return tokenString, nil
}

func SendConfirmationEmail(emailAddress string, userID string) (err error) {
	ctx := context.Background()
	// TODO: make distinction between new non-director verified user and registered user updating his email
	user, ok := pg.GetMedicalExpert(ctx, userID)
	if !user.Verified {
		jwt, err := makeEmailJWT(emailAddress, userID, "true")
		if err != nil {
			return err
		}
		url := fmt.Sprintf("https://girc.app/confirmemail?key=%s", jwt)
		e := email.NewEmail()
		e.From = "admin@globalirc.org"
		e.To = []string{emailAddress}
		e.Subject = "Please Confirm Your Email Address"
		//	e.Text = []byte("Text Body is, of course, supported!")
		html := fmt.Sprintf("<a href=%s>Click here to confirm email.</a>", url)
		e.HTML = []byte(html)
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "gircapp3@gmail.com", emailPassword, "smtp.gmail.com"))

		return nil
	} else { // then user is verified
		ctx := context.Background()
		// TODO: make distinction between new non-director verified user and registered user updating his email
		_, ok := pg.GetMedicalExpert(ctx , userID)
		if !ok {
		jwt, err := makeEmailJWT(emailAddress, userID, "false")
		if err != nil {
			return err
		}
		url := fmt.Sprintf("https://girc.app/confirmemail?key=%s", jwt)
		e := email.NewEmail()
		e.From = "admin@globalirc.org"
		e.To = []string{emailAddress}
		e.Subject = "Please Confirm Your Email Address"
		//	e.Text = []byte("Text Body is, of course, supported!")
		html := fmt.Sprintf("<a href=%s>Click here to confirm email</a>", url)
		e.HTML = []byte(html)
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "gircapp3@gmail.com", emailPassword, "smtp.gmail.com"))

		return nil

	}
}

func DecodeJWT(jwtValue string) (email string, userID string, newUser bool, err error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	hmacSecret := []byte(jwtSecret)
	token, err := jwt.Parse(jwtValue, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims["foo"], claims["nbf"])
		email := fmt.Sprintf("%s", claims["email"])
		userID := fmt.Sprintf("%s", claims["userID"])
		newUser := fmt.Sprintf("%s", claims["newUser"])
		
		return email, userID, newUser, nil
	} else {
		return "", "", err
	}

}

func SendDirectorVerificationEmail(userName string, userEmail string, userID string) {
	jwt, err := makeEmailJWT(userEmail, userID)
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("https://girc.app/directorverifybyemail?key=%s", jwt)
	e := email.NewEmail()
	e.From = "admin@globalirc.org"
	e.To = whiteListerEmails
	e.Subject = "A New User Has Confirmed Their Email"
	//	e.Text = []byte("Text Body is, of course, supported!")
	html := fmt.Sprintf("<a href=%s>link text</a>", url)
	e.HTML = []byte(html)
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "gircapp3@gmail.com", emailPassword, "smtp.gmail.com"))
}
