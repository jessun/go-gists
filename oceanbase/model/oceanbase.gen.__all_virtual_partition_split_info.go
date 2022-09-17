package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualPartitionSplitInfoMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionSplitInfoMgr open func
func AllVirtualPartitionSplitInfoMgr(db *gorm.DB) *_AllVirtualPartitionSplitInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionSplitInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionSplitInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_split_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionSplitInfoMgr) GetTableName() string {
	return "__all_virtual_partition_split_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionSplitInfoMgr) Reset() *_AllVirtualPartitionSplitInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionSplitInfoMgr) Get() (result AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionSplitInfoMgr) Gets() (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionSplitInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionSplitInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionSplitInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionSplitInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithZone zone获取
func (obj *_AllVirtualPartitionSplitInfoMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionSplitInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPartitionSplitInfoMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSplitState split_state获取
func (obj *_AllVirtualPartitionSplitInfoMgr) WithSplitState(splitState string) Option {
	return optionFunc(func(o *options) { o.query["split_state"] = splitState })
}

// WithMergeVersion merge_version获取
func (obj *_AllVirtualPartitionSplitInfoMgr) WithMergeVersion(mergeVersion int64) Option {
	return optionFunc(func(o *options) { o.query["merge_version"] = mergeVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionSplitInfoMgr) GetByOption(opts ...Option) (result AllVirtualPartitionSplitInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionSplitInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionSplitInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionSplitInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionSplitInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where(options.query)
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
func (obj *_AllVirtualPartitionSplitInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionSplitInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPartitionSplitInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionSplitInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionSplitInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionSplitInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualPartitionSplitInfoMgr) GetFromZone(zone string) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualPartitionSplitInfoMgr) GetBatchFromZone(zones []string) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionSplitInfoMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionSplitInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPartitionSplitInfoMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPartitionSplitInfoMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSplitState 通过split_state获取内容
func (obj *_AllVirtualPartitionSplitInfoMgr) GetFromSplitState(splitState string) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`split_state` = ?", splitState).Find(&results).Error

	return
}

// GetBatchFromSplitState 批量查找
func (obj *_AllVirtualPartitionSplitInfoMgr) GetBatchFromSplitState(splitStates []string) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`split_state` IN (?)", splitStates).Find(&results).Error

	return
}

// GetFromMergeVersion 通过merge_version获取内容
func (obj *_AllVirtualPartitionSplitInfoMgr) GetFromMergeVersion(mergeVersion int64) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`merge_version` = ?", mergeVersion).Find(&results).Error

	return
}

// GetBatchFromMergeVersion 批量查找
func (obj *_AllVirtualPartitionSplitInfoMgr) GetBatchFromMergeVersion(mergeVersions []int64) (results []*AllVirtualPartitionSplitInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSplitInfo{}).Where("`merge_version` IN (?)", mergeVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
