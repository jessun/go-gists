package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualTenantUserFailedLoginStatMgr struct {
	*_BaseMgr
}

// AllVirtualTenantUserFailedLoginStatMgr open func
func AllVirtualTenantUserFailedLoginStatMgr(db *gorm.DB) *_AllVirtualTenantUserFailedLoginStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantUserFailedLoginStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantUserFailedLoginStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_user_failed_login_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetTableName() string {
	return "__all_virtual_tenant_user_failed_login_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) Reset() *_AllVirtualTenantUserFailedLoginStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) Get() (result AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) Gets() (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserID user_id获取
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithUserName user_name获取
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithFailedLoginAttempts failed_login_attempts获取
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) WithFailedLoginAttempts(failedLoginAttempts int64) Option {
	return optionFunc(func(o *options) { o.query["failed_login_attempts"] = failedLoginAttempts })
}

// WithLastFailedLoginSvrIP last_failed_login_svr_ip获取
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) WithLastFailedLoginSvrIP(lastFailedLoginSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["last_failed_login_svr_ip"] = lastFailedLoginSvrIP })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetByOption(opts ...Option) (result AllVirtualTenantUserFailedLoginStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantUserFailedLoginStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where(options.query)
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
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetFromUserID(userID int64) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetBatchFromUserID(userIDs []int64) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetFromUserName(userName string) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetBatchFromUserName(userNames []string) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromFailedLoginAttempts 通过failed_login_attempts获取内容
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetFromFailedLoginAttempts(failedLoginAttempts int64) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`failed_login_attempts` = ?", failedLoginAttempts).Find(&results).Error

	return
}

// GetBatchFromFailedLoginAttempts 批量查找
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetBatchFromFailedLoginAttempts(failedLoginAttemptss []int64) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`failed_login_attempts` IN (?)", failedLoginAttemptss).Find(&results).Error

	return
}

// GetFromLastFailedLoginSvrIP 通过last_failed_login_svr_ip获取内容
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetFromLastFailedLoginSvrIP(lastFailedLoginSvrIP string) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`last_failed_login_svr_ip` = ?", lastFailedLoginSvrIP).Find(&results).Error

	return
}

// GetBatchFromLastFailedLoginSvrIP 批量查找
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) GetBatchFromLastFailedLoginSvrIP(lastFailedLoginSvrIPs []string) (results []*AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`last_failed_login_svr_ip` IN (?)", lastFailedLoginSvrIPs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTenantUserFailedLoginStatMgr) FetchByPrimaryKey(tenantID int64, userID int64) (result AllVirtualTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantUserFailedLoginStat{}).Where("`tenant_id` = ? AND `user_id` = ?", tenantID, userID).First(&result).Error

	return
}
