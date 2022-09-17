package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualDagWarningHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualDagWarningHistoryMgr open func
func AllVirtualDagWarningHistoryMgr(db *gorm.DB) *_AllVirtualDagWarningHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDagWarningHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDagWarningHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_dag_warning_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDagWarningHistoryMgr) GetTableName() string {
	return "__all_virtual_dag_warning_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDagWarningHistoryMgr) Reset() *_AllVirtualDagWarningHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDagWarningHistoryMgr) Get() (result AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDagWarningHistoryMgr) Gets() (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDagWarningHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualDagWarningHistoryMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualDagWarningHistoryMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualDagWarningHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTaskID task_id获取
func (obj *_AllVirtualDagWarningHistoryMgr) WithTaskID(taskID string) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithModule module获取
func (obj *_AllVirtualDagWarningHistoryMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

// WithType type获取
func (obj *_AllVirtualDagWarningHistoryMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithRet ret获取
func (obj *_AllVirtualDagWarningHistoryMgr) WithRet(ret string) Option {
	return optionFunc(func(o *options) { o.query["ret"] = ret })
}

// WithStatus status获取
func (obj *_AllVirtualDagWarningHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualDagWarningHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualDagWarningHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithWarningInfo warning_info获取
func (obj *_AllVirtualDagWarningHistoryMgr) WithWarningInfo(warningInfo string) Option {
	return optionFunc(func(o *options) { o.query["warning_info"] = warningInfo })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDagWarningHistoryMgr) GetByOption(opts ...Option) (result AllVirtualDagWarningHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDagWarningHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualDagWarningHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDagWarningHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDagWarningHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where(options.query)
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
func (obj *_AllVirtualDagWarningHistoryMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualDagWarningHistoryMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualDagWarningHistoryMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualDagWarningHistoryMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualDagWarningHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDagWarningHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容
func (obj *_AllVirtualDagWarningHistoryMgr) GetFromTaskID(taskID string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_AllVirtualDagWarningHistoryMgr) GetBatchFromTaskID(taskIDs []string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromModule 通过module获取内容
func (obj *_AllVirtualDagWarningHistoryMgr) GetFromModule(module string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`module` = ?", module).Find(&results).Error

	return
}

// GetBatchFromModule 批量查找
func (obj *_AllVirtualDagWarningHistoryMgr) GetBatchFromModule(modules []string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`module` IN (?)", modules).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualDagWarningHistoryMgr) GetFromType(_type string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualDagWarningHistoryMgr) GetBatchFromType(_types []string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromRet 通过ret获取内容
func (obj *_AllVirtualDagWarningHistoryMgr) GetFromRet(ret string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`ret` = ?", ret).Find(&results).Error

	return
}

// GetBatchFromRet 批量查找
func (obj *_AllVirtualDagWarningHistoryMgr) GetBatchFromRet(rets []string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`ret` IN (?)", rets).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualDagWarningHistoryMgr) GetFromStatus(status string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualDagWarningHistoryMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualDagWarningHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualDagWarningHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualDagWarningHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualDagWarningHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromWarningInfo 通过warning_info获取内容
func (obj *_AllVirtualDagWarningHistoryMgr) GetFromWarningInfo(warningInfo string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`warning_info` = ?", warningInfo).Find(&results).Error

	return
}

// GetBatchFromWarningInfo 批量查找
func (obj *_AllVirtualDagWarningHistoryMgr) GetBatchFromWarningInfo(warningInfos []string) (results []*AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`warning_info` IN (?)", warningInfos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDagWarningHistoryMgr) FetchByPrimaryKey(svrIP string, svrPort int64, tenantID int64, taskID string) (result AllVirtualDagWarningHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDagWarningHistory{}).Where("`svr_ip` = ? AND `svr_port` = ? AND `tenant_id` = ? AND `task_id` = ?", svrIP, svrPort, tenantID, taskID).First(&result).Error

	return
}
