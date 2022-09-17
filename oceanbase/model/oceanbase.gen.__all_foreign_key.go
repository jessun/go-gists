package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllForeignKeyMgr struct {
	*_BaseMgr
}

// AllForeignKeyMgr open func
func AllForeignKeyMgr(db *gorm.DB) *_AllForeignKeyMgr {
	if db == nil {
		panic(fmt.Errorf("AllForeignKeyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllForeignKeyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_foreign_key"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllForeignKeyMgr) GetTableName() string {
	return "__all_foreign_key"
}

// Reset 重置gorm会话
func (obj *_AllForeignKeyMgr) Reset() *_AllForeignKeyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllForeignKeyMgr) Get() (result AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllForeignKeyMgr) Gets() (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllForeignKeyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllForeignKeyMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllForeignKeyMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllForeignKeyMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithForeignKeyID foreign_key_id获取
func (obj *_AllForeignKeyMgr) WithForeignKeyID(foreignKeyID int64) Option {
	return optionFunc(func(o *options) { o.query["foreign_key_id"] = foreignKeyID })
}

// WithForeignKeyName foreign_key_name获取
func (obj *_AllForeignKeyMgr) WithForeignKeyName(foreignKeyName string) Option {
	return optionFunc(func(o *options) { o.query["foreign_key_name"] = foreignKeyName })
}

// WithChildTableID child_table_id获取
func (obj *_AllForeignKeyMgr) WithChildTableID(childTableID int64) Option {
	return optionFunc(func(o *options) { o.query["child_table_id"] = childTableID })
}

// WithParentTableID parent_table_id获取
func (obj *_AllForeignKeyMgr) WithParentTableID(parentTableID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_table_id"] = parentTableID })
}

// WithUpdateAction update_action获取
func (obj *_AllForeignKeyMgr) WithUpdateAction(updateAction int64) Option {
	return optionFunc(func(o *options) { o.query["update_action"] = updateAction })
}

// WithDeleteAction delete_action获取
func (obj *_AllForeignKeyMgr) WithDeleteAction(deleteAction int64) Option {
	return optionFunc(func(o *options) { o.query["delete_action"] = deleteAction })
}

// WithEnableFlag enable_flag获取
func (obj *_AllForeignKeyMgr) WithEnableFlag(enableFlag int8) Option {
	return optionFunc(func(o *options) { o.query["enable_flag"] = enableFlag })
}

// WithRefCstType ref_cst_type获取
func (obj *_AllForeignKeyMgr) WithRefCstType(refCstType int64) Option {
	return optionFunc(func(o *options) { o.query["ref_cst_type"] = refCstType })
}

// WithRefCstID ref_cst_id获取
func (obj *_AllForeignKeyMgr) WithRefCstID(refCstID int64) Option {
	return optionFunc(func(o *options) { o.query["ref_cst_id"] = refCstID })
}

// WithValidateFlag validate_flag获取
func (obj *_AllForeignKeyMgr) WithValidateFlag(validateFlag int8) Option {
	return optionFunc(func(o *options) { o.query["validate_flag"] = validateFlag })
}

// WithRelyFlag rely_flag获取
func (obj *_AllForeignKeyMgr) WithRelyFlag(relyFlag int8) Option {
	return optionFunc(func(o *options) { o.query["rely_flag"] = relyFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllForeignKeyMgr) GetByOption(opts ...Option) (result AllForeignKey, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllForeignKeyMgr) GetByOptions(opts ...Option) (results []*AllForeignKey, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllForeignKeyMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllForeignKey, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where(options.query)
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
func (obj *_AllForeignKeyMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllForeignKeyMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllForeignKeyMgr) GetFromTenantID(tenantID int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromForeignKeyID 通过foreign_key_id获取内容
func (obj *_AllForeignKeyMgr) GetFromForeignKeyID(foreignKeyID int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`foreign_key_id` = ?", foreignKeyID).Find(&results).Error

	return
}

// GetBatchFromForeignKeyID 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromForeignKeyID(foreignKeyIDs []int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`foreign_key_id` IN (?)", foreignKeyIDs).Find(&results).Error

	return
}

// GetFromForeignKeyName 通过foreign_key_name获取内容
func (obj *_AllForeignKeyMgr) GetFromForeignKeyName(foreignKeyName string) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`foreign_key_name` = ?", foreignKeyName).Find(&results).Error

	return
}

// GetBatchFromForeignKeyName 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromForeignKeyName(foreignKeyNames []string) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`foreign_key_name` IN (?)", foreignKeyNames).Find(&results).Error

	return
}

// GetFromChildTableID 通过child_table_id获取内容
func (obj *_AllForeignKeyMgr) GetFromChildTableID(childTableID int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`child_table_id` = ?", childTableID).Find(&results).Error

	return
}

// GetBatchFromChildTableID 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromChildTableID(childTableIDs []int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`child_table_id` IN (?)", childTableIDs).Find(&results).Error

	return
}

// GetFromParentTableID 通过parent_table_id获取内容
func (obj *_AllForeignKeyMgr) GetFromParentTableID(parentTableID int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`parent_table_id` = ?", parentTableID).Find(&results).Error

	return
}

// GetBatchFromParentTableID 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromParentTableID(parentTableIDs []int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`parent_table_id` IN (?)", parentTableIDs).Find(&results).Error

	return
}

// GetFromUpdateAction 通过update_action获取内容
func (obj *_AllForeignKeyMgr) GetFromUpdateAction(updateAction int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`update_action` = ?", updateAction).Find(&results).Error

	return
}

// GetBatchFromUpdateAction 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromUpdateAction(updateActions []int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`update_action` IN (?)", updateActions).Find(&results).Error

	return
}

// GetFromDeleteAction 通过delete_action获取内容
func (obj *_AllForeignKeyMgr) GetFromDeleteAction(deleteAction int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`delete_action` = ?", deleteAction).Find(&results).Error

	return
}

// GetBatchFromDeleteAction 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromDeleteAction(deleteActions []int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`delete_action` IN (?)", deleteActions).Find(&results).Error

	return
}

// GetFromEnableFlag 通过enable_flag获取内容
func (obj *_AllForeignKeyMgr) GetFromEnableFlag(enableFlag int8) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`enable_flag` = ?", enableFlag).Find(&results).Error

	return
}

// GetBatchFromEnableFlag 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromEnableFlag(enableFlags []int8) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`enable_flag` IN (?)", enableFlags).Find(&results).Error

	return
}

// GetFromRefCstType 通过ref_cst_type获取内容
func (obj *_AllForeignKeyMgr) GetFromRefCstType(refCstType int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`ref_cst_type` = ?", refCstType).Find(&results).Error

	return
}

// GetBatchFromRefCstType 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromRefCstType(refCstTypes []int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`ref_cst_type` IN (?)", refCstTypes).Find(&results).Error

	return
}

// GetFromRefCstID 通过ref_cst_id获取内容
func (obj *_AllForeignKeyMgr) GetFromRefCstID(refCstID int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`ref_cst_id` = ?", refCstID).Find(&results).Error

	return
}

// GetBatchFromRefCstID 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromRefCstID(refCstIDs []int64) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`ref_cst_id` IN (?)", refCstIDs).Find(&results).Error

	return
}

// GetFromValidateFlag 通过validate_flag获取内容
func (obj *_AllForeignKeyMgr) GetFromValidateFlag(validateFlag int8) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`validate_flag` = ?", validateFlag).Find(&results).Error

	return
}

// GetBatchFromValidateFlag 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromValidateFlag(validateFlags []int8) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`validate_flag` IN (?)", validateFlags).Find(&results).Error

	return
}

// GetFromRelyFlag 通过rely_flag获取内容
func (obj *_AllForeignKeyMgr) GetFromRelyFlag(relyFlag int8) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`rely_flag` = ?", relyFlag).Find(&results).Error

	return
}

// GetBatchFromRelyFlag 批量查找
func (obj *_AllForeignKeyMgr) GetBatchFromRelyFlag(relyFlags []int8) (results []*AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`rely_flag` IN (?)", relyFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllForeignKeyMgr) FetchByPrimaryKey(tenantID int64, foreignKeyID int64) (result AllForeignKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKey{}).Where("`tenant_id` = ? AND `foreign_key_id` = ?", tenantID, foreignKeyID).First(&result).Error

	return
}
