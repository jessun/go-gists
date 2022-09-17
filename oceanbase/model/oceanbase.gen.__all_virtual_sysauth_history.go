package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualSysauthHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualSysauthHistoryMgr open func
func AllVirtualSysauthHistoryMgr(db *gorm.DB) *_AllVirtualSysauthHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSysauthHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSysauthHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sysauth_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSysauthHistoryMgr) GetTableName() string {
	return "__all_virtual_sysauth_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSysauthHistoryMgr) Reset() *_AllVirtualSysauthHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSysauthHistoryMgr) Get() (result AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSysauthHistoryMgr) Gets() (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSysauthHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSysauthHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithGranteeID grantee_id获取
func (obj *_AllVirtualSysauthHistoryMgr) WithGranteeID(granteeID int64) Option {
	return optionFunc(func(o *options) { o.query["grantee_id"] = granteeID })
}

// WithPrivID priv_id获取
func (obj *_AllVirtualSysauthHistoryMgr) WithPrivID(privID int64) Option {
	return optionFunc(func(o *options) { o.query["priv_id"] = privID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualSysauthHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSysauthHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSysauthHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualSysauthHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithPrivOption priv_option获取
func (obj *_AllVirtualSysauthHistoryMgr) WithPrivOption(privOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_option"] = privOption })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSysauthHistoryMgr) GetByOption(opts ...Option) (result AllVirtualSysauthHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSysauthHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualSysauthHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSysauthHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSysauthHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where(options.query)
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
func (obj *_AllVirtualSysauthHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSysauthHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromGranteeID 通过grantee_id获取内容
func (obj *_AllVirtualSysauthHistoryMgr) GetFromGranteeID(granteeID int64) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`grantee_id` = ?", granteeID).Find(&results).Error

	return
}

// GetBatchFromGranteeID 批量查找
func (obj *_AllVirtualSysauthHistoryMgr) GetBatchFromGranteeID(granteeIDs []int64) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`grantee_id` IN (?)", granteeIDs).Find(&results).Error

	return
}

// GetFromPrivID 通过priv_id获取内容
func (obj *_AllVirtualSysauthHistoryMgr) GetFromPrivID(privID int64) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`priv_id` = ?", privID).Find(&results).Error

	return
}

// GetBatchFromPrivID 批量查找
func (obj *_AllVirtualSysauthHistoryMgr) GetBatchFromPrivID(privIDs []int64) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`priv_id` IN (?)", privIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualSysauthHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualSysauthHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSysauthHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSysauthHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSysauthHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSysauthHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualSysauthHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualSysauthHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromPrivOption 通过priv_option获取内容
func (obj *_AllVirtualSysauthHistoryMgr) GetFromPrivOption(privOption int64) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`priv_option` = ?", privOption).Find(&results).Error

	return
}

// GetBatchFromPrivOption 批量查找
func (obj *_AllVirtualSysauthHistoryMgr) GetBatchFromPrivOption(privOptions []int64) (results []*AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`priv_option` IN (?)", privOptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSysauthHistoryMgr) FetchByPrimaryKey(tenantID int64, granteeID int64, privID int64, schemaVersion int64) (result AllVirtualSysauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysauthHistory{}).Where("`tenant_id` = ? AND `grantee_id` = ? AND `priv_id` = ? AND `schema_version` = ?", tenantID, granteeID, privID, schemaVersion).First(&result).Error

	return
}
