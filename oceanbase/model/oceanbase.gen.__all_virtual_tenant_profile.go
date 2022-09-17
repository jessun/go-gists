package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualTenantProfileMgr struct {
	*_BaseMgr
}

// AllVirtualTenantProfileMgr open func
func AllVirtualTenantProfileMgr(db *gorm.DB) *_AllVirtualTenantProfileMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantProfileMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantProfileMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_profile"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantProfileMgr) GetTableName() string {
	return "__all_virtual_tenant_profile"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantProfileMgr) Reset() *_AllVirtualTenantProfileMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantProfileMgr) Get() (result AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantProfileMgr) Gets() (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantProfileMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantProfileMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithProfileID profile_id获取
func (obj *_AllVirtualTenantProfileMgr) WithProfileID(profileID int64) Option {
	return optionFunc(func(o *options) { o.query["profile_id"] = profileID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTenantProfileMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTenantProfileMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithProfileName profile_name获取
func (obj *_AllVirtualTenantProfileMgr) WithProfileName(profileName string) Option {
	return optionFunc(func(o *options) { o.query["profile_name"] = profileName })
}

// WithFailedLoginAttempts failed_login_attempts获取
func (obj *_AllVirtualTenantProfileMgr) WithFailedLoginAttempts(failedLoginAttempts int64) Option {
	return optionFunc(func(o *options) { o.query["failed_login_attempts"] = failedLoginAttempts })
}

// WithPasswordLockTime password_lock_time获取
func (obj *_AllVirtualTenantProfileMgr) WithPasswordLockTime(passwordLockTime int64) Option {
	return optionFunc(func(o *options) { o.query["password_lock_time"] = passwordLockTime })
}

// WithPasswordVerifyFunction password_verify_function获取
func (obj *_AllVirtualTenantProfileMgr) WithPasswordVerifyFunction(passwordVerifyFunction string) Option {
	return optionFunc(func(o *options) { o.query["password_verify_function"] = passwordVerifyFunction })
}

// WithPasswordLifeTime password_life_time获取
func (obj *_AllVirtualTenantProfileMgr) WithPasswordLifeTime(passwordLifeTime int64) Option {
	return optionFunc(func(o *options) { o.query["password_life_time"] = passwordLifeTime })
}

// WithPasswordGraceTime password_grace_time获取
func (obj *_AllVirtualTenantProfileMgr) WithPasswordGraceTime(passwordGraceTime int64) Option {
	return optionFunc(func(o *options) { o.query["password_grace_time"] = passwordGraceTime })
}

// WithPasswordReuseTime password_reuse_time获取
func (obj *_AllVirtualTenantProfileMgr) WithPasswordReuseTime(passwordReuseTime int64) Option {
	return optionFunc(func(o *options) { o.query["password_reuse_time"] = passwordReuseTime })
}

// WithPasswordReuseMax password_reuse_max获取
func (obj *_AllVirtualTenantProfileMgr) WithPasswordReuseMax(passwordReuseMax int64) Option {
	return optionFunc(func(o *options) { o.query["password_reuse_max"] = passwordReuseMax })
}

// WithInactiveAccountTime inactive_account_time获取
func (obj *_AllVirtualTenantProfileMgr) WithInactiveAccountTime(inactiveAccountTime int64) Option {
	return optionFunc(func(o *options) { o.query["inactive_account_time"] = inactiveAccountTime })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantProfileMgr) GetByOption(opts ...Option) (result AllVirtualTenantProfile, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantProfileMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantProfile, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantProfileMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantProfile, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where(options.query)
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
func (obj *_AllVirtualTenantProfileMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromProfileID 通过profile_id获取内容
func (obj *_AllVirtualTenantProfileMgr) GetFromProfileID(profileID int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`profile_id` = ?", profileID).Find(&results).Error

	return
}

// GetBatchFromProfileID 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromProfileID(profileIDs []int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`profile_id` IN (?)", profileIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTenantProfileMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTenantProfileMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromProfileName 通过profile_name获取内容
func (obj *_AllVirtualTenantProfileMgr) GetFromProfileName(profileName string) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`profile_name` = ?", profileName).Find(&results).Error

	return
}

// GetBatchFromProfileName 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromProfileName(profileNames []string) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`profile_name` IN (?)", profileNames).Find(&results).Error

	return
}

// GetFromFailedLoginAttempts 通过failed_login_attempts获取内容
func (obj *_AllVirtualTenantProfileMgr) GetFromFailedLoginAttempts(failedLoginAttempts int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`failed_login_attempts` = ?", failedLoginAttempts).Find(&results).Error

	return
}

// GetBatchFromFailedLoginAttempts 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromFailedLoginAttempts(failedLoginAttemptss []int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`failed_login_attempts` IN (?)", failedLoginAttemptss).Find(&results).Error

	return
}

// GetFromPasswordLockTime 通过password_lock_time获取内容
func (obj *_AllVirtualTenantProfileMgr) GetFromPasswordLockTime(passwordLockTime int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`password_lock_time` = ?", passwordLockTime).Find(&results).Error

	return
}

// GetBatchFromPasswordLockTime 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromPasswordLockTime(passwordLockTimes []int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`password_lock_time` IN (?)", passwordLockTimes).Find(&results).Error

	return
}

// GetFromPasswordVerifyFunction 通过password_verify_function获取内容
func (obj *_AllVirtualTenantProfileMgr) GetFromPasswordVerifyFunction(passwordVerifyFunction string) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`password_verify_function` = ?", passwordVerifyFunction).Find(&results).Error

	return
}

// GetBatchFromPasswordVerifyFunction 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromPasswordVerifyFunction(passwordVerifyFunctions []string) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`password_verify_function` IN (?)", passwordVerifyFunctions).Find(&results).Error

	return
}

// GetFromPasswordLifeTime 通过password_life_time获取内容
func (obj *_AllVirtualTenantProfileMgr) GetFromPasswordLifeTime(passwordLifeTime int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`password_life_time` = ?", passwordLifeTime).Find(&results).Error

	return
}

// GetBatchFromPasswordLifeTime 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromPasswordLifeTime(passwordLifeTimes []int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`password_life_time` IN (?)", passwordLifeTimes).Find(&results).Error

	return
}

// GetFromPasswordGraceTime 通过password_grace_time获取内容
func (obj *_AllVirtualTenantProfileMgr) GetFromPasswordGraceTime(passwordGraceTime int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`password_grace_time` = ?", passwordGraceTime).Find(&results).Error

	return
}

// GetBatchFromPasswordGraceTime 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromPasswordGraceTime(passwordGraceTimes []int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`password_grace_time` IN (?)", passwordGraceTimes).Find(&results).Error

	return
}

// GetFromPasswordReuseTime 通过password_reuse_time获取内容
func (obj *_AllVirtualTenantProfileMgr) GetFromPasswordReuseTime(passwordReuseTime int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`password_reuse_time` = ?", passwordReuseTime).Find(&results).Error

	return
}

// GetBatchFromPasswordReuseTime 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromPasswordReuseTime(passwordReuseTimes []int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`password_reuse_time` IN (?)", passwordReuseTimes).Find(&results).Error

	return
}

// GetFromPasswordReuseMax 通过password_reuse_max获取内容
func (obj *_AllVirtualTenantProfileMgr) GetFromPasswordReuseMax(passwordReuseMax int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`password_reuse_max` = ?", passwordReuseMax).Find(&results).Error

	return
}

// GetBatchFromPasswordReuseMax 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromPasswordReuseMax(passwordReuseMaxs []int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`password_reuse_max` IN (?)", passwordReuseMaxs).Find(&results).Error

	return
}

// GetFromInactiveAccountTime 通过inactive_account_time获取内容
func (obj *_AllVirtualTenantProfileMgr) GetFromInactiveAccountTime(inactiveAccountTime int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`inactive_account_time` = ?", inactiveAccountTime).Find(&results).Error

	return
}

// GetBatchFromInactiveAccountTime 批量查找
func (obj *_AllVirtualTenantProfileMgr) GetBatchFromInactiveAccountTime(inactiveAccountTimes []int64) (results []*AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`inactive_account_time` IN (?)", inactiveAccountTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTenantProfileMgr) FetchByPrimaryKey(tenantID int64, profileID int64) (result AllVirtualTenantProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfile{}).Where("`tenant_id` = ? AND `profile_id` = ?", tenantID, profileID).First(&result).Error

	return
}
