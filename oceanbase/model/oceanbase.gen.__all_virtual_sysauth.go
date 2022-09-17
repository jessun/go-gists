package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualSysauthMgr struct {
	*_BaseMgr
}

// AllVirtualSysauthMgr open func
func AllVirtualSysauthMgr(db *gorm.DB) *_AllVirtualSysauthMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSysauthMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSysauthMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sysauth"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSysauthMgr) GetTableName() string {
	return "__all_virtual_sysauth"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSysauthMgr) Reset() *_AllVirtualSysauthMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSysauthMgr) Get() (result AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSysauthMgr) Gets() (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSysauthMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSysauthMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithGranteeID grantee_id获取
func (obj *_AllVirtualSysauthMgr) WithGranteeID(granteeID int64) Option {
	return optionFunc(func(o *options) { o.query["grantee_id"] = granteeID })
}

// WithPrivID priv_id获取
func (obj *_AllVirtualSysauthMgr) WithPrivID(privID int64) Option {
	return optionFunc(func(o *options) { o.query["priv_id"] = privID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSysauthMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSysauthMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithPrivOption priv_option获取
func (obj *_AllVirtualSysauthMgr) WithPrivOption(privOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_option"] = privOption })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSysauthMgr) GetByOption(opts ...Option) (result AllVirtualSysauth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSysauthMgr) GetByOptions(opts ...Option) (results []*AllVirtualSysauth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSysauthMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSysauth, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where(options.query)
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
func (obj *_AllVirtualSysauthMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSysauthMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromGranteeID 通过grantee_id获取内容
func (obj *_AllVirtualSysauthMgr) GetFromGranteeID(granteeID int64) (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`grantee_id` = ?", granteeID).Find(&results).Error

	return
}

// GetBatchFromGranteeID 批量查找
func (obj *_AllVirtualSysauthMgr) GetBatchFromGranteeID(granteeIDs []int64) (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`grantee_id` IN (?)", granteeIDs).Find(&results).Error

	return
}

// GetFromPrivID 通过priv_id获取内容
func (obj *_AllVirtualSysauthMgr) GetFromPrivID(privID int64) (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`priv_id` = ?", privID).Find(&results).Error

	return
}

// GetBatchFromPrivID 批量查找
func (obj *_AllVirtualSysauthMgr) GetBatchFromPrivID(privIDs []int64) (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`priv_id` IN (?)", privIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSysauthMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSysauthMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSysauthMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSysauthMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromPrivOption 通过priv_option获取内容
func (obj *_AllVirtualSysauthMgr) GetFromPrivOption(privOption int64) (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`priv_option` = ?", privOption).Find(&results).Error

	return
}

// GetBatchFromPrivOption 批量查找
func (obj *_AllVirtualSysauthMgr) GetBatchFromPrivOption(privOptions []int64) (results []*AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`priv_option` IN (?)", privOptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSysauthMgr) FetchByPrimaryKey(tenantID int64, granteeID int64, privID int64) (result AllVirtualSysauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauth{}).Where("`tenant_id` = ? AND `grantee_id` = ? AND `priv_id` = ?", tenantID, granteeID, privID).First(&result).Error

	return
}
