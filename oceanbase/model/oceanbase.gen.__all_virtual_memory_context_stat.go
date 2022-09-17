package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualMemoryContextStatMgr struct {
	*_BaseMgr
}

// AllVirtualMemoryContextStatMgr open func
func AllVirtualMemoryContextStatMgr(db *gorm.DB) *_AllVirtualMemoryContextStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualMemoryContextStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualMemoryContextStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_memory_context_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualMemoryContextStatMgr) GetTableName() string {
	return "__all_virtual_memory_context_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualMemoryContextStatMgr) Reset() *_AllVirtualMemoryContextStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualMemoryContextStatMgr) Get() (result AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualMemoryContextStatMgr) Gets() (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualMemoryContextStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualMemoryContextStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualMemoryContextStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithEntity entity获取
func (obj *_AllVirtualMemoryContextStatMgr) WithEntity(entity string) Option {
	return optionFunc(func(o *options) { o.query["entity"] = entity })
}

// WithPEntity p_entity获取
func (obj *_AllVirtualMemoryContextStatMgr) WithPEntity(pEntity string) Option {
	return optionFunc(func(o *options) { o.query["p_entity"] = pEntity })
}

// WithHold hold获取
func (obj *_AllVirtualMemoryContextStatMgr) WithHold(hold int64) Option {
	return optionFunc(func(o *options) { o.query["hold"] = hold })
}

// WithMallocHold malloc_hold获取
func (obj *_AllVirtualMemoryContextStatMgr) WithMallocHold(mallocHold int64) Option {
	return optionFunc(func(o *options) { o.query["malloc_hold"] = mallocHold })
}

// WithMallocUsed malloc_used获取
func (obj *_AllVirtualMemoryContextStatMgr) WithMallocUsed(mallocUsed int64) Option {
	return optionFunc(func(o *options) { o.query["malloc_used"] = mallocUsed })
}

// WithArenaHold arena_hold获取
func (obj *_AllVirtualMemoryContextStatMgr) WithArenaHold(arenaHold int64) Option {
	return optionFunc(func(o *options) { o.query["arena_hold"] = arenaHold })
}

// WithArenaUsed arena_used获取
func (obj *_AllVirtualMemoryContextStatMgr) WithArenaUsed(arenaUsed int64) Option {
	return optionFunc(func(o *options) { o.query["arena_used"] = arenaUsed })
}

// WithCreateTime create_time获取
func (obj *_AllVirtualMemoryContextStatMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithLocation location获取
func (obj *_AllVirtualMemoryContextStatMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualMemoryContextStatMgr) GetByOption(opts ...Option) (result AllVirtualMemoryContextStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualMemoryContextStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualMemoryContextStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualMemoryContextStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualMemoryContextStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where(options.query)
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
func (obj *_AllVirtualMemoryContextStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualMemoryContextStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualMemoryContextStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualMemoryContextStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromEntity 通过entity获取内容
func (obj *_AllVirtualMemoryContextStatMgr) GetFromEntity(entity string) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`entity` = ?", entity).Find(&results).Error

	return
}

// GetBatchFromEntity 批量查找
func (obj *_AllVirtualMemoryContextStatMgr) GetBatchFromEntity(entitys []string) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`entity` IN (?)", entitys).Find(&results).Error

	return
}

// GetFromPEntity 通过p_entity获取内容
func (obj *_AllVirtualMemoryContextStatMgr) GetFromPEntity(pEntity string) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`p_entity` = ?", pEntity).Find(&results).Error

	return
}

// GetBatchFromPEntity 批量查找
func (obj *_AllVirtualMemoryContextStatMgr) GetBatchFromPEntity(pEntitys []string) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`p_entity` IN (?)", pEntitys).Find(&results).Error

	return
}

// GetFromHold 通过hold获取内容
func (obj *_AllVirtualMemoryContextStatMgr) GetFromHold(hold int64) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`hold` = ?", hold).Find(&results).Error

	return
}

// GetBatchFromHold 批量查找
func (obj *_AllVirtualMemoryContextStatMgr) GetBatchFromHold(holds []int64) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`hold` IN (?)", holds).Find(&results).Error

	return
}

// GetFromMallocHold 通过malloc_hold获取内容
func (obj *_AllVirtualMemoryContextStatMgr) GetFromMallocHold(mallocHold int64) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`malloc_hold` = ?", mallocHold).Find(&results).Error

	return
}

// GetBatchFromMallocHold 批量查找
func (obj *_AllVirtualMemoryContextStatMgr) GetBatchFromMallocHold(mallocHolds []int64) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`malloc_hold` IN (?)", mallocHolds).Find(&results).Error

	return
}

// GetFromMallocUsed 通过malloc_used获取内容
func (obj *_AllVirtualMemoryContextStatMgr) GetFromMallocUsed(mallocUsed int64) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`malloc_used` = ?", mallocUsed).Find(&results).Error

	return
}

// GetBatchFromMallocUsed 批量查找
func (obj *_AllVirtualMemoryContextStatMgr) GetBatchFromMallocUsed(mallocUseds []int64) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`malloc_used` IN (?)", mallocUseds).Find(&results).Error

	return
}

// GetFromArenaHold 通过arena_hold获取内容
func (obj *_AllVirtualMemoryContextStatMgr) GetFromArenaHold(arenaHold int64) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`arena_hold` = ?", arenaHold).Find(&results).Error

	return
}

// GetBatchFromArenaHold 批量查找
func (obj *_AllVirtualMemoryContextStatMgr) GetBatchFromArenaHold(arenaHolds []int64) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`arena_hold` IN (?)", arenaHolds).Find(&results).Error

	return
}

// GetFromArenaUsed 通过arena_used获取内容
func (obj *_AllVirtualMemoryContextStatMgr) GetFromArenaUsed(arenaUsed int64) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`arena_used` = ?", arenaUsed).Find(&results).Error

	return
}

// GetBatchFromArenaUsed 批量查找
func (obj *_AllVirtualMemoryContextStatMgr) GetBatchFromArenaUsed(arenaUseds []int64) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`arena_used` IN (?)", arenaUseds).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_AllVirtualMemoryContextStatMgr) GetFromCreateTime(createTime time.Time) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_AllVirtualMemoryContextStatMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromLocation 通过location获取内容
func (obj *_AllVirtualMemoryContextStatMgr) GetFromLocation(location string) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`location` = ?", location).Find(&results).Error

	return
}

// GetBatchFromLocation 批量查找
func (obj *_AllVirtualMemoryContextStatMgr) GetBatchFromLocation(locations []string) (results []*AllVirtualMemoryContextStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryContextStat{}).Where("`location` IN (?)", locations).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
