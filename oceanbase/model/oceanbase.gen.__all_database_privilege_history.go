package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllDatabasePrivilegeHistoryMgr struct {
	*_BaseMgr
}

// AllDatabasePrivilegeHistoryMgr open func
func AllDatabasePrivilegeHistoryMgr(db *gorm.DB) *_AllDatabasePrivilegeHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllDatabasePrivilegeHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllDatabasePrivilegeHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_database_privilege_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllDatabasePrivilegeHistoryMgr) GetTableName() string {
	return "__all_database_privilege_history"
}

// Reset 重置gorm会话
func (obj *_AllDatabasePrivilegeHistoryMgr) Reset() *_AllDatabasePrivilegeHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllDatabasePrivilegeHistoryMgr) Get() (result AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllDatabasePrivilegeHistoryMgr) Gets() (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllDatabasePrivilegeHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserID user_id获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithDatabaseName database_name获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithSchemaVersion schema_version获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithPrivAlter priv_alter获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithPrivAlter(privAlter int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter"] = privAlter })
}

// WithPrivCreate priv_create获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithPrivCreate(privCreate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create"] = privCreate })
}

// WithPrivDelete priv_delete获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithPrivDelete(privDelete int64) Option {
	return optionFunc(func(o *options) { o.query["priv_delete"] = privDelete })
}

// WithPrivDrop priv_drop获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithPrivDrop(privDrop int64) Option {
	return optionFunc(func(o *options) { o.query["priv_drop"] = privDrop })
}

// WithPrivGrantOption priv_grant_option获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithPrivGrantOption(privGrantOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_grant_option"] = privGrantOption })
}

// WithPrivInsert priv_insert获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithPrivInsert(privInsert int64) Option {
	return optionFunc(func(o *options) { o.query["priv_insert"] = privInsert })
}

// WithPrivUpdate priv_update获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithPrivUpdate(privUpdate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_update"] = privUpdate })
}

// WithPrivSelect priv_select获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithPrivSelect(privSelect int64) Option {
	return optionFunc(func(o *options) { o.query["priv_select"] = privSelect })
}

// WithPrivIndex priv_index获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithPrivIndex(privIndex int64) Option {
	return optionFunc(func(o *options) { o.query["priv_index"] = privIndex })
}

// WithPrivCreateView priv_create_view获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithPrivCreateView(privCreateView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_view"] = privCreateView })
}

// WithPrivShowView priv_show_view获取
func (obj *_AllDatabasePrivilegeHistoryMgr) WithPrivShowView(privShowView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_view"] = privShowView })
}

// GetByOption 功能选项模式获取
func (obj *_AllDatabasePrivilegeHistoryMgr) GetByOption(opts ...Option) (result AllDatabasePrivilegeHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllDatabasePrivilegeHistoryMgr) GetByOptions(opts ...Option) (results []*AllDatabasePrivilegeHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllDatabasePrivilegeHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllDatabasePrivilegeHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where(options.query)
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
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromUserID(userID int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromUserID(userIDs []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromDatabaseName(databaseName string) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromPrivAlter 通过priv_alter获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromPrivAlter(privAlter int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_alter` = ?", privAlter).Find(&results).Error

	return
}

// GetBatchFromPrivAlter 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromPrivAlter(privAlters []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_alter` IN (?)", privAlters).Find(&results).Error

	return
}

// GetFromPrivCreate 通过priv_create获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromPrivCreate(privCreate int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_create` = ?", privCreate).Find(&results).Error

	return
}

// GetBatchFromPrivCreate 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromPrivCreate(privCreates []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_create` IN (?)", privCreates).Find(&results).Error

	return
}

// GetFromPrivDelete 通过priv_delete获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromPrivDelete(privDelete int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_delete` = ?", privDelete).Find(&results).Error

	return
}

// GetBatchFromPrivDelete 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromPrivDelete(privDeletes []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_delete` IN (?)", privDeletes).Find(&results).Error

	return
}

// GetFromPrivDrop 通过priv_drop获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromPrivDrop(privDrop int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_drop` = ?", privDrop).Find(&results).Error

	return
}

// GetBatchFromPrivDrop 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromPrivDrop(privDrops []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_drop` IN (?)", privDrops).Find(&results).Error

	return
}

// GetFromPrivGrantOption 通过priv_grant_option获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromPrivGrantOption(privGrantOption int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_grant_option` = ?", privGrantOption).Find(&results).Error

	return
}

// GetBatchFromPrivGrantOption 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromPrivGrantOption(privGrantOptions []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_grant_option` IN (?)", privGrantOptions).Find(&results).Error

	return
}

// GetFromPrivInsert 通过priv_insert获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromPrivInsert(privInsert int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_insert` = ?", privInsert).Find(&results).Error

	return
}

// GetBatchFromPrivInsert 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromPrivInsert(privInserts []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_insert` IN (?)", privInserts).Find(&results).Error

	return
}

// GetFromPrivUpdate 通过priv_update获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromPrivUpdate(privUpdate int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_update` = ?", privUpdate).Find(&results).Error

	return
}

// GetBatchFromPrivUpdate 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromPrivUpdate(privUpdates []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_update` IN (?)", privUpdates).Find(&results).Error

	return
}

// GetFromPrivSelect 通过priv_select获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromPrivSelect(privSelect int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_select` = ?", privSelect).Find(&results).Error

	return
}

// GetBatchFromPrivSelect 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromPrivSelect(privSelects []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_select` IN (?)", privSelects).Find(&results).Error

	return
}

// GetFromPrivIndex 通过priv_index获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromPrivIndex(privIndex int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_index` = ?", privIndex).Find(&results).Error

	return
}

// GetBatchFromPrivIndex 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromPrivIndex(privIndexs []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_index` IN (?)", privIndexs).Find(&results).Error

	return
}

// GetFromPrivCreateView 通过priv_create_view获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromPrivCreateView(privCreateView int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_create_view` = ?", privCreateView).Find(&results).Error

	return
}

// GetBatchFromPrivCreateView 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromPrivCreateView(privCreateViews []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_create_view` IN (?)", privCreateViews).Find(&results).Error

	return
}

// GetFromPrivShowView 通过priv_show_view获取内容
func (obj *_AllDatabasePrivilegeHistoryMgr) GetFromPrivShowView(privShowView int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_show_view` = ?", privShowView).Find(&results).Error

	return
}

// GetBatchFromPrivShowView 批量查找
func (obj *_AllDatabasePrivilegeHistoryMgr) GetBatchFromPrivShowView(privShowViews []int64) (results []*AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`priv_show_view` IN (?)", privShowViews).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllDatabasePrivilegeHistoryMgr) FetchByPrimaryKey(tenantID int64, userID int64, databaseName string, schemaVersion int64) (result AllDatabasePrivilegeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilegeHistory{}).Where("`tenant_id` = ? AND `user_id` = ? AND `database_name` = ? AND `schema_version` = ?", tenantID, userID, databaseName, schemaVersion).First(&result).Error

	return
}
