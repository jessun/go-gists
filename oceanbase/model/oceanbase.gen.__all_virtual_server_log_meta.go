package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualServerLogMetaMgr struct {
	*_BaseMgr
}

// AllVirtualServerLogMetaMgr open func
func AllVirtualServerLogMetaMgr(db *gorm.DB) *_AllVirtualServerLogMetaMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualServerLogMetaMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualServerLogMetaMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_server_log_meta"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualServerLogMetaMgr) GetTableName() string {
	return "__all_virtual_server_log_meta"
}

// Reset 重置gorm会话
func (obj *_AllVirtualServerLogMetaMgr) Reset() *_AllVirtualServerLogMetaMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualServerLogMetaMgr) Get() (result AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualServerLogMetaMgr) Gets() (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualServerLogMetaMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualServerLogMetaMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualServerLogMetaMgr) WithTableID(tableID uint64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualServerLogMetaMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualServerLogMetaMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithStartLogID start_log_id获取
func (obj *_AllVirtualServerLogMetaMgr) WithStartLogID(startLogID uint64) Option {
	return optionFunc(func(o *options) { o.query["start_log_id"] = startLogID })
}

// WithStartLogTimestamp start_log_timestamp获取
func (obj *_AllVirtualServerLogMetaMgr) WithStartLogTimestamp(startLogTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["start_log_timestamp"] = startLogTimestamp })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualServerLogMetaMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualServerLogMetaMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualServerLogMetaMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualServerLogMetaMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithEndLogID end_log_id获取
func (obj *_AllVirtualServerLogMetaMgr) WithEndLogID(endLogID uint64) Option {
	return optionFunc(func(o *options) { o.query["end_log_id"] = endLogID })
}

// WithEndLogTimestamp end_log_timestamp获取
func (obj *_AllVirtualServerLogMetaMgr) WithEndLogTimestamp(endLogTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["end_log_timestamp"] = endLogTimestamp })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualServerLogMetaMgr) GetByOption(opts ...Option) (result AllVirtualServerLogMeta, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualServerLogMetaMgr) GetByOptions(opts ...Option) (results []*AllVirtualServerLogMeta, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualServerLogMetaMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualServerLogMeta, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where(options.query)
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
func (obj *_AllVirtualServerLogMetaMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualServerLogMetaMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualServerLogMetaMgr) GetFromTableID(tableID uint64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualServerLogMetaMgr) GetBatchFromTableID(tableIDs []uint64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualServerLogMetaMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualServerLogMetaMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualServerLogMetaMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualServerLogMetaMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromStartLogID 通过start_log_id获取内容
func (obj *_AllVirtualServerLogMetaMgr) GetFromStartLogID(startLogID uint64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`start_log_id` = ?", startLogID).Find(&results).Error

	return
}

// GetBatchFromStartLogID 批量查找
func (obj *_AllVirtualServerLogMetaMgr) GetBatchFromStartLogID(startLogIDs []uint64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`start_log_id` IN (?)", startLogIDs).Find(&results).Error

	return
}

// GetFromStartLogTimestamp 通过start_log_timestamp获取内容
func (obj *_AllVirtualServerLogMetaMgr) GetFromStartLogTimestamp(startLogTimestamp int64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`start_log_timestamp` = ?", startLogTimestamp).Find(&results).Error

	return
}

// GetBatchFromStartLogTimestamp 批量查找
func (obj *_AllVirtualServerLogMetaMgr) GetBatchFromStartLogTimestamp(startLogTimestamps []int64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`start_log_timestamp` IN (?)", startLogTimestamps).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualServerLogMetaMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualServerLogMetaMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualServerLogMetaMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualServerLogMetaMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualServerLogMetaMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualServerLogMetaMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualServerLogMetaMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualServerLogMetaMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromEndLogID 通过end_log_id获取内容
func (obj *_AllVirtualServerLogMetaMgr) GetFromEndLogID(endLogID uint64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`end_log_id` = ?", endLogID).Find(&results).Error

	return
}

// GetBatchFromEndLogID 批量查找
func (obj *_AllVirtualServerLogMetaMgr) GetBatchFromEndLogID(endLogIDs []uint64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`end_log_id` IN (?)", endLogIDs).Find(&results).Error

	return
}

// GetFromEndLogTimestamp 通过end_log_timestamp获取内容
func (obj *_AllVirtualServerLogMetaMgr) GetFromEndLogTimestamp(endLogTimestamp int64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`end_log_timestamp` = ?", endLogTimestamp).Find(&results).Error

	return
}

// GetBatchFromEndLogTimestamp 批量查找
func (obj *_AllVirtualServerLogMetaMgr) GetBatchFromEndLogTimestamp(endLogTimestamps []int64) (results []*AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`end_log_timestamp` IN (?)", endLogTimestamps).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualServerLogMetaMgr) FetchByPrimaryKey(tenantID int64, tableID uint64, partitionIDx int64, partitionCnt int64, startLogID uint64, startLogTimestamp int64, svrIP string, svrPort int64) (result AllVirtualServerLogMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerLogMeta{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_idx` = ? AND `partition_cnt` = ? AND `start_log_id` = ? AND `start_log_timestamp` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, tableID, partitionIDx, partitionCnt, startLogID, startLogTimestamp, svrIP, svrPort).First(&result).Error

	return
}
