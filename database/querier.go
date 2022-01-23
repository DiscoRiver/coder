// Code generated by sqlc. DO NOT EDIT.

package database

import (
	"context"
)

type querier interface {
	GetAPIKeyByID(ctx context.Context, id string) (APIKey, error)
	GetOrganizationByName(ctx context.Context, name string) (Organization, error)
	GetOrganizationsByUserID(ctx context.Context, userID string) ([]Organization, error)
	GetUserByEmailOrUsername(ctx context.Context, arg GetUserByEmailOrUsernameParams) (User, error)
	GetUserByID(ctx context.Context, id string) (User, error)
	GetUserCount(ctx context.Context) (int64, error)
	InsertAPIKey(ctx context.Context, arg InsertAPIKeyParams) (APIKey, error)
	InsertOrganization(ctx context.Context, arg InsertOrganizationParams) (Organization, error)
	InsertOrganizationMember(ctx context.Context, arg InsertOrganizationMemberParams) (OrganizationMember, error)
	InsertUser(ctx context.Context, arg InsertUserParams) (User, error)
	UpdateAPIKeyByID(ctx context.Context, arg UpdateAPIKeyByIDParams) error
}

var _ querier = (*sqlQuerier)(nil)
