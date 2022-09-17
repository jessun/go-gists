package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllDdlIDMgr struct {
	*_BaseMgr
}

// AllDdlIDMgr open func
func AllDdlIDMgr(db *gorm.DB) *_AllDdlIDMgr {
	if db == nil {
		panic(fmt.Errorf("AllDdlIDMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllDdlIDMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_ddl_id"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllDdlIDMgr) GetTableName() string {
	return "__all_ddl_id"
}

// Reset 重置gorm会话
func (obj *_AllDdlIDMgr) Reset() *_AllDdlIDMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllDdlIDMgr) Get() (result AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllDdlIDMgr) Gets() (results []*AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllDdlIDMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllDdlIDMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllDdlIDMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllDdlIDMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDdlIDStr ddl_id_str获取
func (obj *_AllDdlIDMgr) WithDdlIDStr(ddlIDStr string) Option {
	return optionFunc(func(o *options) { o.query["ddl_id_str"] = ddlIDStr })
}

// WithDdlStmtStr ddl_stmt_str获取
func (obj *_AllDdlIDMgr) WithDdlStmtStr(ddlStmtStr string) Option {
	return optionFunc(func(o *options) { o.query["ddl_stmt_str"] = ddlStmtStr })
}

// GetByOption 功能选项模式获取
func (obj *_AllDdlIDMgr) GetByOption(opts ...Option) (result AllDdlID, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllDdlIDMgr) GetByOptions(opts ...Option) (results []*AllDdlID, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllDdlIDMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllDdlID, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where(options.query)
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
func (obj *_AllDdlIDMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllDdlIDMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllDdlIDMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllDdlIDMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllDdlIDMgr) GetFromTenantID(tenantID int64) (results []*AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllDdlIDMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDdlIDStr 通过ddl_id_str获取内容
func (obj *_AllDdlIDMgr) GetFromDdlIDStr(ddlIDStr string) (results []*AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where("`ddl_id_str` = ?", ddlIDStr).Find(&results).Error

	return
}

// GetBatchFromDdlIDStr 批量查找
func (obj *_AllDdlIDMgr) GetBatchFromDdlIDStr(ddlIDStrs []string) (results []*AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where("`ddl_id_str` IN (?)", ddlIDStrs).Find(&results).Error

	return
}

// GetFromDdlStmtStr 通过ddl_stmt_str获取内容
func (obj *_AllDdlIDMgr) GetFromDdlStmtStr(ddlStmtStr string) (results []*AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where("`ddl_stmt_str` = ?", ddlStmtStr).Find(&results).Error

	return
}

// GetBatchFromDdlStmtStr 批量查找
func (obj *_AllDdlIDMgr) GetBatchFromDdlStmtStr(ddlStmtStrs []string) (results []*AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where("`ddl_stmt_str` IN (?)", ddlStmtStrs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllDdlIDMgr) FetchByPrimaryKey(tenantID int64, ddlIDStr string) (result AllDdlID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlID{}).Where("`tenant_id` = ? AND `ddl_id_str` = ?", tenantID, ddlIDStr).First(&result).Error

	return
}
