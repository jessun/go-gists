package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualForeignKeyHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualForeignKeyHistoryMgr open func
func AllVirtualForeignKeyHistoryMgr(db *gorm.DB) *_AllVirtualForeignKeyHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualForeignKeyHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualForeignKeyHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_foreign_key_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualForeignKeyHistoryMgr) GetTableName() string {
	return "__all_virtual_foreign_key_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualForeignKeyHistoryMgr) Reset() *_AllVirtualForeignKeyHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualForeignKeyHistoryMgr) Get() (result AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualForeignKeyHistoryMgr) Gets() (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualForeignKeyHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithForeignKeyID foreign_key_id获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithForeignKeyID(foreignKeyID int64) Option {
	return optionFunc(func(o *options) { o.query["foreign_key_id"] = foreignKeyID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithForeignKeyName foreign_key_name获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithForeignKeyName(foreignKeyName string) Option {
	return optionFunc(func(o *options) { o.query["foreign_key_name"] = foreignKeyName })
}

// WithChildTableID child_table_id获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithChildTableID(childTableID int64) Option {
	return optionFunc(func(o *options) { o.query["child_table_id"] = childTableID })
}

// WithParentTableID parent_table_id获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithParentTableID(parentTableID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_table_id"] = parentTableID })
}

// WithUpdateAction update_action获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithUpdateAction(updateAction int64) Option {
	return optionFunc(func(o *options) { o.query["update_action"] = updateAction })
}

// WithDeleteAction delete_action获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithDeleteAction(deleteAction int64) Option {
	return optionFunc(func(o *options) { o.query["delete_action"] = deleteAction })
}

// WithEnableFlag enable_flag获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithEnableFlag(enableFlag int8) Option {
	return optionFunc(func(o *options) { o.query["enable_flag"] = enableFlag })
}

// WithRefCstType ref_cst_type获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithRefCstType(refCstType int64) Option {
	return optionFunc(func(o *options) { o.query["ref_cst_type"] = refCstType })
}

// WithRefCstID ref_cst_id获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithRefCstID(refCstID int64) Option {
	return optionFunc(func(o *options) { o.query["ref_cst_id"] = refCstID })
}

// WithValidateFlag validate_flag获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithValidateFlag(validateFlag int8) Option {
	return optionFunc(func(o *options) { o.query["validate_flag"] = validateFlag })
}

// WithRelyFlag rely_flag获取
func (obj *_AllVirtualForeignKeyHistoryMgr) WithRelyFlag(relyFlag int8) Option {
	return optionFunc(func(o *options) { o.query["rely_flag"] = relyFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualForeignKeyHistoryMgr) GetByOption(opts ...Option) (result AllVirtualForeignKeyHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualForeignKeyHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualForeignKeyHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualForeignKeyHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualForeignKeyHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where(options.query)
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
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromForeignKeyID 通过foreign_key_id获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromForeignKeyID(foreignKeyID int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`foreign_key_id` = ?", foreignKeyID).Find(&results).Error

	return
}

// GetBatchFromForeignKeyID 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromForeignKeyID(foreignKeyIDs []int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`foreign_key_id` IN (?)", foreignKeyIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromForeignKeyName 通过foreign_key_name获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromForeignKeyName(foreignKeyName string) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`foreign_key_name` = ?", foreignKeyName).Find(&results).Error

	return
}

// GetBatchFromForeignKeyName 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromForeignKeyName(foreignKeyNames []string) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`foreign_key_name` IN (?)", foreignKeyNames).Find(&results).Error

	return
}

// GetFromChildTableID 通过child_table_id获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromChildTableID(childTableID int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`child_table_id` = ?", childTableID).Find(&results).Error

	return
}

// GetBatchFromChildTableID 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromChildTableID(childTableIDs []int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`child_table_id` IN (?)", childTableIDs).Find(&results).Error

	return
}

// GetFromParentTableID 通过parent_table_id获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromParentTableID(parentTableID int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`parent_table_id` = ?", parentTableID).Find(&results).Error

	return
}

// GetBatchFromParentTableID 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromParentTableID(parentTableIDs []int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`parent_table_id` IN (?)", parentTableIDs).Find(&results).Error

	return
}

// GetFromUpdateAction 通过update_action获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromUpdateAction(updateAction int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`update_action` = ?", updateAction).Find(&results).Error

	return
}

// GetBatchFromUpdateAction 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromUpdateAction(updateActions []int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`update_action` IN (?)", updateActions).Find(&results).Error

	return
}

// GetFromDeleteAction 通过delete_action获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromDeleteAction(deleteAction int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`delete_action` = ?", deleteAction).Find(&results).Error

	return
}

// GetBatchFromDeleteAction 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromDeleteAction(deleteActions []int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`delete_action` IN (?)", deleteActions).Find(&results).Error

	return
}

// GetFromEnableFlag 通过enable_flag获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromEnableFlag(enableFlag int8) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`enable_flag` = ?", enableFlag).Find(&results).Error

	return
}

// GetBatchFromEnableFlag 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromEnableFlag(enableFlags []int8) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`enable_flag` IN (?)", enableFlags).Find(&results).Error

	return
}

// GetFromRefCstType 通过ref_cst_type获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromRefCstType(refCstType int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`ref_cst_type` = ?", refCstType).Find(&results).Error

	return
}

// GetBatchFromRefCstType 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromRefCstType(refCstTypes []int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`ref_cst_type` IN (?)", refCstTypes).Find(&results).Error

	return
}

// GetFromRefCstID 通过ref_cst_id获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromRefCstID(refCstID int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`ref_cst_id` = ?", refCstID).Find(&results).Error

	return
}

// GetBatchFromRefCstID 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromRefCstID(refCstIDs []int64) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`ref_cst_id` IN (?)", refCstIDs).Find(&results).Error

	return
}

// GetFromValidateFlag 通过validate_flag获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromValidateFlag(validateFlag int8) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`validate_flag` = ?", validateFlag).Find(&results).Error

	return
}

// GetBatchFromValidateFlag 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromValidateFlag(validateFlags []int8) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`validate_flag` IN (?)", validateFlags).Find(&results).Error

	return
}

// GetFromRelyFlag 通过rely_flag获取内容
func (obj *_AllVirtualForeignKeyHistoryMgr) GetFromRelyFlag(relyFlag int8) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`rely_flag` = ?", relyFlag).Find(&results).Error

	return
}

// GetBatchFromRelyFlag 批量查找
func (obj *_AllVirtualForeignKeyHistoryMgr) GetBatchFromRelyFlag(relyFlags []int8) (results []*AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`rely_flag` IN (?)", relyFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualForeignKeyHistoryMgr) FetchByPrimaryKey(tenantID int64, foreignKeyID int64, schemaVersion int64) (result AllVirtualForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyHistory{}).Where("`tenant_id` = ? AND `foreign_key_id` = ? AND `schema_version` = ?", tenantID, foreignKeyID, schemaVersion).First(&result).Error

	return
}
