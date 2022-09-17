package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantSysauthHistoryMgr struct {
	*_BaseMgr
}

// AllTenantSysauthHistoryMgr open func
func AllTenantSysauthHistoryMgr(db *gorm.DB) *_AllTenantSysauthHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantSysauthHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantSysauthHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_sysauth_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantSysauthHistoryMgr) GetTableName() string {
	return "__all_tenant_sysauth_history"
}

// Reset 重置gorm会话
func (obj *_AllTenantSysauthHistoryMgr) Reset() *_AllTenantSysauthHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantSysauthHistoryMgr) Get() (result AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantSysauthHistoryMgr) Gets() (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantSysauthHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantSysauthHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantSysauthHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantSysauthHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithGranteeID grantee_id获取
func (obj *_AllTenantSysauthHistoryMgr) WithGranteeID(granteeID int64) Option {
	return optionFunc(func(o *options) { o.query["grantee_id"] = granteeID })
}

// WithPrivID priv_id获取
func (obj *_AllTenantSysauthHistoryMgr) WithPrivID(privID int64) Option {
	return optionFunc(func(o *options) { o.query["priv_id"] = privID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantSysauthHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTenantSysauthHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithPrivOption priv_option获取
func (obj *_AllTenantSysauthHistoryMgr) WithPrivOption(privOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_option"] = privOption })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantSysauthHistoryMgr) GetByOption(opts ...Option) (result AllTenantSysauthHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantSysauthHistoryMgr) GetByOptions(opts ...Option) (results []*AllTenantSysauthHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantSysauthHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantSysauthHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where(options.query)
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
func (obj *_AllTenantSysauthHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantSysauthHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantSysauthHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantSysauthHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantSysauthHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantSysauthHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromGranteeID 通过grantee_id获取内容
func (obj *_AllTenantSysauthHistoryMgr) GetFromGranteeID(granteeID int64) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`grantee_id` = ?", granteeID).Find(&results).Error

	return
}

// GetBatchFromGranteeID 批量查找
func (obj *_AllTenantSysauthHistoryMgr) GetBatchFromGranteeID(granteeIDs []int64) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`grantee_id` IN (?)", granteeIDs).Find(&results).Error

	return
}

// GetFromPrivID 通过priv_id获取内容
func (obj *_AllTenantSysauthHistoryMgr) GetFromPrivID(privID int64) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`priv_id` = ?", privID).Find(&results).Error

	return
}

// GetBatchFromPrivID 批量查找
func (obj *_AllTenantSysauthHistoryMgr) GetBatchFromPrivID(privIDs []int64) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`priv_id` IN (?)", privIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantSysauthHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantSysauthHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTenantSysauthHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTenantSysauthHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromPrivOption 通过priv_option获取内容
func (obj *_AllTenantSysauthHistoryMgr) GetFromPrivOption(privOption int64) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`priv_option` = ?", privOption).Find(&results).Error

	return
}

// GetBatchFromPrivOption 批量查找
func (obj *_AllTenantSysauthHistoryMgr) GetBatchFromPrivOption(privOptions []int64) (results []*AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`priv_option` IN (?)", privOptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantSysauthHistoryMgr) FetchByPrimaryKey(tenantID int64, granteeID int64, privID int64, schemaVersion int64) (result AllTenantSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSysauthHistory{}).Where("`tenant_id` = ? AND `grantee_id` = ? AND `priv_id` = ? AND `schema_version` = ?", tenantID, granteeID, privID, schemaVersion).First(&result).Error

	return
}
