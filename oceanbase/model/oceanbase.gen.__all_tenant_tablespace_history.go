package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantTablespaceHistoryMgr struct {
	*_BaseMgr
}

// AllTenantTablespaceHistoryMgr open func
func AllTenantTablespaceHistoryMgr(db *gorm.DB) *_AllTenantTablespaceHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantTablespaceHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantTablespaceHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_tablespace_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantTablespaceHistoryMgr) GetTableName() string {
	return "__all_tenant_tablespace_history"
}

// Reset 重置gorm会话
func (obj *_AllTenantTablespaceHistoryMgr) Reset() *_AllTenantTablespaceHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantTablespaceHistoryMgr) Get() (result AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantTablespaceHistoryMgr) Gets() (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantTablespaceHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantTablespaceHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantTablespaceHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantTablespaceHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllTenantTablespaceHistoryMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantTablespaceHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTenantTablespaceHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithTablespaceName tablespace_name获取
func (obj *_AllTenantTablespaceHistoryMgr) WithTablespaceName(tablespaceName string) Option {
	return optionFunc(func(o *options) { o.query["tablespace_name"] = tablespaceName })
}

// WithEncryptionName encryption_name获取
func (obj *_AllTenantTablespaceHistoryMgr) WithEncryptionName(encryptionName string) Option {
	return optionFunc(func(o *options) { o.query["encryption_name"] = encryptionName })
}

// WithEncryptKey encrypt_key获取
func (obj *_AllTenantTablespaceHistoryMgr) WithEncryptKey(encryptKey []byte) Option {
	return optionFunc(func(o *options) { o.query["encrypt_key"] = encryptKey })
}

// WithMasterKeyID master_key_id获取
func (obj *_AllTenantTablespaceHistoryMgr) WithMasterKeyID(masterKeyID uint64) Option {
	return optionFunc(func(o *options) { o.query["master_key_id"] = masterKeyID })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantTablespaceHistoryMgr) GetByOption(opts ...Option) (result AllTenantTablespaceHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantTablespaceHistoryMgr) GetByOptions(opts ...Option) (results []*AllTenantTablespaceHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantTablespaceHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantTablespaceHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where(options.query)
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
func (obj *_AllTenantTablespaceHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantTablespaceHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantTablespaceHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantTablespaceHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantTablespaceHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantTablespaceHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllTenantTablespaceHistoryMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllTenantTablespaceHistoryMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantTablespaceHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantTablespaceHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTenantTablespaceHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTenantTablespaceHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromTablespaceName 通过tablespace_name获取内容
func (obj *_AllTenantTablespaceHistoryMgr) GetFromTablespaceName(tablespaceName string) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`tablespace_name` = ?", tablespaceName).Find(&results).Error

	return
}

// GetBatchFromTablespaceName 批量查找
func (obj *_AllTenantTablespaceHistoryMgr) GetBatchFromTablespaceName(tablespaceNames []string) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`tablespace_name` IN (?)", tablespaceNames).Find(&results).Error

	return
}

// GetFromEncryptionName 通过encryption_name获取内容
func (obj *_AllTenantTablespaceHistoryMgr) GetFromEncryptionName(encryptionName string) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`encryption_name` = ?", encryptionName).Find(&results).Error

	return
}

// GetBatchFromEncryptionName 批量查找
func (obj *_AllTenantTablespaceHistoryMgr) GetBatchFromEncryptionName(encryptionNames []string) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`encryption_name` IN (?)", encryptionNames).Find(&results).Error

	return
}

// GetFromEncryptKey 通过encrypt_key获取内容
func (obj *_AllTenantTablespaceHistoryMgr) GetFromEncryptKey(encryptKey []byte) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`encrypt_key` = ?", encryptKey).Find(&results).Error

	return
}

// GetBatchFromEncryptKey 批量查找
func (obj *_AllTenantTablespaceHistoryMgr) GetBatchFromEncryptKey(encryptKeys [][]byte) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`encrypt_key` IN (?)", encryptKeys).Find(&results).Error

	return
}

// GetFromMasterKeyID 通过master_key_id获取内容
func (obj *_AllTenantTablespaceHistoryMgr) GetFromMasterKeyID(masterKeyID uint64) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`master_key_id` = ?", masterKeyID).Find(&results).Error

	return
}

// GetBatchFromMasterKeyID 批量查找
func (obj *_AllTenantTablespaceHistoryMgr) GetBatchFromMasterKeyID(masterKeyIDs []uint64) (results []*AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`master_key_id` IN (?)", masterKeyIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantTablespaceHistoryMgr) FetchByPrimaryKey(tenantID int64, tablespaceID int64, schemaVersion int64) (result AllTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespaceHistory{}).Where("`tenant_id` = ? AND `tablespace_id` = ? AND `schema_version` = ?", tenantID, tablespaceID, schemaVersion).First(&result).Error

	return
}
