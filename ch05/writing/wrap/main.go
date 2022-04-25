package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// error-main
// loadConfigError は設定ファイル読み込みエラーを表すエラー
type loadConfigError struct {
	msg string
	err error
}

func (e *loadConfigError) Error() string {
	return fmt.Sprintf("cannot load config: %s (%s)", e.msg, e.err.Error())
}

func (e *loadConfigError) Unwrap() error {
	return e.err
}

type Config struct {
	// ...
}

func LoadConfig(configFilePath string) (*Config, error) {
	var cfg *Config
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, &loadConfigError{msg: fmt.Sprintf("read file `%s`", configFilePath), err: err}
	}
	if err = json.Unmarshal(data, cfg); err != nil {
		return nil, &loadConfigError{msg: fmt.Sprintf("parse config file `%s`", configFilePath), err: err}
	}
	return cfg, nil
}

// error-main

// error-as-start
func main() {
	if _, err := LoadConfig("non-existing"); err != nil {
		// errors.As の中でエラーをアンラップして詳細なエラーを調べられます
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			// ...
		} else {
			// ...
		}
	}
}

// error-as-end
