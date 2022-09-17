package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualPsItemInfoMgr struct {
	*_BaseMgr
}

// AllVirtualPsItemInfoMgr open func
func AllVirtualPsItemInfoMgr(db *gorm.DB) *_AllVirtualPsItemInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPsItemInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPsItemInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_ps_item_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPsItemInfoMgr) GetTableName() string {
	return "__all_virtual_ps_item_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPsItemInfoMgr) Reset() *_AllVirtualPsItemInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPsItemInfoMgr) Get() (result AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPsItemInfoMgr) Gets() (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPsItemInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPsItemInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPsItemInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPsItemInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithStmtID stmt_id获取
func (obj *_AllVirtualPsItemInfoMgr) WithStmtID(stmtID int64) Option {
	return optionFunc(func(o *options) { o.query["stmt_id"] = stmtID })
}

// WithDbID db_id获取
func (obj *_AllVirtualPsItemInfoMgr) WithDbID(dbID int64) Option {
	return optionFunc(func(o *options) { o.query["db_id"] = dbID })
}

// WithPsSQL ps_sql获取
func (obj *_AllVirtualPsItemInfoMgr) WithPsSQL(psSQL string) Option {
	return optionFunc(func(o *options) { o.query["ps_sql"] = psSQL })
}

// WithParamCount param_count获取
func (obj *_AllVirtualPsItemInfoMgr) WithParamCount(paramCount int64) Option {
	return optionFunc(func(o *options) { o.query["param_count"] = paramCount })
}

// WithStmtItemRefCount stmt_item_ref_count获取
func (obj *_AllVirtualPsItemInfoMgr) WithStmtItemRefCount(stmtItemRefCount int64) Option {
	return optionFunc(func(o *options) { o.query["stmt_item_ref_count"] = stmtItemRefCount })
}

// WithStmtInfoRefCount stmt_info_ref_count获取
func (obj *_AllVirtualPsItemInfoMgr) WithStmtInfoRefCount(stmtInfoRefCount int64) Option {
	return optionFunc(func(o *options) { o.query["stmt_info_ref_count"] = stmtInfoRefCount })
}

// WithMemHold mem_hold获取
func (obj *_AllVirtualPsItemInfoMgr) WithMemHold(memHold int64) Option {
	return optionFunc(func(o *options) { o.query["mem_hold"] = memHold })
}

// WithStmtType stmt_type获取
func (obj *_AllVirtualPsItemInfoMgr) WithStmtType(stmtType int64) Option {
	return optionFunc(func(o *options) { o.query["stmt_type"] = stmtType })
}

// WithChecksum checksum获取
func (obj *_AllVirtualPsItemInfoMgr) WithChecksum(checksum int64) Option {
	return optionFunc(func(o *options) { o.query["checksum"] = checksum })
}

// WithExpired expired获取
func (obj *_AllVirtualPsItemInfoMgr) WithExpired(expired int8) Option {
	return optionFunc(func(o *options) { o.query["expired"] = expired })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPsItemInfoMgr) GetByOption(opts ...Option) (result AllVirtualPsItemInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPsItemInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualPsItemInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPsItemInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPsItemInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where(options.query)
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
func (obj *_AllVirtualPsItemInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPsItemInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPsItemInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromStmtID 通过stmt_id获取内容
func (obj *_AllVirtualPsItemInfoMgr) GetFromStmtID(stmtID int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`stmt_id` = ?", stmtID).Find(&results).Error

	return
}

// GetBatchFromStmtID 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromStmtID(stmtIDs []int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`stmt_id` IN (?)", stmtIDs).Find(&results).Error

	return
}

// GetFromDbID 通过db_id获取内容
func (obj *_AllVirtualPsItemInfoMgr) GetFromDbID(dbID int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`db_id` = ?", dbID).Find(&results).Error

	return
}

// GetBatchFromDbID 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromDbID(dbIDs []int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`db_id` IN (?)", dbIDs).Find(&results).Error

	return
}

// GetFromPsSQL 通过ps_sql获取内容
func (obj *_AllVirtualPsItemInfoMgr) GetFromPsSQL(psSQL string) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`ps_sql` = ?", psSQL).Find(&results).Error

	return
}

// GetBatchFromPsSQL 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromPsSQL(psSQLs []string) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`ps_sql` IN (?)", psSQLs).Find(&results).Error

	return
}

// GetFromParamCount 通过param_count获取内容
func (obj *_AllVirtualPsItemInfoMgr) GetFromParamCount(paramCount int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`param_count` = ?", paramCount).Find(&results).Error

	return
}

// GetBatchFromParamCount 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromParamCount(paramCounts []int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`param_count` IN (?)", paramCounts).Find(&results).Error

	return
}

// GetFromStmtItemRefCount 通过stmt_item_ref_count获取内容
func (obj *_AllVirtualPsItemInfoMgr) GetFromStmtItemRefCount(stmtItemRefCount int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`stmt_item_ref_count` = ?", stmtItemRefCount).Find(&results).Error

	return
}

// GetBatchFromStmtItemRefCount 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromStmtItemRefCount(stmtItemRefCounts []int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`stmt_item_ref_count` IN (?)", stmtItemRefCounts).Find(&results).Error

	return
}

// GetFromStmtInfoRefCount 通过stmt_info_ref_count获取内容
func (obj *_AllVirtualPsItemInfoMgr) GetFromStmtInfoRefCount(stmtInfoRefCount int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`stmt_info_ref_count` = ?", stmtInfoRefCount).Find(&results).Error

	return
}

// GetBatchFromStmtInfoRefCount 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromStmtInfoRefCount(stmtInfoRefCounts []int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`stmt_info_ref_count` IN (?)", stmtInfoRefCounts).Find(&results).Error

	return
}

// GetFromMemHold 通过mem_hold获取内容
func (obj *_AllVirtualPsItemInfoMgr) GetFromMemHold(memHold int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`mem_hold` = ?", memHold).Find(&results).Error

	return
}

// GetBatchFromMemHold 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromMemHold(memHolds []int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`mem_hold` IN (?)", memHolds).Find(&results).Error

	return
}

// GetFromStmtType 通过stmt_type获取内容
func (obj *_AllVirtualPsItemInfoMgr) GetFromStmtType(stmtType int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`stmt_type` = ?", stmtType).Find(&results).Error

	return
}

// GetBatchFromStmtType 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromStmtType(stmtTypes []int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`stmt_type` IN (?)", stmtTypes).Find(&results).Error

	return
}

// GetFromChecksum 通过checksum获取内容
func (obj *_AllVirtualPsItemInfoMgr) GetFromChecksum(checksum int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`checksum` = ?", checksum).Find(&results).Error

	return
}

// GetBatchFromChecksum 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromChecksum(checksums []int64) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`checksum` IN (?)", checksums).Find(&results).Error

	return
}

// GetFromExpired 通过expired获取内容
func (obj *_AllVirtualPsItemInfoMgr) GetFromExpired(expired int8) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`expired` = ?", expired).Find(&results).Error

	return
}

// GetBatchFromExpired 批量查找
func (obj *_AllVirtualPsItemInfoMgr) GetBatchFromExpired(expireds []int8) (results []*AllVirtualPsItemInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPsItemInfo{}).Where("`expired` IN (?)", expireds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
