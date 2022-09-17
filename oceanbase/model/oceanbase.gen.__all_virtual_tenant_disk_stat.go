package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTenantDiskStatMgr struct {
	*_BaseMgr
}

// AllVirtualTenantDiskStatMgr open func
func AllVirtualTenantDiskStatMgr(db *gorm.DB) *_AllVirtualTenantDiskStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantDiskStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantDiskStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_disk_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantDiskStatMgr) GetTableName() string {
	return "__all_virtual_tenant_disk_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantDiskStatMgr) Reset() *_AllVirtualTenantDiskStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantDiskStatMgr) Get() (result AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantDiskStatMgr) Gets() (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantDiskStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantDiskStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTenantDiskStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTenantDiskStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithZone zone获取
func (obj *_AllVirtualTenantDiskStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithBlockType block_type获取
func (obj *_AllVirtualTenantDiskStatMgr) WithBlockType(blockType string) Option {
	return optionFunc(func(o *options) { o.query["block_type"] = blockType })
}

// WithBlockSize block_size获取
func (obj *_AllVirtualTenantDiskStatMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantDiskStatMgr) GetByOption(opts ...Option) (result AllVirtualTenantDiskStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantDiskStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantDiskStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantDiskStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantDiskStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where(options.query)
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
func (obj *_AllVirtualTenantDiskStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantDiskStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualTenantDiskStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTenantDiskStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTenantDiskStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTenantDiskStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualTenantDiskStatMgr) GetFromZone(zone string) (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualTenantDiskStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromBlockType 通过block_type获取内容
func (obj *_AllVirtualTenantDiskStatMgr) GetFromBlockType(blockType string) (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where("`block_type` = ?", blockType).Find(&results).Error

	return
}

// GetBatchFromBlockType 批量查找
func (obj *_AllVirtualTenantDiskStatMgr) GetBatchFromBlockType(blockTypes []string) (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where("`block_type` IN (?)", blockTypes).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllVirtualTenantDiskStatMgr) GetFromBlockSize(blockSize int64) (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllVirtualTenantDiskStatMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllVirtualTenantDiskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantDiskStat{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
