package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TenantVirtualCurrentTenantMgr struct {
	*_BaseMgr
}

// TenantVirtualCurrentTenantMgr open func
func TenantVirtualCurrentTenantMgr(db *gorm.DB) *_TenantVirtualCurrentTenantMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualCurrentTenantMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualCurrentTenantMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_current_tenant"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualCurrentTenantMgr) GetTableName() string {
	return "__tenant_virtual_current_tenant"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualCurrentTenantMgr) Reset() *_TenantVirtualCurrentTenantMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualCurrentTenantMgr) Get() (result TenantVirtualCurrentTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualCurrentTenantMgr) Gets() (results []*TenantVirtualCurrentTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualCurrentTenantMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_TenantVirtualCurrentTenantMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTenantName tenant_name获取
func (obj *_TenantVirtualCurrentTenantMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithCreateStmt create_stmt获取
func (obj *_TenantVirtualCurrentTenantMgr) WithCreateStmt(createStmt string) Option {
	return optionFunc(func(o *options) { o.query["create_stmt"] = createStmt })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualCurrentTenantMgr) GetByOption(opts ...Option) (result TenantVirtualCurrentTenant, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualCurrentTenantMgr) GetByOptions(opts ...Option) (results []*TenantVirtualCurrentTenant, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualCurrentTenantMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualCurrentTenant, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).Where(options.query)
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
func (obj *_TenantVirtualCurrentTenantMgr) GetFromTenantID(tenantID int64) (result TenantVirtualCurrentTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_TenantVirtualCurrentTenantMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*TenantVirtualCurrentTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_TenantVirtualCurrentTenantMgr) GetFromTenantName(tenantName string) (results []*TenantVirtualCurrentTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_TenantVirtualCurrentTenantMgr) GetBatchFromTenantName(tenantNames []string) (results []*TenantVirtualCurrentTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromCreateStmt 通过create_stmt获取内容
func (obj *_TenantVirtualCurrentTenantMgr) GetFromCreateStmt(createStmt string) (results []*TenantVirtualCurrentTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).Where("`create_stmt` = ?", createStmt).Find(&results).Error

	return
}

// GetBatchFromCreateStmt 批量查找
func (obj *_TenantVirtualCurrentTenantMgr) GetBatchFromCreateStmt(createStmts []string) (results []*TenantVirtualCurrentTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).Where("`create_stmt` IN (?)", createStmts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TenantVirtualCurrentTenantMgr) FetchByPrimaryKey(tenantID int64) (result TenantVirtualCurrentTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCurrentTenant{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}
