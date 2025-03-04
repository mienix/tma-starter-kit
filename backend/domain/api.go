package domain

import "context"

type UserAPI interface {
	Start(context.Context) error
}
