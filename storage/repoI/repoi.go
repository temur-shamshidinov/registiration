package repoi

import "registiration/models"

type UserRepoI interface {
	CheckUserNotExists(username, gmail string) (bool, error)
	CreateUser(account models.UserAccount) error
}
