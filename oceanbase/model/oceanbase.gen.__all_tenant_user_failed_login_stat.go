package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantUserFailedLoginStatMgr struct {
	*_BaseMgr
}

// AllTenantUserFailedLoginStatMgr open func
func AllTenantUserFailedLoginStatMgr(db *gorm.DB) *_AllTenantUserFailedLoginStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantUserFailedLoginStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantUserFailedLoginStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_user_failed_login_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantUserFailedLoginStatMgr) GetTableName() string {
	return "__all_tenant_user_failed_login_stat"
}

// Reset 重置gorm会话
func (obj *_AllTenantUserFailedLoginStatMgr) Reset() *_AllTenantUserFailedLoginStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantUserFailedLoginStatMgr) Get() (result AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantUserFailedLoginStatMgr) Gets() (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantUserFailedLoginStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantUserFailedLoginStatMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantUserFailedLoginStatMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantUserFailedLoginStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserID user_id获取
func (obj *_AllTenantUserFailedLoginStatMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUserName user_name获取
func (obj *_AllTenantUserFailedLoginStatMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithFailedLoginAttempts failed_login_attempts获取
func (obj *_AllTenantUserFailedLoginStatMgr) WithFailedLoginAttempts(failedLoginAttempts int64) Option {
	return optionFunc(func(o *options) { o.query["failed_login_attempts"] = failedLoginAttempts })
}

// WithLastFailedLoginSvrIP last_failed_login_svr_ip获取
func (obj *_AllTenantUserFailedLoginStatMgr) WithLastFailedLoginSvrIP(lastFailedLoginSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["last_failed_login_svr_ip"] = lastFailedLoginSvrIP })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantUserFailedLoginStatMgr) GetByOption(opts ...Option) (result AllTenantUserFailedLoginStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantUserFailedLoginStatMgr) GetByOptions(opts ...Option) (results []*AllTenantUserFailedLoginStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantUserFailedLoginStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantUserFailedLoginStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where(options.query)
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
func (obj *_AllTenantUserFailedLoginStatMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantUserFailedLoginStatMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantUserFailedLoginStatMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantUserFailedLoginStatMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantUserFailedLoginStatMgr) GetFromTenantID(tenantID int64) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantUserFailedLoginStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllTenantUserFailedLoginStatMgr) GetFromUserID(userID int64) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllTenantUserFailedLoginStatMgr) GetBatchFromUserID(userIDs []int64) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容
func (obj *_AllTenantUserFailedLoginStatMgr) GetFromUserName(userName string) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_AllTenantUserFailedLoginStatMgr) GetBatchFromUserName(userNames []string) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromFailedLoginAttempts 通过failed_login_attempts获取内容
func (obj *_AllTenantUserFailedLoginStatMgr) GetFromFailedLoginAttempts(failedLoginAttempts int64) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`failed_login_attempts` = ?", failedLoginAttempts).Find(&results).Error

	return
}

// GetBatchFromFailedLoginAttempts 批量查找
func (obj *_AllTenantUserFailedLoginStatMgr) GetBatchFromFailedLoginAttempts(failedLoginAttemptss []int64) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`failed_login_attempts` IN (?)", failedLoginAttemptss).Find(&results).Error

	return
}

// GetFromLastFailedLoginSvrIP 通过last_failed_login_svr_ip获取内容
func (obj *_AllTenantUserFailedLoginStatMgr) GetFromLastFailedLoginSvrIP(lastFailedLoginSvrIP string) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`last_failed_login_svr_ip` = ?", lastFailedLoginSvrIP).Find(&results).Error

	return
}

// GetBatchFromLastFailedLoginSvrIP 批量查找
func (obj *_AllTenantUserFailedLoginStatMgr) GetBatchFromLastFailedLoginSvrIP(lastFailedLoginSvrIPs []string) (results []*AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`last_failed_login_svr_ip` IN (?)", lastFailedLoginSvrIPs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantUserFailedLoginStatMgr) FetchByPrimaryKey(tenantID int64, userID int64) (result AllTenantUserFailedLoginStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantUserFailedLoginStat{}).Where("`tenant_id` = ? AND `user_id` = ?", tenantID, userID).First(&result).Error

	return
}
