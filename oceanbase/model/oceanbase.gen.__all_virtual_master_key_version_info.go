package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualMasterKeyVersionInfoMgr struct {
	*_BaseMgr
}

// AllVirtualMasterKeyVersionInfoMgr open func
func AllVirtualMasterKeyVersionInfoMgr(db *gorm.DB) *_AllVirtualMasterKeyVersionInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualMasterKeyVersionInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualMasterKeyVersionInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_master_key_version_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetTableName() string {
	return "__all_virtual_master_key_version_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualMasterKeyVersionInfoMgr) Reset() *_AllVirtualMasterKeyVersionInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualMasterKeyVersionInfoMgr) Get() (result AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualMasterKeyVersionInfoMgr) Gets() (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualMasterKeyVersionInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualMasterKeyVersionInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualMasterKeyVersionInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualMasterKeyVersionInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithMaxActiveVersion max_active_version获取
func (obj *_AllVirtualMasterKeyVersionInfoMgr) WithMaxActiveVersion(maxActiveVersion int64) Option {
	return optionFunc(func(o *options) { o.query["max_active_version"] = maxActiveVersion })
}

// WithMaxStoredVersion max_stored_version获取
func (obj *_AllVirtualMasterKeyVersionInfoMgr) WithMaxStoredVersion(maxStoredVersion int64) Option {
	return optionFunc(func(o *options) { o.query["max_stored_version"] = maxStoredVersion })
}

// WithExpectVersion expect_version获取
func (obj *_AllVirtualMasterKeyVersionInfoMgr) WithExpectVersion(expectVersion int64) Option {
	return optionFunc(func(o *options) { o.query["expect_version"] = expectVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetByOption(opts ...Option) (result AllVirtualMasterKeyVersionInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualMasterKeyVersionInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualMasterKeyVersionInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where(options.query)
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

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromMaxActiveVersion 通过max_active_version获取内容
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetFromMaxActiveVersion(maxActiveVersion int64) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`max_active_version` = ?", maxActiveVersion).Find(&results).Error

	return
}

// GetBatchFromMaxActiveVersion 批量查找
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetBatchFromMaxActiveVersion(maxActiveVersions []int64) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`max_active_version` IN (?)", maxActiveVersions).Find(&results).Error

	return
}

// GetFromMaxStoredVersion 通过max_stored_version获取内容
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetFromMaxStoredVersion(maxStoredVersion int64) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`max_stored_version` = ?", maxStoredVersion).Find(&results).Error

	return
}

// GetBatchFromMaxStoredVersion 批量查找
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetBatchFromMaxStoredVersion(maxStoredVersions []int64) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`max_stored_version` IN (?)", maxStoredVersions).Find(&results).Error

	return
}

// GetFromExpectVersion 通过expect_version获取内容
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetFromExpectVersion(expectVersion int64) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`expect_version` = ?", expectVersion).Find(&results).Error

	return
}

// GetBatchFromExpectVersion 批量查找
func (obj *_AllVirtualMasterKeyVersionInfoMgr) GetBatchFromExpectVersion(expectVersions []int64) (results []*AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`expect_version` IN (?)", expectVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualMasterKeyVersionInfoMgr) FetchByPrimaryKey(svrIP string, svrPort int64, tenantID int64) (result AllVirtualMasterKeyVersionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMasterKeyVersionInfo{}).Where("`svr_ip` = ? AND `svr_port` = ? AND `tenant_id` = ?", svrIP, svrPort, tenantID).First(&result).Error

	return
}
