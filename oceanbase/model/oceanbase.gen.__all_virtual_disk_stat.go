package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualDiskStatMgr struct {
	*_BaseMgr
}

// AllVirtualDiskStatMgr open func
func AllVirtualDiskStatMgr(db *gorm.DB) *_AllVirtualDiskStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDiskStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDiskStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_disk_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDiskStatMgr) GetTableName() string {
	return "__all_virtual_disk_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDiskStatMgr) Reset() *_AllVirtualDiskStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDiskStatMgr) Get() (result AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDiskStatMgr) Gets() (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDiskStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualDiskStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualDiskStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTotalSize total_size获取
func (obj *_AllVirtualDiskStatMgr) WithTotalSize(totalSize int64) Option {
	return optionFunc(func(o *options) { o.query["total_size"] = totalSize })
}

// WithUsedSize used_size获取
func (obj *_AllVirtualDiskStatMgr) WithUsedSize(usedSize int64) Option {
	return optionFunc(func(o *options) { o.query["used_size"] = usedSize })
}

// WithFreeSize free_size获取
func (obj *_AllVirtualDiskStatMgr) WithFreeSize(freeSize int64) Option {
	return optionFunc(func(o *options) { o.query["free_size"] = freeSize })
}

// WithIsDiskValid is_disk_valid获取
func (obj *_AllVirtualDiskStatMgr) WithIsDiskValid(isDiskValid int64) Option {
	return optionFunc(func(o *options) { o.query["is_disk_valid"] = isDiskValid })
}

// WithDiskErrorBeginTs disk_error_begin_ts获取
func (obj *_AllVirtualDiskStatMgr) WithDiskErrorBeginTs(diskErrorBeginTs int64) Option {
	return optionFunc(func(o *options) { o.query["disk_error_begin_ts"] = diskErrorBeginTs })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDiskStatMgr) GetByOption(opts ...Option) (result AllVirtualDiskStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDiskStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualDiskStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDiskStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDiskStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where(options.query)
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
func (obj *_AllVirtualDiskStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualDiskStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualDiskStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualDiskStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTotalSize 通过total_size获取内容
func (obj *_AllVirtualDiskStatMgr) GetFromTotalSize(totalSize int64) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`total_size` = ?", totalSize).Find(&results).Error

	return
}

// GetBatchFromTotalSize 批量查找
func (obj *_AllVirtualDiskStatMgr) GetBatchFromTotalSize(totalSizes []int64) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`total_size` IN (?)", totalSizes).Find(&results).Error

	return
}

// GetFromUsedSize 通过used_size获取内容
func (obj *_AllVirtualDiskStatMgr) GetFromUsedSize(usedSize int64) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`used_size` = ?", usedSize).Find(&results).Error

	return
}

// GetBatchFromUsedSize 批量查找
func (obj *_AllVirtualDiskStatMgr) GetBatchFromUsedSize(usedSizes []int64) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`used_size` IN (?)", usedSizes).Find(&results).Error

	return
}

// GetFromFreeSize 通过free_size获取内容
func (obj *_AllVirtualDiskStatMgr) GetFromFreeSize(freeSize int64) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`free_size` = ?", freeSize).Find(&results).Error

	return
}

// GetBatchFromFreeSize 批量查找
func (obj *_AllVirtualDiskStatMgr) GetBatchFromFreeSize(freeSizes []int64) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`free_size` IN (?)", freeSizes).Find(&results).Error

	return
}

// GetFromIsDiskValid 通过is_disk_valid获取内容
func (obj *_AllVirtualDiskStatMgr) GetFromIsDiskValid(isDiskValid int64) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`is_disk_valid` = ?", isDiskValid).Find(&results).Error

	return
}

// GetBatchFromIsDiskValid 批量查找
func (obj *_AllVirtualDiskStatMgr) GetBatchFromIsDiskValid(isDiskValids []int64) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`is_disk_valid` IN (?)", isDiskValids).Find(&results).Error

	return
}

// GetFromDiskErrorBeginTs 通过disk_error_begin_ts获取内容
func (obj *_AllVirtualDiskStatMgr) GetFromDiskErrorBeginTs(diskErrorBeginTs int64) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`disk_error_begin_ts` = ?", diskErrorBeginTs).Find(&results).Error

	return
}

// GetBatchFromDiskErrorBeginTs 批量查找
func (obj *_AllVirtualDiskStatMgr) GetBatchFromDiskErrorBeginTs(diskErrorBeginTss []int64) (results []*AllVirtualDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDiskStat{}).Where("`disk_error_begin_ts` IN (?)", diskErrorBeginTss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
