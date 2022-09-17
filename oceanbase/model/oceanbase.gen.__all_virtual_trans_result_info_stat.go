package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTransResultInfoStatMgr struct {
	*_BaseMgr
}

// AllVirtualTransResultInfoStatMgr open func
func AllVirtualTransResultInfoStatMgr(db *gorm.DB) *_AllVirtualTransResultInfoStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTransResultInfoStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTransResultInfoStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_trans_result_info_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTransResultInfoStatMgr) GetTableName() string {
	return "__all_virtual_trans_result_info_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTransResultInfoStatMgr) Reset() *_AllVirtualTransResultInfoStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTransResultInfoStatMgr) Get() (result AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTransResultInfoStatMgr) Gets() (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTransResultInfoStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTransID trans_id获取
func (obj *_AllVirtualTransResultInfoStatMgr) WithTransID(transID string) Option {
	return optionFunc(func(o *options) { o.query["trans_id"] = transID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTransResultInfoStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTransResultInfoStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithPartition partition获取
func (obj *_AllVirtualTransResultInfoStatMgr) WithPartition(partition string) Option {
	return optionFunc(func(o *options) { o.query["partition"] = partition })
}

// WithState state获取
func (obj *_AllVirtualTransResultInfoStatMgr) WithState(state int64) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithCommitVersion commit_version获取
func (obj *_AllVirtualTransResultInfoStatMgr) WithCommitVersion(commitVersion int64) Option {
	return optionFunc(func(o *options) { o.query["commit_version"] = commitVersion })
}

// WithMinLogID min_log_id获取
func (obj *_AllVirtualTransResultInfoStatMgr) WithMinLogID(minLogID int64) Option {
	return optionFunc(func(o *options) { o.query["min_log_id"] = minLogID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTransResultInfoStatMgr) GetByOption(opts ...Option) (result AllVirtualTransResultInfoStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTransResultInfoStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualTransResultInfoStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTransResultInfoStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTransResultInfoStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where(options.query)
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

// GetFromTransID 通过trans_id获取内容
func (obj *_AllVirtualTransResultInfoStatMgr) GetFromTransID(transID string) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`trans_id` = ?", transID).Find(&results).Error

	return
}

// GetBatchFromTransID 批量查找
func (obj *_AllVirtualTransResultInfoStatMgr) GetBatchFromTransID(transIDs []string) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`trans_id` IN (?)", transIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualTransResultInfoStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTransResultInfoStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTransResultInfoStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTransResultInfoStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromPartition 通过partition获取内容
func (obj *_AllVirtualTransResultInfoStatMgr) GetFromPartition(partition string) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`partition` = ?", partition).Find(&results).Error

	return
}

// GetBatchFromPartition 批量查找
func (obj *_AllVirtualTransResultInfoStatMgr) GetBatchFromPartition(partitions []string) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`partition` IN (?)", partitions).Find(&results).Error

	return
}

// GetFromState 通过state获取内容
func (obj *_AllVirtualTransResultInfoStatMgr) GetFromState(state int64) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找
func (obj *_AllVirtualTransResultInfoStatMgr) GetBatchFromState(states []int64) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromCommitVersion 通过commit_version获取内容
func (obj *_AllVirtualTransResultInfoStatMgr) GetFromCommitVersion(commitVersion int64) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`commit_version` = ?", commitVersion).Find(&results).Error

	return
}

// GetBatchFromCommitVersion 批量查找
func (obj *_AllVirtualTransResultInfoStatMgr) GetBatchFromCommitVersion(commitVersions []int64) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`commit_version` IN (?)", commitVersions).Find(&results).Error

	return
}

// GetFromMinLogID 通过min_log_id获取内容
func (obj *_AllVirtualTransResultInfoStatMgr) GetFromMinLogID(minLogID int64) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`min_log_id` = ?", minLogID).Find(&results).Error

	return
}

// GetBatchFromMinLogID 批量查找
func (obj *_AllVirtualTransResultInfoStatMgr) GetBatchFromMinLogID(minLogIDs []int64) (results []*AllVirtualTransResultInfoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransResultInfoStat{}).Where("`min_log_id` IN (?)", minLogIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
