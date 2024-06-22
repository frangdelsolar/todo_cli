package auth_cli

import (
	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/frangdelsolar/todo_cli/pkg/config"
)

var cfg *config.Config

const userKey = "USER_ID"

func init(){
    cfg, _ = config.Load()
}

func IsLoggedIn() bool {

    userId, err := cfg.GetSession(userKey)
    if err != nil {
        return false
    }

    user, err := auth.GetUserById(userId)
    if err != nil {
        return false
    }

    return user.ID > 0
}

func GetUserId() string {
    userId, _ := cfg.GetSession(userKey)
    return userId
}
