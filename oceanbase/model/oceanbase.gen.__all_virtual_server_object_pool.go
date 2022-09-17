package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualServerObjectPoolMgr struct {
	*_BaseMgr
}

// AllVirtualServerObjectPoolMgr open func
func AllVirtualServerObjectPoolMgr(db *gorm.DB) *_AllVirtualServerObjectPoolMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualServerObjectPoolMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualServerObjectPoolMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_server_object_pool"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualServerObjectPoolMgr) GetTableName() string {
	return "__all_virtual_server_object_pool"
}

// Reset 重置gorm会话
func (obj *_AllVirtualServerObjectPoolMgr) Reset() *_AllVirtualServerObjectPoolMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualServerObjectPoolMgr) Get() (result AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualServerObjectPoolMgr) Gets() (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualServerObjectPoolMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualServerObjectPoolMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualServerObjectPoolMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithObjectType object_type获取
func (obj *_AllVirtualServerObjectPoolMgr) WithObjectType(objectType string) Option {
	return optionFunc(func(o *options) { o.query["object_type"] = objectType })
}

// WithArenaID arena_id获取
func (obj *_AllVirtualServerObjectPoolMgr) WithArenaID(arenaID int64) Option {
	return optionFunc(func(o *options) { o.query["arena_id"] = arenaID })
}

// WithLock lock获取
func (obj *_AllVirtualServerObjectPoolMgr) WithLock(lock int64) Option {
	return optionFunc(func(o *options) { o.query["lock"] = lock })
}

// WithBorrowCount borrow_count获取
func (obj *_AllVirtualServerObjectPoolMgr) WithBorrowCount(borrowCount int64) Option {
	return optionFunc(func(o *options) { o.query["borrow_count"] = borrowCount })
}

// WithReturnCount return_count获取
func (obj *_AllVirtualServerObjectPoolMgr) WithReturnCount(returnCount int64) Option {
	return optionFunc(func(o *options) { o.query["return_count"] = returnCount })
}

// WithMissCount miss_count获取
func (obj *_AllVirtualServerObjectPoolMgr) WithMissCount(missCount int64) Option {
	return optionFunc(func(o *options) { o.query["miss_count"] = missCount })
}

// WithMissReturnCount miss_return_count获取
func (obj *_AllVirtualServerObjectPoolMgr) WithMissReturnCount(missReturnCount int64) Option {
	return optionFunc(func(o *options) { o.query["miss_return_count"] = missReturnCount })
}

// WithFreeNum free_num获取
func (obj *_AllVirtualServerObjectPoolMgr) WithFreeNum(freeNum int64) Option {
	return optionFunc(func(o *options) { o.query["free_num"] = freeNum })
}

// WithLastBorrowTs last_borrow_ts获取
func (obj *_AllVirtualServerObjectPoolMgr) WithLastBorrowTs(lastBorrowTs int64) Option {
	return optionFunc(func(o *options) { o.query["last_borrow_ts"] = lastBorrowTs })
}

// WithLastReturnTs last_return_ts获取
func (obj *_AllVirtualServerObjectPoolMgr) WithLastReturnTs(lastReturnTs int64) Option {
	return optionFunc(func(o *options) { o.query["last_return_ts"] = lastReturnTs })
}

// WithLastMissTs last_miss_ts获取
func (obj *_AllVirtualServerObjectPoolMgr) WithLastMissTs(lastMissTs int64) Option {
	return optionFunc(func(o *options) { o.query["last_miss_ts"] = lastMissTs })
}

// WithLastMissReturnTs last_miss_return_ts获取
func (obj *_AllVirtualServerObjectPoolMgr) WithLastMissReturnTs(lastMissReturnTs int64) Option {
	return optionFunc(func(o *options) { o.query["last_miss_return_ts"] = lastMissReturnTs })
}

// WithNext next获取
func (obj *_AllVirtualServerObjectPoolMgr) WithNext(next int64) Option {
	return optionFunc(func(o *options) { o.query["next"] = next })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualServerObjectPoolMgr) GetByOption(opts ...Option) (result AllVirtualServerObjectPool, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualServerObjectPoolMgr) GetByOptions(opts ...Option) (results []*AllVirtualServerObjectPool, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualServerObjectPoolMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualServerObjectPool, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where(options.query)
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

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromObjectType 通过object_type获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromObjectType(objectType string) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`object_type` = ?", objectType).Find(&results).Error

	return
}

// GetBatchFromObjectType 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromObjectType(objectTypes []string) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`object_type` IN (?)", objectTypes).Find(&results).Error

	return
}

// GetFromArenaID 通过arena_id获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromArenaID(arenaID int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`arena_id` = ?", arenaID).Find(&results).Error

	return
}

// GetBatchFromArenaID 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromArenaID(arenaIDs []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`arena_id` IN (?)", arenaIDs).Find(&results).Error

	return
}

// GetFromLock 通过lock获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromLock(lock int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`lock` = ?", lock).Find(&results).Error

	return
}

// GetBatchFromLock 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromLock(locks []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`lock` IN (?)", locks).Find(&results).Error

	return
}

// GetFromBorrowCount 通过borrow_count获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromBorrowCount(borrowCount int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`borrow_count` = ?", borrowCount).Find(&results).Error

	return
}

// GetBatchFromBorrowCount 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromBorrowCount(borrowCounts []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`borrow_count` IN (?)", borrowCounts).Find(&results).Error

	return
}

// GetFromReturnCount 通过return_count获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromReturnCount(returnCount int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`return_count` = ?", returnCount).Find(&results).Error

	return
}

// GetBatchFromReturnCount 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromReturnCount(returnCounts []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`return_count` IN (?)", returnCounts).Find(&results).Error

	return
}

// GetFromMissCount 通过miss_count获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromMissCount(missCount int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`miss_count` = ?", missCount).Find(&results).Error

	return
}

// GetBatchFromMissCount 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromMissCount(missCounts []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`miss_count` IN (?)", missCounts).Find(&results).Error

	return
}

// GetFromMissReturnCount 通过miss_return_count获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromMissReturnCount(missReturnCount int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`miss_return_count` = ?", missReturnCount).Find(&results).Error

	return
}

// GetBatchFromMissReturnCount 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromMissReturnCount(missReturnCounts []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`miss_return_count` IN (?)", missReturnCounts).Find(&results).Error

	return
}

// GetFromFreeNum 通过free_num获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromFreeNum(freeNum int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`free_num` = ?", freeNum).Find(&results).Error

	return
}

// GetBatchFromFreeNum 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromFreeNum(freeNums []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`free_num` IN (?)", freeNums).Find(&results).Error

	return
}

// GetFromLastBorrowTs 通过last_borrow_ts获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromLastBorrowTs(lastBorrowTs int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`last_borrow_ts` = ?", lastBorrowTs).Find(&results).Error

	return
}

// GetBatchFromLastBorrowTs 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromLastBorrowTs(lastBorrowTss []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`last_borrow_ts` IN (?)", lastBorrowTss).Find(&results).Error

	return
}

// GetFromLastReturnTs 通过last_return_ts获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromLastReturnTs(lastReturnTs int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`last_return_ts` = ?", lastReturnTs).Find(&results).Error

	return
}

// GetBatchFromLastReturnTs 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromLastReturnTs(lastReturnTss []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`last_return_ts` IN (?)", lastReturnTss).Find(&results).Error

	return
}

// GetFromLastMissTs 通过last_miss_ts获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromLastMissTs(lastMissTs int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`last_miss_ts` = ?", lastMissTs).Find(&results).Error

	return
}

// GetBatchFromLastMissTs 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromLastMissTs(lastMissTss []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`last_miss_ts` IN (?)", lastMissTss).Find(&results).Error

	return
}

// GetFromLastMissReturnTs 通过last_miss_return_ts获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromLastMissReturnTs(lastMissReturnTs int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`last_miss_return_ts` = ?", lastMissReturnTs).Find(&results).Error

	return
}

// GetBatchFromLastMissReturnTs 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromLastMissReturnTs(lastMissReturnTss []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`last_miss_return_ts` IN (?)", lastMissReturnTss).Find(&results).Error

	return
}

// GetFromNext 通过next获取内容
func (obj *_AllVirtualServerObjectPoolMgr) GetFromNext(next int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`next` = ?", next).Find(&results).Error

	return
}

// GetBatchFromNext 批量查找
func (obj *_AllVirtualServerObjectPoolMgr) GetBatchFromNext(nexts []int64) (results []*AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`next` IN (?)", nexts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualServerObjectPoolMgr) FetchByPrimaryKey(svrIP string, svrPort int64, objectType string, arenaID int64) (result AllVirtualServerObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerObjectPool{}).Where("`svr_ip` = ? AND `svr_port` = ? AND `object_type` = ? AND `arena_id` = ?", svrIP, svrPort, objectType, arenaID).First(&result).Error

	return
}
