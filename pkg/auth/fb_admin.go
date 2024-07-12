package auth

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

var Admin *FirebaseAdmin
var ctx = context.Background()

type FirebaseAdmin struct {
	*firebase.App
	*auth.Client
}

// RegisterUser registers a new user in Firebase with the given name, email, and password.
//
// Parameters:
// - name: the display name of the user.
// - email: the email address of the user.
// - password: the password for the user.
//
// Returns:
// - *auth.UserRecord: the user record of the newly created user.
// - error: an error if the user creation fails.
func (fa *FirebaseAdmin) RegisterUser(name string, email string, password string) (*auth.UserRecord, error) {
	userToCreate := &auth.UserToCreate{}
	userToCreate.DisplayName(name)
	userToCreate.Email(email)
	userToCreate.Password(password)

	log.Trace().Interface("UserToCreate", userToCreate).Msg("UserToCreate")

	return fa.CreateUser(ctx, userToCreate)
}

// VerifyIDToken verifies the Firebase ID token provided by the user.
//
// Parameters:
// - ctx: The context to use for the operation.
// - idToken: The Firebase ID token obtained from the user.
//
// Returns:
// - *auth.Token: The decoded Firebase ID token object if valid.
// - error: An error if the token verification fails.
func (fa *FirebaseAdmin) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	return fa.Client.VerifyIDToken(ctx, idToken)
}

// NewFirebaseAdmin creates a new instance of FirebaseAdmin using the provided AuthConfig.
//
// Parameters:
// - config: The AuthConfig containing the credentials file path.
//
// Returns:
// - *FirebaseAdmin: A pointer to the newly created FirebaseAdmin instance.
// - error: An error if there was a problem initializing the Firebase app or the authentication client.
func NewFirebaseAdmin() (*FirebaseAdmin, error) {

	var err error
	output := FirebaseAdmin{}

	creds, err := google.CredentialsFromJSON(context.Background(), []byte(cfg.FirebaseSecret))
	if err != nil {
		log.Err(err).Msg("error initializing firebase")
		return &output, err
	}

	opt := option.WithCredentials(creds)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Err(err).Msg("error initializing firebase")
		return &output, err
	}

	cli, err := app.Auth(context.Background())
	if err != nil {
		log.Err(err).Msg("error initializing firebase")
		return &output, err
	}

	output.App = app
	output.Client = cli

	Admin = &output

	log.Trace().Interface("Output", output).Msg("Firebase Admin")
	log.Debug().Msg("Initialized Firebase Admin")

	return &output, nil

}

// GetFirebaseAdmin returns a pointer to the FirebaseAdmin instance.
//
// It does not take any parameters.
// It returns a pointer to the FirebaseAdmin struct.
func GetFirebaseAdmin() *FirebaseAdmin {
	return Admin
}
