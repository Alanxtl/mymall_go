package utils

import (
	"context"
	"fmt"
	"reflect"
)

func GetUserIdFromCtx(ctx context.Context) (uint32, error) {
	userId := ctx.Value(SessionUserId)

	if userId == nil {
		return 0, nil // 用户 ID 不存在，返回 nil 错误
	}

	userIdInt, ok := userId.(uint32)
	if !ok {
		return 0, fmt.Errorf("GetUserIdFromCtx assert fail: %v(%s) is not uint32", userId, reflect.TypeOf(userId))
	}

	return userIdInt, nil // 返回用户 ID
}
