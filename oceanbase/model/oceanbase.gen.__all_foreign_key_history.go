package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllForeignKeyHistoryMgr struct {
	*_BaseMgr
}

// AllForeignKeyHistoryMgr open func
func AllForeignKeyHistoryMgr(db *gorm.DB) *_AllForeignKeyHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllForeignKeyHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllForeignKeyHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_foreign_key_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllForeignKeyHistoryMgr) GetTableName() string {
	return "__all_foreign_key_history"
}

// Reset 重置gorm会话
func (obj *_AllForeignKeyHistoryMgr) Reset() *_AllForeignKeyHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllForeignKeyHistoryMgr) Get() (result AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllForeignKeyHistoryMgr) Gets() (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllForeignKeyHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllForeignKeyHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllForeignKeyHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllForeignKeyHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithForeignKeyID foreign_key_id获取
func (obj *_AllForeignKeyHistoryMgr) WithForeignKeyID(foreignKeyID int64) Option {
	return optionFunc(func(o *options) { o.query["foreign_key_id"] = foreignKeyID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllForeignKeyHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllForeignKeyHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithForeignKeyName foreign_key_name获取
func (obj *_AllForeignKeyHistoryMgr) WithForeignKeyName(foreignKeyName string) Option {
	return optionFunc(func(o *options) { o.query["foreign_key_name"] = foreignKeyName })
}

// WithChildTableID child_table_id获取
func (obj *_AllForeignKeyHistoryMgr) WithChildTableID(childTableID int64) Option {
	return optionFunc(func(o *options) { o.query["child_table_id"] = childTableID })
}

// WithParentTableID parent_table_id获取
func (obj *_AllForeignKeyHistoryMgr) WithParentTableID(parentTableID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_table_id"] = parentTableID })
}

// WithUpdateAction update_action获取
func (obj *_AllForeignKeyHistoryMgr) WithUpdateAction(updateAction int64) Option {
	return optionFunc(func(o *options) { o.query["update_action"] = updateAction })
}

// WithDeleteAction delete_action获取
func (obj *_AllForeignKeyHistoryMgr) WithDeleteAction(deleteAction int64) Option {
	return optionFunc(func(o *options) { o.query["delete_action"] = deleteAction })
}

// WithEnableFlag enable_flag获取
func (obj *_AllForeignKeyHistoryMgr) WithEnableFlag(enableFlag int8) Option {
	return optionFunc(func(o *options) { o.query["enable_flag"] = enableFlag })
}

// WithRefCstType ref_cst_type获取
func (obj *_AllForeignKeyHistoryMgr) WithRefCstType(refCstType int64) Option {
	return optionFunc(func(o *options) { o.query["ref_cst_type"] = refCstType })
}

// WithRefCstID ref_cst_id获取
func (obj *_AllForeignKeyHistoryMgr) WithRefCstID(refCstID int64) Option {
	return optionFunc(func(o *options) { o.query["ref_cst_id"] = refCstID })
}

// WithValidateFlag validate_flag获取
func (obj *_AllForeignKeyHistoryMgr) WithValidateFlag(validateFlag int8) Option {
	return optionFunc(func(o *options) { o.query["validate_flag"] = validateFlag })
}

// WithRelyFlag rely_flag获取
func (obj *_AllForeignKeyHistoryMgr) WithRelyFlag(relyFlag int8) Option {
	return optionFunc(func(o *options) { o.query["rely_flag"] = relyFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllForeignKeyHistoryMgr) GetByOption(opts ...Option) (result AllForeignKeyHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllForeignKeyHistoryMgr) GetByOptions(opts ...Option) (results []*AllForeignKeyHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllForeignKeyHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllForeignKeyHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where(options.query)
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
func (obj *_AllForeignKeyHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromForeignKeyID 通过foreign_key_id获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromForeignKeyID(foreignKeyID int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`foreign_key_id` = ?", foreignKeyID).Find(&results).Error

	return
}

// GetBatchFromForeignKeyID 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromForeignKeyID(foreignKeyIDs []int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`foreign_key_id` IN (?)", foreignKeyIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromForeignKeyName 通过foreign_key_name获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromForeignKeyName(foreignKeyName string) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`foreign_key_name` = ?", foreignKeyName).Find(&results).Error

	return
}

// GetBatchFromForeignKeyName 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromForeignKeyName(foreignKeyNames []string) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`foreign_key_name` IN (?)", foreignKeyNames).Find(&results).Error

	return
}

// GetFromChildTableID 通过child_table_id获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromChildTableID(childTableID int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`child_table_id` = ?", childTableID).Find(&results).Error

	return
}

// GetBatchFromChildTableID 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromChildTableID(childTableIDs []int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`child_table_id` IN (?)", childTableIDs).Find(&results).Error

	return
}

// GetFromParentTableID 通过parent_table_id获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromParentTableID(parentTableID int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`parent_table_id` = ?", parentTableID).Find(&results).Error

	return
}

// GetBatchFromParentTableID 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromParentTableID(parentTableIDs []int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`parent_table_id` IN (?)", parentTableIDs).Find(&results).Error

	return
}

// GetFromUpdateAction 通过update_action获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromUpdateAction(updateAction int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`update_action` = ?", updateAction).Find(&results).Error

	return
}

// GetBatchFromUpdateAction 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromUpdateAction(updateActions []int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`update_action` IN (?)", updateActions).Find(&results).Error

	return
}

// GetFromDeleteAction 通过delete_action获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromDeleteAction(deleteAction int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`delete_action` = ?", deleteAction).Find(&results).Error

	return
}

// GetBatchFromDeleteAction 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromDeleteAction(deleteActions []int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`delete_action` IN (?)", deleteActions).Find(&results).Error

	return
}

// GetFromEnableFlag 通过enable_flag获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromEnableFlag(enableFlag int8) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`enable_flag` = ?", enableFlag).Find(&results).Error

	return
}

// GetBatchFromEnableFlag 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromEnableFlag(enableFlags []int8) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`enable_flag` IN (?)", enableFlags).Find(&results).Error

	return
}

// GetFromRefCstType 通过ref_cst_type获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromRefCstType(refCstType int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`ref_cst_type` = ?", refCstType).Find(&results).Error

	return
}

// GetBatchFromRefCstType 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromRefCstType(refCstTypes []int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`ref_cst_type` IN (?)", refCstTypes).Find(&results).Error

	return
}

// GetFromRefCstID 通过ref_cst_id获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromRefCstID(refCstID int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`ref_cst_id` = ?", refCstID).Find(&results).Error

	return
}

// GetBatchFromRefCstID 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromRefCstID(refCstIDs []int64) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`ref_cst_id` IN (?)", refCstIDs).Find(&results).Error

	return
}

// GetFromValidateFlag 通过validate_flag获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromValidateFlag(validateFlag int8) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`validate_flag` = ?", validateFlag).Find(&results).Error

	return
}

// GetBatchFromValidateFlag 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromValidateFlag(validateFlags []int8) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`validate_flag` IN (?)", validateFlags).Find(&results).Error

	return
}

// GetFromRelyFlag 通过rely_flag获取内容
func (obj *_AllForeignKeyHistoryMgr) GetFromRelyFlag(relyFlag int8) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`rely_flag` = ?", relyFlag).Find(&results).Error

	return
}

// GetBatchFromRelyFlag 批量查找
func (obj *_AllForeignKeyHistoryMgr) GetBatchFromRelyFlag(relyFlags []int8) (results []*AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`rely_flag` IN (?)", relyFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllForeignKeyHistoryMgr) FetchByPrimaryKey(tenantID int64, foreignKeyID int64, schemaVersion int64) (result AllForeignKeyHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyHistory{}).Where("`tenant_id` = ? AND `foreign_key_id` = ? AND `schema_version` = ?", tenantID, foreignKeyID, schemaVersion).First(&result).Error

	return
}
