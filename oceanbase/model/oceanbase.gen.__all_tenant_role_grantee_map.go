package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantRoleGranteeMapMgr struct {
	*_BaseMgr
}

// AllTenantRoleGranteeMapMgr open func
func AllTenantRoleGranteeMapMgr(db *gorm.DB) *_AllTenantRoleGranteeMapMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantRoleGranteeMapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantRoleGranteeMapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_role_grantee_map"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantRoleGranteeMapMgr) GetTableName() string {
	return "__all_tenant_role_grantee_map"
}

// Reset 重置gorm会话
func (obj *_AllTenantRoleGranteeMapMgr) Reset() *_AllTenantRoleGranteeMapMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantRoleGranteeMapMgr) Get() (result AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantRoleGranteeMapMgr) Gets() (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantRoleGranteeMapMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantRoleGranteeMapMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantRoleGranteeMapMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantRoleGranteeMapMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithGranteeID grantee_id获取
func (obj *_AllTenantRoleGranteeMapMgr) WithGranteeID(granteeID int64) Option {
	return optionFunc(func(o *options) { o.query["grantee_id"] = granteeID })
}

// WithRoleID role_id获取
func (obj *_AllTenantRoleGranteeMapMgr) WithRoleID(roleID int64) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithAdminOption admin_option获取
func (obj *_AllTenantRoleGranteeMapMgr) WithAdminOption(adminOption int64) Option {
	return optionFunc(func(o *options) { o.query["admin_option"] = adminOption })
}

// WithDisableFlag disable_flag获取
func (obj *_AllTenantRoleGranteeMapMgr) WithDisableFlag(disableFlag int64) Option {
	return optionFunc(func(o *options) { o.query["disable_flag"] = disableFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantRoleGranteeMapMgr) GetByOption(opts ...Option) (result AllTenantRoleGranteeMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantRoleGranteeMapMgr) GetByOptions(opts ...Option) (results []*AllTenantRoleGranteeMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantRoleGranteeMapMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantRoleGranteeMap, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where(options.query)
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
func (obj *_AllTenantRoleGranteeMapMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantRoleGranteeMapMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantRoleGranteeMapMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantRoleGranteeMapMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantRoleGranteeMapMgr) GetFromTenantID(tenantID int64) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantRoleGranteeMapMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromGranteeID 通过grantee_id获取内容
func (obj *_AllTenantRoleGranteeMapMgr) GetFromGranteeID(granteeID int64) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`grantee_id` = ?", granteeID).Find(&results).Error

	return
}

// GetBatchFromGranteeID 批量查找
func (obj *_AllTenantRoleGranteeMapMgr) GetBatchFromGranteeID(granteeIDs []int64) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`grantee_id` IN (?)", granteeIDs).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容
func (obj *_AllTenantRoleGranteeMapMgr) GetFromRoleID(roleID int64) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`role_id` = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量查找
func (obj *_AllTenantRoleGranteeMapMgr) GetBatchFromRoleID(roleIDs []int64) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`role_id` IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromAdminOption 通过admin_option获取内容
func (obj *_AllTenantRoleGranteeMapMgr) GetFromAdminOption(adminOption int64) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`admin_option` = ?", adminOption).Find(&results).Error

	return
}

// GetBatchFromAdminOption 批量查找
func (obj *_AllTenantRoleGranteeMapMgr) GetBatchFromAdminOption(adminOptions []int64) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`admin_option` IN (?)", adminOptions).Find(&results).Error

	return
}

// GetFromDisableFlag 通过disable_flag获取内容
func (obj *_AllTenantRoleGranteeMapMgr) GetFromDisableFlag(disableFlag int64) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`disable_flag` = ?", disableFlag).Find(&results).Error

	return
}

// GetBatchFromDisableFlag 批量查找
func (obj *_AllTenantRoleGranteeMapMgr) GetBatchFromDisableFlag(disableFlags []int64) (results []*AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`disable_flag` IN (?)", disableFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantRoleGranteeMapMgr) FetchByPrimaryKey(tenantID int64, granteeID int64, roleID int64) (result AllTenantRoleGranteeMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMap{}).Where("`tenant_id` = ? AND `grantee_id` = ? AND `role_id` = ?", tenantID, granteeID, roleID).First(&result).Error

	return
}
