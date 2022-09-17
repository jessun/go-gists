package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TenantVirtualConcurrentLimitSQLMgr struct {
	*_BaseMgr
}

// TenantVirtualConcurrentLimitSQLMgr open func
func TenantVirtualConcurrentLimitSQLMgr(db *gorm.DB) *_TenantVirtualConcurrentLimitSQLMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualConcurrentLimitSQLMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualConcurrentLimitSQLMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_concurrent_limit_sql"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetTableName() string {
	return "__tenant_virtual_concurrent_limit_sql"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualConcurrentLimitSQLMgr) Reset() *_TenantVirtualConcurrentLimitSQLMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) Get() (result TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualConcurrentLimitSQLMgr) Gets() (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualConcurrentLimitSQLMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDatabaseID database_id获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithOutlineID outline_id获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) WithOutlineID(outlineID int64) Option {
	return optionFunc(func(o *options) { o.query["outline_id"] = outlineID })
}

// WithDatabaseName database_name获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithOutlineName outline_name获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) WithOutlineName(outlineName string) Option {
	return optionFunc(func(o *options) { o.query["outline_name"] = outlineName })
}

// WithOutlineContent outline_content获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) WithOutlineContent(outlineContent string) Option {
	return optionFunc(func(o *options) { o.query["outline_content"] = outlineContent })
}

// WithVisibleSignature visible_signature获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) WithVisibleSignature(visibleSignature string) Option {
	return optionFunc(func(o *options) { o.query["visible_signature"] = visibleSignature })
}

// WithSQLText sql_text获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// WithConcurrentNum concurrent_num获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) WithConcurrentNum(concurrentNum int64) Option {
	return optionFunc(func(o *options) { o.query["concurrent_num"] = concurrentNum })
}

// WithLimitTarget limit_target获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) WithLimitTarget(limitTarget string) Option {
	return optionFunc(func(o *options) { o.query["limit_target"] = limitTarget })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetByOption(opts ...Option) (result TenantVirtualConcurrentLimitSQL, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetByOptions(opts ...Option) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualConcurrentLimitSQLMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualConcurrentLimitSQL, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where(options.query)
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
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetFromTenantID(tenantID int64) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetFromDatabaseID(databaseID int64) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromOutlineID 通过outline_id获取内容
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetFromOutlineID(outlineID int64) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`outline_id` = ?", outlineID).Find(&results).Error

	return
}

// GetBatchFromOutlineID 批量查找
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetBatchFromOutlineID(outlineIDs []int64) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`outline_id` IN (?)", outlineIDs).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetFromDatabaseName(databaseName string) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromOutlineName 通过outline_name获取内容
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetFromOutlineName(outlineName string) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`outline_name` = ?", outlineName).Find(&results).Error

	return
}

// GetBatchFromOutlineName 批量查找
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetBatchFromOutlineName(outlineNames []string) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`outline_name` IN (?)", outlineNames).Find(&results).Error

	return
}

// GetFromOutlineContent 通过outline_content获取内容
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetFromOutlineContent(outlineContent string) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`outline_content` = ?", outlineContent).Find(&results).Error

	return
}

// GetBatchFromOutlineContent 批量查找
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetBatchFromOutlineContent(outlineContents []string) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`outline_content` IN (?)", outlineContents).Find(&results).Error

	return
}

// GetFromVisibleSignature 通过visible_signature获取内容
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetFromVisibleSignature(visibleSignature string) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`visible_signature` = ?", visibleSignature).Find(&results).Error

	return
}

// GetBatchFromVisibleSignature 批量查找
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetBatchFromVisibleSignature(visibleSignatures []string) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`visible_signature` IN (?)", visibleSignatures).Find(&results).Error

	return
}

// GetFromSQLText 通过sql_text获取内容
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetFromSQLText(sqlText string) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`sql_text` = ?", sqlText).Find(&results).Error

	return
}

// GetBatchFromSQLText 批量查找
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetBatchFromSQLText(sqlTexts []string) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error

	return
}

// GetFromConcurrentNum 通过concurrent_num获取内容
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetFromConcurrentNum(concurrentNum int64) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`concurrent_num` = ?", concurrentNum).Find(&results).Error

	return
}

// GetBatchFromConcurrentNum 批量查找
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetBatchFromConcurrentNum(concurrentNums []int64) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`concurrent_num` IN (?)", concurrentNums).Find(&results).Error

	return
}

// GetFromLimitTarget 通过limit_target获取内容
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetFromLimitTarget(limitTarget string) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`limit_target` = ?", limitTarget).Find(&results).Error

	return
}

// GetBatchFromLimitTarget 批量查找
func (obj *_TenantVirtualConcurrentLimitSQLMgr) GetBatchFromLimitTarget(limitTargets []string) (results []*TenantVirtualConcurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualConcurrentLimitSQL{}).Where("`limit_target` IN (?)", limitTargets).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
