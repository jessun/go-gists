package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualPsStatMgr struct {
	*_BaseMgr
}

// AllVirtualPsStatMgr open func
func AllVirtualPsStatMgr(db *gorm.DB) *_AllVirtualPsStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPsStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPsStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_ps_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPsStatMgr) GetTableName() string {
	return "__all_virtual_ps_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPsStatMgr) Reset() *_AllVirtualPsStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPsStatMgr) Get() (result AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPsStatMgr) Gets() (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPsStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPsStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPsStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPsStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithStmtCount stmt_count获取
func (obj *_AllVirtualPsStatMgr) WithStmtCount(stmtCount int64) Option {
	return optionFunc(func(o *options) { o.query["stmt_count"] = stmtCount })
}

// WithHitCount hit_count获取
func (obj *_AllVirtualPsStatMgr) WithHitCount(hitCount int64) Option {
	return optionFunc(func(o *options) { o.query["hit_count"] = hitCount })
}

// WithAccessCount access_count获取
func (obj *_AllVirtualPsStatMgr) WithAccessCount(accessCount int64) Option {
	return optionFunc(func(o *options) { o.query["access_count"] = accessCount })
}

// WithMemHold mem_hold获取
func (obj *_AllVirtualPsStatMgr) WithMemHold(memHold int64) Option {
	return optionFunc(func(o *options) { o.query["mem_hold"] = memHold })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPsStatMgr) GetByOption(opts ...Option) (result AllVirtualPsStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPsStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualPsStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPsStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPsStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where(options.query)
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
func (obj *_AllVirtualPsStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPsStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPsStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPsStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPsStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPsStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromStmtCount 通过stmt_count获取内容
func (obj *_AllVirtualPsStatMgr) GetFromStmtCount(stmtCount int64) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`stmt_count` = ?", stmtCount).Find(&results).Error

	return
}

// GetBatchFromStmtCount 批量查找
func (obj *_AllVirtualPsStatMgr) GetBatchFromStmtCount(stmtCounts []int64) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`stmt_count` IN (?)", stmtCounts).Find(&results).Error

	return
}

// GetFromHitCount 通过hit_count获取内容
func (obj *_AllVirtualPsStatMgr) GetFromHitCount(hitCount int64) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`hit_count` = ?", hitCount).Find(&results).Error

	return
}

// GetBatchFromHitCount 批量查找
func (obj *_AllVirtualPsStatMgr) GetBatchFromHitCount(hitCounts []int64) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`hit_count` IN (?)", hitCounts).Find(&results).Error

	return
}

// GetFromAccessCount 通过access_count获取内容
func (obj *_AllVirtualPsStatMgr) GetFromAccessCount(accessCount int64) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`access_count` = ?", accessCount).Find(&results).Error

	return
}

// GetBatchFromAccessCount 批量查找
func (obj *_AllVirtualPsStatMgr) GetBatchFromAccessCount(accessCounts []int64) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`access_count` IN (?)", accessCounts).Find(&results).Error

	return
}

// GetFromMemHold 通过mem_hold获取内容
func (obj *_AllVirtualPsStatMgr) GetFromMemHold(memHold int64) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`mem_hold` = ?", memHold).Find(&results).Error

	return
}

// GetBatchFromMemHold 批量查找
func (obj *_AllVirtualPsStatMgr) GetBatchFromMemHold(memHolds []int64) (results []*AllVirtualPsStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsStat{}).Where("`mem_hold` IN (?)", memHolds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
