package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualObrpcStatMgr struct {
	*_BaseMgr
}

// AllVirtualObrpcStatMgr open func
func AllVirtualObrpcStatMgr(db *gorm.DB) *_AllVirtualObrpcStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualObrpcStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualObrpcStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_obrpc_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualObrpcStatMgr) GetTableName() string {
	return "__all_virtual_obrpc_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualObrpcStatMgr) Reset() *_AllVirtualObrpcStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualObrpcStatMgr) Get() (result AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualObrpcStatMgr) Gets() (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualObrpcStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualObrpcStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualObrpcStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualObrpcStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithDestIP dest_ip获取
func (obj *_AllVirtualObrpcStatMgr) WithDestIP(destIP string) Option {
	return optionFunc(func(o *options) { o.query["dest_ip"] = destIP })
}

// WithDestPort dest_port获取
func (obj *_AllVirtualObrpcStatMgr) WithDestPort(destPort int64) Option {
	return optionFunc(func(o *options) { o.query["dest_port"] = destPort })
}

// WithIndex index获取
func (obj *_AllVirtualObrpcStatMgr) WithIndex(index int64) Option {
	return optionFunc(func(o *options) { o.query["index"] = index })
}

// WithZone zone获取
func (obj *_AllVirtualObrpcStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithPcode pcode获取
func (obj *_AllVirtualObrpcStatMgr) WithPcode(pcode int64) Option {
	return optionFunc(func(o *options) { o.query["pcode"] = pcode })
}

// WithPcodeName pcode_name获取
func (obj *_AllVirtualObrpcStatMgr) WithPcodeName(pcodeName string) Option {
	return optionFunc(func(o *options) { o.query["pcode_name"] = pcodeName })
}

// WithCount count获取
func (obj *_AllVirtualObrpcStatMgr) WithCount(count int64) Option {
	return optionFunc(func(o *options) { o.query["count"] = count })
}

// WithTotalTime total_time获取
func (obj *_AllVirtualObrpcStatMgr) WithTotalTime(totalTime int64) Option {
	return optionFunc(func(o *options) { o.query["total_time"] = totalTime })
}

// WithTotalSize total_size获取
func (obj *_AllVirtualObrpcStatMgr) WithTotalSize(totalSize int64) Option {
	return optionFunc(func(o *options) { o.query["total_size"] = totalSize })
}

// WithMaxTime max_time获取
func (obj *_AllVirtualObrpcStatMgr) WithMaxTime(maxTime int64) Option {
	return optionFunc(func(o *options) { o.query["max_time"] = maxTime })
}

// WithMinTime min_time获取
func (obj *_AllVirtualObrpcStatMgr) WithMinTime(minTime int64) Option {
	return optionFunc(func(o *options) { o.query["min_time"] = minTime })
}

// WithMaxSize max_size获取
func (obj *_AllVirtualObrpcStatMgr) WithMaxSize(maxSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_size"] = maxSize })
}

// WithMinSize min_size获取
func (obj *_AllVirtualObrpcStatMgr) WithMinSize(minSize int64) Option {
	return optionFunc(func(o *options) { o.query["min_size"] = minSize })
}

// WithFailure failure获取
func (obj *_AllVirtualObrpcStatMgr) WithFailure(failure int64) Option {
	return optionFunc(func(o *options) { o.query["failure"] = failure })
}

// WithTimeout timeout获取
func (obj *_AllVirtualObrpcStatMgr) WithTimeout(timeout int64) Option {
	return optionFunc(func(o *options) { o.query["timeout"] = timeout })
}

// WithSync sync获取
func (obj *_AllVirtualObrpcStatMgr) WithSync(sync int64) Option {
	return optionFunc(func(o *options) { o.query["sync"] = sync })
}

// WithAsync async获取
func (obj *_AllVirtualObrpcStatMgr) WithAsync(async int64) Option {
	return optionFunc(func(o *options) { o.query["async"] = async })
}

// WithLastTimestamp last_timestamp获取
func (obj *_AllVirtualObrpcStatMgr) WithLastTimestamp(lastTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_timestamp"] = lastTimestamp })
}

// WithIsize isize获取
func (obj *_AllVirtualObrpcStatMgr) WithIsize(isize int64) Option {
	return optionFunc(func(o *options) { o.query["isize"] = isize })
}

// WithIcount icount获取
func (obj *_AllVirtualObrpcStatMgr) WithIcount(icount int64) Option {
	return optionFunc(func(o *options) { o.query["icount"] = icount })
}

// WithNetTime net_time获取
func (obj *_AllVirtualObrpcStatMgr) WithNetTime(netTime int64) Option {
	return optionFunc(func(o *options) { o.query["net_time"] = netTime })
}

// WithWaitTime wait_time获取
func (obj *_AllVirtualObrpcStatMgr) WithWaitTime(waitTime int64) Option {
	return optionFunc(func(o *options) { o.query["wait_time"] = waitTime })
}

// WithQueueTime queue_time获取
func (obj *_AllVirtualObrpcStatMgr) WithQueueTime(queueTime int64) Option {
	return optionFunc(func(o *options) { o.query["queue_time"] = queueTime })
}

// WithProcessTime process_time获取
func (obj *_AllVirtualObrpcStatMgr) WithProcessTime(processTime int64) Option {
	return optionFunc(func(o *options) { o.query["process_time"] = processTime })
}

// WithIlastTimestamp ilast_timestamp获取
func (obj *_AllVirtualObrpcStatMgr) WithIlastTimestamp(ilastTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["ilast_timestamp"] = ilastTimestamp })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualObrpcStatMgr) GetByOption(opts ...Option) (result AllVirtualObrpcStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualObrpcStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualObrpcStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualObrpcStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualObrpcStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where(options.query)
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
func (obj *_AllVirtualObrpcStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromDestIP 通过dest_ip获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromDestIP(destIP string) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`dest_ip` = ?", destIP).Find(&results).Error

	return
}

// GetBatchFromDestIP 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromDestIP(destIPs []string) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`dest_ip` IN (?)", destIPs).Find(&results).Error

	return
}

// GetFromDestPort 通过dest_port获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromDestPort(destPort int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`dest_port` = ?", destPort).Find(&results).Error

	return
}

// GetBatchFromDestPort 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromDestPort(destPorts []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`dest_port` IN (?)", destPorts).Find(&results).Error

	return
}

// GetFromIndex 通过index获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromIndex(index int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`index` = ?", index).Find(&results).Error

	return
}

// GetBatchFromIndex 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromIndex(indexs []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`index` IN (?)", indexs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromZone(zone string) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromPcode 通过pcode获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromPcode(pcode int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`pcode` = ?", pcode).Find(&results).Error

	return
}

// GetBatchFromPcode 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromPcode(pcodes []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`pcode` IN (?)", pcodes).Find(&results).Error

	return
}

// GetFromPcodeName 通过pcode_name获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromPcodeName(pcodeName string) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`pcode_name` = ?", pcodeName).Find(&results).Error

	return
}

// GetBatchFromPcodeName 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromPcodeName(pcodeNames []string) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`pcode_name` IN (?)", pcodeNames).Find(&results).Error

	return
}

// GetFromCount 通过count获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromCount(count int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`count` = ?", count).Find(&results).Error

	return
}

// GetBatchFromCount 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromCount(counts []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`count` IN (?)", counts).Find(&results).Error

	return
}

// GetFromTotalTime 通过total_time获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromTotalTime(totalTime int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`total_time` = ?", totalTime).Find(&results).Error

	return
}

// GetBatchFromTotalTime 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromTotalTime(totalTimes []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`total_time` IN (?)", totalTimes).Find(&results).Error

	return
}

// GetFromTotalSize 通过total_size获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromTotalSize(totalSize int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`total_size` = ?", totalSize).Find(&results).Error

	return
}

// GetBatchFromTotalSize 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromTotalSize(totalSizes []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`total_size` IN (?)", totalSizes).Find(&results).Error

	return
}

// GetFromMaxTime 通过max_time获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromMaxTime(maxTime int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`max_time` = ?", maxTime).Find(&results).Error

	return
}

// GetBatchFromMaxTime 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromMaxTime(maxTimes []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`max_time` IN (?)", maxTimes).Find(&results).Error

	return
}

// GetFromMinTime 通过min_time获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromMinTime(minTime int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`min_time` = ?", minTime).Find(&results).Error

	return
}

// GetBatchFromMinTime 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromMinTime(minTimes []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`min_time` IN (?)", minTimes).Find(&results).Error

	return
}

// GetFromMaxSize 通过max_size获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromMaxSize(maxSize int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`max_size` = ?", maxSize).Find(&results).Error

	return
}

// GetBatchFromMaxSize 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromMaxSize(maxSizes []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`max_size` IN (?)", maxSizes).Find(&results).Error

	return
}

// GetFromMinSize 通过min_size获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromMinSize(minSize int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`min_size` = ?", minSize).Find(&results).Error

	return
}

// GetBatchFromMinSize 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromMinSize(minSizes []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`min_size` IN (?)", minSizes).Find(&results).Error

	return
}

// GetFromFailure 通过failure获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromFailure(failure int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`failure` = ?", failure).Find(&results).Error

	return
}

// GetBatchFromFailure 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromFailure(failures []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`failure` IN (?)", failures).Find(&results).Error

	return
}

// GetFromTimeout 通过timeout获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromTimeout(timeout int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`timeout` = ?", timeout).Find(&results).Error

	return
}

// GetBatchFromTimeout 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromTimeout(timeouts []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`timeout` IN (?)", timeouts).Find(&results).Error

	return
}

// GetFromSync 通过sync获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromSync(sync int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`sync` = ?", sync).Find(&results).Error

	return
}

// GetBatchFromSync 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromSync(syncs []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`sync` IN (?)", syncs).Find(&results).Error

	return
}

// GetFromAsync 通过async获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromAsync(async int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`async` = ?", async).Find(&results).Error

	return
}

// GetBatchFromAsync 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromAsync(asyncs []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`async` IN (?)", asyncs).Find(&results).Error

	return
}

// GetFromLastTimestamp 通过last_timestamp获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromLastTimestamp(lastTimestamp time.Time) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`last_timestamp` = ?", lastTimestamp).Find(&results).Error

	return
}

// GetBatchFromLastTimestamp 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromLastTimestamp(lastTimestamps []time.Time) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`last_timestamp` IN (?)", lastTimestamps).Find(&results).Error

	return
}

// GetFromIsize 通过isize获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromIsize(isize int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`isize` = ?", isize).Find(&results).Error

	return
}

// GetBatchFromIsize 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromIsize(isizes []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`isize` IN (?)", isizes).Find(&results).Error

	return
}

// GetFromIcount 通过icount获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromIcount(icount int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`icount` = ?", icount).Find(&results).Error

	return
}

// GetBatchFromIcount 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromIcount(icounts []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`icount` IN (?)", icounts).Find(&results).Error

	return
}

// GetFromNetTime 通过net_time获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromNetTime(netTime int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`net_time` = ?", netTime).Find(&results).Error

	return
}

// GetBatchFromNetTime 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromNetTime(netTimes []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`net_time` IN (?)", netTimes).Find(&results).Error

	return
}

// GetFromWaitTime 通过wait_time获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromWaitTime(waitTime int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`wait_time` = ?", waitTime).Find(&results).Error

	return
}

// GetBatchFromWaitTime 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromWaitTime(waitTimes []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`wait_time` IN (?)", waitTimes).Find(&results).Error

	return
}

// GetFromQueueTime 通过queue_time获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromQueueTime(queueTime int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`queue_time` = ?", queueTime).Find(&results).Error

	return
}

// GetBatchFromQueueTime 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromQueueTime(queueTimes []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`queue_time` IN (?)", queueTimes).Find(&results).Error

	return
}

// GetFromProcessTime 通过process_time获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromProcessTime(processTime int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`process_time` = ?", processTime).Find(&results).Error

	return
}

// GetBatchFromProcessTime 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromProcessTime(processTimes []int64) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`process_time` IN (?)", processTimes).Find(&results).Error

	return
}

// GetFromIlastTimestamp 通过ilast_timestamp获取内容
func (obj *_AllVirtualObrpcStatMgr) GetFromIlastTimestamp(ilastTimestamp time.Time) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`ilast_timestamp` = ?", ilastTimestamp).Find(&results).Error

	return
}

// GetBatchFromIlastTimestamp 批量查找
func (obj *_AllVirtualObrpcStatMgr) GetBatchFromIlastTimestamp(ilastTimestamps []time.Time) (results []*AllVirtualObrpcStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObrpcStat{}).Where("`ilast_timestamp` IN (?)", ilastTimestamps).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
