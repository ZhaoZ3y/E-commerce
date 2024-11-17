package utils

import "context"

func GetUserIDFromCtx(ctx context.Context) int32 {
	userID := ctx.Value(SessionUserId)
	if userID == nil {
		return 0
	}
	return userID.(int32)
}
