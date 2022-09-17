package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualTenantTablespaceHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualTenantTablespaceHistoryMgr open func
func AllVirtualTenantTablespaceHistoryMgr(db *gorm.DB) *_AllVirtualTenantTablespaceHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantTablespaceHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantTablespaceHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_tablespace_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetTableName() string {
	return "__all_virtual_tenant_tablespace_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantTablespaceHistoryMgr) Reset() *_AllVirtualTenantTablespaceHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) Get() (result AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantTablespaceHistoryMgr) Gets() (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantTablespaceHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithTablespaceName tablespace_name获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) WithTablespaceName(tablespaceName string) Option {
	return optionFunc(func(o *options) { o.query["tablespace_name"] = tablespaceName })
}

// WithEncryptionName encryption_name获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) WithEncryptionName(encryptionName string) Option {
	return optionFunc(func(o *options) { o.query["encryption_name"] = encryptionName })
}

// WithEncryptKey encrypt_key获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) WithEncryptKey(encryptKey []byte) Option {
	return optionFunc(func(o *options) { o.query["encrypt_key"] = encryptKey })
}

// WithMasterKeyID master_key_id获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) WithMasterKeyID(masterKeyID uint64) Option {
	return optionFunc(func(o *options) { o.query["master_key_id"] = masterKeyID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetByOption(opts ...Option) (result AllVirtualTenantTablespaceHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantTablespaceHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantTablespaceHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantTablespaceHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where(options.query)
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
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromTablespaceName 通过tablespace_name获取内容
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetFromTablespaceName(tablespaceName string) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`tablespace_name` = ?", tablespaceName).Find(&results).Error

	return
}

// GetBatchFromTablespaceName 批量查找
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetBatchFromTablespaceName(tablespaceNames []string) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`tablespace_name` IN (?)", tablespaceNames).Find(&results).Error

	return
}

// GetFromEncryptionName 通过encryption_name获取内容
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetFromEncryptionName(encryptionName string) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`encryption_name` = ?", encryptionName).Find(&results).Error

	return
}

// GetBatchFromEncryptionName 批量查找
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetBatchFromEncryptionName(encryptionNames []string) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`encryption_name` IN (?)", encryptionNames).Find(&results).Error

	return
}

// GetFromEncryptKey 通过encrypt_key获取内容
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetFromEncryptKey(encryptKey []byte) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`encrypt_key` = ?", encryptKey).Find(&results).Error

	return
}

// GetBatchFromEncryptKey 批量查找
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetBatchFromEncryptKey(encryptKeys [][]byte) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`encrypt_key` IN (?)", encryptKeys).Find(&results).Error

	return
}

// GetFromMasterKeyID 通过master_key_id获取内容
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetFromMasterKeyID(masterKeyID uint64) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`master_key_id` = ?", masterKeyID).Find(&results).Error

	return
}

// GetBatchFromMasterKeyID 批量查找
func (obj *_AllVirtualTenantTablespaceHistoryMgr) GetBatchFromMasterKeyID(masterKeyIDs []uint64) (results []*AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`master_key_id` IN (?)", masterKeyIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTenantTablespaceHistoryMgr) FetchByPrimaryKey(tenantID int64, tablespaceID int64, schemaVersion int64) (result AllVirtualTenantTablespaceHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespaceHistory{}).Where("`tenant_id` = ? AND `tablespace_id` = ? AND `schema_version` = ?", tenantID, tablespaceID, schemaVersion).First(&result).Error

	return
}
