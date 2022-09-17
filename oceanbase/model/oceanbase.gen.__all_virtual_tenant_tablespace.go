package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualTenantTablespaceMgr struct {
	*_BaseMgr
}

// AllVirtualTenantTablespaceMgr open func
func AllVirtualTenantTablespaceMgr(db *gorm.DB) *_AllVirtualTenantTablespaceMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantTablespaceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantTablespaceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_tablespace"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantTablespaceMgr) GetTableName() string {
	return "__all_virtual_tenant_tablespace"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantTablespaceMgr) Reset() *_AllVirtualTenantTablespaceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantTablespaceMgr) Get() (result AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantTablespaceMgr) Gets() (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantTablespaceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantTablespaceMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllVirtualTenantTablespaceMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTenantTablespaceMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTenantTablespaceMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTablespaceName tablespace_name获取
func (obj *_AllVirtualTenantTablespaceMgr) WithTablespaceName(tablespaceName string) Option {
	return optionFunc(func(o *options) { o.query["tablespace_name"] = tablespaceName })
}

// WithEncryptionName encryption_name获取
func (obj *_AllVirtualTenantTablespaceMgr) WithEncryptionName(encryptionName string) Option {
	return optionFunc(func(o *options) { o.query["encryption_name"] = encryptionName })
}

// WithEncryptKey encrypt_key获取
func (obj *_AllVirtualTenantTablespaceMgr) WithEncryptKey(encryptKey []byte) Option {
	return optionFunc(func(o *options) { o.query["encrypt_key"] = encryptKey })
}

// WithMasterKeyID master_key_id获取
func (obj *_AllVirtualTenantTablespaceMgr) WithMasterKeyID(masterKeyID uint64) Option {
	return optionFunc(func(o *options) { o.query["master_key_id"] = masterKeyID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantTablespaceMgr) GetByOption(opts ...Option) (result AllVirtualTenantTablespace, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantTablespaceMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantTablespace, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantTablespaceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantTablespace, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where(options.query)
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
func (obj *_AllVirtualTenantTablespaceMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantTablespaceMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllVirtualTenantTablespaceMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllVirtualTenantTablespaceMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTenantTablespaceMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTenantTablespaceMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTenantTablespaceMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTenantTablespaceMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTablespaceName 通过tablespace_name获取内容
func (obj *_AllVirtualTenantTablespaceMgr) GetFromTablespaceName(tablespaceName string) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`tablespace_name` = ?", tablespaceName).Find(&results).Error

	return
}

// GetBatchFromTablespaceName 批量查找
func (obj *_AllVirtualTenantTablespaceMgr) GetBatchFromTablespaceName(tablespaceNames []string) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`tablespace_name` IN (?)", tablespaceNames).Find(&results).Error

	return
}

// GetFromEncryptionName 通过encryption_name获取内容
func (obj *_AllVirtualTenantTablespaceMgr) GetFromEncryptionName(encryptionName string) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`encryption_name` = ?", encryptionName).Find(&results).Error

	return
}

// GetBatchFromEncryptionName 批量查找
func (obj *_AllVirtualTenantTablespaceMgr) GetBatchFromEncryptionName(encryptionNames []string) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`encryption_name` IN (?)", encryptionNames).Find(&results).Error

	return
}

// GetFromEncryptKey 通过encrypt_key获取内容
func (obj *_AllVirtualTenantTablespaceMgr) GetFromEncryptKey(encryptKey []byte) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`encrypt_key` = ?", encryptKey).Find(&results).Error

	return
}

// GetBatchFromEncryptKey 批量查找
func (obj *_AllVirtualTenantTablespaceMgr) GetBatchFromEncryptKey(encryptKeys [][]byte) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`encrypt_key` IN (?)", encryptKeys).Find(&results).Error

	return
}

// GetFromMasterKeyID 通过master_key_id获取内容
func (obj *_AllVirtualTenantTablespaceMgr) GetFromMasterKeyID(masterKeyID uint64) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`master_key_id` = ?", masterKeyID).Find(&results).Error

	return
}

// GetBatchFromMasterKeyID 批量查找
func (obj *_AllVirtualTenantTablespaceMgr) GetBatchFromMasterKeyID(masterKeyIDs []uint64) (results []*AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`master_key_id` IN (?)", masterKeyIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTenantTablespaceMgr) FetchByPrimaryKey(tenantID int64, tablespaceID int64) (result AllVirtualTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantTablespace{}).Where("`tenant_id` = ? AND `tablespace_id` = ?", tenantID, tablespaceID).First(&result).Error

	return
}
