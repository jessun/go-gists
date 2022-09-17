package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualTablePrivilegeHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualTablePrivilegeHistoryMgr open func
func AllVirtualTablePrivilegeHistoryMgr(db *gorm.DB) *_AllVirtualTablePrivilegeHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTablePrivilegeHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTablePrivilegeHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_table_privilege_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetTableName() string {
	return "__all_virtual_table_privilege_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTablePrivilegeHistoryMgr) Reset() *_AllVirtualTablePrivilegeHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) Get() (result AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTablePrivilegeHistoryMgr) Gets() (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTablePrivilegeHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserID user_id获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithDatabaseName database_name获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithTableName table_name获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithPrivAlter priv_alter获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithPrivAlter(privAlter int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter"] = privAlter })
}

// WithPrivCreate priv_create获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithPrivCreate(privCreate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create"] = privCreate })
}

// WithPrivDelete priv_delete获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithPrivDelete(privDelete int64) Option {
	return optionFunc(func(o *options) { o.query["priv_delete"] = privDelete })
}

// WithPrivDrop priv_drop获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithPrivDrop(privDrop int64) Option {
	return optionFunc(func(o *options) { o.query["priv_drop"] = privDrop })
}

// WithPrivGrantOption priv_grant_option获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithPrivGrantOption(privGrantOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_grant_option"] = privGrantOption })
}

// WithPrivInsert priv_insert获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithPrivInsert(privInsert int64) Option {
	return optionFunc(func(o *options) { o.query["priv_insert"] = privInsert })
}

// WithPrivUpdate priv_update获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithPrivUpdate(privUpdate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_update"] = privUpdate })
}

// WithPrivSelect priv_select获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithPrivSelect(privSelect int64) Option {
	return optionFunc(func(o *options) { o.query["priv_select"] = privSelect })
}

// WithPrivIndex priv_index获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithPrivIndex(privIndex int64) Option {
	return optionFunc(func(o *options) { o.query["priv_index"] = privIndex })
}

// WithPrivCreateView priv_create_view获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithPrivCreateView(privCreateView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_view"] = privCreateView })
}

// WithPrivShowView priv_show_view获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) WithPrivShowView(privShowView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_view"] = privShowView })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetByOption(opts ...Option) (result AllVirtualTablePrivilegeHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualTablePrivilegeHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTablePrivilegeHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTablePrivilegeHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where(options.query)
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
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromUserID(userID int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromUserID(userIDs []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromDatabaseName(databaseName string) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromTableName(tableName string) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromTableName(tableNames []string) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromPrivAlter 通过priv_alter获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromPrivAlter(privAlter int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_alter` = ?", privAlter).Find(&results).Error

	return
}

// GetBatchFromPrivAlter 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromPrivAlter(privAlters []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_alter` IN (?)", privAlters).Find(&results).Error

	return
}

// GetFromPrivCreate 通过priv_create获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromPrivCreate(privCreate int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_create` = ?", privCreate).Find(&results).Error

	return
}

// GetBatchFromPrivCreate 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromPrivCreate(privCreates []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_create` IN (?)", privCreates).Find(&results).Error

	return
}

// GetFromPrivDelete 通过priv_delete获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromPrivDelete(privDelete int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_delete` = ?", privDelete).Find(&results).Error

	return
}

// GetBatchFromPrivDelete 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromPrivDelete(privDeletes []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_delete` IN (?)", privDeletes).Find(&results).Error

	return
}

// GetFromPrivDrop 通过priv_drop获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromPrivDrop(privDrop int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_drop` = ?", privDrop).Find(&results).Error

	return
}

// GetBatchFromPrivDrop 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromPrivDrop(privDrops []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_drop` IN (?)", privDrops).Find(&results).Error

	return
}

// GetFromPrivGrantOption 通过priv_grant_option获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromPrivGrantOption(privGrantOption int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_grant_option` = ?", privGrantOption).Find(&results).Error

	return
}

// GetBatchFromPrivGrantOption 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromPrivGrantOption(privGrantOptions []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_grant_option` IN (?)", privGrantOptions).Find(&results).Error

	return
}

// GetFromPrivInsert 通过priv_insert获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromPrivInsert(privInsert int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_insert` = ?", privInsert).Find(&results).Error

	return
}

// GetBatchFromPrivInsert 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromPrivInsert(privInserts []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_insert` IN (?)", privInserts).Find(&results).Error

	return
}

// GetFromPrivUpdate 通过priv_update获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromPrivUpdate(privUpdate int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_update` = ?", privUpdate).Find(&results).Error

	return
}

// GetBatchFromPrivUpdate 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromPrivUpdate(privUpdates []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_update` IN (?)", privUpdates).Find(&results).Error

	return
}

// GetFromPrivSelect 通过priv_select获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromPrivSelect(privSelect int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_select` = ?", privSelect).Find(&results).Error

	return
}

// GetBatchFromPrivSelect 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromPrivSelect(privSelects []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_select` IN (?)", privSelects).Find(&results).Error

	return
}

// GetFromPrivIndex 通过priv_index获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromPrivIndex(privIndex int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_index` = ?", privIndex).Find(&results).Error

	return
}

// GetBatchFromPrivIndex 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromPrivIndex(privIndexs []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_index` IN (?)", privIndexs).Find(&results).Error

	return
}

// GetFromPrivCreateView 通过priv_create_view获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromPrivCreateView(privCreateView int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_create_view` = ?", privCreateView).Find(&results).Error

	return
}

// GetBatchFromPrivCreateView 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromPrivCreateView(privCreateViews []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_create_view` IN (?)", privCreateViews).Find(&results).Error

	return
}

// GetFromPrivShowView 通过priv_show_view获取内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetFromPrivShowView(privShowView int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_show_view` = ?", privShowView).Find(&results).Error

	return
}

// GetBatchFromPrivShowView 批量查找
func (obj *_AllVirtualTablePrivilegeHistoryMgr) GetBatchFromPrivShowView(privShowViews []int64) (results []*AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`priv_show_view` IN (?)", privShowViews).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTablePrivilegeHistoryMgr) FetchByPrimaryKey(tenantID int64, userID int64, databaseName string, tableName string, schemaVersion int64) (result AllVirtualTablePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTablePrivilegeHistory{}).Where("`tenant_id` = ? AND `user_id` = ? AND `database_name` = ? AND `table_name` = ? AND `schema_version` = ?", tenantID, userID, databaseName, tableName, schemaVersion).First(&result).Error

	return
}
