package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualGlobalTransactionMgr struct {
	*_BaseMgr
}

// AllVirtualGlobalTransactionMgr open func
func AllVirtualGlobalTransactionMgr(db *gorm.DB) *_AllVirtualGlobalTransactionMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualGlobalTransactionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualGlobalTransactionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_global_transaction"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualGlobalTransactionMgr) GetTableName() string {
	return "__all_virtual_global_transaction"
}

// Reset 重置gorm会话
func (obj *_AllVirtualGlobalTransactionMgr) Reset() *_AllVirtualGlobalTransactionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualGlobalTransactionMgr) Get() (result AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualGlobalTransactionMgr) Gets() (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualGlobalTransactionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualGlobalTransactionMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithGtrid gtrid获取
func (obj *_AllVirtualGlobalTransactionMgr) WithGtrid(gtrid []byte) Option {
	return optionFunc(func(o *options) { o.query["gtrid"] = gtrid })
}

// WithBqual bqual获取
func (obj *_AllVirtualGlobalTransactionMgr) WithBqual(bqual []byte) Option {
	return optionFunc(func(o *options) { o.query["bqual"] = bqual })
}

// WithFormatID format_id获取
func (obj *_AllVirtualGlobalTransactionMgr) WithFormatID(formatID int64) Option {
	return optionFunc(func(o *options) { o.query["format_id"] = formatID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualGlobalTransactionMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualGlobalTransactionMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTransID trans_id获取
func (obj *_AllVirtualGlobalTransactionMgr) WithTransID(transID string) Option {
	return optionFunc(func(o *options) { o.query["trans_id"] = transID })
}

// WithCoordinator coordinator获取
func (obj *_AllVirtualGlobalTransactionMgr) WithCoordinator(coordinator string) Option {
	return optionFunc(func(o *options) { o.query["coordinator"] = coordinator })
}

// WithSchedulerIP scheduler_ip获取
func (obj *_AllVirtualGlobalTransactionMgr) WithSchedulerIP(schedulerIP string) Option {
	return optionFunc(func(o *options) { o.query["scheduler_ip"] = schedulerIP })
}

// WithSchedulerPort scheduler_port获取
func (obj *_AllVirtualGlobalTransactionMgr) WithSchedulerPort(schedulerPort int64) Option {
	return optionFunc(func(o *options) { o.query["scheduler_port"] = schedulerPort })
}

// WithIsReadonly is_readonly获取
func (obj *_AllVirtualGlobalTransactionMgr) WithIsReadonly(isReadonly int8) Option {
	return optionFunc(func(o *options) { o.query["is_readonly"] = isReadonly })
}

// WithState state获取
func (obj *_AllVirtualGlobalTransactionMgr) WithState(state int64) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithEndFlag end_flag获取
func (obj *_AllVirtualGlobalTransactionMgr) WithEndFlag(endFlag int64) Option {
	return optionFunc(func(o *options) { o.query["end_flag"] = endFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualGlobalTransactionMgr) GetByOption(opts ...Option) (result AllVirtualGlobalTransaction, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualGlobalTransactionMgr) GetByOptions(opts ...Option) (results []*AllVirtualGlobalTransaction, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualGlobalTransactionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualGlobalTransaction, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where(options.query)
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
func (obj *_AllVirtualGlobalTransactionMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromGtrid 通过gtrid获取内容
func (obj *_AllVirtualGlobalTransactionMgr) GetFromGtrid(gtrid []byte) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`gtrid` = ?", gtrid).Find(&results).Error

	return
}

// GetBatchFromGtrid 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromGtrid(gtrids [][]byte) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`gtrid` IN (?)", gtrids).Find(&results).Error

	return
}

// GetFromBqual 通过bqual获取内容
func (obj *_AllVirtualGlobalTransactionMgr) GetFromBqual(bqual []byte) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`bqual` = ?", bqual).Find(&results).Error

	return
}

// GetBatchFromBqual 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromBqual(bquals [][]byte) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`bqual` IN (?)", bquals).Find(&results).Error

	return
}

// GetFromFormatID 通过format_id获取内容
func (obj *_AllVirtualGlobalTransactionMgr) GetFromFormatID(formatID int64) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`format_id` = ?", formatID).Find(&results).Error

	return
}

// GetBatchFromFormatID 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromFormatID(formatIDs []int64) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`format_id` IN (?)", formatIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualGlobalTransactionMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualGlobalTransactionMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTransID 通过trans_id获取内容
func (obj *_AllVirtualGlobalTransactionMgr) GetFromTransID(transID string) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`trans_id` = ?", transID).Find(&results).Error

	return
}

// GetBatchFromTransID 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromTransID(transIDs []string) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`trans_id` IN (?)", transIDs).Find(&results).Error

	return
}

// GetFromCoordinator 通过coordinator获取内容
func (obj *_AllVirtualGlobalTransactionMgr) GetFromCoordinator(coordinator string) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`coordinator` = ?", coordinator).Find(&results).Error

	return
}

// GetBatchFromCoordinator 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromCoordinator(coordinators []string) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`coordinator` IN (?)", coordinators).Find(&results).Error

	return
}

// GetFromSchedulerIP 通过scheduler_ip获取内容
func (obj *_AllVirtualGlobalTransactionMgr) GetFromSchedulerIP(schedulerIP string) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`scheduler_ip` = ?", schedulerIP).Find(&results).Error

	return
}

// GetBatchFromSchedulerIP 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromSchedulerIP(schedulerIPs []string) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`scheduler_ip` IN (?)", schedulerIPs).Find(&results).Error

	return
}

// GetFromSchedulerPort 通过scheduler_port获取内容
func (obj *_AllVirtualGlobalTransactionMgr) GetFromSchedulerPort(schedulerPort int64) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`scheduler_port` = ?", schedulerPort).Find(&results).Error

	return
}

// GetBatchFromSchedulerPort 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromSchedulerPort(schedulerPorts []int64) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`scheduler_port` IN (?)", schedulerPorts).Find(&results).Error

	return
}

// GetFromIsReadonly 通过is_readonly获取内容
func (obj *_AllVirtualGlobalTransactionMgr) GetFromIsReadonly(isReadonly int8) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`is_readonly` = ?", isReadonly).Find(&results).Error

	return
}

// GetBatchFromIsReadonly 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromIsReadonly(isReadonlys []int8) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`is_readonly` IN (?)", isReadonlys).Find(&results).Error

	return
}

// GetFromState 通过state获取内容
func (obj *_AllVirtualGlobalTransactionMgr) GetFromState(state int64) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromState(states []int64) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromEndFlag 通过end_flag获取内容
func (obj *_AllVirtualGlobalTransactionMgr) GetFromEndFlag(endFlag int64) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`end_flag` = ?", endFlag).Find(&results).Error

	return
}

// GetBatchFromEndFlag 批量查找
func (obj *_AllVirtualGlobalTransactionMgr) GetBatchFromEndFlag(endFlags []int64) (results []*AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`end_flag` IN (?)", endFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualGlobalTransactionMgr) FetchByPrimaryKey(tenantID int64, gtrid []byte, bqual []byte, formatID int64) (result AllVirtualGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualGlobalTransaction{}).Where("`tenant_id` = ? AND `gtrid` = ? AND `bqual` = ? AND `format_id` = ?", tenantID, gtrid, bqual, formatID).First(&result).Error

	return
}
