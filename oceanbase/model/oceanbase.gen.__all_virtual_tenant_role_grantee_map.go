package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualTenantRoleGranteeMapMgr struct {
	*_BaseMgr
}

// AllVirtualTenantRoleGranteeMapMgr open func
func AllVirtualTenantRoleGranteeMapMgr(db *gorm.DB) *_AllVirtualTenantRoleGranteeMapMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantRoleGranteeMapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantRoleGranteeMapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_role_grantee_map"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetTableName() string {
	return "__all_virtual_tenant_role_grantee_map"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantRoleGranteeMapMgr) Reset() *_AllVirtualTenantRoleGranteeMapMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantRoleGranteeMapMgr) Get() (result AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantRoleGranteeMapMgr) Gets() (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantRoleGranteeMapMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantRoleGranteeMapMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithGranteeID grantee_id获取
func (obj *_AllVirtualTenantRoleGranteeMapMgr) WithGranteeID(granteeID int64) Option {
	return optionFunc(func(o *options) { o.query["grantee_id"] = granteeID })
}

// WithRoleID role_id获取
func (obj *_AllVirtualTenantRoleGranteeMapMgr) WithRoleID(roleID int64) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTenantRoleGranteeMapMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTenantRoleGranteeMapMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithAdminOption admin_option获取
func (obj *_AllVirtualTenantRoleGranteeMapMgr) WithAdminOption(adminOption int64) Option {
	return optionFunc(func(o *options) { o.query["admin_option"] = adminOption })
}

// WithDisableFlag disable_flag获取
func (obj *_AllVirtualTenantRoleGranteeMapMgr) WithDisableFlag(disableFlag int64) Option {
	return optionFunc(func(o *options) { o.query["disable_flag"] = disableFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetByOption(opts ...Option) (result AllVirtualTenantRoleGranteeMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantRoleGranteeMapMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantRoleGranteeMap, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where(options.query)
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
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromGranteeID 通过grantee_id获取内容
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetFromGranteeID(granteeID int64) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`grantee_id` = ?", granteeID).Find(&results).Error

	return
}

// GetBatchFromGranteeID 批量查找
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetBatchFromGranteeID(granteeIDs []int64) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`grantee_id` IN (?)", granteeIDs).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetFromRoleID(roleID int64) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`role_id` = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量查找
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetBatchFromRoleID(roleIDs []int64) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`role_id` IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromAdminOption 通过admin_option获取内容
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetFromAdminOption(adminOption int64) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`admin_option` = ?", adminOption).Find(&results).Error

	return
}

// GetBatchFromAdminOption 批量查找
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetBatchFromAdminOption(adminOptions []int64) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`admin_option` IN (?)", adminOptions).Find(&results).Error

	return
}

// GetFromDisableFlag 通过disable_flag获取内容
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetFromDisableFlag(disableFlag int64) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`disable_flag` = ?", disableFlag).Find(&results).Error

	return
}

// GetBatchFromDisableFlag 批量查找
func (obj *_AllVirtualTenantRoleGranteeMapMgr) GetBatchFromDisableFlag(disableFlags []int64) (results []*AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`disable_flag` IN (?)", disableFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTenantRoleGranteeMapMgr) FetchByPrimaryKey(tenantID int64, granteeID int64, roleID int64) (result AllVirtualTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantRoleGranteeMap{}).Where("`tenant_id` = ? AND `grantee_id` = ? AND `role_id` = ?", tenantID, granteeID, roleID).First(&result).Error

	return
}
