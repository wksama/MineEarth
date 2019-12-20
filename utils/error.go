package utils

import (
	"errors"
	"fmt"
)

func BuildErrMsg(msg string, err error) error {
	return errors.New(fmt.Sprintf("%s：%s", msg, err))
}
