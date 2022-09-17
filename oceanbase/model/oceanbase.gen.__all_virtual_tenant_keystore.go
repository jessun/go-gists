package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualTenantKeystoreMgr struct {
	*_BaseMgr
}

// AllVirtualTenantKeystoreMgr open func
func AllVirtualTenantKeystoreMgr(db *gorm.DB) *_AllVirtualTenantKeystoreMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantKeystoreMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantKeystoreMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_keystore"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantKeystoreMgr) GetTableName() string {
	return "__all_virtual_tenant_keystore"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantKeystoreMgr) Reset() *_AllVirtualTenantKeystoreMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantKeystoreMgr) Get() (result AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantKeystoreMgr) Gets() (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantKeystoreMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantKeystoreMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithKeystoreID keystore_id获取
func (obj *_AllVirtualTenantKeystoreMgr) WithKeystoreID(keystoreID int64) Option {
	return optionFunc(func(o *options) { o.query["keystore_id"] = keystoreID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTenantKeystoreMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTenantKeystoreMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithKeystoreName keystore_name获取
func (obj *_AllVirtualTenantKeystoreMgr) WithKeystoreName(keystoreName string) Option {
	return optionFunc(func(o *options) { o.query["keystore_name"] = keystoreName })
}

// WithPassword password获取
func (obj *_AllVirtualTenantKeystoreMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithStatus status获取
func (obj *_AllVirtualTenantKeystoreMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithMasterKeyID master_key_id获取
func (obj *_AllVirtualTenantKeystoreMgr) WithMasterKeyID(masterKeyID int64) Option {
	return optionFunc(func(o *options) { o.query["master_key_id"] = masterKeyID })
}

// WithMasterKey master_key获取
func (obj *_AllVirtualTenantKeystoreMgr) WithMasterKey(masterKey string) Option {
	return optionFunc(func(o *options) { o.query["master_key"] = masterKey })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantKeystoreMgr) GetByOption(opts ...Option) (result AllVirtualTenantKeystore, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantKeystoreMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantKeystore, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantKeystoreMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantKeystore, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where(options.query)
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
func (obj *_AllVirtualTenantKeystoreMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantKeystoreMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromKeystoreID 通过keystore_id获取内容
func (obj *_AllVirtualTenantKeystoreMgr) GetFromKeystoreID(keystoreID int64) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`keystore_id` = ?", keystoreID).Find(&results).Error

	return
}

// GetBatchFromKeystoreID 批量查找
func (obj *_AllVirtualTenantKeystoreMgr) GetBatchFromKeystoreID(keystoreIDs []int64) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`keystore_id` IN (?)", keystoreIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTenantKeystoreMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTenantKeystoreMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTenantKeystoreMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTenantKeystoreMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromKeystoreName 通过keystore_name获取内容
func (obj *_AllVirtualTenantKeystoreMgr) GetFromKeystoreName(keystoreName string) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`keystore_name` = ?", keystoreName).Find(&results).Error

	return
}

// GetBatchFromKeystoreName 批量查找
func (obj *_AllVirtualTenantKeystoreMgr) GetBatchFromKeystoreName(keystoreNames []string) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`keystore_name` IN (?)", keystoreNames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_AllVirtualTenantKeystoreMgr) GetFromPassword(password string) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_AllVirtualTenantKeystoreMgr) GetBatchFromPassword(passwords []string) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualTenantKeystoreMgr) GetFromStatus(status int64) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualTenantKeystoreMgr) GetBatchFromStatus(statuss []int64) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromMasterKeyID 通过master_key_id获取内容
func (obj *_AllVirtualTenantKeystoreMgr) GetFromMasterKeyID(masterKeyID int64) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`master_key_id` = ?", masterKeyID).Find(&results).Error

	return
}

// GetBatchFromMasterKeyID 批量查找
func (obj *_AllVirtualTenantKeystoreMgr) GetBatchFromMasterKeyID(masterKeyIDs []int64) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`master_key_id` IN (?)", masterKeyIDs).Find(&results).Error

	return
}

// GetFromMasterKey 通过master_key获取内容
func (obj *_AllVirtualTenantKeystoreMgr) GetFromMasterKey(masterKey string) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`master_key` = ?", masterKey).Find(&results).Error

	return
}

// GetBatchFromMasterKey 批量查找
func (obj *_AllVirtualTenantKeystoreMgr) GetBatchFromMasterKey(masterKeys []string) (results []*AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`master_key` IN (?)", masterKeys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTenantKeystoreMgr) FetchByPrimaryKey(tenantID int64, keystoreID int64) (result AllVirtualTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantKeystore{}).Where("`tenant_id` = ? AND `keystore_id` = ?", tenantID, keystoreID).First(&result).Error

	return
}
