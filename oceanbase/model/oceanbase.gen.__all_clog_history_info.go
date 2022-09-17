package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllClogHistoryInfoMgr struct {
	*_BaseMgr
}

// AllClogHistoryInfoMgr open func
func AllClogHistoryInfoMgr(db *gorm.DB) *_AllClogHistoryInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllClogHistoryInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllClogHistoryInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_clog_history_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllClogHistoryInfoMgr) GetTableName() string {
	return "__all_clog_history_info"
}

// Reset 重置gorm会话
func (obj *_AllClogHistoryInfoMgr) Reset() *_AllClogHistoryInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllClogHistoryInfoMgr) Get() (result AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllClogHistoryInfoMgr) Gets() (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllClogHistoryInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableID table_id获取
func (obj *_AllClogHistoryInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllClogHistoryInfoMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllClogHistoryInfoMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithStartLogID start_log_id获取
func (obj *_AllClogHistoryInfoMgr) WithStartLogID(startLogID int64) Option {
	return optionFunc(func(o *options) { o.query["start_log_id"] = startLogID })
}

// WithStartLogTimestamp start_log_timestamp获取
func (obj *_AllClogHistoryInfoMgr) WithStartLogTimestamp(startLogTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["start_log_timestamp"] = startLogTimestamp })
}

// WithSvrIP svr_ip获取
func (obj *_AllClogHistoryInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllClogHistoryInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithEndLogID end_log_id获取
func (obj *_AllClogHistoryInfoMgr) WithEndLogID(endLogID int64) Option {
	return optionFunc(func(o *options) { o.query["end_log_id"] = endLogID })
}

// WithEndLogTimestamp end_log_timestamp获取
func (obj *_AllClogHistoryInfoMgr) WithEndLogTimestamp(endLogTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["end_log_timestamp"] = endLogTimestamp })
}

// GetByOption 功能选项模式获取
func (obj *_AllClogHistoryInfoMgr) GetByOption(opts ...Option) (result AllClogHistoryInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllClogHistoryInfoMgr) GetByOptions(opts ...Option) (results []*AllClogHistoryInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllClogHistoryInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllClogHistoryInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where(options.query)
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

// GetFromTableID 通过table_id获取内容
func (obj *_AllClogHistoryInfoMgr) GetFromTableID(tableID int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllClogHistoryInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllClogHistoryInfoMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllClogHistoryInfoMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllClogHistoryInfoMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllClogHistoryInfoMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromStartLogID 通过start_log_id获取内容
func (obj *_AllClogHistoryInfoMgr) GetFromStartLogID(startLogID int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`start_log_id` = ?", startLogID).Find(&results).Error

	return
}

// GetBatchFromStartLogID 批量查找
func (obj *_AllClogHistoryInfoMgr) GetBatchFromStartLogID(startLogIDs []int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`start_log_id` IN (?)", startLogIDs).Find(&results).Error

	return
}

// GetFromStartLogTimestamp 通过start_log_timestamp获取内容
func (obj *_AllClogHistoryInfoMgr) GetFromStartLogTimestamp(startLogTimestamp int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`start_log_timestamp` = ?", startLogTimestamp).Find(&results).Error

	return
}

// GetBatchFromStartLogTimestamp 批量查找
func (obj *_AllClogHistoryInfoMgr) GetBatchFromStartLogTimestamp(startLogTimestamps []int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`start_log_timestamp` IN (?)", startLogTimestamps).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllClogHistoryInfoMgr) GetFromSvrIP(svrIP string) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllClogHistoryInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllClogHistoryInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllClogHistoryInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromEndLogID 通过end_log_id获取内容
func (obj *_AllClogHistoryInfoMgr) GetFromEndLogID(endLogID int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`end_log_id` = ?", endLogID).Find(&results).Error

	return
}

// GetBatchFromEndLogID 批量查找
func (obj *_AllClogHistoryInfoMgr) GetBatchFromEndLogID(endLogIDs []int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`end_log_id` IN (?)", endLogIDs).Find(&results).Error

	return
}

// GetFromEndLogTimestamp 通过end_log_timestamp获取内容
func (obj *_AllClogHistoryInfoMgr) GetFromEndLogTimestamp(endLogTimestamp int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`end_log_timestamp` = ?", endLogTimestamp).Find(&results).Error

	return
}

// GetBatchFromEndLogTimestamp 批量查找
func (obj *_AllClogHistoryInfoMgr) GetBatchFromEndLogTimestamp(endLogTimestamps []int64) (results []*AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`end_log_timestamp` IN (?)", endLogTimestamps).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllClogHistoryInfoMgr) FetchByPrimaryKey(tableID int64, partitionIDx int64, partitionCnt int64, startLogID int64, startLogTimestamp int64, svrIP string, svrPort int64) (result AllClogHistoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfo{}).Where("`table_id` = ? AND `partition_idx` = ? AND `partition_cnt` = ? AND `start_log_id` = ? AND `start_log_timestamp` = ? AND `svr_ip` = ? AND `svr_port` = ?", tableID, partitionIDx, partitionCnt, startLogID, startLogTimestamp, svrIP, svrPort).First(&result).Error

	return
}
