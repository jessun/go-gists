package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualTenantKeystoreHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualTenantKeystoreHistoryMgr open func
func AllVirtualTenantKeystoreHistoryMgr(db *gorm.DB) *_AllVirtualTenantKeystoreHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantKeystoreHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantKeystoreHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_keystore_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetTableName() string {
	return "__all_virtual_tenant_keystore_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantKeystoreHistoryMgr) Reset() *_AllVirtualTenantKeystoreHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) Get() (result AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantKeystoreHistoryMgr) Gets() (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantKeystoreHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithKeystoreID keystore_id获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) WithKeystoreID(keystoreID int64) Option {
	return optionFunc(func(o *options) { o.query["keystore_id"] = keystoreID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithKeystoreName keystore_name获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) WithKeystoreName(keystoreName string) Option {
	return optionFunc(func(o *options) { o.query["keystore_name"] = keystoreName })
}

// WithPassword password获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithStatus status获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithMasterKeyID master_key_id获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) WithMasterKeyID(masterKeyID int64) Option {
	return optionFunc(func(o *options) { o.query["master_key_id"] = masterKeyID })
}

// WithMasterKey master_key获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) WithMasterKey(masterKey string) Option {
	return optionFunc(func(o *options) { o.query["master_key"] = masterKey })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetByOption(opts ...Option) (result AllVirtualTenantKeystoreHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantKeystoreHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantKeystoreHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantKeystoreHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where(options.query)
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
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromKeystoreID 通过keystore_id获取内容
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetFromKeystoreID(keystoreID int64) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`keystore_id` = ?", keystoreID).Find(&results).Error

	return
}

// GetBatchFromKeystoreID 批量查找
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetBatchFromKeystoreID(keystoreIDs []int64) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`keystore_id` IN (?)", keystoreIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromKeystoreName 通过keystore_name获取内容
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetFromKeystoreName(keystoreName string) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`keystore_name` = ?", keystoreName).Find(&results).Error

	return
}

// GetBatchFromKeystoreName 批量查找
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetBatchFromKeystoreName(keystoreNames []string) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`keystore_name` IN (?)", keystoreNames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetFromPassword(password string) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetBatchFromPassword(passwords []string) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetFromStatus(status int64) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetBatchFromStatus(statuss []int64) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromMasterKeyID 通过master_key_id获取内容
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetFromMasterKeyID(masterKeyID int64) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`master_key_id` = ?", masterKeyID).Find(&results).Error

	return
}

// GetBatchFromMasterKeyID 批量查找
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetBatchFromMasterKeyID(masterKeyIDs []int64) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`master_key_id` IN (?)", masterKeyIDs).Find(&results).Error

	return
}

// GetFromMasterKey 通过master_key获取内容
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetFromMasterKey(masterKey string) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`master_key` = ?", masterKey).Find(&results).Error

	return
}

// GetBatchFromMasterKey 批量查找
func (obj *_AllVirtualTenantKeystoreHistoryMgr) GetBatchFromMasterKey(masterKeys []string) (results []*AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`master_key` IN (?)", masterKeys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTenantKeystoreHistoryMgr) FetchByPrimaryKey(tenantID int64, keystoreID int64, schemaVersion int64) (result AllVirtualTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystoreHistory{}).Where("`tenant_id` = ? AND `keystore_id` = ? AND `schema_version` = ?", tenantID, keystoreID, schemaVersion).First(&result).Error

	return
}
