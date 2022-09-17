package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllBuildIndexParamMgr struct {
	*_BaseMgr
}

// AllBuildIndexParamMgr open func
func AllBuildIndexParamMgr(db *gorm.DB) *_AllBuildIndexParamMgr {
	if db == nil {
		panic(fmt.Errorf("AllBuildIndexParamMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBuildIndexParamMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_build_index_param"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBuildIndexParamMgr) GetTableName() string {
	return "__all_build_index_param"
}

// Reset 重置gorm会话
func (obj *_AllBuildIndexParamMgr) Reset() *_AllBuildIndexParamMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBuildIndexParamMgr) Get() (result AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBuildIndexParamMgr) Gets() (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBuildIndexParamMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBuildIndexParamMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBuildIndexParamMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithJobID job_id获取
func (obj *_AllBuildIndexParamMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllBuildIndexParamMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithExecutionID execution_id获取
func (obj *_AllBuildIndexParamMgr) WithExecutionID(executionID int64) Option {
	return optionFunc(func(o *options) { o.query["execution_id"] = executionID })
}

// WithSeqNo seq_no获取
func (obj *_AllBuildIndexParamMgr) WithSeqNo(seqNo int64) Option {
	return optionFunc(func(o *options) { o.query["seq_no"] = seqNo })
}

// WithParam param获取
func (obj *_AllBuildIndexParamMgr) WithParam(param string) Option {
	return optionFunc(func(o *options) { o.query["param"] = param })
}

// GetByOption 功能选项模式获取
func (obj *_AllBuildIndexParamMgr) GetByOption(opts ...Option) (result AllBuildIndexParam, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBuildIndexParamMgr) GetByOptions(opts ...Option) (results []*AllBuildIndexParam, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBuildIndexParamMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBuildIndexParam, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where(options.query)
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
func (obj *_AllBuildIndexParamMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBuildIndexParamMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBuildIndexParamMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBuildIndexParamMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllBuildIndexParamMgr) GetFromJobID(jobID int64) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllBuildIndexParamMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllBuildIndexParamMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllBuildIndexParamMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromExecutionID 通过execution_id获取内容
func (obj *_AllBuildIndexParamMgr) GetFromExecutionID(executionID int64) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`execution_id` = ?", executionID).Find(&results).Error

	return
}

// GetBatchFromExecutionID 批量查找
func (obj *_AllBuildIndexParamMgr) GetBatchFromExecutionID(executionIDs []int64) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`execution_id` IN (?)", executionIDs).Find(&results).Error

	return
}

// GetFromSeqNo 通过seq_no获取内容
func (obj *_AllBuildIndexParamMgr) GetFromSeqNo(seqNo int64) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`seq_no` = ?", seqNo).Find(&results).Error

	return
}

// GetBatchFromSeqNo 批量查找
func (obj *_AllBuildIndexParamMgr) GetBatchFromSeqNo(seqNos []int64) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`seq_no` IN (?)", seqNos).Find(&results).Error

	return
}

// GetFromParam 通过param获取内容
func (obj *_AllBuildIndexParamMgr) GetFromParam(param string) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`param` = ?", param).Find(&results).Error

	return
}

// GetBatchFromParam 批量查找
func (obj *_AllBuildIndexParamMgr) GetBatchFromParam(params []string) (results []*AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`param` IN (?)", params).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBuildIndexParamMgr) FetchByPrimaryKey(jobID int64, snapshotVersion int64, executionID int64, seqNo int64) (result AllBuildIndexParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBuildIndexParam{}).Where("`job_id` = ? AND `snapshot_version` = ? AND `execution_id` = ? AND `seq_no` = ?", jobID, snapshotVersion, executionID, seqNo).First(&result).Error

	return
}
