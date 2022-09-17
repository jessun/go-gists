package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualConcurrencyObjectPoolMgr struct {
	*_BaseMgr
}

// AllVirtualConcurrencyObjectPoolMgr open func
func AllVirtualConcurrencyObjectPoolMgr(db *gorm.DB) *_AllVirtualConcurrencyObjectPoolMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualConcurrencyObjectPoolMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualConcurrencyObjectPoolMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_concurrency_object_pool"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetTableName() string {
	return "__all_virtual_concurrency_object_pool"
}

// Reset 重置gorm会话
func (obj *_AllVirtualConcurrencyObjectPoolMgr) Reset() *_AllVirtualConcurrencyObjectPoolMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualConcurrencyObjectPoolMgr) Get() (result AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualConcurrencyObjectPoolMgr) Gets() (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualConcurrencyObjectPoolMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualConcurrencyObjectPoolMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualConcurrencyObjectPoolMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithFreeListName free_list_name获取
func (obj *_AllVirtualConcurrencyObjectPoolMgr) WithFreeListName(freeListName string) Option {
	return optionFunc(func(o *options) { o.query["free_list_name"] = freeListName })
}

// WithAllocated allocated获取
func (obj *_AllVirtualConcurrencyObjectPoolMgr) WithAllocated(allocated int64) Option {
	return optionFunc(func(o *options) { o.query["allocated"] = allocated })
}

// WithInUse in_use获取
func (obj *_AllVirtualConcurrencyObjectPoolMgr) WithInUse(inUse int64) Option {
	return optionFunc(func(o *options) { o.query["in_use"] = inUse })
}

// WithCount count获取
func (obj *_AllVirtualConcurrencyObjectPoolMgr) WithCount(count int64) Option {
	return optionFunc(func(o *options) { o.query["count"] = count })
}

// WithTypeSize type_size获取
func (obj *_AllVirtualConcurrencyObjectPoolMgr) WithTypeSize(typeSize int64) Option {
	return optionFunc(func(o *options) { o.query["type_size"] = typeSize })
}

// WithChunkCount chunk_count获取
func (obj *_AllVirtualConcurrencyObjectPoolMgr) WithChunkCount(chunkCount int64) Option {
	return optionFunc(func(o *options) { o.query["chunk_count"] = chunkCount })
}

// WithChunkByteSize chunk_byte_size获取
func (obj *_AllVirtualConcurrencyObjectPoolMgr) WithChunkByteSize(chunkByteSize int64) Option {
	return optionFunc(func(o *options) { o.query["chunk_byte_size"] = chunkByteSize })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetByOption(opts ...Option) (result AllVirtualConcurrencyObjectPool, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetByOptions(opts ...Option) (results []*AllVirtualConcurrencyObjectPool, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualConcurrencyObjectPoolMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualConcurrencyObjectPool, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where(options.query)
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
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromFreeListName 通过free_list_name获取内容
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetFromFreeListName(freeListName string) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`free_list_name` = ?", freeListName).Find(&results).Error

	return
}

// GetBatchFromFreeListName 批量查找
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetBatchFromFreeListName(freeListNames []string) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`free_list_name` IN (?)", freeListNames).Find(&results).Error

	return
}

// GetFromAllocated 通过allocated获取内容
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetFromAllocated(allocated int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`allocated` = ?", allocated).Find(&results).Error

	return
}

// GetBatchFromAllocated 批量查找
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetBatchFromAllocated(allocateds []int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`allocated` IN (?)", allocateds).Find(&results).Error

	return
}

// GetFromInUse 通过in_use获取内容
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetFromInUse(inUse int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`in_use` = ?", inUse).Find(&results).Error

	return
}

// GetBatchFromInUse 批量查找
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetBatchFromInUse(inUses []int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`in_use` IN (?)", inUses).Find(&results).Error

	return
}

// GetFromCount 通过count获取内容
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetFromCount(count int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`count` = ?", count).Find(&results).Error

	return
}

// GetBatchFromCount 批量查找
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetBatchFromCount(counts []int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`count` IN (?)", counts).Find(&results).Error

	return
}

// GetFromTypeSize 通过type_size获取内容
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetFromTypeSize(typeSize int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`type_size` = ?", typeSize).Find(&results).Error

	return
}

// GetBatchFromTypeSize 批量查找
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetBatchFromTypeSize(typeSizes []int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`type_size` IN (?)", typeSizes).Find(&results).Error

	return
}

// GetFromChunkCount 通过chunk_count获取内容
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetFromChunkCount(chunkCount int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`chunk_count` = ?", chunkCount).Find(&results).Error

	return
}

// GetBatchFromChunkCount 批量查找
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetBatchFromChunkCount(chunkCounts []int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`chunk_count` IN (?)", chunkCounts).Find(&results).Error

	return
}

// GetFromChunkByteSize 通过chunk_byte_size获取内容
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetFromChunkByteSize(chunkByteSize int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`chunk_byte_size` = ?", chunkByteSize).Find(&results).Error

	return
}

// GetBatchFromChunkByteSize 批量查找
func (obj *_AllVirtualConcurrencyObjectPoolMgr) GetBatchFromChunkByteSize(chunkByteSizes []int64) (results []*AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`chunk_byte_size` IN (?)", chunkByteSizes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualConcurrencyObjectPoolMgr) FetchByPrimaryKey(svrIP string, svrPort int64) (result AllVirtualConcurrencyObjectPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConcurrencyObjectPool{}).Where("`svr_ip` = ? AND `svr_port` = ?", svrIP, svrPort).First(&result).Error

	return
}
