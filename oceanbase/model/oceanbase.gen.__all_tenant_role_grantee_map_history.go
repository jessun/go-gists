package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantRoleGranteeMapHistoryMgr struct {
	*_BaseMgr
}

// AllTenantRoleGranteeMapHistoryMgr open func
func AllTenantRoleGranteeMapHistoryMgr(db *gorm.DB) *_AllTenantRoleGranteeMapHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantRoleGranteeMapHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantRoleGranteeMapHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_role_grantee_map_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetTableName() string {
	return "__all_tenant_role_grantee_map_history"
}

// Reset 重置gorm会话
func (obj *_AllTenantRoleGranteeMapHistoryMgr) Reset() *_AllTenantRoleGranteeMapHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantRoleGranteeMapHistoryMgr) Get() (result AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantRoleGranteeMapHistoryMgr) Gets() (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantRoleGranteeMapHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantRoleGranteeMapHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantRoleGranteeMapHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantRoleGranteeMapHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithGranteeID grantee_id获取
func (obj *_AllTenantRoleGranteeMapHistoryMgr) WithGranteeID(granteeID int64) Option {
	return optionFunc(func(o *options) { o.query["grantee_id"] = granteeID })
}

// WithRoleID role_id获取
func (obj *_AllTenantRoleGranteeMapHistoryMgr) WithRoleID(roleID int64) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantRoleGranteeMapHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTenantRoleGranteeMapHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithAdminOption admin_option获取
func (obj *_AllTenantRoleGranteeMapHistoryMgr) WithAdminOption(adminOption int64) Option {
	return optionFunc(func(o *options) { o.query["admin_option"] = adminOption })
}

// WithDisableFlag disable_flag获取
func (obj *_AllTenantRoleGranteeMapHistoryMgr) WithDisableFlag(disableFlag int64) Option {
	return optionFunc(func(o *options) { o.query["disable_flag"] = disableFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetByOption(opts ...Option) (result AllTenantRoleGranteeMapHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetByOptions(opts ...Option) (results []*AllTenantRoleGranteeMapHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantRoleGranteeMapHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantRoleGranteeMapHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where(options.query)
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
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromGranteeID 通过grantee_id获取内容
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetFromGranteeID(granteeID int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`grantee_id` = ?", granteeID).Find(&results).Error

	return
}

// GetBatchFromGranteeID 批量查找
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetBatchFromGranteeID(granteeIDs []int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`grantee_id` IN (?)", granteeIDs).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetFromRoleID(roleID int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`role_id` = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量查找
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetBatchFromRoleID(roleIDs []int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`role_id` IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromAdminOption 通过admin_option获取内容
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetFromAdminOption(adminOption int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`admin_option` = ?", adminOption).Find(&results).Error

	return
}

// GetBatchFromAdminOption 批量查找
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetBatchFromAdminOption(adminOptions []int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`admin_option` IN (?)", adminOptions).Find(&results).Error

	return
}

// GetFromDisableFlag 通过disable_flag获取内容
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetFromDisableFlag(disableFlag int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`disable_flag` = ?", disableFlag).Find(&results).Error

	return
}

// GetBatchFromDisableFlag 批量查找
func (obj *_AllTenantRoleGranteeMapHistoryMgr) GetBatchFromDisableFlag(disableFlags []int64) (results []*AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`disable_flag` IN (?)", disableFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantRoleGranteeMapHistoryMgr) FetchByPrimaryKey(tenantID int64, granteeID int64, roleID int64, schemaVersion int64) (result AllTenantRoleGranteeMapHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRoleGranteeMapHistory{}).Where("`tenant_id` = ? AND `grantee_id` = ? AND `role_id` = ? AND `schema_version` = ?", tenantID, granteeID, roleID, schemaVersion).First(&result).Error

	return
}
