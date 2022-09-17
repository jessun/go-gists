package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualServerStatMgr struct {
	*_BaseMgr
}

// AllVirtualServerStatMgr open func
func AllVirtualServerStatMgr(db *gorm.DB) *_AllVirtualServerStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualServerStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualServerStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_server_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualServerStatMgr) GetTableName() string {
	return "__all_virtual_server_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualServerStatMgr) Reset() *_AllVirtualServerStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualServerStatMgr) Get() (result AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualServerStatMgr) Gets() (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualServerStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualServerStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualServerStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithZone zone获取
func (obj *_AllVirtualServerStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithCPUTotal cpu_total获取
func (obj *_AllVirtualServerStatMgr) WithCPUTotal(cpuTotal float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_total"] = cpuTotal })
}

// WithCPUAssigned cpu_assigned获取
func (obj *_AllVirtualServerStatMgr) WithCPUAssigned(cpuAssigned float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_assigned"] = cpuAssigned })
}

// WithCPUAssignedPercent cpu_assigned_percent获取
func (obj *_AllVirtualServerStatMgr) WithCPUAssignedPercent(cpuAssignedPercent int64) Option {
	return optionFunc(func(o *options) { o.query["cpu_assigned_percent"] = cpuAssignedPercent })
}

// WithMemTotal mem_total获取
func (obj *_AllVirtualServerStatMgr) WithMemTotal(memTotal int64) Option {
	return optionFunc(func(o *options) { o.query["mem_total"] = memTotal })
}

// WithMemAssigned mem_assigned获取
func (obj *_AllVirtualServerStatMgr) WithMemAssigned(memAssigned int64) Option {
	return optionFunc(func(o *options) { o.query["mem_assigned"] = memAssigned })
}

// WithMemAssignedPercent mem_assigned_percent获取
func (obj *_AllVirtualServerStatMgr) WithMemAssignedPercent(memAssignedPercent int64) Option {
	return optionFunc(func(o *options) { o.query["mem_assigned_percent"] = memAssignedPercent })
}

// WithDiskTotal disk_total获取
func (obj *_AllVirtualServerStatMgr) WithDiskTotal(diskTotal int64) Option {
	return optionFunc(func(o *options) { o.query["disk_total"] = diskTotal })
}

// WithDiskAssigned disk_assigned获取
func (obj *_AllVirtualServerStatMgr) WithDiskAssigned(diskAssigned int64) Option {
	return optionFunc(func(o *options) { o.query["disk_assigned"] = diskAssigned })
}

// WithDiskAssignedPercent disk_assigned_percent获取
func (obj *_AllVirtualServerStatMgr) WithDiskAssignedPercent(diskAssignedPercent int64) Option {
	return optionFunc(func(o *options) { o.query["disk_assigned_percent"] = diskAssignedPercent })
}

// WithUnitNum unit_num获取
func (obj *_AllVirtualServerStatMgr) WithUnitNum(unitNum int64) Option {
	return optionFunc(func(o *options) { o.query["unit_num"] = unitNum })
}

// WithMigratingUnitNum migrating_unit_num获取
func (obj *_AllVirtualServerStatMgr) WithMigratingUnitNum(migratingUnitNum int64) Option {
	return optionFunc(func(o *options) { o.query["migrating_unit_num"] = migratingUnitNum })
}

// WithMergedVersion merged_version获取
func (obj *_AllVirtualServerStatMgr) WithMergedVersion(mergedVersion int64) Option {
	return optionFunc(func(o *options) { o.query["merged_version"] = mergedVersion })
}

// WithLeaderCount leader_count获取
func (obj *_AllVirtualServerStatMgr) WithLeaderCount(leaderCount int64) Option {
	return optionFunc(func(o *options) { o.query["leader_count"] = leaderCount })
}

// WithLoad load获取
func (obj *_AllVirtualServerStatMgr) WithLoad(load float64) Option {
	return optionFunc(func(o *options) { o.query["load"] = load })
}

// WithCPUWeight cpu_weight获取
func (obj *_AllVirtualServerStatMgr) WithCPUWeight(cpuWeight float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_weight"] = cpuWeight })
}

// WithMemoryWeight memory_weight获取
func (obj *_AllVirtualServerStatMgr) WithMemoryWeight(memoryWeight float64) Option {
	return optionFunc(func(o *options) { o.query["memory_weight"] = memoryWeight })
}

// WithDiskWeight disk_weight获取
func (obj *_AllVirtualServerStatMgr) WithDiskWeight(diskWeight float64) Option {
	return optionFunc(func(o *options) { o.query["disk_weight"] = diskWeight })
}

// WithID id获取
func (obj *_AllVirtualServerStatMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithInnerPort inner_port获取
func (obj *_AllVirtualServerStatMgr) WithInnerPort(innerPort int64) Option {
	return optionFunc(func(o *options) { o.query["inner_port"] = innerPort })
}

// WithBuildVersion build_version获取
func (obj *_AllVirtualServerStatMgr) WithBuildVersion(buildVersion string) Option {
	return optionFunc(func(o *options) { o.query["build_version"] = buildVersion })
}

// WithRegisterTime register_time获取
func (obj *_AllVirtualServerStatMgr) WithRegisterTime(registerTime int64) Option {
	return optionFunc(func(o *options) { o.query["register_time"] = registerTime })
}

// WithLastHeartbeatTime last_heartbeat_time获取
func (obj *_AllVirtualServerStatMgr) WithLastHeartbeatTime(lastHeartbeatTime int64) Option {
	return optionFunc(func(o *options) { o.query["last_heartbeat_time"] = lastHeartbeatTime })
}

// WithBlockMigrateInTime block_migrate_in_time获取
func (obj *_AllVirtualServerStatMgr) WithBlockMigrateInTime(blockMigrateInTime int64) Option {
	return optionFunc(func(o *options) { o.query["block_migrate_in_time"] = blockMigrateInTime })
}

// WithStartServiceTime start_service_time获取
func (obj *_AllVirtualServerStatMgr) WithStartServiceTime(startServiceTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_service_time"] = startServiceTime })
}

// WithLastOfflineTime last_offline_time获取
func (obj *_AllVirtualServerStatMgr) WithLastOfflineTime(lastOfflineTime int64) Option {
	return optionFunc(func(o *options) { o.query["last_offline_time"] = lastOfflineTime })
}

// WithStopTime stop_time获取
func (obj *_AllVirtualServerStatMgr) WithStopTime(stopTime int64) Option {
	return optionFunc(func(o *options) { o.query["stop_time"] = stopTime })
}

// WithForceStopHeartbeat force_stop_heartbeat获取
func (obj *_AllVirtualServerStatMgr) WithForceStopHeartbeat(forceStopHeartbeat int64) Option {
	return optionFunc(func(o *options) { o.query["force_stop_heartbeat"] = forceStopHeartbeat })
}

// WithAdminStatus admin_status获取
func (obj *_AllVirtualServerStatMgr) WithAdminStatus(adminStatus string) Option {
	return optionFunc(func(o *options) { o.query["admin_status"] = adminStatus })
}

// WithHeartbeatStatus heartbeat_status获取
func (obj *_AllVirtualServerStatMgr) WithHeartbeatStatus(heartbeatStatus string) Option {
	return optionFunc(func(o *options) { o.query["heartbeat_status"] = heartbeatStatus })
}

// WithWithRootserver with_rootserver获取
func (obj *_AllVirtualServerStatMgr) WithWithRootserver(withRootserver int64) Option {
	return optionFunc(func(o *options) { o.query["with_rootserver"] = withRootserver })
}

// WithWithPartition with_partition获取
func (obj *_AllVirtualServerStatMgr) WithWithPartition(withPartition int64) Option {
	return optionFunc(func(o *options) { o.query["with_partition"] = withPartition })
}

// WithMemInUse mem_in_use获取
func (obj *_AllVirtualServerStatMgr) WithMemInUse(memInUse int64) Option {
	return optionFunc(func(o *options) { o.query["mem_in_use"] = memInUse })
}

// WithDiskInUse disk_in_use获取
func (obj *_AllVirtualServerStatMgr) WithDiskInUse(diskInUse int64) Option {
	return optionFunc(func(o *options) { o.query["disk_in_use"] = diskInUse })
}

// WithClockDeviation clock_deviation获取
func (obj *_AllVirtualServerStatMgr) WithClockDeviation(clockDeviation int64) Option {
	return optionFunc(func(o *options) { o.query["clock_deviation"] = clockDeviation })
}

// WithHeartbeatLatency heartbeat_latency获取
func (obj *_AllVirtualServerStatMgr) WithHeartbeatLatency(heartbeatLatency int64) Option {
	return optionFunc(func(o *options) { o.query["heartbeat_latency"] = heartbeatLatency })
}

// WithClockSyncStatus clock_sync_status获取
func (obj *_AllVirtualServerStatMgr) WithClockSyncStatus(clockSyncStatus string) Option {
	return optionFunc(func(o *options) { o.query["clock_sync_status"] = clockSyncStatus })
}

// WithCPUCapacity cpu_capacity获取
func (obj *_AllVirtualServerStatMgr) WithCPUCapacity(cpuCapacity float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_capacity"] = cpuCapacity })
}

// WithCPUMaxAssigned cpu_max_assigned获取
func (obj *_AllVirtualServerStatMgr) WithCPUMaxAssigned(cpuMaxAssigned float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_max_assigned"] = cpuMaxAssigned })
}

// WithMemCapacity mem_capacity获取
func (obj *_AllVirtualServerStatMgr) WithMemCapacity(memCapacity int64) Option {
	return optionFunc(func(o *options) { o.query["mem_capacity"] = memCapacity })
}

// WithMemMaxAssigned mem_max_assigned获取
func (obj *_AllVirtualServerStatMgr) WithMemMaxAssigned(memMaxAssigned int64) Option {
	return optionFunc(func(o *options) { o.query["mem_max_assigned"] = memMaxAssigned })
}

// WithSslKeyExpiredTime ssl_key_expired_time获取
func (obj *_AllVirtualServerStatMgr) WithSslKeyExpiredTime(sslKeyExpiredTime int64) Option {
	return optionFunc(func(o *options) { o.query["ssl_key_expired_time"] = sslKeyExpiredTime })
}

// WithDiskActual disk_actual获取
func (obj *_AllVirtualServerStatMgr) WithDiskActual(diskActual int64) Option {
	return optionFunc(func(o *options) { o.query["disk_actual"] = diskActual })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualServerStatMgr) GetByOption(opts ...Option) (result AllVirtualServerStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualServerStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualServerStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualServerStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualServerStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualServerStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualServerStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualServerStatMgr) GetFromZone(zone string) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromCPUTotal 通过cpu_total获取内容
func (obj *_AllVirtualServerStatMgr) GetFromCPUTotal(cpuTotal float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`cpu_total` = ?", cpuTotal).Find(&results).Error

	return
}

// GetBatchFromCPUTotal 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromCPUTotal(cpuTotals []float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`cpu_total` IN (?)", cpuTotals).Find(&results).Error

	return
}

// GetFromCPUAssigned 通过cpu_assigned获取内容
func (obj *_AllVirtualServerStatMgr) GetFromCPUAssigned(cpuAssigned float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`cpu_assigned` = ?", cpuAssigned).Find(&results).Error

	return
}

// GetBatchFromCPUAssigned 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromCPUAssigned(cpuAssigneds []float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`cpu_assigned` IN (?)", cpuAssigneds).Find(&results).Error

	return
}

// GetFromCPUAssignedPercent 通过cpu_assigned_percent获取内容
func (obj *_AllVirtualServerStatMgr) GetFromCPUAssignedPercent(cpuAssignedPercent int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`cpu_assigned_percent` = ?", cpuAssignedPercent).Find(&results).Error

	return
}

// GetBatchFromCPUAssignedPercent 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromCPUAssignedPercent(cpuAssignedPercents []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`cpu_assigned_percent` IN (?)", cpuAssignedPercents).Find(&results).Error

	return
}

// GetFromMemTotal 通过mem_total获取内容
func (obj *_AllVirtualServerStatMgr) GetFromMemTotal(memTotal int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`mem_total` = ?", memTotal).Find(&results).Error

	return
}

// GetBatchFromMemTotal 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromMemTotal(memTotals []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`mem_total` IN (?)", memTotals).Find(&results).Error

	return
}

// GetFromMemAssigned 通过mem_assigned获取内容
func (obj *_AllVirtualServerStatMgr) GetFromMemAssigned(memAssigned int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`mem_assigned` = ?", memAssigned).Find(&results).Error

	return
}

// GetBatchFromMemAssigned 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromMemAssigned(memAssigneds []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`mem_assigned` IN (?)", memAssigneds).Find(&results).Error

	return
}

// GetFromMemAssignedPercent 通过mem_assigned_percent获取内容
func (obj *_AllVirtualServerStatMgr) GetFromMemAssignedPercent(memAssignedPercent int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`mem_assigned_percent` = ?", memAssignedPercent).Find(&results).Error

	return
}

// GetBatchFromMemAssignedPercent 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromMemAssignedPercent(memAssignedPercents []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`mem_assigned_percent` IN (?)", memAssignedPercents).Find(&results).Error

	return
}

// GetFromDiskTotal 通过disk_total获取内容
func (obj *_AllVirtualServerStatMgr) GetFromDiskTotal(diskTotal int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`disk_total` = ?", diskTotal).Find(&results).Error

	return
}

// GetBatchFromDiskTotal 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromDiskTotal(diskTotals []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`disk_total` IN (?)", diskTotals).Find(&results).Error

	return
}

// GetFromDiskAssigned 通过disk_assigned获取内容
func (obj *_AllVirtualServerStatMgr) GetFromDiskAssigned(diskAssigned int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`disk_assigned` = ?", diskAssigned).Find(&results).Error

	return
}

// GetBatchFromDiskAssigned 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromDiskAssigned(diskAssigneds []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`disk_assigned` IN (?)", diskAssigneds).Find(&results).Error

	return
}

// GetFromDiskAssignedPercent 通过disk_assigned_percent获取内容
func (obj *_AllVirtualServerStatMgr) GetFromDiskAssignedPercent(diskAssignedPercent int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`disk_assigned_percent` = ?", diskAssignedPercent).Find(&results).Error

	return
}

// GetBatchFromDiskAssignedPercent 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromDiskAssignedPercent(diskAssignedPercents []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`disk_assigned_percent` IN (?)", diskAssignedPercents).Find(&results).Error

	return
}

// GetFromUnitNum 通过unit_num获取内容
func (obj *_AllVirtualServerStatMgr) GetFromUnitNum(unitNum int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`unit_num` = ?", unitNum).Find(&results).Error

	return
}

// GetBatchFromUnitNum 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromUnitNum(unitNums []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`unit_num` IN (?)", unitNums).Find(&results).Error

	return
}

// GetFromMigratingUnitNum 通过migrating_unit_num获取内容
func (obj *_AllVirtualServerStatMgr) GetFromMigratingUnitNum(migratingUnitNum int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`migrating_unit_num` = ?", migratingUnitNum).Find(&results).Error

	return
}

// GetBatchFromMigratingUnitNum 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromMigratingUnitNum(migratingUnitNums []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`migrating_unit_num` IN (?)", migratingUnitNums).Find(&results).Error

	return
}

// GetFromMergedVersion 通过merged_version获取内容
func (obj *_AllVirtualServerStatMgr) GetFromMergedVersion(mergedVersion int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`merged_version` = ?", mergedVersion).Find(&results).Error

	return
}

// GetBatchFromMergedVersion 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromMergedVersion(mergedVersions []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`merged_version` IN (?)", mergedVersions).Find(&results).Error

	return
}

// GetFromLeaderCount 通过leader_count获取内容
func (obj *_AllVirtualServerStatMgr) GetFromLeaderCount(leaderCount int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`leader_count` = ?", leaderCount).Find(&results).Error

	return
}

// GetBatchFromLeaderCount 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromLeaderCount(leaderCounts []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`leader_count` IN (?)", leaderCounts).Find(&results).Error

	return
}

// GetFromLoad 通过load获取内容
func (obj *_AllVirtualServerStatMgr) GetFromLoad(load float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`load` = ?", load).Find(&results).Error

	return
}

// GetBatchFromLoad 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromLoad(loads []float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`load` IN (?)", loads).Find(&results).Error

	return
}

// GetFromCPUWeight 通过cpu_weight获取内容
func (obj *_AllVirtualServerStatMgr) GetFromCPUWeight(cpuWeight float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`cpu_weight` = ?", cpuWeight).Find(&results).Error

	return
}

// GetBatchFromCPUWeight 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromCPUWeight(cpuWeights []float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`cpu_weight` IN (?)", cpuWeights).Find(&results).Error

	return
}

// GetFromMemoryWeight 通过memory_weight获取内容
func (obj *_AllVirtualServerStatMgr) GetFromMemoryWeight(memoryWeight float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`memory_weight` = ?", memoryWeight).Find(&results).Error

	return
}

// GetBatchFromMemoryWeight 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromMemoryWeight(memoryWeights []float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`memory_weight` IN (?)", memoryWeights).Find(&results).Error

	return
}

// GetFromDiskWeight 通过disk_weight获取内容
func (obj *_AllVirtualServerStatMgr) GetFromDiskWeight(diskWeight float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`disk_weight` = ?", diskWeight).Find(&results).Error

	return
}

// GetBatchFromDiskWeight 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromDiskWeight(diskWeights []float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`disk_weight` IN (?)", diskWeights).Find(&results).Error

	return
}

// GetFromID 通过id获取内容
func (obj *_AllVirtualServerStatMgr) GetFromID(id int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromID(ids []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromInnerPort 通过inner_port获取内容
func (obj *_AllVirtualServerStatMgr) GetFromInnerPort(innerPort int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`inner_port` = ?", innerPort).Find(&results).Error

	return
}

// GetBatchFromInnerPort 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromInnerPort(innerPorts []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`inner_port` IN (?)", innerPorts).Find(&results).Error

	return
}

// GetFromBuildVersion 通过build_version获取内容
func (obj *_AllVirtualServerStatMgr) GetFromBuildVersion(buildVersion string) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`build_version` = ?", buildVersion).Find(&results).Error

	return
}

// GetBatchFromBuildVersion 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromBuildVersion(buildVersions []string) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`build_version` IN (?)", buildVersions).Find(&results).Error

	return
}

// GetFromRegisterTime 通过register_time获取内容
func (obj *_AllVirtualServerStatMgr) GetFromRegisterTime(registerTime int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`register_time` = ?", registerTime).Find(&results).Error

	return
}

// GetBatchFromRegisterTime 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromRegisterTime(registerTimes []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`register_time` IN (?)", registerTimes).Find(&results).Error

	return
}

// GetFromLastHeartbeatTime 通过last_heartbeat_time获取内容
func (obj *_AllVirtualServerStatMgr) GetFromLastHeartbeatTime(lastHeartbeatTime int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`last_heartbeat_time` = ?", lastHeartbeatTime).Find(&results).Error

	return
}

// GetBatchFromLastHeartbeatTime 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromLastHeartbeatTime(lastHeartbeatTimes []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`last_heartbeat_time` IN (?)", lastHeartbeatTimes).Find(&results).Error

	return
}

// GetFromBlockMigrateInTime 通过block_migrate_in_time获取内容
func (obj *_AllVirtualServerStatMgr) GetFromBlockMigrateInTime(blockMigrateInTime int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`block_migrate_in_time` = ?", blockMigrateInTime).Find(&results).Error

	return
}

// GetBatchFromBlockMigrateInTime 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromBlockMigrateInTime(blockMigrateInTimes []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`block_migrate_in_time` IN (?)", blockMigrateInTimes).Find(&results).Error

	return
}

// GetFromStartServiceTime 通过start_service_time获取内容
func (obj *_AllVirtualServerStatMgr) GetFromStartServiceTime(startServiceTime int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`start_service_time` = ?", startServiceTime).Find(&results).Error

	return
}

// GetBatchFromStartServiceTime 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromStartServiceTime(startServiceTimes []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`start_service_time` IN (?)", startServiceTimes).Find(&results).Error

	return
}

// GetFromLastOfflineTime 通过last_offline_time获取内容
func (obj *_AllVirtualServerStatMgr) GetFromLastOfflineTime(lastOfflineTime int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`last_offline_time` = ?", lastOfflineTime).Find(&results).Error

	return
}

// GetBatchFromLastOfflineTime 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromLastOfflineTime(lastOfflineTimes []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`last_offline_time` IN (?)", lastOfflineTimes).Find(&results).Error

	return
}

// GetFromStopTime 通过stop_time获取内容
func (obj *_AllVirtualServerStatMgr) GetFromStopTime(stopTime int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`stop_time` = ?", stopTime).Find(&results).Error

	return
}

// GetBatchFromStopTime 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromStopTime(stopTimes []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`stop_time` IN (?)", stopTimes).Find(&results).Error

	return
}

// GetFromForceStopHeartbeat 通过force_stop_heartbeat获取内容
func (obj *_AllVirtualServerStatMgr) GetFromForceStopHeartbeat(forceStopHeartbeat int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`force_stop_heartbeat` = ?", forceStopHeartbeat).Find(&results).Error

	return
}

// GetBatchFromForceStopHeartbeat 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromForceStopHeartbeat(forceStopHeartbeats []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`force_stop_heartbeat` IN (?)", forceStopHeartbeats).Find(&results).Error

	return
}

// GetFromAdminStatus 通过admin_status获取内容
func (obj *_AllVirtualServerStatMgr) GetFromAdminStatus(adminStatus string) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`admin_status` = ?", adminStatus).Find(&results).Error

	return
}

// GetBatchFromAdminStatus 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromAdminStatus(adminStatuss []string) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`admin_status` IN (?)", adminStatuss).Find(&results).Error

	return
}

// GetFromHeartbeatStatus 通过heartbeat_status获取内容
func (obj *_AllVirtualServerStatMgr) GetFromHeartbeatStatus(heartbeatStatus string) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`heartbeat_status` = ?", heartbeatStatus).Find(&results).Error

	return
}

// GetBatchFromHeartbeatStatus 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromHeartbeatStatus(heartbeatStatuss []string) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`heartbeat_status` IN (?)", heartbeatStatuss).Find(&results).Error

	return
}

// GetFromWithRootserver 通过with_rootserver获取内容
func (obj *_AllVirtualServerStatMgr) GetFromWithRootserver(withRootserver int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`with_rootserver` = ?", withRootserver).Find(&results).Error

	return
}

// GetBatchFromWithRootserver 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromWithRootserver(withRootservers []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`with_rootserver` IN (?)", withRootservers).Find(&results).Error

	return
}

// GetFromWithPartition 通过with_partition获取内容
func (obj *_AllVirtualServerStatMgr) GetFromWithPartition(withPartition int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`with_partition` = ?", withPartition).Find(&results).Error

	return
}

// GetBatchFromWithPartition 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromWithPartition(withPartitions []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`with_partition` IN (?)", withPartitions).Find(&results).Error

	return
}

// GetFromMemInUse 通过mem_in_use获取内容
func (obj *_AllVirtualServerStatMgr) GetFromMemInUse(memInUse int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`mem_in_use` = ?", memInUse).Find(&results).Error

	return
}

// GetBatchFromMemInUse 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromMemInUse(memInUses []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`mem_in_use` IN (?)", memInUses).Find(&results).Error

	return
}

// GetFromDiskInUse 通过disk_in_use获取内容
func (obj *_AllVirtualServerStatMgr) GetFromDiskInUse(diskInUse int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`disk_in_use` = ?", diskInUse).Find(&results).Error

	return
}

// GetBatchFromDiskInUse 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromDiskInUse(diskInUses []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`disk_in_use` IN (?)", diskInUses).Find(&results).Error

	return
}

// GetFromClockDeviation 通过clock_deviation获取内容
func (obj *_AllVirtualServerStatMgr) GetFromClockDeviation(clockDeviation int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`clock_deviation` = ?", clockDeviation).Find(&results).Error

	return
}

// GetBatchFromClockDeviation 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromClockDeviation(clockDeviations []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`clock_deviation` IN (?)", clockDeviations).Find(&results).Error

	return
}

// GetFromHeartbeatLatency 通过heartbeat_latency获取内容
func (obj *_AllVirtualServerStatMgr) GetFromHeartbeatLatency(heartbeatLatency int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`heartbeat_latency` = ?", heartbeatLatency).Find(&results).Error

	return
}

// GetBatchFromHeartbeatLatency 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromHeartbeatLatency(heartbeatLatencys []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`heartbeat_latency` IN (?)", heartbeatLatencys).Find(&results).Error

	return
}

// GetFromClockSyncStatus 通过clock_sync_status获取内容
func (obj *_AllVirtualServerStatMgr) GetFromClockSyncStatus(clockSyncStatus string) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`clock_sync_status` = ?", clockSyncStatus).Find(&results).Error

	return
}

// GetBatchFromClockSyncStatus 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromClockSyncStatus(clockSyncStatuss []string) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`clock_sync_status` IN (?)", clockSyncStatuss).Find(&results).Error

	return
}

// GetFromCPUCapacity 通过cpu_capacity获取内容
func (obj *_AllVirtualServerStatMgr) GetFromCPUCapacity(cpuCapacity float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`cpu_capacity` = ?", cpuCapacity).Find(&results).Error

	return
}

// GetBatchFromCPUCapacity 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromCPUCapacity(cpuCapacitys []float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`cpu_capacity` IN (?)", cpuCapacitys).Find(&results).Error

	return
}

// GetFromCPUMaxAssigned 通过cpu_max_assigned获取内容
func (obj *_AllVirtualServerStatMgr) GetFromCPUMaxAssigned(cpuMaxAssigned float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`cpu_max_assigned` = ?", cpuMaxAssigned).Find(&results).Error

	return
}

// GetBatchFromCPUMaxAssigned 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromCPUMaxAssigned(cpuMaxAssigneds []float64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`cpu_max_assigned` IN (?)", cpuMaxAssigneds).Find(&results).Error

	return
}

// GetFromMemCapacity 通过mem_capacity获取内容
func (obj *_AllVirtualServerStatMgr) GetFromMemCapacity(memCapacity int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`mem_capacity` = ?", memCapacity).Find(&results).Error

	return
}

// GetBatchFromMemCapacity 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromMemCapacity(memCapacitys []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`mem_capacity` IN (?)", memCapacitys).Find(&results).Error

	return
}

// GetFromMemMaxAssigned 通过mem_max_assigned获取内容
func (obj *_AllVirtualServerStatMgr) GetFromMemMaxAssigned(memMaxAssigned int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`mem_max_assigned` = ?", memMaxAssigned).Find(&results).Error

	return
}

// GetBatchFromMemMaxAssigned 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromMemMaxAssigned(memMaxAssigneds []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`mem_max_assigned` IN (?)", memMaxAssigneds).Find(&results).Error

	return
}

// GetFromSslKeyExpiredTime 通过ssl_key_expired_time获取内容
func (obj *_AllVirtualServerStatMgr) GetFromSslKeyExpiredTime(sslKeyExpiredTime int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`ssl_key_expired_time` = ?", sslKeyExpiredTime).Find(&results).Error

	return
}

// GetBatchFromSslKeyExpiredTime 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromSslKeyExpiredTime(sslKeyExpiredTimes []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`ssl_key_expired_time` IN (?)", sslKeyExpiredTimes).Find(&results).Error

	return
}

// GetFromDiskActual 通过disk_actual获取内容
func (obj *_AllVirtualServerStatMgr) GetFromDiskActual(diskActual int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`disk_actual` = ?", diskActual).Find(&results).Error

	return
}

// GetBatchFromDiskActual 批量查找
func (obj *_AllVirtualServerStatMgr) GetBatchFromDiskActual(diskActuals []int64) (results []*AllVirtualServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerStat{}).Where("`disk_actual` IN (?)", diskActuals).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
