package util

import (
	"context"

	"actiontech.cloud/universe/ucommon/v4/log"
	"actiontech.cloud/universe/ucommon/v4/utrace"
)

func TraceUserInfo(ctx context.Context, msg string, args ...interface{}) { //NO_DUPL_CHECK
	stage := utrace.ULogStage(ctx)
	defer stage.Exit()

	utrace.TraceLog(ctx).Infof(msg, args...)
	log.UserInfo(stage, msg, args...)
}

func TraceUserWarn(ctx context.Context, msg string, args ...interface{}) {
	stage := utrace.ULogStage(ctx)
	defer stage.Exit()

	utrace.TraceLog(ctx).Warnf(msg, args...)
	log.UserWarn(stage, msg, args...)
}

func TraceUserError(ctx context.Context, msg string, args ...interface{}) {
	stage := utrace.ULogStage(ctx)
	defer stage.Exit()

	utrace.TraceLog(ctx).Errorf(msg, args...)
	log.UserError(stage, msg, args...)
}

func TraceDetail(ctx context.Context, msg string, args ...interface{}) {
	stage := utrace.ULogStage(ctx)
	defer stage.Exit()

	utrace.TraceLog(ctx).Infof(msg, args...)
	log.Detail(stage, msg, args...)
}

func TraceDebug(ctx context.Context, msg string, args ...interface{}) {
	stage := utrace.ULogStage(ctx)
	defer stage.Exit()

	utrace.TraceLog(ctx).Debugf(msg, args...)
	log.Debug(stage, msg, args...)
}

func TraceErrorf(ctx context.Context, msg string, args ...interface{}) {
	stage := utrace.ULogStage(ctx)
	defer stage.Exit()

	utrace.TraceLog(ctx).Errorf(msg, args...)
	log.Errorf(stage, msg, args...)
}
