package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualTenantParameterInfoMgr struct {
	*_BaseMgr
}

// AllVirtualTenantParameterInfoMgr open func
func AllVirtualTenantParameterInfoMgr(db *gorm.DB) *_AllVirtualTenantParameterInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantParameterInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantParameterInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_parameter_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantParameterInfoMgr) GetTableName() string {
	return "__all_virtual_tenant_parameter_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantParameterInfoMgr) Reset() *_AllVirtualTenantParameterInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantParameterInfoMgr) Get() (result AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantParameterInfoMgr) Gets() (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantParameterInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithZone zone获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithSvrType svr_type获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithSvrType(svrType string) Option {
	return optionFunc(func(o *options) { o.query["svr_type"] = svrType })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithName name获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDataType data_type获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithDataType(dataType string) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithInfo info获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithSection section获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithSection(section string) Option {
	return optionFunc(func(o *options) { o.query["section"] = section })
}

// WithScope scope获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithScope(scope string) Option {
	return optionFunc(func(o *options) { o.query["scope"] = scope })
}

// WithSource source获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithEditLevel edit_level获取
func (obj *_AllVirtualTenantParameterInfoMgr) WithEditLevel(editLevel string) Option {
	return optionFunc(func(o *options) { o.query["edit_level"] = editLevel })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantParameterInfoMgr) GetByOption(opts ...Option) (result AllVirtualTenantParameterInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantParameterInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantParameterInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantParameterInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantParameterInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where(options.query)
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
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromZone(zone string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromZone(zones []string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromSvrType 通过svr_type获取内容
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromSvrType(svrType string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`svr_type` = ?", svrType).Find(&results).Error

	return
}

// GetBatchFromSvrType 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromSvrType(svrTypes []string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`svr_type` IN (?)", svrTypes).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromName(name string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromName(names []string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromDataType(dataType string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromDataType(dataTypes []string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromValue(value string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromValue(values []string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromInfo(info string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromInfo(infos []string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromSection 通过section获取内容
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromSection(section string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`section` = ?", section).Find(&results).Error

	return
}

// GetBatchFromSection 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromSection(sections []string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`section` IN (?)", sections).Find(&results).Error

	return
}

// GetFromScope 通过scope获取内容
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromScope(scope string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`scope` = ?", scope).Find(&results).Error

	return
}

// GetBatchFromScope 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromScope(scopes []string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`scope` IN (?)", scopes).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromSource(source string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromSource(sources []string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromEditLevel 通过edit_level获取内容
func (obj *_AllVirtualTenantParameterInfoMgr) GetFromEditLevel(editLevel string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`edit_level` = ?", editLevel).Find(&results).Error

	return
}

// GetBatchFromEditLevel 批量查找
func (obj *_AllVirtualTenantParameterInfoMgr) GetBatchFromEditLevel(editLevels []string) (results []*AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`edit_level` IN (?)", editLevels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTenantParameterInfoMgr) FetchByPrimaryKey(tenantID int64, zone string, svrType string, svrIP string, svrPort int64, name string) (result AllVirtualTenantParameterInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterInfo{}).Where("`tenant_id` = ? AND `zone` = ? AND `svr_type` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `name` = ?", tenantID, zone, svrType, svrIP, svrPort, name).First(&result).Error

	return
}
