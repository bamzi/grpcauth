package grpcauth

import jwt "github.com/dgrijalva/jwt-go"

// Option helps chnaging Options
type Option func(*GRPCAuth) error

// OptCertificate you need this if you want to create grpc server
func OptCertificate(certPath, keyPath string) Option {
	return func(grpcauth *GRPCAuth) error {
		return grpcauth.Security.loadCertificate(certPath, keyPath)
	}
}

// OptCertificateAuthorityPublicKey you need this if you want to use token
func OptCertificateAuthorityPublicKey(certPubPath string) Option {
	return func(grpcauth *GRPCAuth) error {
		return grpcauth.Security.loadCertificateAuthorityPublicKey(certPubPath)
	}
}

// OptCertificateAuthority you need this if you want to create grpc client
func OptCertificateAuthority(caPath string) Option {
	return func(grpcauth *GRPCAuth) error {
		return grpcauth.Security.loadCertificateAuthority(caPath)
	}
}

// GRPCAuth is a base object
type GRPCAuth struct {
	Security *Security
	Grpc     *Grpc
}

// Token parse jwt token based on public key
func (x *GRPCAuth) Token(token Token) (*jwt.Token, error) {
	return x.Security.ParseJwt(token)
}

// New creates grpcauth object based on list of options
func New(options ...Option) (*GRPCAuth, error) {
	security := Security{}

	grpcauth := GRPCAuth{
		Security: &security,
		Grpc: &Grpc{
			security: &security,
		},
	}

	for _, option := range options {
		if err := option(&grpcauth); err != nil {
			return nil, err
		}
	}

	return &grpcauth, nil
}
