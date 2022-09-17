package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllAcquiredSnapshotMgr struct {
	*_BaseMgr
}

// AllAcquiredSnapshotMgr open func
func AllAcquiredSnapshotMgr(db *gorm.DB) *_AllAcquiredSnapshotMgr {
	if db == nil {
		panic(fmt.Errorf("AllAcquiredSnapshotMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllAcquiredSnapshotMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_acquired_snapshot"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllAcquiredSnapshotMgr) GetTableName() string {
	return "__all_acquired_snapshot"
}

// Reset 重置gorm会话
func (obj *_AllAcquiredSnapshotMgr) Reset() *_AllAcquiredSnapshotMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllAcquiredSnapshotMgr) Get() (result AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllAcquiredSnapshotMgr) Gets() (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllAcquiredSnapshotMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllAcquiredSnapshotMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithSnapshotType snapshot_type获取
func (obj *_AllAcquiredSnapshotMgr) WithSnapshotType(snapshotType int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_type"] = snapshotType })
}

// WithSnapshotTs snapshot_ts获取
func (obj *_AllAcquiredSnapshotMgr) WithSnapshotTs(snapshotTs int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_ts"] = snapshotTs })
}

// WithSchemaVersion schema_version获取
func (obj *_AllAcquiredSnapshotMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithTenantID tenant_id获取
func (obj *_AllAcquiredSnapshotMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllAcquiredSnapshotMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithExtraInfo extra_info获取
func (obj *_AllAcquiredSnapshotMgr) WithExtraInfo(extraInfo string) Option {
	return optionFunc(func(o *options) { o.query["extra_info"] = extraInfo })
}

// GetByOption 功能选项模式获取
func (obj *_AllAcquiredSnapshotMgr) GetByOption(opts ...Option) (result AllAcquiredSnapshot, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllAcquiredSnapshotMgr) GetByOptions(opts ...Option) (results []*AllAcquiredSnapshot, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllAcquiredSnapshotMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllAcquiredSnapshot, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where(options.query)
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
func (obj *_AllAcquiredSnapshotMgr) GetFromGmtCreate(gmtCreate time.Time) (result AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`gmt_create` = ?", gmtCreate).First(&result).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllAcquiredSnapshotMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromSnapshotType 通过snapshot_type获取内容
func (obj *_AllAcquiredSnapshotMgr) GetFromSnapshotType(snapshotType int64) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`snapshot_type` = ?", snapshotType).Find(&results).Error

	return
}

// GetBatchFromSnapshotType 批量查找
func (obj *_AllAcquiredSnapshotMgr) GetBatchFromSnapshotType(snapshotTypes []int64) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`snapshot_type` IN (?)", snapshotTypes).Find(&results).Error

	return
}

// GetFromSnapshotTs 通过snapshot_ts获取内容
func (obj *_AllAcquiredSnapshotMgr) GetFromSnapshotTs(snapshotTs int64) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`snapshot_ts` = ?", snapshotTs).Find(&results).Error

	return
}

// GetBatchFromSnapshotTs 批量查找
func (obj *_AllAcquiredSnapshotMgr) GetBatchFromSnapshotTs(snapshotTss []int64) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`snapshot_ts` IN (?)", snapshotTss).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllAcquiredSnapshotMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllAcquiredSnapshotMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllAcquiredSnapshotMgr) GetFromTenantID(tenantID int64) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllAcquiredSnapshotMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllAcquiredSnapshotMgr) GetFromTableID(tableID int64) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllAcquiredSnapshotMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromExtraInfo 通过extra_info获取内容
func (obj *_AllAcquiredSnapshotMgr) GetFromExtraInfo(extraInfo string) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`extra_info` = ?", extraInfo).Find(&results).Error

	return
}

// GetBatchFromExtraInfo 批量查找
func (obj *_AllAcquiredSnapshotMgr) GetBatchFromExtraInfo(extraInfos []string) (results []*AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`extra_info` IN (?)", extraInfos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllAcquiredSnapshotMgr) FetchByPrimaryKey(gmtCreate time.Time) (result AllAcquiredSnapshot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllAcquiredSnapshot{}).Where("`gmt_create` = ?", gmtCreate).First(&result).Error

	return
}
