package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllElectionEventHistoryMgr struct {
	*_BaseMgr
}

// AllElectionEventHistoryMgr open func
func AllElectionEventHistoryMgr(db *gorm.DB) *_AllElectionEventHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllElectionEventHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllElectionEventHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_election_event_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllElectionEventHistoryMgr) GetTableName() string {
	return "__all_election_event_history"
}

// Reset 重置gorm会话
func (obj *_AllElectionEventHistoryMgr) Reset() *_AllElectionEventHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllElectionEventHistoryMgr) Get() (result AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllElectionEventHistoryMgr) Gets() (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllElectionEventHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllElectionEventHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithTableID table_id获取
func (obj *_AllElectionEventHistoryMgr) WithTableID(tableID uint64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllElectionEventHistoryMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllElectionEventHistoryMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithSvrIP svr_ip获取
func (obj *_AllElectionEventHistoryMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllElectionEventHistoryMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithEvent event获取
func (obj *_AllElectionEventHistoryMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["event"] = event })
}

// WithLeader leader获取
func (obj *_AllElectionEventHistoryMgr) WithLeader(leader string) Option {
	return optionFunc(func(o *options) { o.query["leader"] = leader })
}

// WithInfo info获取
func (obj *_AllElectionEventHistoryMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// GetByOption 功能选项模式获取
func (obj *_AllElectionEventHistoryMgr) GetByOption(opts ...Option) (result AllElectionEventHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllElectionEventHistoryMgr) GetByOptions(opts ...Option) (results []*AllElectionEventHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllElectionEventHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllElectionEventHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where(options.query)
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
func (obj *_AllElectionEventHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllElectionEventHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllElectionEventHistoryMgr) GetFromTableID(tableID uint64) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllElectionEventHistoryMgr) GetBatchFromTableID(tableIDs []uint64) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllElectionEventHistoryMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllElectionEventHistoryMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllElectionEventHistoryMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllElectionEventHistoryMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllElectionEventHistoryMgr) GetFromSvrIP(svrIP string) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllElectionEventHistoryMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllElectionEventHistoryMgr) GetFromSvrPort(svrPort int64) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllElectionEventHistoryMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromEvent 通过event获取内容
func (obj *_AllElectionEventHistoryMgr) GetFromEvent(event string) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`event` = ?", event).Find(&results).Error

	return
}

// GetBatchFromEvent 批量查找
func (obj *_AllElectionEventHistoryMgr) GetBatchFromEvent(events []string) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`event` IN (?)", events).Find(&results).Error

	return
}

// GetFromLeader 通过leader获取内容
func (obj *_AllElectionEventHistoryMgr) GetFromLeader(leader string) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`leader` = ?", leader).Find(&results).Error

	return
}

// GetBatchFromLeader 批量查找
func (obj *_AllElectionEventHistoryMgr) GetBatchFromLeader(leaders []string) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`leader` IN (?)", leaders).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllElectionEventHistoryMgr) GetFromInfo(info string) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllElectionEventHistoryMgr) GetBatchFromInfo(infos []string) (results []*AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllElectionEventHistoryMgr) FetchByPrimaryKey(gmtCreate time.Time, tableID uint64, partitionIDx int64, partitionCnt int64, svrIP string, svrPort int64) (result AllElectionEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllElectionEventHistory{}).Where("`gmt_create` = ? AND `table_id` = ? AND `partition_idx` = ? AND `partition_cnt` = ? AND `svr_ip` = ? AND `svr_port` = ?", gmtCreate, tableID, partitionIDx, partitionCnt, svrIP, svrPort).First(&result).Error

	return
}
