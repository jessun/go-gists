package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualSynonymHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualSynonymHistoryMgr open func
func AllVirtualSynonymHistoryMgr(db *gorm.DB) *_AllVirtualSynonymHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSynonymHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSynonymHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_synonym_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSynonymHistoryMgr) GetTableName() string {
	return "__all_virtual_synonym_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSynonymHistoryMgr) Reset() *_AllVirtualSynonymHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSynonymHistoryMgr) Get() (result AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSynonymHistoryMgr) Gets() (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSynonymHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSynonymHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSynonymID synonym_id获取
func (obj *_AllVirtualSynonymHistoryMgr) WithSynonymID(synonymID int64) Option {
	return optionFunc(func(o *options) { o.query["synonym_id"] = synonymID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualSynonymHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSynonymHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSynonymHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualSynonymHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualSynonymHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSynonymName synonym_name获取
func (obj *_AllVirtualSynonymHistoryMgr) WithSynonymName(synonymName string) Option {
	return optionFunc(func(o *options) { o.query["synonym_name"] = synonymName })
}

// WithObjectName object_name获取
func (obj *_AllVirtualSynonymHistoryMgr) WithObjectName(objectName string) Option {
	return optionFunc(func(o *options) { o.query["object_name"] = objectName })
}

// WithObjectDatabaseID object_database_id获取
func (obj *_AllVirtualSynonymHistoryMgr) WithObjectDatabaseID(objectDatabaseID int64) Option {
	return optionFunc(func(o *options) { o.query["object_database_id"] = objectDatabaseID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSynonymHistoryMgr) GetByOption(opts ...Option) (result AllVirtualSynonymHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSynonymHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualSynonymHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSynonymHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSynonymHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where(options.query)
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
func (obj *_AllVirtualSynonymHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSynonymHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSynonymID 通过synonym_id获取内容
func (obj *_AllVirtualSynonymHistoryMgr) GetFromSynonymID(synonymID int64) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`synonym_id` = ?", synonymID).Find(&results).Error

	return
}

// GetBatchFromSynonymID 批量查找
func (obj *_AllVirtualSynonymHistoryMgr) GetBatchFromSynonymID(synonymIDs []int64) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`synonym_id` IN (?)", synonymIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualSynonymHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualSynonymHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSynonymHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSynonymHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSynonymHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSynonymHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualSynonymHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualSynonymHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualSynonymHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualSynonymHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSynonymName 通过synonym_name获取内容
func (obj *_AllVirtualSynonymHistoryMgr) GetFromSynonymName(synonymName string) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`synonym_name` = ?", synonymName).Find(&results).Error

	return
}

// GetBatchFromSynonymName 批量查找
func (obj *_AllVirtualSynonymHistoryMgr) GetBatchFromSynonymName(synonymNames []string) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`synonym_name` IN (?)", synonymNames).Find(&results).Error

	return
}

// GetFromObjectName 通过object_name获取内容
func (obj *_AllVirtualSynonymHistoryMgr) GetFromObjectName(objectName string) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`object_name` = ?", objectName).Find(&results).Error

	return
}

// GetBatchFromObjectName 批量查找
func (obj *_AllVirtualSynonymHistoryMgr) GetBatchFromObjectName(objectNames []string) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`object_name` IN (?)", objectNames).Find(&results).Error

	return
}

// GetFromObjectDatabaseID 通过object_database_id获取内容
func (obj *_AllVirtualSynonymHistoryMgr) GetFromObjectDatabaseID(objectDatabaseID int64) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`object_database_id` = ?", objectDatabaseID).Find(&results).Error

	return
}

// GetBatchFromObjectDatabaseID 批量查找
func (obj *_AllVirtualSynonymHistoryMgr) GetBatchFromObjectDatabaseID(objectDatabaseIDs []int64) (results []*AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`object_database_id` IN (?)", objectDatabaseIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSynonymHistoryMgr) FetchByPrimaryKey(tenantID int64, synonymID int64, schemaVersion int64) (result AllVirtualSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonymHistory{}).Where("`tenant_id` = ? AND `synonym_id` = ? AND `schema_version` = ?", tenantID, synonymID, schemaVersion).First(&result).Error

	return
}
