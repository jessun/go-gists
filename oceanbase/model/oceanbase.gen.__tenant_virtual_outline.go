package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TenantVirtualOutlineMgr struct {
	*_BaseMgr
}

// TenantVirtualOutlineMgr open func
func TenantVirtualOutlineMgr(db *gorm.DB) *_TenantVirtualOutlineMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualOutlineMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualOutlineMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_outline"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualOutlineMgr) GetTableName() string {
	return "__tenant_virtual_outline"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualOutlineMgr) Reset() *_TenantVirtualOutlineMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualOutlineMgr) Get() (result TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualOutlineMgr) Gets() (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualOutlineMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_TenantVirtualOutlineMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDatabaseID database_id获取
func (obj *_TenantVirtualOutlineMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithOutlineID outline_id获取
func (obj *_TenantVirtualOutlineMgr) WithOutlineID(outlineID int64) Option {
	return optionFunc(func(o *options) { o.query["outline_id"] = outlineID })
}

// WithDatabaseName database_name获取
func (obj *_TenantVirtualOutlineMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithOutlineName outline_name获取
func (obj *_TenantVirtualOutlineMgr) WithOutlineName(outlineName string) Option {
	return optionFunc(func(o *options) { o.query["outline_name"] = outlineName })
}

// WithVisibleSignature visible_signature获取
func (obj *_TenantVirtualOutlineMgr) WithVisibleSignature(visibleSignature string) Option {
	return optionFunc(func(o *options) { o.query["visible_signature"] = visibleSignature })
}

// WithSQLText sql_text获取
func (obj *_TenantVirtualOutlineMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// WithOutlineTarget outline_target获取
func (obj *_TenantVirtualOutlineMgr) WithOutlineTarget(outlineTarget string) Option {
	return optionFunc(func(o *options) { o.query["outline_target"] = outlineTarget })
}

// WithOutlineSQL outline_sql获取
func (obj *_TenantVirtualOutlineMgr) WithOutlineSQL(outlineSQL string) Option {
	return optionFunc(func(o *options) { o.query["outline_sql"] = outlineSQL })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualOutlineMgr) GetByOption(opts ...Option) (result TenantVirtualOutline, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualOutlineMgr) GetByOptions(opts ...Option) (results []*TenantVirtualOutline, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualOutlineMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualOutline, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where(options.query)
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
func (obj *_TenantVirtualOutlineMgr) GetFromTenantID(tenantID int64) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_TenantVirtualOutlineMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_TenantVirtualOutlineMgr) GetFromDatabaseID(databaseID int64) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_TenantVirtualOutlineMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromOutlineID 通过outline_id获取内容
func (obj *_TenantVirtualOutlineMgr) GetFromOutlineID(outlineID int64) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`outline_id` = ?", outlineID).Find(&results).Error

	return
}

// GetBatchFromOutlineID 批量查找
func (obj *_TenantVirtualOutlineMgr) GetBatchFromOutlineID(outlineIDs []int64) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`outline_id` IN (?)", outlineIDs).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_TenantVirtualOutlineMgr) GetFromDatabaseName(databaseName string) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_TenantVirtualOutlineMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromOutlineName 通过outline_name获取内容
func (obj *_TenantVirtualOutlineMgr) GetFromOutlineName(outlineName string) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`outline_name` = ?", outlineName).Find(&results).Error

	return
}

// GetBatchFromOutlineName 批量查找
func (obj *_TenantVirtualOutlineMgr) GetBatchFromOutlineName(outlineNames []string) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`outline_name` IN (?)", outlineNames).Find(&results).Error

	return
}

// GetFromVisibleSignature 通过visible_signature获取内容
func (obj *_TenantVirtualOutlineMgr) GetFromVisibleSignature(visibleSignature string) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`visible_signature` = ?", visibleSignature).Find(&results).Error

	return
}

// GetBatchFromVisibleSignature 批量查找
func (obj *_TenantVirtualOutlineMgr) GetBatchFromVisibleSignature(visibleSignatures []string) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`visible_signature` IN (?)", visibleSignatures).Find(&results).Error

	return
}

// GetFromSQLText 通过sql_text获取内容
func (obj *_TenantVirtualOutlineMgr) GetFromSQLText(sqlText string) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`sql_text` = ?", sqlText).Find(&results).Error

	return
}

// GetBatchFromSQLText 批量查找
func (obj *_TenantVirtualOutlineMgr) GetBatchFromSQLText(sqlTexts []string) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error

	return
}

// GetFromOutlineTarget 通过outline_target获取内容
func (obj *_TenantVirtualOutlineMgr) GetFromOutlineTarget(outlineTarget string) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`outline_target` = ?", outlineTarget).Find(&results).Error

	return
}

// GetBatchFromOutlineTarget 批量查找
func (obj *_TenantVirtualOutlineMgr) GetBatchFromOutlineTarget(outlineTargets []string) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`outline_target` IN (?)", outlineTargets).Find(&results).Error

	return
}

// GetFromOutlineSQL 通过outline_sql获取内容
func (obj *_TenantVirtualOutlineMgr) GetFromOutlineSQL(outlineSQL string) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`outline_sql` = ?", outlineSQL).Find(&results).Error

	return
}

// GetBatchFromOutlineSQL 批量查找
func (obj *_TenantVirtualOutlineMgr) GetBatchFromOutlineSQL(outlineSQLs []string) (results []*TenantVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualOutline{}).Where("`outline_sql` IN (?)", outlineSQLs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
