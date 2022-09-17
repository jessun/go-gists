package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualTenantProfileHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualTenantProfileHistoryMgr open func
func AllVirtualTenantProfileHistoryMgr(db *gorm.DB) *_AllVirtualTenantProfileHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantProfileHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantProfileHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_profile_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantProfileHistoryMgr) GetTableName() string {
	return "__all_virtual_tenant_profile_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantProfileHistoryMgr) Reset() *_AllVirtualTenantProfileHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantProfileHistoryMgr) Get() (result AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantProfileHistoryMgr) Gets() (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantProfileHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithProfileID profile_id获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithProfileID(profileID int64) Option {
	return optionFunc(func(o *options) { o.query["profile_id"] = profileID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithProfileName profile_name获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithProfileName(profileName string) Option {
	return optionFunc(func(o *options) { o.query["profile_name"] = profileName })
}

// WithFailedLoginAttempts failed_login_attempts获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithFailedLoginAttempts(failedLoginAttempts int64) Option {
	return optionFunc(func(o *options) { o.query["failed_login_attempts"] = failedLoginAttempts })
}

// WithPasswordLockTime password_lock_time获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithPasswordLockTime(passwordLockTime int64) Option {
	return optionFunc(func(o *options) { o.query["password_lock_time"] = passwordLockTime })
}

// WithPasswordVerifyFunction password_verify_function获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithPasswordVerifyFunction(passwordVerifyFunction string) Option {
	return optionFunc(func(o *options) { o.query["password_verify_function"] = passwordVerifyFunction })
}

// WithPasswordLifeTime password_life_time获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithPasswordLifeTime(passwordLifeTime int64) Option {
	return optionFunc(func(o *options) { o.query["password_life_time"] = passwordLifeTime })
}

// WithPasswordGraceTime password_grace_time获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithPasswordGraceTime(passwordGraceTime int64) Option {
	return optionFunc(func(o *options) { o.query["password_grace_time"] = passwordGraceTime })
}

// WithPasswordReuseTime password_reuse_time获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithPasswordReuseTime(passwordReuseTime int64) Option {
	return optionFunc(func(o *options) { o.query["password_reuse_time"] = passwordReuseTime })
}

// WithPasswordReuseMax password_reuse_max获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithPasswordReuseMax(passwordReuseMax int64) Option {
	return optionFunc(func(o *options) { o.query["password_reuse_max"] = passwordReuseMax })
}

// WithInactiveAccountTime inactive_account_time获取
func (obj *_AllVirtualTenantProfileHistoryMgr) WithInactiveAccountTime(inactiveAccountTime int64) Option {
	return optionFunc(func(o *options) { o.query["inactive_account_time"] = inactiveAccountTime })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantProfileHistoryMgr) GetByOption(opts ...Option) (result AllVirtualTenantProfileHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantProfileHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantProfileHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantProfileHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantProfileHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where(options.query)
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
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromProfileID 通过profile_id获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromProfileID(profileID int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`profile_id` = ?", profileID).Find(&results).Error

	return
}

// GetBatchFromProfileID 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromProfileID(profileIDs []int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`profile_id` IN (?)", profileIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromProfileName 通过profile_name获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromProfileName(profileName string) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`profile_name` = ?", profileName).Find(&results).Error

	return
}

// GetBatchFromProfileName 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromProfileName(profileNames []string) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`profile_name` IN (?)", profileNames).Find(&results).Error

	return
}

// GetFromFailedLoginAttempts 通过failed_login_attempts获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromFailedLoginAttempts(failedLoginAttempts int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`failed_login_attempts` = ?", failedLoginAttempts).Find(&results).Error

	return
}

// GetBatchFromFailedLoginAttempts 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromFailedLoginAttempts(failedLoginAttemptss []int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`failed_login_attempts` IN (?)", failedLoginAttemptss).Find(&results).Error

	return
}

// GetFromPasswordLockTime 通过password_lock_time获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromPasswordLockTime(passwordLockTime int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`password_lock_time` = ?", passwordLockTime).Find(&results).Error

	return
}

// GetBatchFromPasswordLockTime 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromPasswordLockTime(passwordLockTimes []int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`password_lock_time` IN (?)", passwordLockTimes).Find(&results).Error

	return
}

// GetFromPasswordVerifyFunction 通过password_verify_function获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromPasswordVerifyFunction(passwordVerifyFunction string) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`password_verify_function` = ?", passwordVerifyFunction).Find(&results).Error

	return
}

// GetBatchFromPasswordVerifyFunction 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromPasswordVerifyFunction(passwordVerifyFunctions []string) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`password_verify_function` IN (?)", passwordVerifyFunctions).Find(&results).Error

	return
}

// GetFromPasswordLifeTime 通过password_life_time获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromPasswordLifeTime(passwordLifeTime int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`password_life_time` = ?", passwordLifeTime).Find(&results).Error

	return
}

// GetBatchFromPasswordLifeTime 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromPasswordLifeTime(passwordLifeTimes []int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`password_life_time` IN (?)", passwordLifeTimes).Find(&results).Error

	return
}

// GetFromPasswordGraceTime 通过password_grace_time获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromPasswordGraceTime(passwordGraceTime int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`password_grace_time` = ?", passwordGraceTime).Find(&results).Error

	return
}

// GetBatchFromPasswordGraceTime 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromPasswordGraceTime(passwordGraceTimes []int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`password_grace_time` IN (?)", passwordGraceTimes).Find(&results).Error

	return
}

// GetFromPasswordReuseTime 通过password_reuse_time获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromPasswordReuseTime(passwordReuseTime int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`password_reuse_time` = ?", passwordReuseTime).Find(&results).Error

	return
}

// GetBatchFromPasswordReuseTime 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromPasswordReuseTime(passwordReuseTimes []int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`password_reuse_time` IN (?)", passwordReuseTimes).Find(&results).Error

	return
}

// GetFromPasswordReuseMax 通过password_reuse_max获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromPasswordReuseMax(passwordReuseMax int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`password_reuse_max` = ?", passwordReuseMax).Find(&results).Error

	return
}

// GetBatchFromPasswordReuseMax 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromPasswordReuseMax(passwordReuseMaxs []int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`password_reuse_max` IN (?)", passwordReuseMaxs).Find(&results).Error

	return
}

// GetFromInactiveAccountTime 通过inactive_account_time获取内容
func (obj *_AllVirtualTenantProfileHistoryMgr) GetFromInactiveAccountTime(inactiveAccountTime int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`inactive_account_time` = ?", inactiveAccountTime).Find(&results).Error

	return
}

// GetBatchFromInactiveAccountTime 批量查找
func (obj *_AllVirtualTenantProfileHistoryMgr) GetBatchFromInactiveAccountTime(inactiveAccountTimes []int64) (results []*AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`inactive_account_time` IN (?)", inactiveAccountTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTenantProfileHistoryMgr) FetchByPrimaryKey(tenantID int64, profileID int64, schemaVersion int64) (result AllVirtualTenantProfileHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantProfileHistory{}).Where("`tenant_id` = ? AND `profile_id` = ? AND `schema_version` = ?", tenantID, profileID, schemaVersion).First(&result).Error

	return
}
