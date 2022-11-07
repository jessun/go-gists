package mysql

const (
	AccelerateReplicationModeOnce   = "once" // 一次性复制加速，当复制延迟低于阈值时会自动关闭，且不会再次开启
	AccelerateReplicationModeAlways = "always" // 永久性复制加速，不会自动关闭，只能手动关闭
)
