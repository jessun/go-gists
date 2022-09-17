package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualIoStatMgr struct {
	*_BaseMgr
}

// AllVirtualIoStatMgr open func
func AllVirtualIoStatMgr(db *gorm.DB) *_AllVirtualIoStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualIoStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualIoStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_io_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualIoStatMgr) GetTableName() string {
	return "__all_virtual_io_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualIoStatMgr) Reset() *_AllVirtualIoStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualIoStatMgr) Get() (result AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualIoStatMgr) Gets() (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualIoStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualIoStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualIoStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithFd fd获取
func (obj *_AllVirtualIoStatMgr) WithFd(fd int64) Option {
	return optionFunc(func(o *options) { o.query["fd"] = fd })
}

// WithDiskType disk_type获取
func (obj *_AllVirtualIoStatMgr) WithDiskType(diskType string) Option {
	return optionFunc(func(o *options) { o.query["disk_type"] = diskType })
}

// WithSysIoUpLimitInMb sys_io_up_limit_in_mb获取
func (obj *_AllVirtualIoStatMgr) WithSysIoUpLimitInMb(sysIoUpLimitInMb int64) Option {
	return optionFunc(func(o *options) { o.query["sys_io_up_limit_in_mb"] = sysIoUpLimitInMb })
}

// WithSysIoBandwidthInMb sys_io_bandwidth_in_mb获取
func (obj *_AllVirtualIoStatMgr) WithSysIoBandwidthInMb(sysIoBandwidthInMb int64) Option {
	return optionFunc(func(o *options) { o.query["sys_io_bandwidth_in_mb"] = sysIoBandwidthInMb })
}

// WithSysIoLowWatermarkInMb sys_io_low_watermark_in_mb获取
func (obj *_AllVirtualIoStatMgr) WithSysIoLowWatermarkInMb(sysIoLowWatermarkInMb int64) Option {
	return optionFunc(func(o *options) { o.query["sys_io_low_watermark_in_mb"] = sysIoLowWatermarkInMb })
}

// WithSysIoHighWatermarkInMb sys_io_high_watermark_in_mb获取
func (obj *_AllVirtualIoStatMgr) WithSysIoHighWatermarkInMb(sysIoHighWatermarkInMb int64) Option {
	return optionFunc(func(o *options) { o.query["sys_io_high_watermark_in_mb"] = sysIoHighWatermarkInMb })
}

// WithIoBenchResult io_bench_result获取
func (obj *_AllVirtualIoStatMgr) WithIoBenchResult(ioBenchResult string) Option {
	return optionFunc(func(o *options) { o.query["io_bench_result"] = ioBenchResult })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualIoStatMgr) GetByOption(opts ...Option) (result AllVirtualIoStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualIoStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualIoStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualIoStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualIoStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where(options.query)
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
func (obj *_AllVirtualIoStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualIoStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualIoStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualIoStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromFd 通过fd获取内容
func (obj *_AllVirtualIoStatMgr) GetFromFd(fd int64) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`fd` = ?", fd).Find(&results).Error

	return
}

// GetBatchFromFd 批量查找
func (obj *_AllVirtualIoStatMgr) GetBatchFromFd(fds []int64) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`fd` IN (?)", fds).Find(&results).Error

	return
}

// GetFromDiskType 通过disk_type获取内容
func (obj *_AllVirtualIoStatMgr) GetFromDiskType(diskType string) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`disk_type` = ?", diskType).Find(&results).Error

	return
}

// GetBatchFromDiskType 批量查找
func (obj *_AllVirtualIoStatMgr) GetBatchFromDiskType(diskTypes []string) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`disk_type` IN (?)", diskTypes).Find(&results).Error

	return
}

// GetFromSysIoUpLimitInMb 通过sys_io_up_limit_in_mb获取内容
func (obj *_AllVirtualIoStatMgr) GetFromSysIoUpLimitInMb(sysIoUpLimitInMb int64) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`sys_io_up_limit_in_mb` = ?", sysIoUpLimitInMb).Find(&results).Error

	return
}

// GetBatchFromSysIoUpLimitInMb 批量查找
func (obj *_AllVirtualIoStatMgr) GetBatchFromSysIoUpLimitInMb(sysIoUpLimitInMbs []int64) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`sys_io_up_limit_in_mb` IN (?)", sysIoUpLimitInMbs).Find(&results).Error

	return
}

// GetFromSysIoBandwidthInMb 通过sys_io_bandwidth_in_mb获取内容
func (obj *_AllVirtualIoStatMgr) GetFromSysIoBandwidthInMb(sysIoBandwidthInMb int64) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`sys_io_bandwidth_in_mb` = ?", sysIoBandwidthInMb).Find(&results).Error

	return
}

// GetBatchFromSysIoBandwidthInMb 批量查找
func (obj *_AllVirtualIoStatMgr) GetBatchFromSysIoBandwidthInMb(sysIoBandwidthInMbs []int64) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`sys_io_bandwidth_in_mb` IN (?)", sysIoBandwidthInMbs).Find(&results).Error

	return
}

// GetFromSysIoLowWatermarkInMb 通过sys_io_low_watermark_in_mb获取内容
func (obj *_AllVirtualIoStatMgr) GetFromSysIoLowWatermarkInMb(sysIoLowWatermarkInMb int64) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`sys_io_low_watermark_in_mb` = ?", sysIoLowWatermarkInMb).Find(&results).Error

	return
}

// GetBatchFromSysIoLowWatermarkInMb 批量查找
func (obj *_AllVirtualIoStatMgr) GetBatchFromSysIoLowWatermarkInMb(sysIoLowWatermarkInMbs []int64) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`sys_io_low_watermark_in_mb` IN (?)", sysIoLowWatermarkInMbs).Find(&results).Error

	return
}

// GetFromSysIoHighWatermarkInMb 通过sys_io_high_watermark_in_mb获取内容
func (obj *_AllVirtualIoStatMgr) GetFromSysIoHighWatermarkInMb(sysIoHighWatermarkInMb int64) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`sys_io_high_watermark_in_mb` = ?", sysIoHighWatermarkInMb).Find(&results).Error

	return
}

// GetBatchFromSysIoHighWatermarkInMb 批量查找
func (obj *_AllVirtualIoStatMgr) GetBatchFromSysIoHighWatermarkInMb(sysIoHighWatermarkInMbs []int64) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`sys_io_high_watermark_in_mb` IN (?)", sysIoHighWatermarkInMbs).Find(&results).Error

	return
}

// GetFromIoBenchResult 通过io_bench_result获取内容
func (obj *_AllVirtualIoStatMgr) GetFromIoBenchResult(ioBenchResult string) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`io_bench_result` = ?", ioBenchResult).Find(&results).Error

	return
}

// GetBatchFromIoBenchResult 批量查找
func (obj *_AllVirtualIoStatMgr) GetBatchFromIoBenchResult(ioBenchResults []string) (results []*AllVirtualIoStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualIoStat{}).Where("`io_bench_result` IN (?)", ioBenchResults).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
