package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllDatabasePrivilegeMgr struct {
	*_BaseMgr
}

// AllDatabasePrivilegeMgr open func
func AllDatabasePrivilegeMgr(db *gorm.DB) *_AllDatabasePrivilegeMgr {
	if db == nil {
		panic(fmt.Errorf("AllDatabasePrivilegeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllDatabasePrivilegeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_database_privilege"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllDatabasePrivilegeMgr) GetTableName() string {
	return "__all_database_privilege"
}

// Reset 重置gorm会话
func (obj *_AllDatabasePrivilegeMgr) Reset() *_AllDatabasePrivilegeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllDatabasePrivilegeMgr) Get() (result AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllDatabasePrivilegeMgr) Gets() (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllDatabasePrivilegeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllDatabasePrivilegeMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllDatabasePrivilegeMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllDatabasePrivilegeMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserID user_id获取
func (obj *_AllDatabasePrivilegeMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithDatabaseName database_name获取
func (obj *_AllDatabasePrivilegeMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithPrivAlter priv_alter获取
func (obj *_AllDatabasePrivilegeMgr) WithPrivAlter(privAlter int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter"] = privAlter })
}

// WithPrivCreate priv_create获取
func (obj *_AllDatabasePrivilegeMgr) WithPrivCreate(privCreate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create"] = privCreate })
}

// WithPrivDelete priv_delete获取
func (obj *_AllDatabasePrivilegeMgr) WithPrivDelete(privDelete int64) Option {
	return optionFunc(func(o *options) { o.query["priv_delete"] = privDelete })
}

// WithPrivDrop priv_drop获取
func (obj *_AllDatabasePrivilegeMgr) WithPrivDrop(privDrop int64) Option {
	return optionFunc(func(o *options) { o.query["priv_drop"] = privDrop })
}

// WithPrivGrantOption priv_grant_option获取
func (obj *_AllDatabasePrivilegeMgr) WithPrivGrantOption(privGrantOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_grant_option"] = privGrantOption })
}

// WithPrivInsert priv_insert获取
func (obj *_AllDatabasePrivilegeMgr) WithPrivInsert(privInsert int64) Option {
	return optionFunc(func(o *options) { o.query["priv_insert"] = privInsert })
}

// WithPrivUpdate priv_update获取
func (obj *_AllDatabasePrivilegeMgr) WithPrivUpdate(privUpdate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_update"] = privUpdate })
}

// WithPrivSelect priv_select获取
func (obj *_AllDatabasePrivilegeMgr) WithPrivSelect(privSelect int64) Option {
	return optionFunc(func(o *options) { o.query["priv_select"] = privSelect })
}

// WithPrivIndex priv_index获取
func (obj *_AllDatabasePrivilegeMgr) WithPrivIndex(privIndex int64) Option {
	return optionFunc(func(o *options) { o.query["priv_index"] = privIndex })
}

// WithPrivCreateView priv_create_view获取
func (obj *_AllDatabasePrivilegeMgr) WithPrivCreateView(privCreateView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_view"] = privCreateView })
}

// WithPrivShowView priv_show_view获取
func (obj *_AllDatabasePrivilegeMgr) WithPrivShowView(privShowView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_view"] = privShowView })
}

// GetByOption 功能选项模式获取
func (obj *_AllDatabasePrivilegeMgr) GetByOption(opts ...Option) (result AllDatabasePrivilege, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllDatabasePrivilegeMgr) GetByOptions(opts ...Option) (results []*AllDatabasePrivilege, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllDatabasePrivilegeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllDatabasePrivilege, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where(options.query)
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
func (obj *_AllDatabasePrivilegeMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromTenantID(tenantID int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromUserID(userID int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromUserID(userIDs []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromDatabaseName(databaseName string) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromPrivAlter 通过priv_alter获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromPrivAlter(privAlter int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_alter` = ?", privAlter).Find(&results).Error

	return
}

// GetBatchFromPrivAlter 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromPrivAlter(privAlters []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_alter` IN (?)", privAlters).Find(&results).Error

	return
}

// GetFromPrivCreate 通过priv_create获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromPrivCreate(privCreate int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_create` = ?", privCreate).Find(&results).Error

	return
}

// GetBatchFromPrivCreate 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromPrivCreate(privCreates []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_create` IN (?)", privCreates).Find(&results).Error

	return
}

// GetFromPrivDelete 通过priv_delete获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromPrivDelete(privDelete int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_delete` = ?", privDelete).Find(&results).Error

	return
}

// GetBatchFromPrivDelete 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromPrivDelete(privDeletes []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_delete` IN (?)", privDeletes).Find(&results).Error

	return
}

// GetFromPrivDrop 通过priv_drop获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromPrivDrop(privDrop int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_drop` = ?", privDrop).Find(&results).Error

	return
}

// GetBatchFromPrivDrop 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromPrivDrop(privDrops []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_drop` IN (?)", privDrops).Find(&results).Error

	return
}

// GetFromPrivGrantOption 通过priv_grant_option获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromPrivGrantOption(privGrantOption int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_grant_option` = ?", privGrantOption).Find(&results).Error

	return
}

// GetBatchFromPrivGrantOption 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromPrivGrantOption(privGrantOptions []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_grant_option` IN (?)", privGrantOptions).Find(&results).Error

	return
}

// GetFromPrivInsert 通过priv_insert获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromPrivInsert(privInsert int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_insert` = ?", privInsert).Find(&results).Error

	return
}

// GetBatchFromPrivInsert 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromPrivInsert(privInserts []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_insert` IN (?)", privInserts).Find(&results).Error

	return
}

// GetFromPrivUpdate 通过priv_update获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromPrivUpdate(privUpdate int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_update` = ?", privUpdate).Find(&results).Error

	return
}

// GetBatchFromPrivUpdate 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromPrivUpdate(privUpdates []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_update` IN (?)", privUpdates).Find(&results).Error

	return
}

// GetFromPrivSelect 通过priv_select获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromPrivSelect(privSelect int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_select` = ?", privSelect).Find(&results).Error

	return
}

// GetBatchFromPrivSelect 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromPrivSelect(privSelects []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_select` IN (?)", privSelects).Find(&results).Error

	return
}

// GetFromPrivIndex 通过priv_index获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromPrivIndex(privIndex int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_index` = ?", privIndex).Find(&results).Error

	return
}

// GetBatchFromPrivIndex 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromPrivIndex(privIndexs []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_index` IN (?)", privIndexs).Find(&results).Error

	return
}

// GetFromPrivCreateView 通过priv_create_view获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromPrivCreateView(privCreateView int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_create_view` = ?", privCreateView).Find(&results).Error

	return
}

// GetBatchFromPrivCreateView 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromPrivCreateView(privCreateViews []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_create_view` IN (?)", privCreateViews).Find(&results).Error

	return
}

// GetFromPrivShowView 通过priv_show_view获取内容
func (obj *_AllDatabasePrivilegeMgr) GetFromPrivShowView(privShowView int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_show_view` = ?", privShowView).Find(&results).Error

	return
}

// GetBatchFromPrivShowView 批量查找
func (obj *_AllDatabasePrivilegeMgr) GetBatchFromPrivShowView(privShowViews []int64) (results []*AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`priv_show_view` IN (?)", privShowViews).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllDatabasePrivilegeMgr) FetchByPrimaryKey(tenantID int64, userID int64, databaseName string) (result AllDatabasePrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabasePrivilege{}).Where("`tenant_id` = ? AND `user_id` = ? AND `database_name` = ?", tenantID, userID, databaseName).First(&result).Error

	return
}
