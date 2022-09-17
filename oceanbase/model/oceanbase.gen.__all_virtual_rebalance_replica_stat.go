package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualRebalanceReplicaStatMgr struct {
	*_BaseMgr
}

// AllVirtualRebalanceReplicaStatMgr open func
func AllVirtualRebalanceReplicaStatMgr(db *gorm.DB) *_AllVirtualRebalanceReplicaStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRebalanceReplicaStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRebalanceReplicaStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_rebalance_replica_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetTableName() string {
	return "__all_virtual_rebalance_replica_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRebalanceReplicaStatMgr) Reset() *_AllVirtualRebalanceReplicaStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) Get() (result AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRebalanceReplicaStatMgr) Gets() (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRebalanceReplicaStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithCPUUsage cpu_usage获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithCPUUsage(cpuUsage float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_usage"] = cpuUsage })
}

// WithDiskUsage disk_usage获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithDiskUsage(diskUsage float64) Option {
	return optionFunc(func(o *options) { o.query["disk_usage"] = diskUsage })
}

// WithIopsUsage iops_usage获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithIopsUsage(iopsUsage float64) Option {
	return optionFunc(func(o *options) { o.query["iops_usage"] = iopsUsage })
}

// WithMemoryUsage memory_usage获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithMemoryUsage(memoryUsage float64) Option {
	return optionFunc(func(o *options) { o.query["memory_usage"] = memoryUsage })
}

// WithNetPacketUsage net_packet_usage获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithNetPacketUsage(netPacketUsage float64) Option {
	return optionFunc(func(o *options) { o.query["net_packet_usage"] = netPacketUsage })
}

// WithNetThroughputUsage net_throughput_usage获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithNetThroughputUsage(netThroughputUsage float64) Option {
	return optionFunc(func(o *options) { o.query["net_throughput_usage"] = netThroughputUsage })
}

// WithZone zone获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithUnitID unit_id获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetByOption(opts ...Option) (result AllVirtualRebalanceReplicaStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualRebalanceReplicaStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRebalanceReplicaStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRebalanceReplicaStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where(options.query)
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

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromCPUUsage 通过cpu_usage获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromCPUUsage(cpuUsage float64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`cpu_usage` = ?", cpuUsage).Find(&results).Error

	return
}

// GetBatchFromCPUUsage 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromCPUUsage(cpuUsages []float64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`cpu_usage` IN (?)", cpuUsages).Find(&results).Error

	return
}

// GetFromDiskUsage 通过disk_usage获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromDiskUsage(diskUsage float64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`disk_usage` = ?", diskUsage).Find(&results).Error

	return
}

// GetBatchFromDiskUsage 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromDiskUsage(diskUsages []float64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`disk_usage` IN (?)", diskUsages).Find(&results).Error

	return
}

// GetFromIopsUsage 通过iops_usage获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromIopsUsage(iopsUsage float64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`iops_usage` = ?", iopsUsage).Find(&results).Error

	return
}

// GetBatchFromIopsUsage 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromIopsUsage(iopsUsages []float64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`iops_usage` IN (?)", iopsUsages).Find(&results).Error

	return
}

// GetFromMemoryUsage 通过memory_usage获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromMemoryUsage(memoryUsage float64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`memory_usage` = ?", memoryUsage).Find(&results).Error

	return
}

// GetBatchFromMemoryUsage 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromMemoryUsage(memoryUsages []float64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`memory_usage` IN (?)", memoryUsages).Find(&results).Error

	return
}

// GetFromNetPacketUsage 通过net_packet_usage获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromNetPacketUsage(netPacketUsage float64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`net_packet_usage` = ?", netPacketUsage).Find(&results).Error

	return
}

// GetBatchFromNetPacketUsage 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromNetPacketUsage(netPacketUsages []float64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`net_packet_usage` IN (?)", netPacketUsages).Find(&results).Error

	return
}

// GetFromNetThroughputUsage 通过net_throughput_usage获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromNetThroughputUsage(netThroughputUsage float64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`net_throughput_usage` = ?", netThroughputUsage).Find(&results).Error

	return
}

// GetBatchFromNetThroughputUsage 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromNetThroughputUsage(netThroughputUsages []float64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`net_throughput_usage` IN (?)", netThroughputUsages).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromZone(zone string) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetFromUnitID(unitID int64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllVirtualRebalanceReplicaStatMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllVirtualRebalanceReplicaStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceReplicaStat{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
