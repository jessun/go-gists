package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllIndexScheduleTaskMgr struct {
	*_BaseMgr
}

// AllIndexScheduleTaskMgr open func
func AllIndexScheduleTaskMgr(db *gorm.DB) *_AllIndexScheduleTaskMgr {
	if db == nil {
		panic(fmt.Errorf("AllIndexScheduleTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllIndexScheduleTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_index_schedule_task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllIndexScheduleTaskMgr) GetTableName() string {
	return "__all_index_schedule_task"
}

// Reset 重置gorm会话
func (obj *_AllIndexScheduleTaskMgr) Reset() *_AllIndexScheduleTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllIndexScheduleTaskMgr) Get() (result AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllIndexScheduleTaskMgr) Gets() (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllIndexScheduleTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllIndexScheduleTaskMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllIndexScheduleTaskMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllIndexScheduleTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIndexTableID index_table_id获取
func (obj *_AllIndexScheduleTaskMgr) WithIndexTableID(indexTableID int64) Option {
	return optionFunc(func(o *options) { o.query["index_table_id"] = indexTableID })
}

// WithPartitionID partition_id获取
func (obj *_AllIndexScheduleTaskMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllIndexScheduleTaskMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllIndexScheduleTaskMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithFrozenVersion frozen_version获取
func (obj *_AllIndexScheduleTaskMgr) WithFrozenVersion(frozenVersion int64) Option {
	return optionFunc(func(o *options) { o.query["frozen_version"] = frozenVersion })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllIndexScheduleTaskMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllIndexScheduleTaskMgr) GetByOption(opts ...Option) (result AllIndexScheduleTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllIndexScheduleTaskMgr) GetByOptions(opts ...Option) (results []*AllIndexScheduleTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllIndexScheduleTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllIndexScheduleTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where(options.query)
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

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllIndexScheduleTaskMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllIndexScheduleTaskMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllIndexScheduleTaskMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllIndexScheduleTaskMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllIndexScheduleTaskMgr) GetFromTenantID(tenantID int64) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllIndexScheduleTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIndexTableID 通过index_table_id获取内容
func (obj *_AllIndexScheduleTaskMgr) GetFromIndexTableID(indexTableID int64) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`index_table_id` = ?", indexTableID).Find(&results).Error

	return
}

// GetBatchFromIndexTableID 批量查找
func (obj *_AllIndexScheduleTaskMgr) GetBatchFromIndexTableID(indexTableIDs []int64) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`index_table_id` IN (?)", indexTableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllIndexScheduleTaskMgr) GetFromPartitionID(partitionID int64) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllIndexScheduleTaskMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllIndexScheduleTaskMgr) GetFromSvrIP(svrIP string) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllIndexScheduleTaskMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllIndexScheduleTaskMgr) GetFromSvrPort(svrPort int64) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllIndexScheduleTaskMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromFrozenVersion 通过frozen_version获取内容
func (obj *_AllIndexScheduleTaskMgr) GetFromFrozenVersion(frozenVersion int64) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`frozen_version` = ?", frozenVersion).Find(&results).Error

	return
}

// GetBatchFromFrozenVersion 批量查找
func (obj *_AllIndexScheduleTaskMgr) GetBatchFromFrozenVersion(frozenVersions []int64) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`frozen_version` IN (?)", frozenVersions).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllIndexScheduleTaskMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllIndexScheduleTaskMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllIndexScheduleTaskMgr) FetchByPrimaryKey(tenantID int64, indexTableID int64, partitionID int64) (result AllIndexScheduleTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexScheduleTask{}).Where("`tenant_id` = ? AND `index_table_id` = ? AND `partition_id` = ?", tenantID, indexTableID, partitionID).First(&result).Error

	return
}
