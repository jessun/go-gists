package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTableStatMgr struct {
	*_BaseMgr
}

// AllTableStatMgr open func
func AllTableStatMgr(db *gorm.DB) *_AllTableStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllTableStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTableStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_table_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTableStatMgr) GetTableName() string {
	return "__all_table_stat"
}

// Reset 重置gorm会话
func (obj *_AllTableStatMgr) Reset() *_AllTableStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTableStatMgr) Get() (result AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTableStatMgr) Gets() (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTableStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTableStatMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTableStatMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTableStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllTableStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllTableStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithObjectType object_type获取
func (obj *_AllTableStatMgr) WithObjectType(objectType int64) Option {
	return optionFunc(func(o *options) { o.query["object_type"] = objectType })
}

// WithLastAnalyzed last_analyzed获取
func (obj *_AllTableStatMgr) WithLastAnalyzed(lastAnalyzed time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_analyzed"] = lastAnalyzed })
}

// WithSstableRowCnt sstable_row_cnt获取
func (obj *_AllTableStatMgr) WithSstableRowCnt(sstableRowCnt int64) Option {
	return optionFunc(func(o *options) { o.query["sstable_row_cnt"] = sstableRowCnt })
}

// WithSstableAvgRowLen sstable_avg_row_len获取
func (obj *_AllTableStatMgr) WithSstableAvgRowLen(sstableAvgRowLen int64) Option {
	return optionFunc(func(o *options) { o.query["sstable_avg_row_len"] = sstableAvgRowLen })
}

// WithMacroBlkCnt macro_blk_cnt获取
func (obj *_AllTableStatMgr) WithMacroBlkCnt(macroBlkCnt int64) Option {
	return optionFunc(func(o *options) { o.query["macro_blk_cnt"] = macroBlkCnt })
}

// WithMicroBlkCnt micro_blk_cnt获取
func (obj *_AllTableStatMgr) WithMicroBlkCnt(microBlkCnt int64) Option {
	return optionFunc(func(o *options) { o.query["micro_blk_cnt"] = microBlkCnt })
}

// WithMemtableRowCnt memtable_row_cnt获取
func (obj *_AllTableStatMgr) WithMemtableRowCnt(memtableRowCnt int64) Option {
	return optionFunc(func(o *options) { o.query["memtable_row_cnt"] = memtableRowCnt })
}

// WithMemtableAvgRowLen memtable_avg_row_len获取
func (obj *_AllTableStatMgr) WithMemtableAvgRowLen(memtableAvgRowLen int64) Option {
	return optionFunc(func(o *options) { o.query["memtable_avg_row_len"] = memtableAvgRowLen })
}

// GetByOption 功能选项模式获取
func (obj *_AllTableStatMgr) GetByOption(opts ...Option) (result AllTableStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTableStatMgr) GetByOptions(opts ...Option) (results []*AllTableStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTableStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTableStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where(options.query)
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
func (obj *_AllTableStatMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTableStatMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTableStatMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTableStatMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTableStatMgr) GetFromTenantID(tenantID int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTableStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllTableStatMgr) GetFromTableID(tableID int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllTableStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllTableStatMgr) GetFromPartitionID(partitionID int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllTableStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromObjectType 通过object_type获取内容
func (obj *_AllTableStatMgr) GetFromObjectType(objectType int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`object_type` = ?", objectType).Find(&results).Error

	return
}

// GetBatchFromObjectType 批量查找
func (obj *_AllTableStatMgr) GetBatchFromObjectType(objectTypes []int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`object_type` IN (?)", objectTypes).Find(&results).Error

	return
}

// GetFromLastAnalyzed 通过last_analyzed获取内容
func (obj *_AllTableStatMgr) GetFromLastAnalyzed(lastAnalyzed time.Time) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`last_analyzed` = ?", lastAnalyzed).Find(&results).Error

	return
}

// GetBatchFromLastAnalyzed 批量查找
func (obj *_AllTableStatMgr) GetBatchFromLastAnalyzed(lastAnalyzeds []time.Time) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`last_analyzed` IN (?)", lastAnalyzeds).Find(&results).Error

	return
}

// GetFromSstableRowCnt 通过sstable_row_cnt获取内容
func (obj *_AllTableStatMgr) GetFromSstableRowCnt(sstableRowCnt int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`sstable_row_cnt` = ?", sstableRowCnt).Find(&results).Error

	return
}

// GetBatchFromSstableRowCnt 批量查找
func (obj *_AllTableStatMgr) GetBatchFromSstableRowCnt(sstableRowCnts []int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`sstable_row_cnt` IN (?)", sstableRowCnts).Find(&results).Error

	return
}

// GetFromSstableAvgRowLen 通过sstable_avg_row_len获取内容
func (obj *_AllTableStatMgr) GetFromSstableAvgRowLen(sstableAvgRowLen int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`sstable_avg_row_len` = ?", sstableAvgRowLen).Find(&results).Error

	return
}

// GetBatchFromSstableAvgRowLen 批量查找
func (obj *_AllTableStatMgr) GetBatchFromSstableAvgRowLen(sstableAvgRowLens []int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`sstable_avg_row_len` IN (?)", sstableAvgRowLens).Find(&results).Error

	return
}

// GetFromMacroBlkCnt 通过macro_blk_cnt获取内容
func (obj *_AllTableStatMgr) GetFromMacroBlkCnt(macroBlkCnt int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`macro_blk_cnt` = ?", macroBlkCnt).Find(&results).Error

	return
}

// GetBatchFromMacroBlkCnt 批量查找
func (obj *_AllTableStatMgr) GetBatchFromMacroBlkCnt(macroBlkCnts []int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`macro_blk_cnt` IN (?)", macroBlkCnts).Find(&results).Error

	return
}

// GetFromMicroBlkCnt 通过micro_blk_cnt获取内容
func (obj *_AllTableStatMgr) GetFromMicroBlkCnt(microBlkCnt int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`micro_blk_cnt` = ?", microBlkCnt).Find(&results).Error

	return
}

// GetBatchFromMicroBlkCnt 批量查找
func (obj *_AllTableStatMgr) GetBatchFromMicroBlkCnt(microBlkCnts []int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`micro_blk_cnt` IN (?)", microBlkCnts).Find(&results).Error

	return
}

// GetFromMemtableRowCnt 通过memtable_row_cnt获取内容
func (obj *_AllTableStatMgr) GetFromMemtableRowCnt(memtableRowCnt int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`memtable_row_cnt` = ?", memtableRowCnt).Find(&results).Error

	return
}

// GetBatchFromMemtableRowCnt 批量查找
func (obj *_AllTableStatMgr) GetBatchFromMemtableRowCnt(memtableRowCnts []int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`memtable_row_cnt` IN (?)", memtableRowCnts).Find(&results).Error

	return
}

// GetFromMemtableAvgRowLen 通过memtable_avg_row_len获取内容
func (obj *_AllTableStatMgr) GetFromMemtableAvgRowLen(memtableAvgRowLen int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`memtable_avg_row_len` = ?", memtableAvgRowLen).Find(&results).Error

	return
}

// GetBatchFromMemtableAvgRowLen 批量查找
func (obj *_AllTableStatMgr) GetBatchFromMemtableAvgRowLen(memtableAvgRowLens []int64) (results []*AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`memtable_avg_row_len` IN (?)", memtableAvgRowLens).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTableStatMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64) (result AllTableStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableStat{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ?", tenantID, tableID, partitionID).First(&result).Error

	return
}
