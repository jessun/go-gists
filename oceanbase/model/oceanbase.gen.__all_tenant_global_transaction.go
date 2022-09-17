package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantGlobalTransactionMgr struct {
	*_BaseMgr
}

// AllTenantGlobalTransactionMgr open func
func AllTenantGlobalTransactionMgr(db *gorm.DB) *_AllTenantGlobalTransactionMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantGlobalTransactionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantGlobalTransactionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_global_transaction"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantGlobalTransactionMgr) GetTableName() string {
	return "__all_tenant_global_transaction"
}

// Reset 重置gorm会话
func (obj *_AllTenantGlobalTransactionMgr) Reset() *_AllTenantGlobalTransactionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantGlobalTransactionMgr) Get() (result AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantGlobalTransactionMgr) Gets() (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantGlobalTransactionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantGlobalTransactionMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantGlobalTransactionMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithGtrid gtrid获取
func (obj *_AllTenantGlobalTransactionMgr) WithGtrid(gtrid []byte) Option {
	return optionFunc(func(o *options) { o.query["gtrid"] = gtrid })
}

// WithBqual bqual获取
func (obj *_AllTenantGlobalTransactionMgr) WithBqual(bqual []byte) Option {
	return optionFunc(func(o *options) { o.query["bqual"] = bqual })
}

// WithFormatID format_id获取
func (obj *_AllTenantGlobalTransactionMgr) WithFormatID(formatID int64) Option {
	return optionFunc(func(o *options) { o.query["format_id"] = formatID })
}

// WithTransID trans_id获取
func (obj *_AllTenantGlobalTransactionMgr) WithTransID(transID string) Option {
	return optionFunc(func(o *options) { o.query["trans_id"] = transID })
}

// WithCoordinator coordinator获取
func (obj *_AllTenantGlobalTransactionMgr) WithCoordinator(coordinator string) Option {
	return optionFunc(func(o *options) { o.query["coordinator"] = coordinator })
}

// WithSchedulerIP scheduler_ip获取
func (obj *_AllTenantGlobalTransactionMgr) WithSchedulerIP(schedulerIP string) Option {
	return optionFunc(func(o *options) { o.query["scheduler_ip"] = schedulerIP })
}

// WithSchedulerPort scheduler_port获取
func (obj *_AllTenantGlobalTransactionMgr) WithSchedulerPort(schedulerPort int64) Option {
	return optionFunc(func(o *options) { o.query["scheduler_port"] = schedulerPort })
}

// WithIsReadonly is_readonly获取
func (obj *_AllTenantGlobalTransactionMgr) WithIsReadonly(isReadonly int8) Option {
	return optionFunc(func(o *options) { o.query["is_readonly"] = isReadonly })
}

// WithState state获取
func (obj *_AllTenantGlobalTransactionMgr) WithState(state int64) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithEndFlag end_flag获取
func (obj *_AllTenantGlobalTransactionMgr) WithEndFlag(endFlag int64) Option {
	return optionFunc(func(o *options) { o.query["end_flag"] = endFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantGlobalTransactionMgr) GetByOption(opts ...Option) (result AllTenantGlobalTransaction, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantGlobalTransactionMgr) GetByOptions(opts ...Option) (results []*AllTenantGlobalTransaction, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantGlobalTransactionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantGlobalTransaction, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where(options.query)
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

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllTenantGlobalTransactionMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantGlobalTransactionMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantGlobalTransactionMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantGlobalTransactionMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromGtrid 通过gtrid获取内容
func (obj *_AllTenantGlobalTransactionMgr) GetFromGtrid(gtrid []byte) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`gtrid` = ?", gtrid).Find(&results).Error

	return
}

// GetBatchFromGtrid 批量查找
func (obj *_AllTenantGlobalTransactionMgr) GetBatchFromGtrid(gtrids [][]byte) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`gtrid` IN (?)", gtrids).Find(&results).Error

	return
}

// GetFromBqual 通过bqual获取内容
func (obj *_AllTenantGlobalTransactionMgr) GetFromBqual(bqual []byte) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`bqual` = ?", bqual).Find(&results).Error

	return
}

// GetBatchFromBqual 批量查找
func (obj *_AllTenantGlobalTransactionMgr) GetBatchFromBqual(bquals [][]byte) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`bqual` IN (?)", bquals).Find(&results).Error

	return
}

// GetFromFormatID 通过format_id获取内容
func (obj *_AllTenantGlobalTransactionMgr) GetFromFormatID(formatID int64) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`format_id` = ?", formatID).Find(&results).Error

	return
}

// GetBatchFromFormatID 批量查找
func (obj *_AllTenantGlobalTransactionMgr) GetBatchFromFormatID(formatIDs []int64) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`format_id` IN (?)", formatIDs).Find(&results).Error

	return
}

// GetFromTransID 通过trans_id获取内容
func (obj *_AllTenantGlobalTransactionMgr) GetFromTransID(transID string) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`trans_id` = ?", transID).Find(&results).Error

	return
}

// GetBatchFromTransID 批量查找
func (obj *_AllTenantGlobalTransactionMgr) GetBatchFromTransID(transIDs []string) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`trans_id` IN (?)", transIDs).Find(&results).Error

	return
}

// GetFromCoordinator 通过coordinator获取内容
func (obj *_AllTenantGlobalTransactionMgr) GetFromCoordinator(coordinator string) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`coordinator` = ?", coordinator).Find(&results).Error

	return
}

// GetBatchFromCoordinator 批量查找
func (obj *_AllTenantGlobalTransactionMgr) GetBatchFromCoordinator(coordinators []string) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`coordinator` IN (?)", coordinators).Find(&results).Error

	return
}

// GetFromSchedulerIP 通过scheduler_ip获取内容
func (obj *_AllTenantGlobalTransactionMgr) GetFromSchedulerIP(schedulerIP string) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`scheduler_ip` = ?", schedulerIP).Find(&results).Error

	return
}

// GetBatchFromSchedulerIP 批量查找
func (obj *_AllTenantGlobalTransactionMgr) GetBatchFromSchedulerIP(schedulerIPs []string) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`scheduler_ip` IN (?)", schedulerIPs).Find(&results).Error

	return
}

// GetFromSchedulerPort 通过scheduler_port获取内容
func (obj *_AllTenantGlobalTransactionMgr) GetFromSchedulerPort(schedulerPort int64) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`scheduler_port` = ?", schedulerPort).Find(&results).Error

	return
}

// GetBatchFromSchedulerPort 批量查找
func (obj *_AllTenantGlobalTransactionMgr) GetBatchFromSchedulerPort(schedulerPorts []int64) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`scheduler_port` IN (?)", schedulerPorts).Find(&results).Error

	return
}

// GetFromIsReadonly 通过is_readonly获取内容
func (obj *_AllTenantGlobalTransactionMgr) GetFromIsReadonly(isReadonly int8) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`is_readonly` = ?", isReadonly).Find(&results).Error

	return
}

// GetBatchFromIsReadonly 批量查找
func (obj *_AllTenantGlobalTransactionMgr) GetBatchFromIsReadonly(isReadonlys []int8) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`is_readonly` IN (?)", isReadonlys).Find(&results).Error

	return
}

// GetFromState 通过state获取内容
func (obj *_AllTenantGlobalTransactionMgr) GetFromState(state int64) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找
func (obj *_AllTenantGlobalTransactionMgr) GetBatchFromState(states []int64) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromEndFlag 通过end_flag获取内容
func (obj *_AllTenantGlobalTransactionMgr) GetFromEndFlag(endFlag int64) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`end_flag` = ?", endFlag).Find(&results).Error

	return
}

// GetBatchFromEndFlag 批量查找
func (obj *_AllTenantGlobalTransactionMgr) GetBatchFromEndFlag(endFlags []int64) (results []*AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`end_flag` IN (?)", endFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantGlobalTransactionMgr) FetchByPrimaryKey(gtrid []byte, bqual []byte, formatID int64) (result AllTenantGlobalTransaction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGlobalTransaction{}).Where("`gtrid` = ? AND `bqual` = ? AND `format_id` = ?", gtrid, bqual, formatID).First(&result).Error

	return
}
