package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualDatabasePrivilegeMgr struct {
	*_BaseMgr
}

// AllVirtualDatabasePrivilegeMgr open func
func AllVirtualDatabasePrivilegeMgr(db *gorm.DB) *_AllVirtualDatabasePrivilegeMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDatabasePrivilegeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDatabasePrivilegeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_database_privilege"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDatabasePrivilegeMgr) GetTableName() string {
	return "__all_virtual_database_privilege"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDatabasePrivilegeMgr) Reset() *_AllVirtualDatabasePrivilegeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDatabasePrivilegeMgr) Get() (result AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDatabasePrivilegeMgr) Gets() (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDatabasePrivilegeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserID user_id获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithDatabaseName database_name获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithPrivAlter priv_alter获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithPrivAlter(privAlter int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter"] = privAlter })
}

// WithPrivCreate priv_create获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithPrivCreate(privCreate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create"] = privCreate })
}

// WithPrivDelete priv_delete获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithPrivDelete(privDelete int64) Option {
	return optionFunc(func(o *options) { o.query["priv_delete"] = privDelete })
}

// WithPrivDrop priv_drop获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithPrivDrop(privDrop int64) Option {
	return optionFunc(func(o *options) { o.query["priv_drop"] = privDrop })
}

// WithPrivGrantOption priv_grant_option获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithPrivGrantOption(privGrantOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_grant_option"] = privGrantOption })
}

// WithPrivInsert priv_insert获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithPrivInsert(privInsert int64) Option {
	return optionFunc(func(o *options) { o.query["priv_insert"] = privInsert })
}

// WithPrivUpdate priv_update获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithPrivUpdate(privUpdate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_update"] = privUpdate })
}

// WithPrivSelect priv_select获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithPrivSelect(privSelect int64) Option {
	return optionFunc(func(o *options) { o.query["priv_select"] = privSelect })
}

// WithPrivIndex priv_index获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithPrivIndex(privIndex int64) Option {
	return optionFunc(func(o *options) { o.query["priv_index"] = privIndex })
}

// WithPrivCreateView priv_create_view获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithPrivCreateView(privCreateView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_view"] = privCreateView })
}

// WithPrivShowView priv_show_view获取
func (obj *_AllVirtualDatabasePrivilegeMgr) WithPrivShowView(privShowView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_view"] = privShowView })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDatabasePrivilegeMgr) GetByOption(opts ...Option) (result AllVirtualDatabasePrivilege, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDatabasePrivilegeMgr) GetByOptions(opts ...Option) (results []*AllVirtualDatabasePrivilege, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDatabasePrivilegeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDatabasePrivilege, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where(options.query)
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
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromUserID(userID int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromUserID(userIDs []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromDatabaseName(databaseName string) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromPrivAlter 通过priv_alter获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromPrivAlter(privAlter int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_alter` = ?", privAlter).Find(&results).Error

	return
}

// GetBatchFromPrivAlter 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromPrivAlter(privAlters []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_alter` IN (?)", privAlters).Find(&results).Error

	return
}

// GetFromPrivCreate 通过priv_create获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromPrivCreate(privCreate int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_create` = ?", privCreate).Find(&results).Error

	return
}

// GetBatchFromPrivCreate 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromPrivCreate(privCreates []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_create` IN (?)", privCreates).Find(&results).Error

	return
}

// GetFromPrivDelete 通过priv_delete获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromPrivDelete(privDelete int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_delete` = ?", privDelete).Find(&results).Error

	return
}

// GetBatchFromPrivDelete 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromPrivDelete(privDeletes []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_delete` IN (?)", privDeletes).Find(&results).Error

	return
}

// GetFromPrivDrop 通过priv_drop获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromPrivDrop(privDrop int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_drop` = ?", privDrop).Find(&results).Error

	return
}

// GetBatchFromPrivDrop 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromPrivDrop(privDrops []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_drop` IN (?)", privDrops).Find(&results).Error

	return
}

// GetFromPrivGrantOption 通过priv_grant_option获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromPrivGrantOption(privGrantOption int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_grant_option` = ?", privGrantOption).Find(&results).Error

	return
}

// GetBatchFromPrivGrantOption 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromPrivGrantOption(privGrantOptions []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_grant_option` IN (?)", privGrantOptions).Find(&results).Error

	return
}

// GetFromPrivInsert 通过priv_insert获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromPrivInsert(privInsert int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_insert` = ?", privInsert).Find(&results).Error

	return
}

// GetBatchFromPrivInsert 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromPrivInsert(privInserts []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_insert` IN (?)", privInserts).Find(&results).Error

	return
}

// GetFromPrivUpdate 通过priv_update获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromPrivUpdate(privUpdate int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_update` = ?", privUpdate).Find(&results).Error

	return
}

// GetBatchFromPrivUpdate 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromPrivUpdate(privUpdates []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_update` IN (?)", privUpdates).Find(&results).Error

	return
}

// GetFromPrivSelect 通过priv_select获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromPrivSelect(privSelect int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_select` = ?", privSelect).Find(&results).Error

	return
}

// GetBatchFromPrivSelect 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromPrivSelect(privSelects []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_select` IN (?)", privSelects).Find(&results).Error

	return
}

// GetFromPrivIndex 通过priv_index获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromPrivIndex(privIndex int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_index` = ?", privIndex).Find(&results).Error

	return
}

// GetBatchFromPrivIndex 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromPrivIndex(privIndexs []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_index` IN (?)", privIndexs).Find(&results).Error

	return
}

// GetFromPrivCreateView 通过priv_create_view获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromPrivCreateView(privCreateView int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_create_view` = ?", privCreateView).Find(&results).Error

	return
}

// GetBatchFromPrivCreateView 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromPrivCreateView(privCreateViews []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_create_view` IN (?)", privCreateViews).Find(&results).Error

	return
}

// GetFromPrivShowView 通过priv_show_view获取内容
func (obj *_AllVirtualDatabasePrivilegeMgr) GetFromPrivShowView(privShowView int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_show_view` = ?", privShowView).Find(&results).Error

	return
}

// GetBatchFromPrivShowView 批量查找
func (obj *_AllVirtualDatabasePrivilegeMgr) GetBatchFromPrivShowView(privShowViews []int64) (results []*AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`priv_show_view` IN (?)", privShowViews).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDatabasePrivilegeMgr) FetchByPrimaryKey(tenantID int64, userID int64, databaseName string) (result AllVirtualDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabasePrivilege{}).Where("`tenant_id` = ? AND `user_id` = ? AND `database_name` = ?", tenantID, userID, databaseName).First(&result).Error

	return
}
