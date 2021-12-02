// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	errors "github.com/tkeel-io/kit/errors"
	codes "google.golang.org/grpc/codes"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the ego package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

var oauth2ErrUnknown *errors.TError
var oauth2ErrSecretNotMatch *errors.TError
var oauth2ErrClientIdAlreadyExists *errors.TError
var oauth2ErrInvaildPluginId *errors.TError
var oauth2ErrInternalStore *errors.TError

func init() {
	oauth2ErrUnknown = errors.New(int(codes.Unknown), "oauth2.v1.OAUTH2_ERR_UNKNOWN", Error_OAUTH2_ERR_UNKNOWN.String())
	errors.Register(oauth2ErrUnknown)
	oauth2ErrSecretNotMatch = errors.New(int(codes.InvalidArgument), "oauth2.v1.OAUTH2_ERR_SECRET_NOT_MATCH", Error_OAUTH2_ERR_SECRET_NOT_MATCH.String())
	errors.Register(oauth2ErrSecretNotMatch)
	oauth2ErrClientIdAlreadyExists = errors.New(int(codes.AlreadyExists), "oauth2.v1.OAUTH2_ERR_CLIENT_ID_ALREADY_EXISTS", Error_OAUTH2_ERR_CLIENT_ID_ALREADY_EXISTS.String())
	errors.Register(oauth2ErrClientIdAlreadyExists)
	oauth2ErrInvaildPluginId = errors.New(int(codes.InvalidArgument), "oauth2.v1.OAUTH2_ERR_INVAILD_PLUGIN_ID", Error_OAUTH2_ERR_INVAILD_PLUGIN_ID.String())
	errors.Register(oauth2ErrInvaildPluginId)
	oauth2ErrInternalStore = errors.New(int(codes.Internal), "oauth2.v1.OAUTH2_ERR_INTERNAL_STORE", Error_OAUTH2_ERR_INTERNAL_STORE.String())
	errors.Register(oauth2ErrInternalStore)
}

func Oauth2ErrUnknown() errors.Error {
	return oauth2ErrUnknown
}

func Oauth2ErrSecretNotMatch() errors.Error {
	return oauth2ErrSecretNotMatch
}

func Oauth2ErrClientIdAlreadyExists() errors.Error {
	return oauth2ErrClientIdAlreadyExists
}

func Oauth2ErrInvaildPluginId() errors.Error {
	return oauth2ErrInvaildPluginId
}

func Oauth2ErrInternalStore() errors.Error {
	return oauth2ErrInternalStore
}
