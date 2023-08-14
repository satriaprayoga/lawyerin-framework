package config

import (
	"fmt"

	"github.com/satriaprayoga/lawyerin-framework/pkg/utils"
)

func CheckDotEnv(path string) error {
	err := utils.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}
