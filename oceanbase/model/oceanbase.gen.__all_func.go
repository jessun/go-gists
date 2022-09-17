package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllFuncMgr struct {
	*_BaseMgr
}

// AllFuncMgr open func
func AllFuncMgr(db *gorm.DB) *_AllFuncMgr {
	if db == nil {
		panic(fmt.Errorf("AllFuncMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllFuncMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_func"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllFuncMgr) GetTableName() string {
	return "__all_func"
}

// Reset 重置gorm会话
func (obj *_AllFuncMgr) Reset() *_AllFuncMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllFuncMgr) Get() (result AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllFuncMgr) Gets() (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllFuncMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllFuncMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllFuncMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllFuncMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithName name获取
func (obj *_AllFuncMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithRet ret获取
func (obj *_AllFuncMgr) WithRet(ret int64) Option {
	return optionFunc(func(o *options) { o.query["ret"] = ret })
}

// WithDl dl获取
func (obj *_AllFuncMgr) WithDl(dl string) Option {
	return optionFunc(func(o *options) { o.query["dl"] = dl })
}

// WithUdfID udf_id获取
func (obj *_AllFuncMgr) WithUdfID(udfID int64) Option {
	return optionFunc(func(o *options) { o.query["udf_id"] = udfID })
}

// WithType type获取
func (obj *_AllFuncMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_AllFuncMgr) GetByOption(opts ...Option) (result AllFunc, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllFuncMgr) GetByOptions(opts ...Option) (results []*AllFunc, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllFuncMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllFunc, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where(options.query)
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
func (obj *_AllFuncMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllFuncMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllFuncMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllFuncMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllFuncMgr) GetFromTenantID(tenantID int64) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllFuncMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllFuncMgr) GetFromName(name string) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllFuncMgr) GetBatchFromName(names []string) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromRet 通过ret获取内容
func (obj *_AllFuncMgr) GetFromRet(ret int64) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`ret` = ?", ret).Find(&results).Error

	return
}

// GetBatchFromRet 批量查找
func (obj *_AllFuncMgr) GetBatchFromRet(rets []int64) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`ret` IN (?)", rets).Find(&results).Error

	return
}

// GetFromDl 通过dl获取内容
func (obj *_AllFuncMgr) GetFromDl(dl string) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`dl` = ?", dl).Find(&results).Error

	return
}

// GetBatchFromDl 批量查找
func (obj *_AllFuncMgr) GetBatchFromDl(dls []string) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`dl` IN (?)", dls).Find(&results).Error

	return
}

// GetFromUdfID 通过udf_id获取内容
func (obj *_AllFuncMgr) GetFromUdfID(udfID int64) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`udf_id` = ?", udfID).Find(&results).Error

	return
}

// GetBatchFromUdfID 批量查找
func (obj *_AllFuncMgr) GetBatchFromUdfID(udfIDs []int64) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`udf_id` IN (?)", udfIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllFuncMgr) GetFromType(_type int64) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllFuncMgr) GetBatchFromType(_types []int64) (results []*AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllFuncMgr) FetchByPrimaryKey(tenantID int64, name string) (result AllFunc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFunc{}).Where("`tenant_id` = ? AND `name` = ?", tenantID, name).First(&result).Error

	return
}
