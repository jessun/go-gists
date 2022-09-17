package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantSysauthMgr struct {
	*_BaseMgr
}

// AllTenantSysauthMgr open func
func AllTenantSysauthMgr(db *gorm.DB) *_AllTenantSysauthMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantSysauthMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantSysauthMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_sysauth"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantSysauthMgr) GetTableName() string {
	return "__all_tenant_sysauth"
}

// Reset 重置gorm会话
func (obj *_AllTenantSysauthMgr) Reset() *_AllTenantSysauthMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantSysauthMgr) Get() (result AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantSysauthMgr) Gets() (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantSysauthMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantSysauthMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantSysauthMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantSysauthMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithGranteeID grantee_id获取
func (obj *_AllTenantSysauthMgr) WithGranteeID(granteeID int64) Option {
	return optionFunc(func(o *options) { o.query["grantee_id"] = granteeID })
}

// WithPrivID priv_id获取
func (obj *_AllTenantSysauthMgr) WithPrivID(privID int64) Option {
	return optionFunc(func(o *options) { o.query["priv_id"] = privID })
}

// WithPrivOption priv_option获取
func (obj *_AllTenantSysauthMgr) WithPrivOption(privOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_option"] = privOption })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantSysauthMgr) GetByOption(opts ...Option) (result AllTenantSysauth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantSysauthMgr) GetByOptions(opts ...Option) (results []*AllTenantSysauth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantSysauthMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantSysauth, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where(options.query)
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
func (obj *_AllTenantSysauthMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantSysauthMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantSysauthMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantSysauthMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantSysauthMgr) GetFromTenantID(tenantID int64) (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantSysauthMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromGranteeID 通过grantee_id获取内容
func (obj *_AllTenantSysauthMgr) GetFromGranteeID(granteeID int64) (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`grantee_id` = ?", granteeID).Find(&results).Error

	return
}

// GetBatchFromGranteeID 批量查找
func (obj *_AllTenantSysauthMgr) GetBatchFromGranteeID(granteeIDs []int64) (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`grantee_id` IN (?)", granteeIDs).Find(&results).Error

	return
}

// GetFromPrivID 通过priv_id获取内容
func (obj *_AllTenantSysauthMgr) GetFromPrivID(privID int64) (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`priv_id` = ?", privID).Find(&results).Error

	return
}

// GetBatchFromPrivID 批量查找
func (obj *_AllTenantSysauthMgr) GetBatchFromPrivID(privIDs []int64) (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`priv_id` IN (?)", privIDs).Find(&results).Error

	return
}

// GetFromPrivOption 通过priv_option获取内容
func (obj *_AllTenantSysauthMgr) GetFromPrivOption(privOption int64) (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`priv_option` = ?", privOption).Find(&results).Error

	return
}

// GetBatchFromPrivOption 批量查找
func (obj *_AllTenantSysauthMgr) GetBatchFromPrivOption(privOptions []int64) (results []*AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`priv_option` IN (?)", privOptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantSysauthMgr) FetchByPrimaryKey(tenantID int64, granteeID int64, privID int64) (result AllTenantSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauth{}).Where("`tenant_id` = ? AND `grantee_id` = ? AND `priv_id` = ?", tenantID, granteeID, privID).First(&result).Error

	return
}
