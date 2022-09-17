package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualKvcacheInfoMgr struct {
	*_BaseMgr
}

// AllVirtualKvcacheInfoMgr open func
func AllVirtualKvcacheInfoMgr(db *gorm.DB) *_AllVirtualKvcacheInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualKvcacheInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualKvcacheInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_kvcache_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualKvcacheInfoMgr) GetTableName() string {
	return "__all_virtual_kvcache_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualKvcacheInfoMgr) Reset() *_AllVirtualKvcacheInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualKvcacheInfoMgr) Get() (result AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualKvcacheInfoMgr) Gets() (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualKvcacheInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualKvcacheInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualKvcacheInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualKvcacheInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithCacheName cache_name获取
func (obj *_AllVirtualKvcacheInfoMgr) WithCacheName(cacheName string) Option {
	return optionFunc(func(o *options) { o.query["cache_name"] = cacheName })
}

// WithCacheID cache_id获取
func (obj *_AllVirtualKvcacheInfoMgr) WithCacheID(cacheID int64) Option {
	return optionFunc(func(o *options) { o.query["cache_id"] = cacheID })
}

// WithPriority priority获取
func (obj *_AllVirtualKvcacheInfoMgr) WithPriority(priority int64) Option {
	return optionFunc(func(o *options) { o.query["priority"] = priority })
}

// WithCacheSize cache_size获取
func (obj *_AllVirtualKvcacheInfoMgr) WithCacheSize(cacheSize int64) Option {
	return optionFunc(func(o *options) { o.query["cache_size"] = cacheSize })
}

// WithCacheStoreSize cache_store_size获取
func (obj *_AllVirtualKvcacheInfoMgr) WithCacheStoreSize(cacheStoreSize int64) Option {
	return optionFunc(func(o *options) { o.query["cache_store_size"] = cacheStoreSize })
}

// WithCacheMapSize cache_map_size获取
func (obj *_AllVirtualKvcacheInfoMgr) WithCacheMapSize(cacheMapSize int64) Option {
	return optionFunc(func(o *options) { o.query["cache_map_size"] = cacheMapSize })
}

// WithKvCnt kv_cnt获取
func (obj *_AllVirtualKvcacheInfoMgr) WithKvCnt(kvCnt int64) Option {
	return optionFunc(func(o *options) { o.query["kv_cnt"] = kvCnt })
}

// WithHitRatio hit_ratio获取
func (obj *_AllVirtualKvcacheInfoMgr) WithHitRatio(hitRatio float64) Option {
	return optionFunc(func(o *options) { o.query["hit_ratio"] = hitRatio })
}

// WithTotalPutCnt total_put_cnt获取
func (obj *_AllVirtualKvcacheInfoMgr) WithTotalPutCnt(totalPutCnt int64) Option {
	return optionFunc(func(o *options) { o.query["total_put_cnt"] = totalPutCnt })
}

// WithTotalHitCnt total_hit_cnt获取
func (obj *_AllVirtualKvcacheInfoMgr) WithTotalHitCnt(totalHitCnt int64) Option {
	return optionFunc(func(o *options) { o.query["total_hit_cnt"] = totalHitCnt })
}

// WithTotalMissCnt total_miss_cnt获取
func (obj *_AllVirtualKvcacheInfoMgr) WithTotalMissCnt(totalMissCnt int64) Option {
	return optionFunc(func(o *options) { o.query["total_miss_cnt"] = totalMissCnt })
}

// WithHoldSize hold_size获取
func (obj *_AllVirtualKvcacheInfoMgr) WithHoldSize(holdSize int64) Option {
	return optionFunc(func(o *options) { o.query["hold_size"] = holdSize })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualKvcacheInfoMgr) GetByOption(opts ...Option) (result AllVirtualKvcacheInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualKvcacheInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualKvcacheInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualKvcacheInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualKvcacheInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where(options.query)
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
func (obj *_AllVirtualKvcacheInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromCacheName 通过cache_name获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromCacheName(cacheName string) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`cache_name` = ?", cacheName).Find(&results).Error

	return
}

// GetBatchFromCacheName 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromCacheName(cacheNames []string) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`cache_name` IN (?)", cacheNames).Find(&results).Error

	return
}

// GetFromCacheID 通过cache_id获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromCacheID(cacheID int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`cache_id` = ?", cacheID).Find(&results).Error

	return
}

// GetBatchFromCacheID 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromCacheID(cacheIDs []int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`cache_id` IN (?)", cacheIDs).Find(&results).Error

	return
}

// GetFromPriority 通过priority获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromPriority(priority int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`priority` = ?", priority).Find(&results).Error

	return
}

// GetBatchFromPriority 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromPriority(prioritys []int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`priority` IN (?)", prioritys).Find(&results).Error

	return
}

// GetFromCacheSize 通过cache_size获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromCacheSize(cacheSize int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`cache_size` = ?", cacheSize).Find(&results).Error

	return
}

// GetBatchFromCacheSize 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromCacheSize(cacheSizes []int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`cache_size` IN (?)", cacheSizes).Find(&results).Error

	return
}

// GetFromCacheStoreSize 通过cache_store_size获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromCacheStoreSize(cacheStoreSize int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`cache_store_size` = ?", cacheStoreSize).Find(&results).Error

	return
}

// GetBatchFromCacheStoreSize 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromCacheStoreSize(cacheStoreSizes []int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`cache_store_size` IN (?)", cacheStoreSizes).Find(&results).Error

	return
}

// GetFromCacheMapSize 通过cache_map_size获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromCacheMapSize(cacheMapSize int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`cache_map_size` = ?", cacheMapSize).Find(&results).Error

	return
}

// GetBatchFromCacheMapSize 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromCacheMapSize(cacheMapSizes []int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`cache_map_size` IN (?)", cacheMapSizes).Find(&results).Error

	return
}

// GetFromKvCnt 通过kv_cnt获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromKvCnt(kvCnt int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`kv_cnt` = ?", kvCnt).Find(&results).Error

	return
}

// GetBatchFromKvCnt 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromKvCnt(kvCnts []int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`kv_cnt` IN (?)", kvCnts).Find(&results).Error

	return
}

// GetFromHitRatio 通过hit_ratio获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromHitRatio(hitRatio float64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`hit_ratio` = ?", hitRatio).Find(&results).Error

	return
}

// GetBatchFromHitRatio 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromHitRatio(hitRatios []float64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`hit_ratio` IN (?)", hitRatios).Find(&results).Error

	return
}

// GetFromTotalPutCnt 通过total_put_cnt获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromTotalPutCnt(totalPutCnt int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`total_put_cnt` = ?", totalPutCnt).Find(&results).Error

	return
}

// GetBatchFromTotalPutCnt 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromTotalPutCnt(totalPutCnts []int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`total_put_cnt` IN (?)", totalPutCnts).Find(&results).Error

	return
}

// GetFromTotalHitCnt 通过total_hit_cnt获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromTotalHitCnt(totalHitCnt int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`total_hit_cnt` = ?", totalHitCnt).Find(&results).Error

	return
}

// GetBatchFromTotalHitCnt 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromTotalHitCnt(totalHitCnts []int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`total_hit_cnt` IN (?)", totalHitCnts).Find(&results).Error

	return
}

// GetFromTotalMissCnt 通过total_miss_cnt获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromTotalMissCnt(totalMissCnt int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`total_miss_cnt` = ?", totalMissCnt).Find(&results).Error

	return
}

// GetBatchFromTotalMissCnt 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromTotalMissCnt(totalMissCnts []int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`total_miss_cnt` IN (?)", totalMissCnts).Find(&results).Error

	return
}

// GetFromHoldSize 通过hold_size获取内容
func (obj *_AllVirtualKvcacheInfoMgr) GetFromHoldSize(holdSize int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`hold_size` = ?", holdSize).Find(&results).Error

	return
}

// GetBatchFromHoldSize 批量查找
func (obj *_AllVirtualKvcacheInfoMgr) GetBatchFromHoldSize(holdSizes []int64) (results []*AllVirtualKvcacheInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvcacheInfo{}).Where("`hold_size` IN (?)", holdSizes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
