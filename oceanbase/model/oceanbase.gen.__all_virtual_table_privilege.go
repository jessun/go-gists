package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualTablePrivilegeMgr struct {
	*_BaseMgr
}

// AllVirtualTablePrivilegeMgr open func
func AllVirtualTablePrivilegeMgr(db *gorm.DB) *_AllVirtualTablePrivilegeMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTablePrivilegeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTablePrivilegeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_table_privilege"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTablePrivilegeMgr) GetTableName() string {
	return "__all_virtual_table_privilege"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTablePrivilegeMgr) Reset() *_AllVirtualTablePrivilegeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTablePrivilegeMgr) Get() (result AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTablePrivilegeMgr) Gets() (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTablePrivilegeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTablePrivilegeMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserID user_id获取
func (obj *_AllVirtualTablePrivilegeMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithDatabaseName database_name获取
func (obj *_AllVirtualTablePrivilegeMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithTableName table_name获取
func (obj *_AllVirtualTablePrivilegeMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTablePrivilegeMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTablePrivilegeMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithPrivAlter priv_alter获取
func (obj *_AllVirtualTablePrivilegeMgr) WithPrivAlter(privAlter int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter"] = privAlter })
}

// WithPrivCreate priv_create获取
func (obj *_AllVirtualTablePrivilegeMgr) WithPrivCreate(privCreate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create"] = privCreate })
}

// WithPrivDelete priv_delete获取
func (obj *_AllVirtualTablePrivilegeMgr) WithPrivDelete(privDelete int64) Option {
	return optionFunc(func(o *options) { o.query["priv_delete"] = privDelete })
}

// WithPrivDrop priv_drop获取
func (obj *_AllVirtualTablePrivilegeMgr) WithPrivDrop(privDrop int64) Option {
	return optionFunc(func(o *options) { o.query["priv_drop"] = privDrop })
}

// WithPrivGrantOption priv_grant_option获取
func (obj *_AllVirtualTablePrivilegeMgr) WithPrivGrantOption(privGrantOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_grant_option"] = privGrantOption })
}

// WithPrivInsert priv_insert获取
func (obj *_AllVirtualTablePrivilegeMgr) WithPrivInsert(privInsert int64) Option {
	return optionFunc(func(o *options) { o.query["priv_insert"] = privInsert })
}

// WithPrivUpdate priv_update获取
func (obj *_AllVirtualTablePrivilegeMgr) WithPrivUpdate(privUpdate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_update"] = privUpdate })
}

// WithPrivSelect priv_select获取
func (obj *_AllVirtualTablePrivilegeMgr) WithPrivSelect(privSelect int64) Option {
	return optionFunc(func(o *options) { o.query["priv_select"] = privSelect })
}

// WithPrivIndex priv_index获取
func (obj *_AllVirtualTablePrivilegeMgr) WithPrivIndex(privIndex int64) Option {
	return optionFunc(func(o *options) { o.query["priv_index"] = privIndex })
}

// WithPrivCreateView priv_create_view获取
func (obj *_AllVirtualTablePrivilegeMgr) WithPrivCreateView(privCreateView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_view"] = privCreateView })
}

// WithPrivShowView priv_show_view获取
func (obj *_AllVirtualTablePrivilegeMgr) WithPrivShowView(privShowView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_view"] = privShowView })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTablePrivilegeMgr) GetByOption(opts ...Option) (result AllVirtualTablePrivilege, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTablePrivilegeMgr) GetByOptions(opts ...Option) (results []*AllVirtualTablePrivilege, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTablePrivilegeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTablePrivilege, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where(options.query)
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
func (obj *_AllVirtualTablePrivilegeMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromUserID(userID int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromUserID(userIDs []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromDatabaseName(databaseName string) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromTableName(tableName string) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromTableName(tableNames []string) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromPrivAlter 通过priv_alter获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromPrivAlter(privAlter int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_alter` = ?", privAlter).Find(&results).Error

	return
}

// GetBatchFromPrivAlter 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromPrivAlter(privAlters []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_alter` IN (?)", privAlters).Find(&results).Error

	return
}

// GetFromPrivCreate 通过priv_create获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromPrivCreate(privCreate int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_create` = ?", privCreate).Find(&results).Error

	return
}

// GetBatchFromPrivCreate 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromPrivCreate(privCreates []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_create` IN (?)", privCreates).Find(&results).Error

	return
}

// GetFromPrivDelete 通过priv_delete获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromPrivDelete(privDelete int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_delete` = ?", privDelete).Find(&results).Error

	return
}

// GetBatchFromPrivDelete 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromPrivDelete(privDeletes []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_delete` IN (?)", privDeletes).Find(&results).Error

	return
}

// GetFromPrivDrop 通过priv_drop获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromPrivDrop(privDrop int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_drop` = ?", privDrop).Find(&results).Error

	return
}

// GetBatchFromPrivDrop 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromPrivDrop(privDrops []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_drop` IN (?)", privDrops).Find(&results).Error

	return
}

// GetFromPrivGrantOption 通过priv_grant_option获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromPrivGrantOption(privGrantOption int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_grant_option` = ?", privGrantOption).Find(&results).Error

	return
}

// GetBatchFromPrivGrantOption 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromPrivGrantOption(privGrantOptions []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_grant_option` IN (?)", privGrantOptions).Find(&results).Error

	return
}

// GetFromPrivInsert 通过priv_insert获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromPrivInsert(privInsert int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_insert` = ?", privInsert).Find(&results).Error

	return
}

// GetBatchFromPrivInsert 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromPrivInsert(privInserts []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_insert` IN (?)", privInserts).Find(&results).Error

	return
}

// GetFromPrivUpdate 通过priv_update获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromPrivUpdate(privUpdate int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_update` = ?", privUpdate).Find(&results).Error

	return
}

// GetBatchFromPrivUpdate 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromPrivUpdate(privUpdates []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_update` IN (?)", privUpdates).Find(&results).Error

	return
}

// GetFromPrivSelect 通过priv_select获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromPrivSelect(privSelect int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_select` = ?", privSelect).Find(&results).Error

	return
}

// GetBatchFromPrivSelect 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromPrivSelect(privSelects []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_select` IN (?)", privSelects).Find(&results).Error

	return
}

// GetFromPrivIndex 通过priv_index获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromPrivIndex(privIndex int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_index` = ?", privIndex).Find(&results).Error

	return
}

// GetBatchFromPrivIndex 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromPrivIndex(privIndexs []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_index` IN (?)", privIndexs).Find(&results).Error

	return
}

// GetFromPrivCreateView 通过priv_create_view获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromPrivCreateView(privCreateView int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_create_view` = ?", privCreateView).Find(&results).Error

	return
}

// GetBatchFromPrivCreateView 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromPrivCreateView(privCreateViews []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_create_view` IN (?)", privCreateViews).Find(&results).Error

	return
}

// GetFromPrivShowView 通过priv_show_view获取内容
func (obj *_AllVirtualTablePrivilegeMgr) GetFromPrivShowView(privShowView int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_show_view` = ?", privShowView).Find(&results).Error

	return
}

// GetBatchFromPrivShowView 批量查找
func (obj *_AllVirtualTablePrivilegeMgr) GetBatchFromPrivShowView(privShowViews []int64) (results []*AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`priv_show_view` IN (?)", privShowViews).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTablePrivilegeMgr) FetchByPrimaryKey(tenantID int64, userID int64, databaseName string, tableName string) (result AllVirtualTablePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilege{}).Where("`tenant_id` = ? AND `user_id` = ? AND `database_name` = ? AND `table_name` = ?", tenantID, userID, databaseName, tableName).First(&result).Error

	return
}
