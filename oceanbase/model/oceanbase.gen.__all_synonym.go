package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllSynonymMgr struct {
	*_BaseMgr
}

// AllSynonymMgr open func
func AllSynonymMgr(db *gorm.DB) *_AllSynonymMgr {
	if db == nil {
		panic(fmt.Errorf("AllSynonymMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSynonymMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_synonym"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSynonymMgr) GetTableName() string {
	return "__all_synonym"
}

// Reset 重置gorm会话
func (obj *_AllSynonymMgr) Reset() *_AllSynonymMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSynonymMgr) Get() (result AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSynonymMgr) Gets() (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSynonymMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSynonymMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSynonymMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllSynonymMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSynonymID synonym_id获取
func (obj *_AllSynonymMgr) WithSynonymID(synonymID int64) Option {
	return optionFunc(func(o *options) { o.query["synonym_id"] = synonymID })
}

// WithDatabaseID database_id获取
func (obj *_AllSynonymMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllSynonymMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithSynonymName synonym_name获取
func (obj *_AllSynonymMgr) WithSynonymName(synonymName string) Option {
	return optionFunc(func(o *options) { o.query["synonym_name"] = synonymName })
}

// WithObjectName object_name获取
func (obj *_AllSynonymMgr) WithObjectName(objectName string) Option {
	return optionFunc(func(o *options) { o.query["object_name"] = objectName })
}

// WithObjectDatabaseID object_database_id获取
func (obj *_AllSynonymMgr) WithObjectDatabaseID(objectDatabaseID int64) Option {
	return optionFunc(func(o *options) { o.query["object_database_id"] = objectDatabaseID })
}

// GetByOption 功能选项模式获取
func (obj *_AllSynonymMgr) GetByOption(opts ...Option) (result AllSynonym, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSynonymMgr) GetByOptions(opts ...Option) (results []*AllSynonym, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSynonymMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSynonym, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where(options.query)
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
func (obj *_AllSynonymMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSynonymMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSynonymMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSynonymMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllSynonymMgr) GetFromTenantID(tenantID int64) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllSynonymMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSynonymID 通过synonym_id获取内容
func (obj *_AllSynonymMgr) GetFromSynonymID(synonymID int64) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`synonym_id` = ?", synonymID).Find(&results).Error

	return
}

// GetBatchFromSynonymID 批量查找
func (obj *_AllSynonymMgr) GetBatchFromSynonymID(synonymIDs []int64) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`synonym_id` IN (?)", synonymIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllSynonymMgr) GetFromDatabaseID(databaseID int64) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllSynonymMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllSynonymMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllSynonymMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromSynonymName 通过synonym_name获取内容
func (obj *_AllSynonymMgr) GetFromSynonymName(synonymName string) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`synonym_name` = ?", synonymName).Find(&results).Error

	return
}

// GetBatchFromSynonymName 批量查找
func (obj *_AllSynonymMgr) GetBatchFromSynonymName(synonymNames []string) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`synonym_name` IN (?)", synonymNames).Find(&results).Error

	return
}

// GetFromObjectName 通过object_name获取内容
func (obj *_AllSynonymMgr) GetFromObjectName(objectName string) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`object_name` = ?", objectName).Find(&results).Error

	return
}

// GetBatchFromObjectName 批量查找
func (obj *_AllSynonymMgr) GetBatchFromObjectName(objectNames []string) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`object_name` IN (?)", objectNames).Find(&results).Error

	return
}

// GetFromObjectDatabaseID 通过object_database_id获取内容
func (obj *_AllSynonymMgr) GetFromObjectDatabaseID(objectDatabaseID int64) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`object_database_id` = ?", objectDatabaseID).Find(&results).Error

	return
}

// GetBatchFromObjectDatabaseID 批量查找
func (obj *_AllSynonymMgr) GetBatchFromObjectDatabaseID(objectDatabaseIDs []int64) (results []*AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`object_database_id` IN (?)", objectDatabaseIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSynonymMgr) FetchByPrimaryKey(tenantID int64, synonymID int64) (result AllSynonym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSynonym{}).Where("`tenant_id` = ? AND `synonym_id` = ?", tenantID, synonymID).First(&result).Error

	return
}
