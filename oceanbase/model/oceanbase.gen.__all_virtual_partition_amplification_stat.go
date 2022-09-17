package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualPartitionAmplificationStatMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionAmplificationStatMgr open func
func AllVirtualPartitionAmplificationStatMgr(db *gorm.DB) *_AllVirtualPartitionAmplificationStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionAmplificationStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionAmplificationStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_amplification_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetTableName() string {
	return "__all_virtual_partition_amplification_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionAmplificationStatMgr) Reset() *_AllVirtualPartitionAmplificationStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) Get() (result AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionAmplificationStatMgr) Gets() (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionAmplificationStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithDirtyRatio1 dirty_ratio_1获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithDirtyRatio1(dirtyRatio1 int64) Option {
	return optionFunc(func(o *options) { o.query["dirty_ratio_1"] = dirtyRatio1 })
}

// WithDirtyRatio3 dirty_ratio_3获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithDirtyRatio3(dirtyRatio3 int64) Option {
	return optionFunc(func(o *options) { o.query["dirty_ratio_3"] = dirtyRatio3 })
}

// WithDirtyRatio5 dirty_ratio_5获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithDirtyRatio5(dirtyRatio5 int64) Option {
	return optionFunc(func(o *options) { o.query["dirty_ratio_5"] = dirtyRatio5 })
}

// WithDirtyRatio10 dirty_ratio_10获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithDirtyRatio10(dirtyRatio10 int64) Option {
	return optionFunc(func(o *options) { o.query["dirty_ratio_10"] = dirtyRatio10 })
}

// WithDirtyRatio15 dirty_ratio_15获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithDirtyRatio15(dirtyRatio15 int64) Option {
	return optionFunc(func(o *options) { o.query["dirty_ratio_15"] = dirtyRatio15 })
}

// WithDirtyRatio20 dirty_ratio_20获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithDirtyRatio20(dirtyRatio20 int64) Option {
	return optionFunc(func(o *options) { o.query["dirty_ratio_20"] = dirtyRatio20 })
}

// WithDirtyRatio30 dirty_ratio_30获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithDirtyRatio30(dirtyRatio30 int64) Option {
	return optionFunc(func(o *options) { o.query["dirty_ratio_30"] = dirtyRatio30 })
}

// WithDirtyRatio50 dirty_ratio_50获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithDirtyRatio50(dirtyRatio50 int64) Option {
	return optionFunc(func(o *options) { o.query["dirty_ratio_50"] = dirtyRatio50 })
}

// WithDirtyRatio75 dirty_ratio_75获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithDirtyRatio75(dirtyRatio75 int64) Option {
	return optionFunc(func(o *options) { o.query["dirty_ratio_75"] = dirtyRatio75 })
}

// WithDirtyRatio100 dirty_ratio_100获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithDirtyRatio100(dirtyRatio100 int64) Option {
	return optionFunc(func(o *options) { o.query["dirty_ratio_100"] = dirtyRatio100 })
}

// WithMacroBlockCnt macro_block_cnt获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) WithMacroBlockCnt(macroBlockCnt int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_cnt"] = macroBlockCnt })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetByOption(opts ...Option) (result AllVirtualPartitionAmplificationStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionAmplificationStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionAmplificationStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionAmplificationStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where(options.query)
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
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromDirtyRatio1 通过dirty_ratio_1获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromDirtyRatio1(dirtyRatio1 int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_1` = ?", dirtyRatio1).Find(&results).Error

	return
}

// GetBatchFromDirtyRatio1 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromDirtyRatio1(dirtyRatio1s []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_1` IN (?)", dirtyRatio1s).Find(&results).Error

	return
}

// GetFromDirtyRatio3 通过dirty_ratio_3获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromDirtyRatio3(dirtyRatio3 int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_3` = ?", dirtyRatio3).Find(&results).Error

	return
}

// GetBatchFromDirtyRatio3 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromDirtyRatio3(dirtyRatio3s []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_3` IN (?)", dirtyRatio3s).Find(&results).Error

	return
}

// GetFromDirtyRatio5 通过dirty_ratio_5获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromDirtyRatio5(dirtyRatio5 int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_5` = ?", dirtyRatio5).Find(&results).Error

	return
}

// GetBatchFromDirtyRatio5 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromDirtyRatio5(dirtyRatio5s []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_5` IN (?)", dirtyRatio5s).Find(&results).Error

	return
}

// GetFromDirtyRatio10 通过dirty_ratio_10获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromDirtyRatio10(dirtyRatio10 int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_10` = ?", dirtyRatio10).Find(&results).Error

	return
}

// GetBatchFromDirtyRatio10 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromDirtyRatio10(dirtyRatio10s []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_10` IN (?)", dirtyRatio10s).Find(&results).Error

	return
}

// GetFromDirtyRatio15 通过dirty_ratio_15获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromDirtyRatio15(dirtyRatio15 int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_15` = ?", dirtyRatio15).Find(&results).Error

	return
}

// GetBatchFromDirtyRatio15 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromDirtyRatio15(dirtyRatio15s []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_15` IN (?)", dirtyRatio15s).Find(&results).Error

	return
}

// GetFromDirtyRatio20 通过dirty_ratio_20获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromDirtyRatio20(dirtyRatio20 int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_20` = ?", dirtyRatio20).Find(&results).Error

	return
}

// GetBatchFromDirtyRatio20 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromDirtyRatio20(dirtyRatio20s []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_20` IN (?)", dirtyRatio20s).Find(&results).Error

	return
}

// GetFromDirtyRatio30 通过dirty_ratio_30获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromDirtyRatio30(dirtyRatio30 int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_30` = ?", dirtyRatio30).Find(&results).Error

	return
}

// GetBatchFromDirtyRatio30 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromDirtyRatio30(dirtyRatio30s []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_30` IN (?)", dirtyRatio30s).Find(&results).Error

	return
}

// GetFromDirtyRatio50 通过dirty_ratio_50获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromDirtyRatio50(dirtyRatio50 int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_50` = ?", dirtyRatio50).Find(&results).Error

	return
}

// GetBatchFromDirtyRatio50 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromDirtyRatio50(dirtyRatio50s []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_50` IN (?)", dirtyRatio50s).Find(&results).Error

	return
}

// GetFromDirtyRatio75 通过dirty_ratio_75获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromDirtyRatio75(dirtyRatio75 int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_75` = ?", dirtyRatio75).Find(&results).Error

	return
}

// GetBatchFromDirtyRatio75 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromDirtyRatio75(dirtyRatio75s []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_75` IN (?)", dirtyRatio75s).Find(&results).Error

	return
}

// GetFromDirtyRatio100 通过dirty_ratio_100获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromDirtyRatio100(dirtyRatio100 int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_100` = ?", dirtyRatio100).Find(&results).Error

	return
}

// GetBatchFromDirtyRatio100 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromDirtyRatio100(dirtyRatio100s []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`dirty_ratio_100` IN (?)", dirtyRatio100s).Find(&results).Error

	return
}

// GetFromMacroBlockCnt 通过macro_block_cnt获取内容
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetFromMacroBlockCnt(macroBlockCnt int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`macro_block_cnt` = ?", macroBlockCnt).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCnt 批量查找
func (obj *_AllVirtualPartitionAmplificationStatMgr) GetBatchFromMacroBlockCnt(macroBlockCnts []int64) (results []*AllVirtualPartitionAmplificationStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAmplificationStat{}).Where("`macro_block_cnt` IN (?)", macroBlockCnts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
