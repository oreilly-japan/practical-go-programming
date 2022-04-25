package main

// error-main
import (
	"os"

	"go.uber.org/multierr"
)

func run() error {
	var files []*os.File

	// multi-error-start
	var merr error
	for _, file := range files {
		err := importFile(file)
		if err != nil {
			// エラーが発生した場合は multierr に error を格納して、次の処理を継続
			merr = multierr.Append(merr, err)
		}
	}

	// エラーが一度でも発生していれば、呼び出し元にエラーを返却
	if merr != nil {
		return merr
	}

	// multi-error-end
	return nil
}

func importFile(file *os.File) error {
	return nil
}
