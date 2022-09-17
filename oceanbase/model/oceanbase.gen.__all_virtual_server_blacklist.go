package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualServerBlacklistMgr struct {
	*_BaseMgr
}

// AllVirtualServerBlacklistMgr open func
func AllVirtualServerBlacklistMgr(db *gorm.DB) *_AllVirtualServerBlacklistMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualServerBlacklistMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualServerBlacklistMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_server_blacklist"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualServerBlacklistMgr) GetTableName() string {
	return "__all_virtual_server_blacklist"
}

// Reset 重置gorm会话
func (obj *_AllVirtualServerBlacklistMgr) Reset() *_AllVirtualServerBlacklistMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualServerBlacklistMgr) Get() (result AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualServerBlacklistMgr) Gets() (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualServerBlacklistMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualServerBlacklistMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualServerBlacklistMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithDstIP dst_ip获取
func (obj *_AllVirtualServerBlacklistMgr) WithDstIP(dstIP string) Option {
	return optionFunc(func(o *options) { o.query["dst_ip"] = dstIP })
}

// WithDstPort dst_port获取
func (obj *_AllVirtualServerBlacklistMgr) WithDstPort(dstPort int64) Option {
	return optionFunc(func(o *options) { o.query["dst_port"] = dstPort })
}

// WithIsInBlacklist is_in_blacklist获取
func (obj *_AllVirtualServerBlacklistMgr) WithIsInBlacklist(isInBlacklist int8) Option {
	return optionFunc(func(o *options) { o.query["is_in_blacklist"] = isInBlacklist })
}

// WithIsClockdiffError is_clockdiff_error获取
func (obj *_AllVirtualServerBlacklistMgr) WithIsClockdiffError(isClockdiffError int8) Option {
	return optionFunc(func(o *options) { o.query["is_clockdiff_error"] = isClockdiffError })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualServerBlacklistMgr) GetByOption(opts ...Option) (result AllVirtualServerBlacklist, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualServerBlacklistMgr) GetByOptions(opts ...Option) (results []*AllVirtualServerBlacklist, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualServerBlacklistMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualServerBlacklist, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where(options.query)
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
func (obj *_AllVirtualServerBlacklistMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualServerBlacklistMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualServerBlacklistMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualServerBlacklistMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromDstIP 通过dst_ip获取内容
func (obj *_AllVirtualServerBlacklistMgr) GetFromDstIP(dstIP string) (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where("`dst_ip` = ?", dstIP).Find(&results).Error

	return
}

// GetBatchFromDstIP 批量查找
func (obj *_AllVirtualServerBlacklistMgr) GetBatchFromDstIP(dstIPs []string) (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where("`dst_ip` IN (?)", dstIPs).Find(&results).Error

	return
}

// GetFromDstPort 通过dst_port获取内容
func (obj *_AllVirtualServerBlacklistMgr) GetFromDstPort(dstPort int64) (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where("`dst_port` = ?", dstPort).Find(&results).Error

	return
}

// GetBatchFromDstPort 批量查找
func (obj *_AllVirtualServerBlacklistMgr) GetBatchFromDstPort(dstPorts []int64) (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where("`dst_port` IN (?)", dstPorts).Find(&results).Error

	return
}

// GetFromIsInBlacklist 通过is_in_blacklist获取内容
func (obj *_AllVirtualServerBlacklistMgr) GetFromIsInBlacklist(isInBlacklist int8) (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where("`is_in_blacklist` = ?", isInBlacklist).Find(&results).Error

	return
}

// GetBatchFromIsInBlacklist 批量查找
func (obj *_AllVirtualServerBlacklistMgr) GetBatchFromIsInBlacklist(isInBlacklists []int8) (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where("`is_in_blacklist` IN (?)", isInBlacklists).Find(&results).Error

	return
}

// GetFromIsClockdiffError 通过is_clockdiff_error获取内容
func (obj *_AllVirtualServerBlacklistMgr) GetFromIsClockdiffError(isClockdiffError int8) (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where("`is_clockdiff_error` = ?", isClockdiffError).Find(&results).Error

	return
}

// GetBatchFromIsClockdiffError 批量查找
func (obj *_AllVirtualServerBlacklistMgr) GetBatchFromIsClockdiffError(isClockdiffErrors []int8) (results []*AllVirtualServerBlacklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBlacklist{}).Where("`is_clockdiff_error` IN (?)", isClockdiffErrors).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
