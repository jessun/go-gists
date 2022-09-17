package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CdbCkptHistoryMgr struct {
	*_BaseMgr
}

// CdbCkptHistoryMgr open func
func CdbCkptHistoryMgr(db *gorm.DB) *_CdbCkptHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbCkptHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbCkptHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_CKPT_HISTORY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbCkptHistoryMgr) GetTableName() string {
	return "CDB_CKPT_HISTORY"
}

// Reset 重置gorm会话
func (obj *_CdbCkptHistoryMgr) Reset() *_CdbCkptHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbCkptHistoryMgr) Get() (result CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbCkptHistoryMgr) Gets() (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbCkptHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP SVR_IP获取
func (obj *_CdbCkptHistoryMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取
func (obj *_CdbCkptHistoryMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbCkptHistoryMgr) WithTenantID(tenantID string) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithCheckpointSnapshot CHECKPOINT_SNAPSHOT获取
func (obj *_CdbCkptHistoryMgr) WithCheckpointSnapshot(checkpointSnapshot string) Option {
	return optionFunc(func(o *options) { o.query["CHECKPOINT_SNAPSHOT"] = checkpointSnapshot })
}

// WithCheckpointType CHECKPOINT_TYPE获取
func (obj *_CdbCkptHistoryMgr) WithCheckpointType(checkpointType string) Option {
	return optionFunc(func(o *options) { o.query["CHECKPOINT_TYPE"] = checkpointType })
}

// WithCheckpointClusterVersion CHECKPOINT_CLUSTER_VERSION获取
func (obj *_CdbCkptHistoryMgr) WithCheckpointClusterVersion(checkpointClusterVersion string) Option {
	return optionFunc(func(o *options) { o.query["CHECKPOINT_CLUSTER_VERSION"] = checkpointClusterVersion })
}

// WithStartTime START_TIME获取
func (obj *_CdbCkptHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithFinishTime FINISH_TIME获取
func (obj *_CdbCkptHistoryMgr) WithFinishTime(finishTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["FINISH_TIME"] = finishTime })
}

// GetByOption 功能选项模式获取
func (obj *_CdbCkptHistoryMgr) GetByOption(opts ...Option) (result CdbCkptHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbCkptHistoryMgr) GetByOptions(opts ...Option) (results []*CdbCkptHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbCkptHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbCkptHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where(options.query)
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

// GetFromSvrIP 通过SVR_IP获取内容
func (obj *_CdbCkptHistoryMgr) GetFromSvrIP(svrIP string) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_CdbCkptHistoryMgr) GetBatchFromSvrIP(svrIPs []string) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过SVR_PORT获取内容
func (obj *_CdbCkptHistoryMgr) GetFromSvrPort(svrPort int64) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_CdbCkptHistoryMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbCkptHistoryMgr) GetFromTenantID(tenantID string) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbCkptHistoryMgr) GetBatchFromTenantID(tenantIDs []string) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromCheckpointSnapshot 通过CHECKPOINT_SNAPSHOT获取内容
func (obj *_CdbCkptHistoryMgr) GetFromCheckpointSnapshot(checkpointSnapshot string) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`CHECKPOINT_SNAPSHOT` = ?", checkpointSnapshot).Find(&results).Error

	return
}

// GetBatchFromCheckpointSnapshot 批量查找
func (obj *_CdbCkptHistoryMgr) GetBatchFromCheckpointSnapshot(checkpointSnapshots []string) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`CHECKPOINT_SNAPSHOT` IN (?)", checkpointSnapshots).Find(&results).Error

	return
}

// GetFromCheckpointType 通过CHECKPOINT_TYPE获取内容
func (obj *_CdbCkptHistoryMgr) GetFromCheckpointType(checkpointType string) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`CHECKPOINT_TYPE` = ?", checkpointType).Find(&results).Error

	return
}

// GetBatchFromCheckpointType 批量查找
func (obj *_CdbCkptHistoryMgr) GetBatchFromCheckpointType(checkpointTypes []string) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`CHECKPOINT_TYPE` IN (?)", checkpointTypes).Find(&results).Error

	return
}

// GetFromCheckpointClusterVersion 通过CHECKPOINT_CLUSTER_VERSION获取内容
func (obj *_CdbCkptHistoryMgr) GetFromCheckpointClusterVersion(checkpointClusterVersion string) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`CHECKPOINT_CLUSTER_VERSION` = ?", checkpointClusterVersion).Find(&results).Error

	return
}

// GetBatchFromCheckpointClusterVersion 批量查找
func (obj *_CdbCkptHistoryMgr) GetBatchFromCheckpointClusterVersion(checkpointClusterVersions []string) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`CHECKPOINT_CLUSTER_VERSION` IN (?)", checkpointClusterVersions).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbCkptHistoryMgr) GetFromStartTime(startTime time.Time) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbCkptHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromFinishTime 通过FINISH_TIME获取内容
func (obj *_CdbCkptHistoryMgr) GetFromFinishTime(finishTime time.Time) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`FINISH_TIME` = ?", finishTime).Find(&results).Error

	return
}

// GetBatchFromFinishTime 批量查找
func (obj *_CdbCkptHistoryMgr) GetBatchFromFinishTime(finishTimes []time.Time) (results []*CdbCkptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbCkptHistory{}).Where("`FINISH_TIME` IN (?)", finishTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
