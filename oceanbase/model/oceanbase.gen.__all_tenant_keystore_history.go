package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantKeystoreHistoryMgr struct {
	*_BaseMgr
}

// AllTenantKeystoreHistoryMgr open func
func AllTenantKeystoreHistoryMgr(db *gorm.DB) *_AllTenantKeystoreHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantKeystoreHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantKeystoreHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_keystore_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantKeystoreHistoryMgr) GetTableName() string {
	return "__all_tenant_keystore_history"
}

// Reset 重置gorm会话
func (obj *_AllTenantKeystoreHistoryMgr) Reset() *_AllTenantKeystoreHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantKeystoreHistoryMgr) Get() (result AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantKeystoreHistoryMgr) Gets() (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantKeystoreHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantKeystoreHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantKeystoreHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantKeystoreHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithKeystoreID keystore_id获取
func (obj *_AllTenantKeystoreHistoryMgr) WithKeystoreID(keystoreID int64) Option {
	return optionFunc(func(o *options) { o.query["keystore_id"] = keystoreID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantKeystoreHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTenantKeystoreHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithKeystoreName keystore_name获取
func (obj *_AllTenantKeystoreHistoryMgr) WithKeystoreName(keystoreName string) Option {
	return optionFunc(func(o *options) { o.query["keystore_name"] = keystoreName })
}

// WithPassword password获取
func (obj *_AllTenantKeystoreHistoryMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithStatus status获取
func (obj *_AllTenantKeystoreHistoryMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithMasterKeyID master_key_id获取
func (obj *_AllTenantKeystoreHistoryMgr) WithMasterKeyID(masterKeyID int64) Option {
	return optionFunc(func(o *options) { o.query["master_key_id"] = masterKeyID })
}

// WithMasterKey master_key获取
func (obj *_AllTenantKeystoreHistoryMgr) WithMasterKey(masterKey string) Option {
	return optionFunc(func(o *options) { o.query["master_key"] = masterKey })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantKeystoreHistoryMgr) GetByOption(opts ...Option) (result AllTenantKeystoreHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantKeystoreHistoryMgr) GetByOptions(opts ...Option) (results []*AllTenantKeystoreHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantKeystoreHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantKeystoreHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where(options.query)
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
func (obj *_AllTenantKeystoreHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantKeystoreHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantKeystoreHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantKeystoreHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantKeystoreHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantKeystoreHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromKeystoreID 通过keystore_id获取内容
func (obj *_AllTenantKeystoreHistoryMgr) GetFromKeystoreID(keystoreID int64) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`keystore_id` = ?", keystoreID).Find(&results).Error

	return
}

// GetBatchFromKeystoreID 批量查找
func (obj *_AllTenantKeystoreHistoryMgr) GetBatchFromKeystoreID(keystoreIDs []int64) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`keystore_id` IN (?)", keystoreIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantKeystoreHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantKeystoreHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTenantKeystoreHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTenantKeystoreHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromKeystoreName 通过keystore_name获取内容
func (obj *_AllTenantKeystoreHistoryMgr) GetFromKeystoreName(keystoreName string) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`keystore_name` = ?", keystoreName).Find(&results).Error

	return
}

// GetBatchFromKeystoreName 批量查找
func (obj *_AllTenantKeystoreHistoryMgr) GetBatchFromKeystoreName(keystoreNames []string) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`keystore_name` IN (?)", keystoreNames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_AllTenantKeystoreHistoryMgr) GetFromPassword(password string) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_AllTenantKeystoreHistoryMgr) GetBatchFromPassword(passwords []string) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllTenantKeystoreHistoryMgr) GetFromStatus(status int64) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllTenantKeystoreHistoryMgr) GetBatchFromStatus(statuss []int64) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromMasterKeyID 通过master_key_id获取内容
func (obj *_AllTenantKeystoreHistoryMgr) GetFromMasterKeyID(masterKeyID int64) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`master_key_id` = ?", masterKeyID).Find(&results).Error

	return
}

// GetBatchFromMasterKeyID 批量查找
func (obj *_AllTenantKeystoreHistoryMgr) GetBatchFromMasterKeyID(masterKeyIDs []int64) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`master_key_id` IN (?)", masterKeyIDs).Find(&results).Error

	return
}

// GetFromMasterKey 通过master_key获取内容
func (obj *_AllTenantKeystoreHistoryMgr) GetFromMasterKey(masterKey string) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`master_key` = ?", masterKey).Find(&results).Error

	return
}

// GetBatchFromMasterKey 批量查找
func (obj *_AllTenantKeystoreHistoryMgr) GetBatchFromMasterKey(masterKeys []string) (results []*AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`master_key` IN (?)", masterKeys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantKeystoreHistoryMgr) FetchByPrimaryKey(tenantID int64, keystoreID int64, schemaVersion int64) (result AllTenantKeystoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystoreHistory{}).Where("`tenant_id` = ? AND `keystore_id` = ? AND `schema_version` = ?", tenantID, keystoreID, schemaVersion).First(&result).Error

	return
}
