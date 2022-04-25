package main

import (
	"context"
	"fmt"
)

func run() error {
	ctx := context.Background()
	email := "test@sample.com"

	// error-no-wrap-start
	user, err := getInvitedUserWithEmail(ctx, email)
	if err != nil {
		// 呼び出し先で発生したエラーをそのまま呼び出し元に返却
		return err
	}
	// error-no-wrap-end

	_ = user

	return nil
}

func runImprove() error {
	ctx := context.Background()
	email := "test@sample.com"

	// error-wrap-start
	user, err := getInvitedUserWithEmail(ctx, email)
	if err != nil {
		// 呼び出し先で発生したエラーをラップし、付加情報を付与して呼び出し元に返却
		return fmt.Errorf("fail to get invited user with email(%s): %w", email, err)
	}
	// error-wrap-end

	_ = user

	return nil
}

func getInvitedUserWithEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}

type User struct{}
