package mysql

import (
	"fmt"
)

func NewPool(user, password, host, port, name string) (string, error) {
	connCfg := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)

	return connCfg, nil
}
