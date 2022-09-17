package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualSynonymMgr struct {
	*_BaseMgr
}

// AllVirtualSynonymMgr open func
func AllVirtualSynonymMgr(db *gorm.DB) *_AllVirtualSynonymMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSynonymMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSynonymMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_synonym"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSynonymMgr) GetTableName() string {
	return "__all_virtual_synonym"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSynonymMgr) Reset() *_AllVirtualSynonymMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSynonymMgr) Get() (result AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSynonymMgr) Gets() (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSynonymMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSynonymMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSynonymID synonym_id获取
func (obj *_AllVirtualSynonymMgr) WithSynonymID(synonymID int64) Option {
	return optionFunc(func(o *options) { o.query["synonym_id"] = synonymID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSynonymMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSynonymMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualSynonymMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualSynonymMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithSynonymName synonym_name获取
func (obj *_AllVirtualSynonymMgr) WithSynonymName(synonymName string) Option {
	return optionFunc(func(o *options) { o.query["synonym_name"] = synonymName })
}

// WithObjectName object_name获取
func (obj *_AllVirtualSynonymMgr) WithObjectName(objectName string) Option {
	return optionFunc(func(o *options) { o.query["object_name"] = objectName })
}

// WithObjectDatabaseID object_database_id获取
func (obj *_AllVirtualSynonymMgr) WithObjectDatabaseID(objectDatabaseID int64) Option {
	return optionFunc(func(o *options) { o.query["object_database_id"] = objectDatabaseID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSynonymMgr) GetByOption(opts ...Option) (result AllVirtualSynonym, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSynonymMgr) GetByOptions(opts ...Option) (results []*AllVirtualSynonym, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSynonymMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSynonym, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where(options.query)
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
func (obj *_AllVirtualSynonymMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSynonymMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSynonymID 通过synonym_id获取内容
func (obj *_AllVirtualSynonymMgr) GetFromSynonymID(synonymID int64) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`synonym_id` = ?", synonymID).Find(&results).Error

	return
}

// GetBatchFromSynonymID 批量查找
func (obj *_AllVirtualSynonymMgr) GetBatchFromSynonymID(synonymIDs []int64) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`synonym_id` IN (?)", synonymIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSynonymMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSynonymMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSynonymMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSynonymMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualSynonymMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualSynonymMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualSynonymMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualSynonymMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromSynonymName 通过synonym_name获取内容
func (obj *_AllVirtualSynonymMgr) GetFromSynonymName(synonymName string) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`synonym_name` = ?", synonymName).Find(&results).Error

	return
}

// GetBatchFromSynonymName 批量查找
func (obj *_AllVirtualSynonymMgr) GetBatchFromSynonymName(synonymNames []string) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`synonym_name` IN (?)", synonymNames).Find(&results).Error

	return
}

// GetFromObjectName 通过object_name获取内容
func (obj *_AllVirtualSynonymMgr) GetFromObjectName(objectName string) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`object_name` = ?", objectName).Find(&results).Error

	return
}

// GetBatchFromObjectName 批量查找
func (obj *_AllVirtualSynonymMgr) GetBatchFromObjectName(objectNames []string) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`object_name` IN (?)", objectNames).Find(&results).Error

	return
}

// GetFromObjectDatabaseID 通过object_database_id获取内容
func (obj *_AllVirtualSynonymMgr) GetFromObjectDatabaseID(objectDatabaseID int64) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`object_database_id` = ?", objectDatabaseID).Find(&results).Error

	return
}

// GetBatchFromObjectDatabaseID 批量查找
func (obj *_AllVirtualSynonymMgr) GetBatchFromObjectDatabaseID(objectDatabaseIDs []int64) (results []*AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`object_database_id` IN (?)", objectDatabaseIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSynonymMgr) FetchByPrimaryKey(tenantID int64, synonymID int64) (result AllVirtualSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSynonym{}).Where("`tenant_id` = ? AND `synonym_id` = ?", tenantID, synonymID).First(&result).Error

	return
}
