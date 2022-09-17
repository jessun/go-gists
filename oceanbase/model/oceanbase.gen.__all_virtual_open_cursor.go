package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualOpenCursorMgr struct {
	*_BaseMgr
}

// AllVirtualOpenCursorMgr open func
func AllVirtualOpenCursorMgr(db *gorm.DB) *_AllVirtualOpenCursorMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualOpenCursorMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualOpenCursorMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_open_cursor"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualOpenCursorMgr) GetTableName() string {
	return "__all_virtual_open_cursor"
}

// Reset 重置gorm会话
func (obj *_AllVirtualOpenCursorMgr) Reset() *_AllVirtualOpenCursorMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualOpenCursorMgr) Get() (result AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualOpenCursorMgr) Gets() (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualOpenCursorMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取
func (obj *_AllVirtualOpenCursorMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithSvrIP SVR_IP获取
func (obj *_AllVirtualOpenCursorMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取
func (obj *_AllVirtualOpenCursorMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithSaddr SADDR获取
func (obj *_AllVirtualOpenCursorMgr) WithSaddr(saddr string) Option {
	return optionFunc(func(o *options) { o.query["SADDR"] = saddr })
}

// WithSid SID获取
func (obj *_AllVirtualOpenCursorMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["SID"] = sid })
}

// WithUserName USER_NAME获取
func (obj *_AllVirtualOpenCursorMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["USER_NAME"] = userName })
}

// WithAddress ADDRESS获取
func (obj *_AllVirtualOpenCursorMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["ADDRESS"] = address })
}

// WithHashValue HASH_VALUE获取
func (obj *_AllVirtualOpenCursorMgr) WithHashValue(hashValue int64) Option {
	return optionFunc(func(o *options) { o.query["HASH_VALUE"] = hashValue })
}

// WithSQLID SQL_ID获取
func (obj *_AllVirtualOpenCursorMgr) WithSQLID(sqlID string) Option {
	return optionFunc(func(o *options) { o.query["SQL_ID"] = sqlID })
}

// WithSQLText SQL_TEXT获取
func (obj *_AllVirtualOpenCursorMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["SQL_TEXT"] = sqlText })
}

// WithLastSQLActiveTime LAST_SQL_ACTIVE_TIME获取
func (obj *_AllVirtualOpenCursorMgr) WithLastSQLActiveTime(lastSQLActiveTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["LAST_SQL_ACTIVE_TIME"] = lastSQLActiveTime })
}

// WithSQLExecID SQL_EXEC_ID获取
func (obj *_AllVirtualOpenCursorMgr) WithSQLExecID(sqlExecID int64) Option {
	return optionFunc(func(o *options) { o.query["SQL_EXEC_ID"] = sqlExecID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualOpenCursorMgr) GetByOption(opts ...Option) (result AllVirtualOpenCursor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualOpenCursorMgr) GetByOptions(opts ...Option) (results []*AllVirtualOpenCursor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualOpenCursorMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualOpenCursor, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where(options.query)
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

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_AllVirtualOpenCursorMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualOpenCursorMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过SVR_IP获取内容
func (obj *_AllVirtualOpenCursorMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualOpenCursorMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过SVR_PORT获取内容
func (obj *_AllVirtualOpenCursorMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualOpenCursorMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromSaddr 通过SADDR获取内容
func (obj *_AllVirtualOpenCursorMgr) GetFromSaddr(saddr string) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SADDR` = ?", saddr).Find(&results).Error

	return
}

// GetBatchFromSaddr 批量查找
func (obj *_AllVirtualOpenCursorMgr) GetBatchFromSaddr(saddrs []string) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SADDR` IN (?)", saddrs).Find(&results).Error

	return
}

// GetFromSid 通过SID获取内容
func (obj *_AllVirtualOpenCursorMgr) GetFromSid(sid int64) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SID` = ?", sid).Find(&results).Error

	return
}

// GetBatchFromSid 批量查找
func (obj *_AllVirtualOpenCursorMgr) GetBatchFromSid(sids []int64) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SID` IN (?)", sids).Find(&results).Error

	return
}

// GetFromUserName 通过USER_NAME获取内容
func (obj *_AllVirtualOpenCursorMgr) GetFromUserName(userName string) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`USER_NAME` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_AllVirtualOpenCursorMgr) GetBatchFromUserName(userNames []string) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`USER_NAME` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromAddress 通过ADDRESS获取内容
func (obj *_AllVirtualOpenCursorMgr) GetFromAddress(address string) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`ADDRESS` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找
func (obj *_AllVirtualOpenCursorMgr) GetBatchFromAddress(addresss []string) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`ADDRESS` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromHashValue 通过HASH_VALUE获取内容
func (obj *_AllVirtualOpenCursorMgr) GetFromHashValue(hashValue int64) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`HASH_VALUE` = ?", hashValue).Find(&results).Error

	return
}

// GetBatchFromHashValue 批量查找
func (obj *_AllVirtualOpenCursorMgr) GetBatchFromHashValue(hashValues []int64) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`HASH_VALUE` IN (?)", hashValues).Find(&results).Error

	return
}

// GetFromSQLID 通过SQL_ID获取内容
func (obj *_AllVirtualOpenCursorMgr) GetFromSQLID(sqlID string) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SQL_ID` = ?", sqlID).Find(&results).Error

	return
}

// GetBatchFromSQLID 批量查找
func (obj *_AllVirtualOpenCursorMgr) GetBatchFromSQLID(sqlIDs []string) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SQL_ID` IN (?)", sqlIDs).Find(&results).Error

	return
}

// GetFromSQLText 通过SQL_TEXT获取内容
func (obj *_AllVirtualOpenCursorMgr) GetFromSQLText(sqlText string) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SQL_TEXT` = ?", sqlText).Find(&results).Error

	return
}

// GetBatchFromSQLText 批量查找
func (obj *_AllVirtualOpenCursorMgr) GetBatchFromSQLText(sqlTexts []string) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SQL_TEXT` IN (?)", sqlTexts).Find(&results).Error

	return
}

// GetFromLastSQLActiveTime 通过LAST_SQL_ACTIVE_TIME获取内容
func (obj *_AllVirtualOpenCursorMgr) GetFromLastSQLActiveTime(lastSQLActiveTime time.Time) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`LAST_SQL_ACTIVE_TIME` = ?", lastSQLActiveTime).Find(&results).Error

	return
}

// GetBatchFromLastSQLActiveTime 批量查找
func (obj *_AllVirtualOpenCursorMgr) GetBatchFromLastSQLActiveTime(lastSQLActiveTimes []time.Time) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`LAST_SQL_ACTIVE_TIME` IN (?)", lastSQLActiveTimes).Find(&results).Error

	return
}

// GetFromSQLExecID 通过SQL_EXEC_ID获取内容
func (obj *_AllVirtualOpenCursorMgr) GetFromSQLExecID(sqlExecID int64) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SQL_EXEC_ID` = ?", sqlExecID).Find(&results).Error

	return
}

// GetBatchFromSQLExecID 批量查找
func (obj *_AllVirtualOpenCursorMgr) GetBatchFromSQLExecID(sqlExecIDs []int64) (results []*AllVirtualOpenCursor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOpenCursor{}).Where("`SQL_EXEC_ID` IN (?)", sqlExecIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
