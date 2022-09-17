package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllSynonymHistoryMgr struct {
	*_BaseMgr
}

// AllSynonymHistoryMgr open func
func AllSynonymHistoryMgr(db *gorm.DB) *_AllSynonymHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllSynonymHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSynonymHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_synonym_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSynonymHistoryMgr) GetTableName() string {
	return "__all_synonym_history"
}

// Reset 重置gorm会话
func (obj *_AllSynonymHistoryMgr) Reset() *_AllSynonymHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSynonymHistoryMgr) Get() (result AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSynonymHistoryMgr) Gets() (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSynonymHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSynonymHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSynonymHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllSynonymHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSynonymID synonym_id获取
func (obj *_AllSynonymHistoryMgr) WithSynonymID(synonymID int64) Option {
	return optionFunc(func(o *options) { o.query["synonym_id"] = synonymID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllSynonymHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllSynonymHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseID database_id获取
func (obj *_AllSynonymHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSynonymName synonym_name获取
func (obj *_AllSynonymHistoryMgr) WithSynonymName(synonymName string) Option {
	return optionFunc(func(o *options) { o.query["synonym_name"] = synonymName })
}

// WithObjectName object_name获取
func (obj *_AllSynonymHistoryMgr) WithObjectName(objectName string) Option {
	return optionFunc(func(o *options) { o.query["object_name"] = objectName })
}

// WithObjectDatabaseID object_database_id获取
func (obj *_AllSynonymHistoryMgr) WithObjectDatabaseID(objectDatabaseID int64) Option {
	return optionFunc(func(o *options) { o.query["object_database_id"] = objectDatabaseID })
}

// GetByOption 功能选项模式获取
func (obj *_AllSynonymHistoryMgr) GetByOption(opts ...Option) (result AllSynonymHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSynonymHistoryMgr) GetByOptions(opts ...Option) (results []*AllSynonymHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSynonymHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSynonymHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where(options.query)
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
func (obj *_AllSynonymHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSynonymHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSynonymHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSynonymHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllSynonymHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllSynonymHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSynonymID 通过synonym_id获取内容
func (obj *_AllSynonymHistoryMgr) GetFromSynonymID(synonymID int64) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`synonym_id` = ?", synonymID).Find(&results).Error

	return
}

// GetBatchFromSynonymID 批量查找
func (obj *_AllSynonymHistoryMgr) GetBatchFromSynonymID(synonymIDs []int64) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`synonym_id` IN (?)", synonymIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllSynonymHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllSynonymHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllSynonymHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllSynonymHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllSynonymHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllSynonymHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSynonymName 通过synonym_name获取内容
func (obj *_AllSynonymHistoryMgr) GetFromSynonymName(synonymName string) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`synonym_name` = ?", synonymName).Find(&results).Error

	return
}

// GetBatchFromSynonymName 批量查找
func (obj *_AllSynonymHistoryMgr) GetBatchFromSynonymName(synonymNames []string) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`synonym_name` IN (?)", synonymNames).Find(&results).Error

	return
}

// GetFromObjectName 通过object_name获取内容
func (obj *_AllSynonymHistoryMgr) GetFromObjectName(objectName string) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`object_name` = ?", objectName).Find(&results).Error

	return
}

// GetBatchFromObjectName 批量查找
func (obj *_AllSynonymHistoryMgr) GetBatchFromObjectName(objectNames []string) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`object_name` IN (?)", objectNames).Find(&results).Error

	return
}

// GetFromObjectDatabaseID 通过object_database_id获取内容
func (obj *_AllSynonymHistoryMgr) GetFromObjectDatabaseID(objectDatabaseID int64) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`object_database_id` = ?", objectDatabaseID).Find(&results).Error

	return
}

// GetBatchFromObjectDatabaseID 批量查找
func (obj *_AllSynonymHistoryMgr) GetBatchFromObjectDatabaseID(objectDatabaseIDs []int64) (results []*AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`object_database_id` IN (?)", objectDatabaseIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSynonymHistoryMgr) FetchByPrimaryKey(tenantID int64, synonymID int64, schemaVersion int64) (result AllSynonymHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonymHistory{}).Where("`tenant_id` = ? AND `synonym_id` = ? AND `schema_version` = ?", tenantID, synonymID, schemaVersion).First(&result).Error

	return
}
