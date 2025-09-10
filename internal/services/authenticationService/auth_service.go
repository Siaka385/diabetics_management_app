package authentication

import (
	repository "diawise/internal/repository/userRegistrationLoginRepository"
)

type Signup_service struct {
	SignUpRepo *repository.Signup_repository
}

func NewRegistration(store *repository.Signup_repository) *Signup_service {
	return &Signup_service{SignUpRepo: store}
}

type Signin_service struct {
	SignInRepo *repository.Signin_repository
}

func NewLogin(store *repository.Signin_repository) *Signin_service {
	return &Signin_service{SignInRepo: store}
}
