package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualLongOpsStatusMgr struct {
	*_BaseMgr
}

// AllVirtualLongOpsStatusMgr open func
func AllVirtualLongOpsStatusMgr(db *gorm.DB) *_AllVirtualLongOpsStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualLongOpsStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualLongOpsStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_long_ops_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualLongOpsStatusMgr) GetTableName() string {
	return "__all_virtual_long_ops_status"
}

// Reset 重置gorm会话
func (obj *_AllVirtualLongOpsStatusMgr) Reset() *_AllVirtualLongOpsStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualLongOpsStatusMgr) Get() (result AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualLongOpsStatusMgr) Gets() (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualLongOpsStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualLongOpsStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSid sid获取
func (obj *_AllVirtualLongOpsStatusMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["sid"] = sid })
}

// WithOpName op_name获取
func (obj *_AllVirtualLongOpsStatusMgr) WithOpName(opName string) Option {
	return optionFunc(func(o *options) { o.query["op_name"] = opName })
}

// WithTarget target获取
func (obj *_AllVirtualLongOpsStatusMgr) WithTarget(target string) Option {
	return optionFunc(func(o *options) { o.query["target"] = target })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualLongOpsStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualLongOpsStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithStartTime start_time获取
func (obj *_AllVirtualLongOpsStatusMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithFinishTime finish_time获取
func (obj *_AllVirtualLongOpsStatusMgr) WithFinishTime(finishTime int64) Option {
	return optionFunc(func(o *options) { o.query["finish_time"] = finishTime })
}

// WithElapsedTime elapsed_time获取
func (obj *_AllVirtualLongOpsStatusMgr) WithElapsedTime(elapsedTime int64) Option {
	return optionFunc(func(o *options) { o.query["elapsed_time"] = elapsedTime })
}

// WithRemainingTime remaining_time获取
func (obj *_AllVirtualLongOpsStatusMgr) WithRemainingTime(remainingTime int64) Option {
	return optionFunc(func(o *options) { o.query["remaining_time"] = remainingTime })
}

// WithLastUpdateTime last_update_time获取
func (obj *_AllVirtualLongOpsStatusMgr) WithLastUpdateTime(lastUpdateTime int64) Option {
	return optionFunc(func(o *options) { o.query["last_update_time"] = lastUpdateTime })
}

// WithPercentage percentage获取
func (obj *_AllVirtualLongOpsStatusMgr) WithPercentage(percentage int64) Option {
	return optionFunc(func(o *options) { o.query["percentage"] = percentage })
}

// WithMessage message获取
func (obj *_AllVirtualLongOpsStatusMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualLongOpsStatusMgr) GetByOption(opts ...Option) (result AllVirtualLongOpsStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualLongOpsStatusMgr) GetByOptions(opts ...Option) (results []*AllVirtualLongOpsStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualLongOpsStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualLongOpsStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where(options.query)
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
func (obj *_AllVirtualLongOpsStatusMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSid 通过sid获取内容
func (obj *_AllVirtualLongOpsStatusMgr) GetFromSid(sid int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`sid` = ?", sid).Find(&results).Error

	return
}

// GetBatchFromSid 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromSid(sids []int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`sid` IN (?)", sids).Find(&results).Error

	return
}

// GetFromOpName 通过op_name获取内容
func (obj *_AllVirtualLongOpsStatusMgr) GetFromOpName(opName string) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`op_name` = ?", opName).Find(&results).Error

	return
}

// GetBatchFromOpName 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromOpName(opNames []string) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`op_name` IN (?)", opNames).Find(&results).Error

	return
}

// GetFromTarget 通过target获取内容
func (obj *_AllVirtualLongOpsStatusMgr) GetFromTarget(target string) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`target` = ?", target).Find(&results).Error

	return
}

// GetBatchFromTarget 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromTarget(targets []string) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`target` IN (?)", targets).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualLongOpsStatusMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualLongOpsStatusMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllVirtualLongOpsStatusMgr) GetFromStartTime(startTime int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromStartTime(startTimes []int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromFinishTime 通过finish_time获取内容
func (obj *_AllVirtualLongOpsStatusMgr) GetFromFinishTime(finishTime int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`finish_time` = ?", finishTime).Find(&results).Error

	return
}

// GetBatchFromFinishTime 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromFinishTime(finishTimes []int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`finish_time` IN (?)", finishTimes).Find(&results).Error

	return
}

// GetFromElapsedTime 通过elapsed_time获取内容
func (obj *_AllVirtualLongOpsStatusMgr) GetFromElapsedTime(elapsedTime int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`elapsed_time` = ?", elapsedTime).Find(&results).Error

	return
}

// GetBatchFromElapsedTime 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromElapsedTime(elapsedTimes []int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`elapsed_time` IN (?)", elapsedTimes).Find(&results).Error

	return
}

// GetFromRemainingTime 通过remaining_time获取内容
func (obj *_AllVirtualLongOpsStatusMgr) GetFromRemainingTime(remainingTime int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`remaining_time` = ?", remainingTime).Find(&results).Error

	return
}

// GetBatchFromRemainingTime 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromRemainingTime(remainingTimes []int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`remaining_time` IN (?)", remainingTimes).Find(&results).Error

	return
}

// GetFromLastUpdateTime 通过last_update_time获取内容
func (obj *_AllVirtualLongOpsStatusMgr) GetFromLastUpdateTime(lastUpdateTime int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`last_update_time` = ?", lastUpdateTime).Find(&results).Error

	return
}

// GetBatchFromLastUpdateTime 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromLastUpdateTime(lastUpdateTimes []int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`last_update_time` IN (?)", lastUpdateTimes).Find(&results).Error

	return
}

// GetFromPercentage 通过percentage获取内容
func (obj *_AllVirtualLongOpsStatusMgr) GetFromPercentage(percentage int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`percentage` = ?", percentage).Find(&results).Error

	return
}

// GetBatchFromPercentage 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromPercentage(percentages []int64) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`percentage` IN (?)", percentages).Find(&results).Error

	return
}

// GetFromMessage 通过message获取内容
func (obj *_AllVirtualLongOpsStatusMgr) GetFromMessage(message string) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`message` = ?", message).Find(&results).Error

	return
}

// GetBatchFromMessage 批量查找
func (obj *_AllVirtualLongOpsStatusMgr) GetBatchFromMessage(messages []string) (results []*AllVirtualLongOpsStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLongOpsStatus{}).Where("`message` IN (?)", messages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
