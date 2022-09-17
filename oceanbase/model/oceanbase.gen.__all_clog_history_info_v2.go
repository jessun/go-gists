package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllClogHistoryInfoV2Mgr struct {
	*_BaseMgr
}

// AllClogHistoryInfoV2Mgr open func
func AllClogHistoryInfoV2Mgr(db *gorm.DB) *_AllClogHistoryInfoV2Mgr {
	if db == nil {
		panic(fmt.Errorf("AllClogHistoryInfoV2Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllClogHistoryInfoV2Mgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_clog_history_info_v2"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllClogHistoryInfoV2Mgr) GetTableName() string {
	return "__all_clog_history_info_v2"
}

// Reset 重置gorm会话
func (obj *_AllClogHistoryInfoV2Mgr) Reset() *_AllClogHistoryInfoV2Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllClogHistoryInfoV2Mgr) Get() (result AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllClogHistoryInfoV2Mgr) Gets() (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllClogHistoryInfoV2Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllClogHistoryInfoV2Mgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllClogHistoryInfoV2Mgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTableID table_id获取
func (obj *_AllClogHistoryInfoV2Mgr) WithTableID(tableID uint64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllClogHistoryInfoV2Mgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllClogHistoryInfoV2Mgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithStartLogID start_log_id获取
func (obj *_AllClogHistoryInfoV2Mgr) WithStartLogID(startLogID uint64) Option {
	return optionFunc(func(o *options) { o.query["start_log_id"] = startLogID })
}

// WithStartLogTimestamp start_log_timestamp获取
func (obj *_AllClogHistoryInfoV2Mgr) WithStartLogTimestamp(startLogTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["start_log_timestamp"] = startLogTimestamp })
}

// WithSvrIP svr_ip获取
func (obj *_AllClogHistoryInfoV2Mgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllClogHistoryInfoV2Mgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithEndLogID end_log_id获取
func (obj *_AllClogHistoryInfoV2Mgr) WithEndLogID(endLogID uint64) Option {
	return optionFunc(func(o *options) { o.query["end_log_id"] = endLogID })
}

// WithEndLogTimestamp end_log_timestamp获取
func (obj *_AllClogHistoryInfoV2Mgr) WithEndLogTimestamp(endLogTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["end_log_timestamp"] = endLogTimestamp })
}

// GetByOption 功能选项模式获取
func (obj *_AllClogHistoryInfoV2Mgr) GetByOption(opts ...Option) (result AllClogHistoryInfoV2, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllClogHistoryInfoV2Mgr) GetByOptions(opts ...Option) (results []*AllClogHistoryInfoV2, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllClogHistoryInfoV2Mgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllClogHistoryInfoV2, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where(options.query)
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
func (obj *_AllClogHistoryInfoV2Mgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllClogHistoryInfoV2Mgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllClogHistoryInfoV2Mgr) GetFromGmtModified(gmtModified time.Time) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllClogHistoryInfoV2Mgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllClogHistoryInfoV2Mgr) GetFromTableID(tableID uint64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllClogHistoryInfoV2Mgr) GetBatchFromTableID(tableIDs []uint64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllClogHistoryInfoV2Mgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllClogHistoryInfoV2Mgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllClogHistoryInfoV2Mgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllClogHistoryInfoV2Mgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromStartLogID 通过start_log_id获取内容
func (obj *_AllClogHistoryInfoV2Mgr) GetFromStartLogID(startLogID uint64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`start_log_id` = ?", startLogID).Find(&results).Error

	return
}

// GetBatchFromStartLogID 批量查找
func (obj *_AllClogHistoryInfoV2Mgr) GetBatchFromStartLogID(startLogIDs []uint64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`start_log_id` IN (?)", startLogIDs).Find(&results).Error

	return
}

// GetFromStartLogTimestamp 通过start_log_timestamp获取内容
func (obj *_AllClogHistoryInfoV2Mgr) GetFromStartLogTimestamp(startLogTimestamp int64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`start_log_timestamp` = ?", startLogTimestamp).Find(&results).Error

	return
}

// GetBatchFromStartLogTimestamp 批量查找
func (obj *_AllClogHistoryInfoV2Mgr) GetBatchFromStartLogTimestamp(startLogTimestamps []int64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`start_log_timestamp` IN (?)", startLogTimestamps).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllClogHistoryInfoV2Mgr) GetFromSvrIP(svrIP string) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllClogHistoryInfoV2Mgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllClogHistoryInfoV2Mgr) GetFromSvrPort(svrPort int64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllClogHistoryInfoV2Mgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromEndLogID 通过end_log_id获取内容
func (obj *_AllClogHistoryInfoV2Mgr) GetFromEndLogID(endLogID uint64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`end_log_id` = ?", endLogID).Find(&results).Error

	return
}

// GetBatchFromEndLogID 批量查找
func (obj *_AllClogHistoryInfoV2Mgr) GetBatchFromEndLogID(endLogIDs []uint64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`end_log_id` IN (?)", endLogIDs).Find(&results).Error

	return
}

// GetFromEndLogTimestamp 通过end_log_timestamp获取内容
func (obj *_AllClogHistoryInfoV2Mgr) GetFromEndLogTimestamp(endLogTimestamp int64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`end_log_timestamp` = ?", endLogTimestamp).Find(&results).Error

	return
}

// GetBatchFromEndLogTimestamp 批量查找
func (obj *_AllClogHistoryInfoV2Mgr) GetBatchFromEndLogTimestamp(endLogTimestamps []int64) (results []*AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`end_log_timestamp` IN (?)", endLogTimestamps).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllClogHistoryInfoV2Mgr) FetchByPrimaryKey(tableID uint64, partitionIDx int64, partitionCnt int64, startLogID uint64, startLogTimestamp int64, svrIP string, svrPort int64) (result AllClogHistoryInfoV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllClogHistoryInfoV2{}).Where("`table_id` = ? AND `partition_idx` = ? AND `partition_cnt` = ? AND `start_log_id` = ? AND `start_log_timestamp` = ? AND `svr_ip` = ? AND `svr_port` = ?", tableID, partitionIDx, partitionCnt, startLogID, startLogTimestamp, svrIP, svrPort).First(&result).Error

	return
}
