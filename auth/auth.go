package auth

import (
	"context"
	"os"

	"github.com/nedpals/supabase-go"
)

const BaseAuthURL = "https://bvutljexdthpferjwsgz.supabase.co/auth/v1/recover"

// https://<project_ref>.supabase.co/rest/v1/
const ResetPasswordEndpoint = "auth/v1/recover"

var Client *supabase.Client

func Init() error {
	sbHost := os.Getenv("SUPABASE_URL_TEST")
	if os.Getenv("ENV") == "production" {
		sbHost = os.Getenv("SUPABASE_URL_PROD")
	}

	sbSecret := os.Getenv("SUPABASE_SECRET_TEST")
	if os.Getenv("ENV") == "production" {
		sbSecret = os.Getenv("SUPABASE_SECRET_PROD")
	}

	Client = supabase.CreateClient(sbHost, sbSecret)
	return nil
}

func Signup(ctx context.Context, email, password string) (*supabase.AuthenticatedDetails, error) {
	credentials := supabase.UserCredentials{
		Email:    email,
		Password: password,
	}
	_, err := Client.Auth.SignUp(ctx, credentials)
	if err != nil {
		return nil, err
	}
	authUser, err := Client.Auth.SignIn(ctx, credentials)
	if err != nil {
		return nil, err
	}
	return authUser, err
}
