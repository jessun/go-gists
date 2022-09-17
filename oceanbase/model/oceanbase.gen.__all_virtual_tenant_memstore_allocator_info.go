package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTenantMemstoreAllocatorInfoMgr struct {
	*_BaseMgr
}

// AllVirtualTenantMemstoreAllocatorInfoMgr open func
func AllVirtualTenantMemstoreAllocatorInfoMgr(db *gorm.DB) *_AllVirtualTenantMemstoreAllocatorInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantMemstoreAllocatorInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantMemstoreAllocatorInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_memstore_allocator_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetTableName() string {
	return "__all_virtual_tenant_memstore_allocator_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) Reset() *_AllVirtualTenantMemstoreAllocatorInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) Get() (result AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) Gets() (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithMtBaseVersion mt_base_version获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) WithMtBaseVersion(mtBaseVersion int64) Option {
	return optionFunc(func(o *options) { o.query["mt_base_version"] = mtBaseVersion })
}

// WithRetireClock retire_clock获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) WithRetireClock(retireClock int64) Option {
	return optionFunc(func(o *options) { o.query["retire_clock"] = retireClock })
}

// WithMtIsFrozen mt_is_frozen获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) WithMtIsFrozen(mtIsFrozen int64) Option {
	return optionFunc(func(o *options) { o.query["mt_is_frozen"] = mtIsFrozen })
}

// WithMtProtectionClock mt_protection_clock获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) WithMtProtectionClock(mtProtectionClock int64) Option {
	return optionFunc(func(o *options) { o.query["mt_protection_clock"] = mtProtectionClock })
}

// WithMtSnapshotVersion mt_snapshot_version获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) WithMtSnapshotVersion(mtSnapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["mt_snapshot_version"] = mtSnapshotVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetByOption(opts ...Option) (result AllVirtualTenantMemstoreAllocatorInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantMemstoreAllocatorInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where(options.query)
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
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetFromTableID(tableID int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromMtBaseVersion 通过mt_base_version获取内容
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetFromMtBaseVersion(mtBaseVersion int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`mt_base_version` = ?", mtBaseVersion).Find(&results).Error

	return
}

// GetBatchFromMtBaseVersion 批量查找
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetBatchFromMtBaseVersion(mtBaseVersions []int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`mt_base_version` IN (?)", mtBaseVersions).Find(&results).Error

	return
}

// GetFromRetireClock 通过retire_clock获取内容
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetFromRetireClock(retireClock int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`retire_clock` = ?", retireClock).Find(&results).Error

	return
}

// GetBatchFromRetireClock 批量查找
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetBatchFromRetireClock(retireClocks []int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`retire_clock` IN (?)", retireClocks).Find(&results).Error

	return
}

// GetFromMtIsFrozen 通过mt_is_frozen获取内容
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetFromMtIsFrozen(mtIsFrozen int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`mt_is_frozen` = ?", mtIsFrozen).Find(&results).Error

	return
}

// GetBatchFromMtIsFrozen 批量查找
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetBatchFromMtIsFrozen(mtIsFrozens []int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`mt_is_frozen` IN (?)", mtIsFrozens).Find(&results).Error

	return
}

// GetFromMtProtectionClock 通过mt_protection_clock获取内容
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetFromMtProtectionClock(mtProtectionClock int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`mt_protection_clock` = ?", mtProtectionClock).Find(&results).Error

	return
}

// GetBatchFromMtProtectionClock 批量查找
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetBatchFromMtProtectionClock(mtProtectionClocks []int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`mt_protection_clock` IN (?)", mtProtectionClocks).Find(&results).Error

	return
}

// GetFromMtSnapshotVersion 通过mt_snapshot_version获取内容
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetFromMtSnapshotVersion(mtSnapshotVersion int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`mt_snapshot_version` = ?", mtSnapshotVersion).Find(&results).Error

	return
}

// GetBatchFromMtSnapshotVersion 批量查找
func (obj *_AllVirtualTenantMemstoreAllocatorInfoMgr) GetBatchFromMtSnapshotVersion(mtSnapshotVersions []int64) (results []*AllVirtualTenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreAllocatorInfo{}).Where("`mt_snapshot_version` IN (?)", mtSnapshotVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
