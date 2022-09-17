package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TenantVirtualWarningMgr struct {
	*_BaseMgr
}

// TenantVirtualWarningMgr open func
func TenantVirtualWarningMgr(db *gorm.DB) *_TenantVirtualWarningMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualWarningMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualWarningMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_warning"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualWarningMgr) GetTableName() string {
	return "__tenant_virtual_warning"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualWarningMgr) Reset() *_TenantVirtualWarningMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualWarningMgr) Get() (result TenantVirtualWarning, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualWarning{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualWarningMgr) Gets() (results []*TenantVirtualWarning, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualWarning{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualWarningMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualWarning{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithLevel level获取
func (obj *_TenantVirtualWarningMgr) WithLevel(level string) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithCode code获取
func (obj *_TenantVirtualWarningMgr) WithCode(code int64) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithMessage message获取
func (obj *_TenantVirtualWarningMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualWarningMgr) GetByOption(opts ...Option) (result TenantVirtualWarning, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualWarning{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualWarningMgr) GetByOptions(opts ...Option) (results []*TenantVirtualWarning, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualWarning{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualWarningMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualWarning, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualWarning{}).Where(options.query)
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

// GetFromLevel 通过level获取内容
func (obj *_TenantVirtualWarningMgr) GetFromLevel(level string) (results []*TenantVirtualWarning, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualWarning{}).Where("`level` = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量查找
func (obj *_TenantVirtualWarningMgr) GetBatchFromLevel(levels []string) (results []*TenantVirtualWarning, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualWarning{}).Where("`level` IN (?)", levels).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容
func (obj *_TenantVirtualWarningMgr) GetFromCode(code int64) (results []*TenantVirtualWarning, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualWarning{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找
func (obj *_TenantVirtualWarningMgr) GetBatchFromCode(codes []int64) (results []*TenantVirtualWarning, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualWarning{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromMessage 通过message获取内容
func (obj *_TenantVirtualWarningMgr) GetFromMessage(message string) (results []*TenantVirtualWarning, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualWarning{}).Where("`message` = ?", message).Find(&results).Error

	return
}

// GetBatchFromMessage 批量查找
func (obj *_TenantVirtualWarningMgr) GetBatchFromMessage(messages []string) (results []*TenantVirtualWarning, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualWarning{}).Where("`message` IN (?)", messages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
