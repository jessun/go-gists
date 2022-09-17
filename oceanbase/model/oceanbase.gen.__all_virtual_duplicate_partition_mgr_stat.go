package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualDuplicatePartitionMgrStatMgr struct {
	*_BaseMgr
}

// AllVirtualDuplicatePartitionMgrStatMgr open func
func AllVirtualDuplicatePartitionMgrStatMgr(db *gorm.DB) *_AllVirtualDuplicatePartitionMgrStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDuplicatePartitionMgrStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDuplicatePartitionMgrStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_duplicate_partition_mgr_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetTableName() string {
	return "__all_virtual_duplicate_partition_mgr_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) Reset() *_AllVirtualDuplicatePartitionMgrStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) Get() (result AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) Gets() (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTableID table_id获取
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithPartitionLeaseList partition_lease_list获取
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) WithPartitionLeaseList(partitionLeaseList string) Option {
	return optionFunc(func(o *options) { o.query["partition_lease_list"] = partitionLeaseList })
}

// WithIsMaster is_master获取
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) WithIsMaster(isMaster int64) Option {
	return optionFunc(func(o *options) { o.query["is_master"] = isMaster })
}

// WithCurLogID cur_log_id获取
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) WithCurLogID(curLogID int64) Option {
	return optionFunc(func(o *options) { o.query["cur_log_id"] = curLogID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetByOption(opts ...Option) (result AllVirtualDuplicatePartitionMgrStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDuplicatePartitionMgrStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where(options.query)
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
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromPartitionLeaseList 通过partition_lease_list获取内容
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetFromPartitionLeaseList(partitionLeaseList string) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`partition_lease_list` = ?", partitionLeaseList).Find(&results).Error

	return
}

// GetBatchFromPartitionLeaseList 批量查找
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetBatchFromPartitionLeaseList(partitionLeaseLists []string) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`partition_lease_list` IN (?)", partitionLeaseLists).Find(&results).Error

	return
}

// GetFromIsMaster 通过is_master获取内容
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetFromIsMaster(isMaster int64) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`is_master` = ?", isMaster).Find(&results).Error

	return
}

// GetBatchFromIsMaster 批量查找
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetBatchFromIsMaster(isMasters []int64) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`is_master` IN (?)", isMasters).Find(&results).Error

	return
}

// GetFromCurLogID 通过cur_log_id获取内容
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetFromCurLogID(curLogID int64) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`cur_log_id` = ?", curLogID).Find(&results).Error

	return
}

// GetBatchFromCurLogID 批量查找
func (obj *_AllVirtualDuplicatePartitionMgrStatMgr) GetBatchFromCurLogID(curLogIDs []int64) (results []*AllVirtualDuplicatePartitionMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDuplicatePartitionMgrStat{}).Where("`cur_log_id` IN (?)", curLogIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
