package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualFuncMgr struct {
	*_BaseMgr
}

// AllVirtualFuncMgr open func
func AllVirtualFuncMgr(db *gorm.DB) *_AllVirtualFuncMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualFuncMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualFuncMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_func"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualFuncMgr) GetTableName() string {
	return "__all_virtual_func"
}

// Reset 重置gorm会话
func (obj *_AllVirtualFuncMgr) Reset() *_AllVirtualFuncMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualFuncMgr) Get() (result AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualFuncMgr) Gets() (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualFuncMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualFuncMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithName name获取
func (obj *_AllVirtualFuncMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualFuncMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualFuncMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithRet ret获取
func (obj *_AllVirtualFuncMgr) WithRet(ret int64) Option {
	return optionFunc(func(o *options) { o.query["ret"] = ret })
}

// WithDl dl获取
func (obj *_AllVirtualFuncMgr) WithDl(dl string) Option {
	return optionFunc(func(o *options) { o.query["dl"] = dl })
}

// WithUdfID udf_id获取
func (obj *_AllVirtualFuncMgr) WithUdfID(udfID int64) Option {
	return optionFunc(func(o *options) { o.query["udf_id"] = udfID })
}

// WithType type获取
func (obj *_AllVirtualFuncMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualFuncMgr) GetByOption(opts ...Option) (result AllVirtualFunc, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualFuncMgr) GetByOptions(opts ...Option) (results []*AllVirtualFunc, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualFuncMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualFunc, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where(options.query)
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
func (obj *_AllVirtualFuncMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualFuncMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualFuncMgr) GetFromName(name string) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualFuncMgr) GetBatchFromName(names []string) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualFuncMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualFuncMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualFuncMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualFuncMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromRet 通过ret获取内容
func (obj *_AllVirtualFuncMgr) GetFromRet(ret int64) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`ret` = ?", ret).Find(&results).Error

	return
}

// GetBatchFromRet 批量查找
func (obj *_AllVirtualFuncMgr) GetBatchFromRet(rets []int64) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`ret` IN (?)", rets).Find(&results).Error

	return
}

// GetFromDl 通过dl获取内容
func (obj *_AllVirtualFuncMgr) GetFromDl(dl string) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`dl` = ?", dl).Find(&results).Error

	return
}

// GetBatchFromDl 批量查找
func (obj *_AllVirtualFuncMgr) GetBatchFromDl(dls []string) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`dl` IN (?)", dls).Find(&results).Error

	return
}

// GetFromUdfID 通过udf_id获取内容
func (obj *_AllVirtualFuncMgr) GetFromUdfID(udfID int64) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`udf_id` = ?", udfID).Find(&results).Error

	return
}

// GetBatchFromUdfID 批量查找
func (obj *_AllVirtualFuncMgr) GetBatchFromUdfID(udfIDs []int64) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`udf_id` IN (?)", udfIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualFuncMgr) GetFromType(_type int64) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualFuncMgr) GetBatchFromType(_types []int64) (results []*AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualFuncMgr) FetchByPrimaryKey(tenantID int64, name string) (result AllVirtualFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFunc{}).Where("`tenant_id` = ? AND `name` = ?", tenantID, name).First(&result).Error

	return
}
