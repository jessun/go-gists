package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualForeignKeyMgr struct {
	*_BaseMgr
}

// AllVirtualForeignKeyMgr open func
func AllVirtualForeignKeyMgr(db *gorm.DB) *_AllVirtualForeignKeyMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualForeignKeyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualForeignKeyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_foreign_key"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualForeignKeyMgr) GetTableName() string {
	return "__all_virtual_foreign_key"
}

// Reset 重置gorm会话
func (obj *_AllVirtualForeignKeyMgr) Reset() *_AllVirtualForeignKeyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualForeignKeyMgr) Get() (result AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualForeignKeyMgr) Gets() (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualForeignKeyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualForeignKeyMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithForeignKeyID foreign_key_id获取
func (obj *_AllVirtualForeignKeyMgr) WithForeignKeyID(foreignKeyID int64) Option {
	return optionFunc(func(o *options) { o.query["foreign_key_id"] = foreignKeyID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualForeignKeyMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualForeignKeyMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithForeignKeyName foreign_key_name获取
func (obj *_AllVirtualForeignKeyMgr) WithForeignKeyName(foreignKeyName string) Option {
	return optionFunc(func(o *options) { o.query["foreign_key_name"] = foreignKeyName })
}

// WithChildTableID child_table_id获取
func (obj *_AllVirtualForeignKeyMgr) WithChildTableID(childTableID int64) Option {
	return optionFunc(func(o *options) { o.query["child_table_id"] = childTableID })
}

// WithParentTableID parent_table_id获取
func (obj *_AllVirtualForeignKeyMgr) WithParentTableID(parentTableID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_table_id"] = parentTableID })
}

// WithUpdateAction update_action获取
func (obj *_AllVirtualForeignKeyMgr) WithUpdateAction(updateAction int64) Option {
	return optionFunc(func(o *options) { o.query["update_action"] = updateAction })
}

// WithDeleteAction delete_action获取
func (obj *_AllVirtualForeignKeyMgr) WithDeleteAction(deleteAction int64) Option {
	return optionFunc(func(o *options) { o.query["delete_action"] = deleteAction })
}

// WithEnableFlag enable_flag获取
func (obj *_AllVirtualForeignKeyMgr) WithEnableFlag(enableFlag int8) Option {
	return optionFunc(func(o *options) { o.query["enable_flag"] = enableFlag })
}

// WithRefCstType ref_cst_type获取
func (obj *_AllVirtualForeignKeyMgr) WithRefCstType(refCstType int64) Option {
	return optionFunc(func(o *options) { o.query["ref_cst_type"] = refCstType })
}

// WithRefCstID ref_cst_id获取
func (obj *_AllVirtualForeignKeyMgr) WithRefCstID(refCstID int64) Option {
	return optionFunc(func(o *options) { o.query["ref_cst_id"] = refCstID })
}

// WithValidateFlag validate_flag获取
func (obj *_AllVirtualForeignKeyMgr) WithValidateFlag(validateFlag int8) Option {
	return optionFunc(func(o *options) { o.query["validate_flag"] = validateFlag })
}

// WithRelyFlag rely_flag获取
func (obj *_AllVirtualForeignKeyMgr) WithRelyFlag(relyFlag int8) Option {
	return optionFunc(func(o *options) { o.query["rely_flag"] = relyFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualForeignKeyMgr) GetByOption(opts ...Option) (result AllVirtualForeignKey, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualForeignKeyMgr) GetByOptions(opts ...Option) (results []*AllVirtualForeignKey, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualForeignKeyMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualForeignKey, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where(options.query)
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
func (obj *_AllVirtualForeignKeyMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromForeignKeyID 通过foreign_key_id获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromForeignKeyID(foreignKeyID int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`foreign_key_id` = ?", foreignKeyID).Find(&results).Error

	return
}

// GetBatchFromForeignKeyID 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromForeignKeyID(foreignKeyIDs []int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`foreign_key_id` IN (?)", foreignKeyIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromForeignKeyName 通过foreign_key_name获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromForeignKeyName(foreignKeyName string) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`foreign_key_name` = ?", foreignKeyName).Find(&results).Error

	return
}

// GetBatchFromForeignKeyName 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromForeignKeyName(foreignKeyNames []string) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`foreign_key_name` IN (?)", foreignKeyNames).Find(&results).Error

	return
}

// GetFromChildTableID 通过child_table_id获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromChildTableID(childTableID int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`child_table_id` = ?", childTableID).Find(&results).Error

	return
}

// GetBatchFromChildTableID 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromChildTableID(childTableIDs []int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`child_table_id` IN (?)", childTableIDs).Find(&results).Error

	return
}

// GetFromParentTableID 通过parent_table_id获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromParentTableID(parentTableID int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`parent_table_id` = ?", parentTableID).Find(&results).Error

	return
}

// GetBatchFromParentTableID 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromParentTableID(parentTableIDs []int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`parent_table_id` IN (?)", parentTableIDs).Find(&results).Error

	return
}

// GetFromUpdateAction 通过update_action获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromUpdateAction(updateAction int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`update_action` = ?", updateAction).Find(&results).Error

	return
}

// GetBatchFromUpdateAction 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromUpdateAction(updateActions []int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`update_action` IN (?)", updateActions).Find(&results).Error

	return
}

// GetFromDeleteAction 通过delete_action获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromDeleteAction(deleteAction int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`delete_action` = ?", deleteAction).Find(&results).Error

	return
}

// GetBatchFromDeleteAction 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromDeleteAction(deleteActions []int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`delete_action` IN (?)", deleteActions).Find(&results).Error

	return
}

// GetFromEnableFlag 通过enable_flag获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromEnableFlag(enableFlag int8) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`enable_flag` = ?", enableFlag).Find(&results).Error

	return
}

// GetBatchFromEnableFlag 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromEnableFlag(enableFlags []int8) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`enable_flag` IN (?)", enableFlags).Find(&results).Error

	return
}

// GetFromRefCstType 通过ref_cst_type获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromRefCstType(refCstType int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`ref_cst_type` = ?", refCstType).Find(&results).Error

	return
}

// GetBatchFromRefCstType 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromRefCstType(refCstTypes []int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`ref_cst_type` IN (?)", refCstTypes).Find(&results).Error

	return
}

// GetFromRefCstID 通过ref_cst_id获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromRefCstID(refCstID int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`ref_cst_id` = ?", refCstID).Find(&results).Error

	return
}

// GetBatchFromRefCstID 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromRefCstID(refCstIDs []int64) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`ref_cst_id` IN (?)", refCstIDs).Find(&results).Error

	return
}

// GetFromValidateFlag 通过validate_flag获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromValidateFlag(validateFlag int8) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`validate_flag` = ?", validateFlag).Find(&results).Error

	return
}

// GetBatchFromValidateFlag 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromValidateFlag(validateFlags []int8) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`validate_flag` IN (?)", validateFlags).Find(&results).Error

	return
}

// GetFromRelyFlag 通过rely_flag获取内容
func (obj *_AllVirtualForeignKeyMgr) GetFromRelyFlag(relyFlag int8) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`rely_flag` = ?", relyFlag).Find(&results).Error

	return
}

// GetBatchFromRelyFlag 批量查找
func (obj *_AllVirtualForeignKeyMgr) GetBatchFromRelyFlag(relyFlags []int8) (results []*AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`rely_flag` IN (?)", relyFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualForeignKeyMgr) FetchByPrimaryKey(tenantID int64, foreignKeyID int64) (result AllVirtualForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKey{}).Where("`tenant_id` = ? AND `foreign_key_id` = ?", tenantID, foreignKeyID).First(&result).Error

	return
}
