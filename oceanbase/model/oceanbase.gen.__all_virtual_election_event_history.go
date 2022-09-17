package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualElectionEventHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualElectionEventHistoryMgr open func
func AllVirtualElectionEventHistoryMgr(db *gorm.DB) *_AllVirtualElectionEventHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualElectionEventHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualElectionEventHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_election_event_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualElectionEventHistoryMgr) GetTableName() string {
	return "__all_virtual_election_event_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualElectionEventHistoryMgr) Reset() *_AllVirtualElectionEventHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualElectionEventHistoryMgr) Get() (result AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualElectionEventHistoryMgr) Gets() (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualElectionEventHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualElectionEventHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualElectionEventHistoryMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualElectionEventHistoryMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTableID table_id获取
func (obj *_AllVirtualElectionEventHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualElectionEventHistoryMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithEvent event获取
func (obj *_AllVirtualElectionEventHistoryMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["event"] = event })
}

// WithLeader leader获取
func (obj *_AllVirtualElectionEventHistoryMgr) WithLeader(leader string) Option {
	return optionFunc(func(o *options) { o.query["leader"] = leader })
}

// WithInfo info获取
func (obj *_AllVirtualElectionEventHistoryMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualElectionEventHistoryMgr) GetByOption(opts ...Option) (result AllVirtualElectionEventHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualElectionEventHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualElectionEventHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualElectionEventHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualElectionEventHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where(options.query)
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
func (obj *_AllVirtualElectionEventHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualElectionEventHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualElectionEventHistoryMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualElectionEventHistoryMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualElectionEventHistoryMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualElectionEventHistoryMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualElectionEventHistoryMgr) GetFromTableID(tableID int64) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualElectionEventHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualElectionEventHistoryMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualElectionEventHistoryMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromEvent 通过event获取内容
func (obj *_AllVirtualElectionEventHistoryMgr) GetFromEvent(event string) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`event` = ?", event).Find(&results).Error

	return
}

// GetBatchFromEvent 批量查找
func (obj *_AllVirtualElectionEventHistoryMgr) GetBatchFromEvent(events []string) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`event` IN (?)", events).Find(&results).Error

	return
}

// GetFromLeader 通过leader获取内容
func (obj *_AllVirtualElectionEventHistoryMgr) GetFromLeader(leader string) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`leader` = ?", leader).Find(&results).Error

	return
}

// GetBatchFromLeader 批量查找
func (obj *_AllVirtualElectionEventHistoryMgr) GetBatchFromLeader(leaders []string) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`leader` IN (?)", leaders).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllVirtualElectionEventHistoryMgr) GetFromInfo(info string) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllVirtualElectionEventHistoryMgr) GetBatchFromInfo(infos []string) (results []*AllVirtualElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionEventHistory{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
