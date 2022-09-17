package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualDtlFirstCachedBufferMgr struct {
	*_BaseMgr
}

// AllVirtualDtlFirstCachedBufferMgr open func
func AllVirtualDtlFirstCachedBufferMgr(db *gorm.DB) *_AllVirtualDtlFirstCachedBufferMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDtlFirstCachedBufferMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDtlFirstCachedBufferMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_dtl_first_cached_buffer"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetTableName() string {
	return "__all_virtual_dtl_first_cached_buffer"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDtlFirstCachedBufferMgr) Reset() *_AllVirtualDtlFirstCachedBufferMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDtlFirstCachedBufferMgr) Get() (result AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDtlFirstCachedBufferMgr) Gets() (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDtlFirstCachedBufferMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualDtlFirstCachedBufferMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualDtlFirstCachedBufferMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualDtlFirstCachedBufferMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithChannelID channel_id获取
func (obj *_AllVirtualDtlFirstCachedBufferMgr) WithChannelID(channelID int64) Option {
	return optionFunc(func(o *options) { o.query["channel_id"] = channelID })
}

// WithCalcedVal calced_val获取
func (obj *_AllVirtualDtlFirstCachedBufferMgr) WithCalcedVal(calcedVal int64) Option {
	return optionFunc(func(o *options) { o.query["calced_val"] = calcedVal })
}

// WithBufferPoolID buffer_pool_id获取
func (obj *_AllVirtualDtlFirstCachedBufferMgr) WithBufferPoolID(bufferPoolID int64) Option {
	return optionFunc(func(o *options) { o.query["buffer_pool_id"] = bufferPoolID })
}

// WithTimeoutTs timeout_ts获取
func (obj *_AllVirtualDtlFirstCachedBufferMgr) WithTimeoutTs(timeoutTs time.Time) Option {
	return optionFunc(func(o *options) { o.query["timeout_ts"] = timeoutTs })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetByOption(opts ...Option) (result AllVirtualDtlFirstCachedBuffer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetByOptions(opts ...Option) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDtlFirstCachedBufferMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDtlFirstCachedBuffer, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where(options.query)
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
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromChannelID 通过channel_id获取内容
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetFromChannelID(channelID int64) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`channel_id` = ?", channelID).Find(&results).Error

	return
}

// GetBatchFromChannelID 批量查找
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetBatchFromChannelID(channelIDs []int64) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`channel_id` IN (?)", channelIDs).Find(&results).Error

	return
}

// GetFromCalcedVal 通过calced_val获取内容
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetFromCalcedVal(calcedVal int64) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`calced_val` = ?", calcedVal).Find(&results).Error

	return
}

// GetBatchFromCalcedVal 批量查找
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetBatchFromCalcedVal(calcedVals []int64) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`calced_val` IN (?)", calcedVals).Find(&results).Error

	return
}

// GetFromBufferPoolID 通过buffer_pool_id获取内容
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetFromBufferPoolID(bufferPoolID int64) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`buffer_pool_id` = ?", bufferPoolID).Find(&results).Error

	return
}

// GetBatchFromBufferPoolID 批量查找
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetBatchFromBufferPoolID(bufferPoolIDs []int64) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`buffer_pool_id` IN (?)", bufferPoolIDs).Find(&results).Error

	return
}

// GetFromTimeoutTs 通过timeout_ts获取内容
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetFromTimeoutTs(timeoutTs time.Time) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`timeout_ts` = ?", timeoutTs).Find(&results).Error

	return
}

// GetBatchFromTimeoutTs 批量查找
func (obj *_AllVirtualDtlFirstCachedBufferMgr) GetBatchFromTimeoutTs(timeoutTss []time.Time) (results []*AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`timeout_ts` IN (?)", timeoutTss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDtlFirstCachedBufferMgr) FetchByPrimaryKey(svrIP string, svrPort int64) (result AllVirtualDtlFirstCachedBuffer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlFirstCachedBuffer{}).Where("`svr_ip` = ? AND `svr_port` = ?", svrIP, svrPort).First(&result).Error

	return
}
