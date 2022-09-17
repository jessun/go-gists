package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantKeystoreMgr struct {
	*_BaseMgr
}

// AllTenantKeystoreMgr open func
func AllTenantKeystoreMgr(db *gorm.DB) *_AllTenantKeystoreMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantKeystoreMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantKeystoreMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_keystore"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantKeystoreMgr) GetTableName() string {
	return "__all_tenant_keystore"
}

// Reset 重置gorm会话
func (obj *_AllTenantKeystoreMgr) Reset() *_AllTenantKeystoreMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantKeystoreMgr) Get() (result AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantKeystoreMgr) Gets() (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantKeystoreMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantKeystoreMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantKeystoreMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantKeystoreMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithKeystoreID keystore_id获取
func (obj *_AllTenantKeystoreMgr) WithKeystoreID(keystoreID int64) Option {
	return optionFunc(func(o *options) { o.query["keystore_id"] = keystoreID })
}

// WithKeystoreName keystore_name获取
func (obj *_AllTenantKeystoreMgr) WithKeystoreName(keystoreName string) Option {
	return optionFunc(func(o *options) { o.query["keystore_name"] = keystoreName })
}

// WithPassword password获取
func (obj *_AllTenantKeystoreMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithStatus status获取
func (obj *_AllTenantKeystoreMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithMasterKeyID master_key_id获取
func (obj *_AllTenantKeystoreMgr) WithMasterKeyID(masterKeyID int64) Option {
	return optionFunc(func(o *options) { o.query["master_key_id"] = masterKeyID })
}

// WithMasterKey master_key获取
func (obj *_AllTenantKeystoreMgr) WithMasterKey(masterKey string) Option {
	return optionFunc(func(o *options) { o.query["master_key"] = masterKey })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantKeystoreMgr) GetByOption(opts ...Option) (result AllTenantKeystore, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantKeystoreMgr) GetByOptions(opts ...Option) (results []*AllTenantKeystore, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantKeystoreMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantKeystore, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where(options.query)
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
func (obj *_AllTenantKeystoreMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantKeystoreMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantKeystoreMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantKeystoreMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantKeystoreMgr) GetFromTenantID(tenantID int64) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantKeystoreMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromKeystoreID 通过keystore_id获取内容
func (obj *_AllTenantKeystoreMgr) GetFromKeystoreID(keystoreID int64) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`keystore_id` = ?", keystoreID).Find(&results).Error

	return
}

// GetBatchFromKeystoreID 批量查找
func (obj *_AllTenantKeystoreMgr) GetBatchFromKeystoreID(keystoreIDs []int64) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`keystore_id` IN (?)", keystoreIDs).Find(&results).Error

	return
}

// GetFromKeystoreName 通过keystore_name获取内容
func (obj *_AllTenantKeystoreMgr) GetFromKeystoreName(keystoreName string) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`keystore_name` = ?", keystoreName).Find(&results).Error

	return
}

// GetBatchFromKeystoreName 批量查找
func (obj *_AllTenantKeystoreMgr) GetBatchFromKeystoreName(keystoreNames []string) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`keystore_name` IN (?)", keystoreNames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_AllTenantKeystoreMgr) GetFromPassword(password string) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_AllTenantKeystoreMgr) GetBatchFromPassword(passwords []string) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllTenantKeystoreMgr) GetFromStatus(status int64) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllTenantKeystoreMgr) GetBatchFromStatus(statuss []int64) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromMasterKeyID 通过master_key_id获取内容
func (obj *_AllTenantKeystoreMgr) GetFromMasterKeyID(masterKeyID int64) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`master_key_id` = ?", masterKeyID).Find(&results).Error

	return
}

// GetBatchFromMasterKeyID 批量查找
func (obj *_AllTenantKeystoreMgr) GetBatchFromMasterKeyID(masterKeyIDs []int64) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`master_key_id` IN (?)", masterKeyIDs).Find(&results).Error

	return
}

// GetFromMasterKey 通过master_key获取内容
func (obj *_AllTenantKeystoreMgr) GetFromMasterKey(masterKey string) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`master_key` = ?", masterKey).Find(&results).Error

	return
}

// GetBatchFromMasterKey 批量查找
func (obj *_AllTenantKeystoreMgr) GetBatchFromMasterKey(masterKeys []string) (results []*AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`master_key` IN (?)", masterKeys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantKeystoreMgr) FetchByPrimaryKey(tenantID int64, keystoreID int64) (result AllTenantKeystore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantKeystore{}).Where("`tenant_id` = ? AND `keystore_id` = ?", tenantID, keystoreID).First(&result).Error

	return
}
