package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTableModificationsMgr struct {
	*_BaseMgr
}

// AllVirtualTableModificationsMgr open func
func AllVirtualTableModificationsMgr(db *gorm.DB) *_AllVirtualTableModificationsMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTableModificationsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTableModificationsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_table_modifications"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTableModificationsMgr) GetTableName() string {
	return "__all_virtual_table_modifications"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTableModificationsMgr) Reset() *_AllVirtualTableModificationsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTableModificationsMgr) Get() (result AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTableModificationsMgr) Gets() (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTableModificationsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTableModificationsMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTableModificationsMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualTableModificationsMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualTableModificationsMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualTableModificationsMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithInsertRowCount insert_row_count获取
func (obj *_AllVirtualTableModificationsMgr) WithInsertRowCount(insertRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["insert_row_count"] = insertRowCount })
}

// WithUpdateRowCount update_row_count获取
func (obj *_AllVirtualTableModificationsMgr) WithUpdateRowCount(updateRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["update_row_count"] = updateRowCount })
}

// WithDeleteRowCount delete_row_count获取
func (obj *_AllVirtualTableModificationsMgr) WithDeleteRowCount(deleteRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["delete_row_count"] = deleteRowCount })
}

// WithMaxSnapshotVersion max_snapshot_version获取
func (obj *_AllVirtualTableModificationsMgr) WithMaxSnapshotVersion(maxSnapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["max_snapshot_version"] = maxSnapshotVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTableModificationsMgr) GetByOption(opts ...Option) (result AllVirtualTableModifications, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTableModificationsMgr) GetByOptions(opts ...Option) (results []*AllVirtualTableModifications, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTableModificationsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTableModifications, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where(options.query)
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
func (obj *_AllVirtualTableModificationsMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTableModificationsMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTableModificationsMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTableModificationsMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualTableModificationsMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTableModificationsMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualTableModificationsMgr) GetFromTableID(tableID int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualTableModificationsMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualTableModificationsMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualTableModificationsMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromInsertRowCount 通过insert_row_count获取内容
func (obj *_AllVirtualTableModificationsMgr) GetFromInsertRowCount(insertRowCount int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`insert_row_count` = ?", insertRowCount).Find(&results).Error

	return
}

// GetBatchFromInsertRowCount 批量查找
func (obj *_AllVirtualTableModificationsMgr) GetBatchFromInsertRowCount(insertRowCounts []int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`insert_row_count` IN (?)", insertRowCounts).Find(&results).Error

	return
}

// GetFromUpdateRowCount 通过update_row_count获取内容
func (obj *_AllVirtualTableModificationsMgr) GetFromUpdateRowCount(updateRowCount int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`update_row_count` = ?", updateRowCount).Find(&results).Error

	return
}

// GetBatchFromUpdateRowCount 批量查找
func (obj *_AllVirtualTableModificationsMgr) GetBatchFromUpdateRowCount(updateRowCounts []int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`update_row_count` IN (?)", updateRowCounts).Find(&results).Error

	return
}

// GetFromDeleteRowCount 通过delete_row_count获取内容
func (obj *_AllVirtualTableModificationsMgr) GetFromDeleteRowCount(deleteRowCount int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`delete_row_count` = ?", deleteRowCount).Find(&results).Error

	return
}

// GetBatchFromDeleteRowCount 批量查找
func (obj *_AllVirtualTableModificationsMgr) GetBatchFromDeleteRowCount(deleteRowCounts []int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`delete_row_count` IN (?)", deleteRowCounts).Find(&results).Error

	return
}

// GetFromMaxSnapshotVersion 通过max_snapshot_version获取内容
func (obj *_AllVirtualTableModificationsMgr) GetFromMaxSnapshotVersion(maxSnapshotVersion int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`max_snapshot_version` = ?", maxSnapshotVersion).Find(&results).Error

	return
}

// GetBatchFromMaxSnapshotVersion 批量查找
func (obj *_AllVirtualTableModificationsMgr) GetBatchFromMaxSnapshotVersion(maxSnapshotVersions []int64) (results []*AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`max_snapshot_version` IN (?)", maxSnapshotVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTableModificationsMgr) FetchByPrimaryKey(svrIP string, svrPort int64, tenantID int64, tableID int64, partitionID int64) (result AllVirtualTableModifications, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableModifications{}).Where("`svr_ip` = ? AND `svr_port` = ? AND `tenant_id` = ? AND `table_id` = ? AND `partition_id` = ?", svrIP, svrPort, tenantID, tableID, partitionID).First(&result).Error

	return
}
