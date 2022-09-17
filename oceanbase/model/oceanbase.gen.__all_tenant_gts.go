package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantGtsMgr struct {
	*_BaseMgr
}

// AllTenantGtsMgr open func
func AllTenantGtsMgr(db *gorm.DB) *_AllTenantGtsMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantGtsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantGtsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_gts"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantGtsMgr) GetTableName() string {
	return "__all_tenant_gts"
}

// Reset 重置gorm会话
func (obj *_AllTenantGtsMgr) Reset() *_AllTenantGtsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantGtsMgr) Get() (result AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantGtsMgr) Gets() (results []*AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantGtsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantGtsMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantGtsMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantGtsMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithGtsID gts_id获取
func (obj *_AllTenantGtsMgr) WithGtsID(gtsID int64) Option {
	return optionFunc(func(o *options) { o.query["gts_id"] = gtsID })
}

// WithOrigGtsID orig_gts_id获取
func (obj *_AllTenantGtsMgr) WithOrigGtsID(origGtsID int64) Option {
	return optionFunc(func(o *options) { o.query["orig_gts_id"] = origGtsID })
}

// WithGtsInvalidTs gts_invalid_ts获取
func (obj *_AllTenantGtsMgr) WithGtsInvalidTs(gtsInvalidTs int64) Option {
	return optionFunc(func(o *options) { o.query["gts_invalid_ts"] = gtsInvalidTs })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantGtsMgr) GetByOption(opts ...Option) (result AllTenantGts, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantGtsMgr) GetByOptions(opts ...Option) (results []*AllTenantGts, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantGtsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantGts, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where(options.query)
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
func (obj *_AllTenantGtsMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantGtsMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantGtsMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantGtsMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantGtsMgr) GetFromTenantID(tenantID int64) (result AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantGtsMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromGtsID 通过gts_id获取内容
func (obj *_AllTenantGtsMgr) GetFromGtsID(gtsID int64) (results []*AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`gts_id` = ?", gtsID).Find(&results).Error

	return
}

// GetBatchFromGtsID 批量查找
func (obj *_AllTenantGtsMgr) GetBatchFromGtsID(gtsIDs []int64) (results []*AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`gts_id` IN (?)", gtsIDs).Find(&results).Error

	return
}

// GetFromOrigGtsID 通过orig_gts_id获取内容
func (obj *_AllTenantGtsMgr) GetFromOrigGtsID(origGtsID int64) (results []*AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`orig_gts_id` = ?", origGtsID).Find(&results).Error

	return
}

// GetBatchFromOrigGtsID 批量查找
func (obj *_AllTenantGtsMgr) GetBatchFromOrigGtsID(origGtsIDs []int64) (results []*AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`orig_gts_id` IN (?)", origGtsIDs).Find(&results).Error

	return
}

// GetFromGtsInvalidTs 通过gts_invalid_ts获取内容
func (obj *_AllTenantGtsMgr) GetFromGtsInvalidTs(gtsInvalidTs int64) (results []*AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`gts_invalid_ts` = ?", gtsInvalidTs).Find(&results).Error

	return
}

// GetBatchFromGtsInvalidTs 批量查找
func (obj *_AllTenantGtsMgr) GetBatchFromGtsInvalidTs(gtsInvalidTss []int64) (results []*AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`gts_invalid_ts` IN (?)", gtsInvalidTss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantGtsMgr) FetchByPrimaryKey(tenantID int64) (result AllTenantGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGts{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}
