package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualTableStatMgr struct {
	*_BaseMgr
}

// AllVirtualTableStatMgr open func
func AllVirtualTableStatMgr(db *gorm.DB) *_AllVirtualTableStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTableStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTableStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_table_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTableStatMgr) GetTableName() string {
	return "__all_virtual_table_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTableStatMgr) Reset() *_AllVirtualTableStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTableStatMgr) Get() (result AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTableStatMgr) Gets() (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTableStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTableStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualTableStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualTableStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTableStatMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTableStatMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithObjectType object_type获取
func (obj *_AllVirtualTableStatMgr) WithObjectType(objectType int64) Option {
	return optionFunc(func(o *options) { o.query["object_type"] = objectType })
}

// WithLastAnalyzed last_analyzed获取
func (obj *_AllVirtualTableStatMgr) WithLastAnalyzed(lastAnalyzed time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_analyzed"] = lastAnalyzed })
}

// WithSstableRowCnt sstable_row_cnt获取
func (obj *_AllVirtualTableStatMgr) WithSstableRowCnt(sstableRowCnt int64) Option {
	return optionFunc(func(o *options) { o.query["sstable_row_cnt"] = sstableRowCnt })
}

// WithSstableAvgRowLen sstable_avg_row_len获取
func (obj *_AllVirtualTableStatMgr) WithSstableAvgRowLen(sstableAvgRowLen int64) Option {
	return optionFunc(func(o *options) { o.query["sstable_avg_row_len"] = sstableAvgRowLen })
}

// WithMacroBlkCnt macro_blk_cnt获取
func (obj *_AllVirtualTableStatMgr) WithMacroBlkCnt(macroBlkCnt int64) Option {
	return optionFunc(func(o *options) { o.query["macro_blk_cnt"] = macroBlkCnt })
}

// WithMicroBlkCnt micro_blk_cnt获取
func (obj *_AllVirtualTableStatMgr) WithMicroBlkCnt(microBlkCnt int64) Option {
	return optionFunc(func(o *options) { o.query["micro_blk_cnt"] = microBlkCnt })
}

// WithMemtableRowCnt memtable_row_cnt获取
func (obj *_AllVirtualTableStatMgr) WithMemtableRowCnt(memtableRowCnt int64) Option {
	return optionFunc(func(o *options) { o.query["memtable_row_cnt"] = memtableRowCnt })
}

// WithMemtableAvgRowLen memtable_avg_row_len获取
func (obj *_AllVirtualTableStatMgr) WithMemtableAvgRowLen(memtableAvgRowLen int64) Option {
	return optionFunc(func(o *options) { o.query["memtable_avg_row_len"] = memtableAvgRowLen })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTableStatMgr) GetByOption(opts ...Option) (result AllVirtualTableStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTableStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualTableStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTableStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTableStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where(options.query)
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
func (obj *_AllVirtualTableStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualTableStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualTableStatMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTableStatMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTableStatMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromObjectType 通过object_type获取内容
func (obj *_AllVirtualTableStatMgr) GetFromObjectType(objectType int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`object_type` = ?", objectType).Find(&results).Error

	return
}

// GetBatchFromObjectType 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromObjectType(objectTypes []int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`object_type` IN (?)", objectTypes).Find(&results).Error

	return
}

// GetFromLastAnalyzed 通过last_analyzed获取内容
func (obj *_AllVirtualTableStatMgr) GetFromLastAnalyzed(lastAnalyzed time.Time) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`last_analyzed` = ?", lastAnalyzed).Find(&results).Error

	return
}

// GetBatchFromLastAnalyzed 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromLastAnalyzed(lastAnalyzeds []time.Time) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`last_analyzed` IN (?)", lastAnalyzeds).Find(&results).Error

	return
}

// GetFromSstableRowCnt 通过sstable_row_cnt获取内容
func (obj *_AllVirtualTableStatMgr) GetFromSstableRowCnt(sstableRowCnt int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`sstable_row_cnt` = ?", sstableRowCnt).Find(&results).Error

	return
}

// GetBatchFromSstableRowCnt 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromSstableRowCnt(sstableRowCnts []int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`sstable_row_cnt` IN (?)", sstableRowCnts).Find(&results).Error

	return
}

// GetFromSstableAvgRowLen 通过sstable_avg_row_len获取内容
func (obj *_AllVirtualTableStatMgr) GetFromSstableAvgRowLen(sstableAvgRowLen int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`sstable_avg_row_len` = ?", sstableAvgRowLen).Find(&results).Error

	return
}

// GetBatchFromSstableAvgRowLen 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromSstableAvgRowLen(sstableAvgRowLens []int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`sstable_avg_row_len` IN (?)", sstableAvgRowLens).Find(&results).Error

	return
}

// GetFromMacroBlkCnt 通过macro_blk_cnt获取内容
func (obj *_AllVirtualTableStatMgr) GetFromMacroBlkCnt(macroBlkCnt int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`macro_blk_cnt` = ?", macroBlkCnt).Find(&results).Error

	return
}

// GetBatchFromMacroBlkCnt 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromMacroBlkCnt(macroBlkCnts []int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`macro_blk_cnt` IN (?)", macroBlkCnts).Find(&results).Error

	return
}

// GetFromMicroBlkCnt 通过micro_blk_cnt获取内容
func (obj *_AllVirtualTableStatMgr) GetFromMicroBlkCnt(microBlkCnt int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`micro_blk_cnt` = ?", microBlkCnt).Find(&results).Error

	return
}

// GetBatchFromMicroBlkCnt 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromMicroBlkCnt(microBlkCnts []int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`micro_blk_cnt` IN (?)", microBlkCnts).Find(&results).Error

	return
}

// GetFromMemtableRowCnt 通过memtable_row_cnt获取内容
func (obj *_AllVirtualTableStatMgr) GetFromMemtableRowCnt(memtableRowCnt int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`memtable_row_cnt` = ?", memtableRowCnt).Find(&results).Error

	return
}

// GetBatchFromMemtableRowCnt 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromMemtableRowCnt(memtableRowCnts []int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`memtable_row_cnt` IN (?)", memtableRowCnts).Find(&results).Error

	return
}

// GetFromMemtableAvgRowLen 通过memtable_avg_row_len获取内容
func (obj *_AllVirtualTableStatMgr) GetFromMemtableAvgRowLen(memtableAvgRowLen int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`memtable_avg_row_len` = ?", memtableAvgRowLen).Find(&results).Error

	return
}

// GetBatchFromMemtableAvgRowLen 批量查找
func (obj *_AllVirtualTableStatMgr) GetBatchFromMemtableAvgRowLen(memtableAvgRowLens []int64) (results []*AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`memtable_avg_row_len` IN (?)", memtableAvgRowLens).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTableStatMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64) (result AllVirtualTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableStat{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ?", tenantID, tableID, partitionID).First(&result).Error

	return
}
