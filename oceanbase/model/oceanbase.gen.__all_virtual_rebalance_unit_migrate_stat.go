package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualRebalanceUnitMigrateStatMgr struct {
	*_BaseMgr
}

// AllVirtualRebalanceUnitMigrateStatMgr open func
func AllVirtualRebalanceUnitMigrateStatMgr(db *gorm.DB) *_AllVirtualRebalanceUnitMigrateStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRebalanceUnitMigrateStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRebalanceUnitMigrateStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_rebalance_unit_migrate_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetTableName() string {
	return "__all_virtual_rebalance_unit_migrate_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) Reset() *_AllVirtualRebalanceUnitMigrateStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) Get() (result AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) Gets() (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUnitID unit_id获取
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithZone zone获取
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithSrcSvrIP src_svr_ip获取
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) WithSrcSvrIP(srcSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["src_svr_ip"] = srcSvrIP })
}

// WithSrcSvrPort src_svr_port获取
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) WithSrcSvrPort(srcSvrPort int64) Option {
	return optionFunc(func(o *options) { o.query["src_svr_port"] = srcSvrPort })
}

// WithDstSvrIP dst_svr_ip获取
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) WithDstSvrIP(dstSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["dst_svr_ip"] = dstSvrIP })
}

// WithDstSvrPort dst_svr_port获取
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) WithDstSvrPort(dstSvrPort int64) Option {
	return optionFunc(func(o *options) { o.query["dst_svr_port"] = dstSvrPort })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetByOption(opts ...Option) (result AllVirtualRebalanceUnitMigrateStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRebalanceUnitMigrateStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where(options.query)
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

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetFromUnitID(unitID int64) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetFromZone(zone string) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromSrcSvrIP 通过src_svr_ip获取内容
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetFromSrcSvrIP(srcSvrIP string) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where("`src_svr_ip` = ?", srcSvrIP).Find(&results).Error

	return
}

// GetBatchFromSrcSvrIP 批量查找
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetBatchFromSrcSvrIP(srcSvrIPs []string) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where("`src_svr_ip` IN (?)", srcSvrIPs).Find(&results).Error

	return
}

// GetFromSrcSvrPort 通过src_svr_port获取内容
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetFromSrcSvrPort(srcSvrPort int64) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where("`src_svr_port` = ?", srcSvrPort).Find(&results).Error

	return
}

// GetBatchFromSrcSvrPort 批量查找
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetBatchFromSrcSvrPort(srcSvrPorts []int64) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where("`src_svr_port` IN (?)", srcSvrPorts).Find(&results).Error

	return
}

// GetFromDstSvrIP 通过dst_svr_ip获取内容
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetFromDstSvrIP(dstSvrIP string) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where("`dst_svr_ip` = ?", dstSvrIP).Find(&results).Error

	return
}

// GetBatchFromDstSvrIP 批量查找
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetBatchFromDstSvrIP(dstSvrIPs []string) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where("`dst_svr_ip` IN (?)", dstSvrIPs).Find(&results).Error

	return
}

// GetFromDstSvrPort 通过dst_svr_port获取内容
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetFromDstSvrPort(dstSvrPort int64) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where("`dst_svr_port` = ?", dstSvrPort).Find(&results).Error

	return
}

// GetBatchFromDstSvrPort 批量查找
func (obj *_AllVirtualRebalanceUnitMigrateStatMgr) GetBatchFromDstSvrPort(dstSvrPorts []int64) (results []*AllVirtualRebalanceUnitMigrateStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitMigrateStat{}).Where("`dst_svr_port` IN (?)", dstSvrPorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
