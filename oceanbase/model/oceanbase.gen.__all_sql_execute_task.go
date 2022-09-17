package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllSQLExecuteTaskMgr struct {
	*_BaseMgr
}

// AllSQLExecuteTaskMgr open func
func AllSQLExecuteTaskMgr(db *gorm.DB) *_AllSQLExecuteTaskMgr {
	if db == nil {
		panic(fmt.Errorf("AllSQLExecuteTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSQLExecuteTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_sql_execute_task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSQLExecuteTaskMgr) GetTableName() string {
	return "__all_sql_execute_task"
}

// Reset 重置gorm会话
func (obj *_AllSQLExecuteTaskMgr) Reset() *_AllSQLExecuteTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSQLExecuteTaskMgr) Get() (result AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSQLExecuteTaskMgr) Gets() (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSQLExecuteTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSQLExecuteTaskMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSQLExecuteTaskMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithJobID job_id获取
func (obj *_AllSQLExecuteTaskMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithExecutionID execution_id获取
func (obj *_AllSQLExecuteTaskMgr) WithExecutionID(executionID int64) Option {
	return optionFunc(func(o *options) { o.query["execution_id"] = executionID })
}

// WithSQLJobID sql_job_id获取
func (obj *_AllSQLExecuteTaskMgr) WithSQLJobID(sqlJobID int64) Option {
	return optionFunc(func(o *options) { o.query["sql_job_id"] = sqlJobID })
}

// WithTaskID task_id获取
func (obj *_AllSQLExecuteTaskMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithSvrIP svr_ip获取
func (obj *_AllSQLExecuteTaskMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllSQLExecuteTaskMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithSliceCount slice_count获取
func (obj *_AllSQLExecuteTaskMgr) WithSliceCount(sliceCount int64) Option {
	return optionFunc(func(o *options) { o.query["slice_count"] = sliceCount })
}

// WithTaskStat task_stat获取
func (obj *_AllSQLExecuteTaskMgr) WithTaskStat(taskStat string) Option {
	return optionFunc(func(o *options) { o.query["task_stat"] = taskStat })
}

// WithTaskResult task_result获取
func (obj *_AllSQLExecuteTaskMgr) WithTaskResult(taskResult int64) Option {
	return optionFunc(func(o *options) { o.query["task_result"] = taskResult })
}

// WithTaskInfo task_info获取
func (obj *_AllSQLExecuteTaskMgr) WithTaskInfo(taskInfo string) Option {
	return optionFunc(func(o *options) { o.query["task_info"] = taskInfo })
}

// GetByOption 功能选项模式获取
func (obj *_AllSQLExecuteTaskMgr) GetByOption(opts ...Option) (result AllSQLExecuteTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSQLExecuteTaskMgr) GetByOptions(opts ...Option) (results []*AllSQLExecuteTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSQLExecuteTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSQLExecuteTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where(options.query)
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
func (obj *_AllSQLExecuteTaskMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSQLExecuteTaskMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSQLExecuteTaskMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSQLExecuteTaskMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllSQLExecuteTaskMgr) GetFromJobID(jobID int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllSQLExecuteTaskMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromExecutionID 通过execution_id获取内容
func (obj *_AllSQLExecuteTaskMgr) GetFromExecutionID(executionID int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`execution_id` = ?", executionID).Find(&results).Error

	return
}

// GetBatchFromExecutionID 批量查找
func (obj *_AllSQLExecuteTaskMgr) GetBatchFromExecutionID(executionIDs []int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`execution_id` IN (?)", executionIDs).Find(&results).Error

	return
}

// GetFromSQLJobID 通过sql_job_id获取内容
func (obj *_AllSQLExecuteTaskMgr) GetFromSQLJobID(sqlJobID int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`sql_job_id` = ?", sqlJobID).Find(&results).Error

	return
}

// GetBatchFromSQLJobID 批量查找
func (obj *_AllSQLExecuteTaskMgr) GetBatchFromSQLJobID(sqlJobIDs []int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`sql_job_id` IN (?)", sqlJobIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容
func (obj *_AllSQLExecuteTaskMgr) GetFromTaskID(taskID int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_AllSQLExecuteTaskMgr) GetBatchFromTaskID(taskIDs []int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllSQLExecuteTaskMgr) GetFromSvrIP(svrIP string) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllSQLExecuteTaskMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllSQLExecuteTaskMgr) GetFromSvrPort(svrPort int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllSQLExecuteTaskMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromSliceCount 通过slice_count获取内容
func (obj *_AllSQLExecuteTaskMgr) GetFromSliceCount(sliceCount int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`slice_count` = ?", sliceCount).Find(&results).Error

	return
}

// GetBatchFromSliceCount 批量查找
func (obj *_AllSQLExecuteTaskMgr) GetBatchFromSliceCount(sliceCounts []int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`slice_count` IN (?)", sliceCounts).Find(&results).Error

	return
}

// GetFromTaskStat 通过task_stat获取内容
func (obj *_AllSQLExecuteTaskMgr) GetFromTaskStat(taskStat string) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`task_stat` = ?", taskStat).Find(&results).Error

	return
}

// GetBatchFromTaskStat 批量查找
func (obj *_AllSQLExecuteTaskMgr) GetBatchFromTaskStat(taskStats []string) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`task_stat` IN (?)", taskStats).Find(&results).Error

	return
}

// GetFromTaskResult 通过task_result获取内容
func (obj *_AllSQLExecuteTaskMgr) GetFromTaskResult(taskResult int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`task_result` = ?", taskResult).Find(&results).Error

	return
}

// GetBatchFromTaskResult 批量查找
func (obj *_AllSQLExecuteTaskMgr) GetBatchFromTaskResult(taskResults []int64) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`task_result` IN (?)", taskResults).Find(&results).Error

	return
}

// GetFromTaskInfo 通过task_info获取内容
func (obj *_AllSQLExecuteTaskMgr) GetFromTaskInfo(taskInfo string) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`task_info` = ?", taskInfo).Find(&results).Error

	return
}

// GetBatchFromTaskInfo 批量查找
func (obj *_AllSQLExecuteTaskMgr) GetBatchFromTaskInfo(taskInfos []string) (results []*AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`task_info` IN (?)", taskInfos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSQLExecuteTaskMgr) FetchByPrimaryKey(jobID int64, executionID int64, sqlJobID int64, taskID int64, svrIP string, svrPort int64) (result AllSQLExecuteTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSQLExecuteTask{}).Where("`job_id` = ? AND `execution_id` = ? AND `sql_job_id` = ? AND `task_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", jobID, executionID, sqlJobID, taskID, svrIP, svrPort).First(&result).Error

	return
}
