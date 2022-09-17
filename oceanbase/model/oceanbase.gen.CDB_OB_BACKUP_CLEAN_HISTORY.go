package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CdbObBackupCleanHistoryMgr struct {
	*_BaseMgr
}

// CdbObBackupCleanHistoryMgr open func
func CdbObBackupCleanHistoryMgr(db *gorm.DB) *_CdbObBackupCleanHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupCleanHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupCleanHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_CLEAN_HISTORY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupCleanHistoryMgr) GetTableName() string {
	return "CDB_OB_BACKUP_CLEAN_HISTORY"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupCleanHistoryMgr) Reset() *_CdbObBackupCleanHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupCleanHistoryMgr) Get() (result CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupCleanHistoryMgr) Gets() (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupCleanHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupCleanHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithBsKey BS_KEY获取
func (obj *_CdbObBackupCleanHistoryMgr) WithBsKey(bsKey int64) Option {
	return optionFunc(func(o *options) { o.query["BS_KEY"] = bsKey })
}

// WithCopyID COPY_ID获取
func (obj *_CdbObBackupCleanHistoryMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["COPY_ID"] = copyID })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupCleanHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithEndTime END_TIME获取
func (obj *_CdbObBackupCleanHistoryMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["END_TIME"] = endTime })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupCleanHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithType TYPE获取
func (obj *_CdbObBackupCleanHistoryMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["TYPE"] = _type })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupCleanHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithParameter PARAMETER获取
func (obj *_CdbObBackupCleanHistoryMgr) WithParameter(parameter string) Option {
	return optionFunc(func(o *options) { o.query["PARAMETER"] = parameter })
}

// WithErrorMsg ERROR_MSG获取
func (obj *_CdbObBackupCleanHistoryMgr) WithErrorMsg(errorMsg string) Option {
	return optionFunc(func(o *options) { o.query["ERROR_MSG"] = errorMsg })
}

// WithComment COMMENT获取
func (obj *_CdbObBackupCleanHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["COMMENT"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupCleanHistoryMgr) GetByOption(opts ...Option) (result CdbObBackupCleanHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupCleanHistoryMgr) GetByOptions(opts ...Option) (results []*CdbObBackupCleanHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupCleanHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupCleanHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where(options.query)
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
func (obj *_CdbObBackupCleanHistoryMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupCleanHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromBsKey 通过BS_KEY获取内容
func (obj *_CdbObBackupCleanHistoryMgr) GetFromBsKey(bsKey int64) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`BS_KEY` = ?", bsKey).Find(&results).Error

	return
}

// GetBatchFromBsKey 批量查找
func (obj *_CdbObBackupCleanHistoryMgr) GetBatchFromBsKey(bsKeys []int64) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`BS_KEY` IN (?)", bsKeys).Find(&results).Error

	return
}

// GetFromCopyID 通过COPY_ID获取内容
func (obj *_CdbObBackupCleanHistoryMgr) GetFromCopyID(copyID int64) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`COPY_ID` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_CdbObBackupCleanHistoryMgr) GetBatchFromCopyID(copyIDs []int64) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`COPY_ID` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupCleanHistoryMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupCleanHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过END_TIME获取内容
func (obj *_CdbObBackupCleanHistoryMgr) GetFromEndTime(endTime time.Time) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`END_TIME` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_CdbObBackupCleanHistoryMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`END_TIME` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupCleanHistoryMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupCleanHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromType 通过TYPE获取内容
func (obj *_CdbObBackupCleanHistoryMgr) GetFromType(_type string) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`TYPE` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_CdbObBackupCleanHistoryMgr) GetBatchFromType(_types []string) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`TYPE` IN (?)", _types).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupCleanHistoryMgr) GetFromStatus(status string) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupCleanHistoryMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromParameter 通过PARAMETER获取内容
func (obj *_CdbObBackupCleanHistoryMgr) GetFromParameter(parameter string) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`PARAMETER` = ?", parameter).Find(&results).Error

	return
}

// GetBatchFromParameter 批量查找
func (obj *_CdbObBackupCleanHistoryMgr) GetBatchFromParameter(parameters []string) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`PARAMETER` IN (?)", parameters).Find(&results).Error

	return
}

// GetFromErrorMsg 通过ERROR_MSG获取内容
func (obj *_CdbObBackupCleanHistoryMgr) GetFromErrorMsg(errorMsg string) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`ERROR_MSG` = ?", errorMsg).Find(&results).Error

	return
}

// GetBatchFromErrorMsg 批量查找
func (obj *_CdbObBackupCleanHistoryMgr) GetBatchFromErrorMsg(errorMsgs []string) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`ERROR_MSG` IN (?)", errorMsgs).Find(&results).Error

	return
}

// GetFromComment 通过COMMENT获取内容
func (obj *_CdbObBackupCleanHistoryMgr) GetFromComment(comment string) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`COMMENT` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_CdbObBackupCleanHistoryMgr) GetBatchFromComment(comments []string) (results []*CdbObBackupCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupCleanHistory{}).Where("`COMMENT` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
