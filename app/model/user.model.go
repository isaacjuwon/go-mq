
package model

import (
    "fusossafuoye.ng/app/dao"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type UserModel struct {
    dao.User
}

func (u *UserModel) BeforeCreate(db *gorm.DB) error {
    // Generate UUID
    u.ID = uuid.New()

    // Hash password
    password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
    if err != nil {
        return err
    }
    u.Password = string(password)

    return nil
}

