package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantTablespaceMgr struct {
	*_BaseMgr
}

// AllTenantTablespaceMgr open func
func AllTenantTablespaceMgr(db *gorm.DB) *_AllTenantTablespaceMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantTablespaceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantTablespaceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_tablespace"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantTablespaceMgr) GetTableName() string {
	return "__all_tenant_tablespace"
}

// Reset 重置gorm会话
func (obj *_AllTenantTablespaceMgr) Reset() *_AllTenantTablespaceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantTablespaceMgr) Get() (result AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantTablespaceMgr) Gets() (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantTablespaceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantTablespaceMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantTablespaceMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantTablespaceMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllTenantTablespaceMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithTablespaceName tablespace_name获取
func (obj *_AllTenantTablespaceMgr) WithTablespaceName(tablespaceName string) Option {
	return optionFunc(func(o *options) { o.query["tablespace_name"] = tablespaceName })
}

// WithEncryptionName encryption_name获取
func (obj *_AllTenantTablespaceMgr) WithEncryptionName(encryptionName string) Option {
	return optionFunc(func(o *options) { o.query["encryption_name"] = encryptionName })
}

// WithEncryptKey encrypt_key获取
func (obj *_AllTenantTablespaceMgr) WithEncryptKey(encryptKey []byte) Option {
	return optionFunc(func(o *options) { o.query["encrypt_key"] = encryptKey })
}

// WithMasterKeyID master_key_id获取
func (obj *_AllTenantTablespaceMgr) WithMasterKeyID(masterKeyID uint64) Option {
	return optionFunc(func(o *options) { o.query["master_key_id"] = masterKeyID })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantTablespaceMgr) GetByOption(opts ...Option) (result AllTenantTablespace, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantTablespaceMgr) GetByOptions(opts ...Option) (results []*AllTenantTablespace, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantTablespaceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantTablespace, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where(options.query)
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
func (obj *_AllTenantTablespaceMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantTablespaceMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantTablespaceMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantTablespaceMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantTablespaceMgr) GetFromTenantID(tenantID int64) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantTablespaceMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllTenantTablespaceMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllTenantTablespaceMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromTablespaceName 通过tablespace_name获取内容
func (obj *_AllTenantTablespaceMgr) GetFromTablespaceName(tablespaceName string) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`tablespace_name` = ?", tablespaceName).Find(&results).Error

	return
}

// GetBatchFromTablespaceName 批量查找
func (obj *_AllTenantTablespaceMgr) GetBatchFromTablespaceName(tablespaceNames []string) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`tablespace_name` IN (?)", tablespaceNames).Find(&results).Error

	return
}

// GetFromEncryptionName 通过encryption_name获取内容
func (obj *_AllTenantTablespaceMgr) GetFromEncryptionName(encryptionName string) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`encryption_name` = ?", encryptionName).Find(&results).Error

	return
}

// GetBatchFromEncryptionName 批量查找
func (obj *_AllTenantTablespaceMgr) GetBatchFromEncryptionName(encryptionNames []string) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`encryption_name` IN (?)", encryptionNames).Find(&results).Error

	return
}

// GetFromEncryptKey 通过encrypt_key获取内容
func (obj *_AllTenantTablespaceMgr) GetFromEncryptKey(encryptKey []byte) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`encrypt_key` = ?", encryptKey).Find(&results).Error

	return
}

// GetBatchFromEncryptKey 批量查找
func (obj *_AllTenantTablespaceMgr) GetBatchFromEncryptKey(encryptKeys [][]byte) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`encrypt_key` IN (?)", encryptKeys).Find(&results).Error

	return
}

// GetFromMasterKeyID 通过master_key_id获取内容
func (obj *_AllTenantTablespaceMgr) GetFromMasterKeyID(masterKeyID uint64) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`master_key_id` = ?", masterKeyID).Find(&results).Error

	return
}

// GetBatchFromMasterKeyID 批量查找
func (obj *_AllTenantTablespaceMgr) GetBatchFromMasterKeyID(masterKeyIDs []uint64) (results []*AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`master_key_id` IN (?)", masterKeyIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantTablespaceMgr) FetchByPrimaryKey(tenantID int64, tablespaceID int64) (result AllTenantTablespace, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTablespace{}).Where("`tenant_id` = ? AND `tablespace_id` = ?", tenantID, tablespaceID).First(&result).Error

	return
}
