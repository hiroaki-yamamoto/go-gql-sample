// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package pub

import (
	"fmt"
	"io"
	"strconv"

	"github.com/hiroaki-yamamoto/go-gql-sample/backend/prisma"
)

type AggregateUser struct {
	Count int `json:"count"`
}

type BatchPayload struct {
	Count string `json:"count"`
}

type Error struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

type Mutation struct {
	CreateUser      prisma.User  `json:"createUser"`
	UpdateUser      *prisma.User `json:"updateUser"`
	UpdateManyUsers BatchPayload `json:"updateManyUsers"`
	UpsertUser      prisma.User  `json:"upsertUser"`
	DeleteUser      *prisma.User `json:"deleteUser"`
	DeleteManyUsers BatchPayload `json:"deleteManyUsers"`
}

type Node interface {
	IsNode()
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type Query struct {
	User            *prisma.User   `json:"user"`
	Users           []*prisma.User `json:"users"`
	UsersConnection UserConnection `json:"usersConnection"`
	Node            Node           `json:"node"`
}

type Session struct {
	ID string `json:"id"`
}

func (Session) IsSessionAndStatus() {}

type SessionAndStatus interface {
	IsSessionAndStatus()
}

type Status struct {
	Ok     bool    `json:"ok"`
	Errors []Error `json:"errors"`
}

func (Status) IsUserAndStatus()    {}
func (Status) IsSessionAndStatus() {}

type UserAndStatus interface {
	IsUserAndStatus()
}

type UserConnection struct {
	PageInfo  PageInfo      `json:"pageInfo"`
	Edges     []*UserEdge   `json:"edges"`
	Aggregate AggregateUser `json:"aggregate"`
}

type UserCreateInput struct {
	ID       *string `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
}

type UserEdge struct {
	Node   prisma.User `json:"node"`
	Cursor string      `json:"cursor"`
}

type UserPreviousValues struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSubscriptionPayload struct {
	Mutation       MutationType        `json:"mutation"`
	Node           *prisma.User        `json:"node"`
	UpdatedFields  []string            `json:"updatedFields"`
	PreviousValues *UserPreviousValues `json:"previousValues"`
}

type UserSubscriptionWhereInput struct {
	MutationIn                 []MutationType               `json:"mutation_in"`
	UpdatedFieldsContains      *string                      `json:"updatedFields_contains"`
	UpdatedFieldsContainsEvery []string                     `json:"updatedFields_contains_every"`
	UpdatedFieldsContainsSome  []string                     `json:"updatedFields_contains_some"`
	Node                       *UserWhereInput              `json:"node"`
	AND                        []UserSubscriptionWhereInput `json:"AND"`
	OR                         []UserSubscriptionWhereInput `json:"OR"`
	NOT                        []UserSubscriptionWhereInput `json:"NOT"`
}

type UserUpdateInput struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type UserUpdateManyMutationInput struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type UserWhereInput struct {
	ID                    *string          `json:"id"`
	IDNot                 *string          `json:"id_not"`
	IDIn                  []string         `json:"id_in"`
	IDNotIn               []string         `json:"id_not_in"`
	IDLt                  *string          `json:"id_lt"`
	IDLte                 *string          `json:"id_lte"`
	IDGt                  *string          `json:"id_gt"`
	IDGte                 *string          `json:"id_gte"`
	IDContains            *string          `json:"id_contains"`
	IDNotContains         *string          `json:"id_not_contains"`
	IDStartsWith          *string          `json:"id_starts_with"`
	IDNotStartsWith       *string          `json:"id_not_starts_with"`
	IDEndsWith            *string          `json:"id_ends_with"`
	IDNotEndsWith         *string          `json:"id_not_ends_with"`
	Username              *string          `json:"username"`
	UsernameNot           *string          `json:"username_not"`
	UsernameIn            []string         `json:"username_in"`
	UsernameNotIn         []string         `json:"username_not_in"`
	UsernameLt            *string          `json:"username_lt"`
	UsernameLte           *string          `json:"username_lte"`
	UsernameGt            *string          `json:"username_gt"`
	UsernameGte           *string          `json:"username_gte"`
	UsernameContains      *string          `json:"username_contains"`
	UsernameNotContains   *string          `json:"username_not_contains"`
	UsernameStartsWith    *string          `json:"username_starts_with"`
	UsernameNotStartsWith *string          `json:"username_not_starts_with"`
	UsernameEndsWith      *string          `json:"username_ends_with"`
	UsernameNotEndsWith   *string          `json:"username_not_ends_with"`
	Password              *string          `json:"password"`
	PasswordNot           *string          `json:"password_not"`
	PasswordIn            []string         `json:"password_in"`
	PasswordNotIn         []string         `json:"password_not_in"`
	PasswordLt            *string          `json:"password_lt"`
	PasswordLte           *string          `json:"password_lte"`
	PasswordGt            *string          `json:"password_gt"`
	PasswordGte           *string          `json:"password_gte"`
	PasswordContains      *string          `json:"password_contains"`
	PasswordNotContains   *string          `json:"password_not_contains"`
	PasswordStartsWith    *string          `json:"password_starts_with"`
	PasswordNotStartsWith *string          `json:"password_not_starts_with"`
	PasswordEndsWith      *string          `json:"password_ends_with"`
	PasswordNotEndsWith   *string          `json:"password_not_ends_with"`
	AND                   []UserWhereInput `json:"AND"`
	OR                    []UserWhereInput `json:"OR"`
	NOT                   []UserWhereInput `json:"NOT"`
}

type UserWhereUniqueInput struct {
	ID       *string `json:"id"`
	Username *string `json:"username"`
}

type MutationType string

const (
	MutationTypeCreated MutationType = "CREATED"
	MutationTypeUpdated MutationType = "UPDATED"
	MutationTypeDeleted MutationType = "DELETED"
)

func (e MutationType) IsValid() bool {
	switch e {
	case MutationTypeCreated, MutationTypeUpdated, MutationTypeDeleted:
		return true
	}
	return false
}

func (e MutationType) String() string {
	return string(e)
}

func (e *MutationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MutationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MutationType", str)
	}
	return nil
}

func (e MutationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserOrderByInput string

const (
	UserOrderByInputIdAsc        UserOrderByInput = "id_ASC"
	UserOrderByInputIdDesc       UserOrderByInput = "id_DESC"
	UserOrderByInputUsernameAsc  UserOrderByInput = "username_ASC"
	UserOrderByInputUsernameDesc UserOrderByInput = "username_DESC"
	UserOrderByInputPasswordAsc  UserOrderByInput = "password_ASC"
	UserOrderByInputPasswordDesc UserOrderByInput = "password_DESC"
)

func (e UserOrderByInput) IsValid() bool {
	switch e {
	case UserOrderByInputIdAsc, UserOrderByInputIdDesc, UserOrderByInputUsernameAsc, UserOrderByInputUsernameDesc, UserOrderByInputPasswordAsc, UserOrderByInputPasswordDesc:
		return true
	}
	return false
}

func (e UserOrderByInput) String() string {
	return string(e)
}

func (e *UserOrderByInput) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserOrderByInput(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserOrderByInput", str)
	}
	return nil
}

func (e UserOrderByInput) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
