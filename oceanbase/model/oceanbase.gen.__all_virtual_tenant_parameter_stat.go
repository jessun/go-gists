package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTenantParameterStatMgr struct {
	*_BaseMgr
}

// AllVirtualTenantParameterStatMgr open func
func AllVirtualTenantParameterStatMgr(db *gorm.DB) *_AllVirtualTenantParameterStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantParameterStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantParameterStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_parameter_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantParameterStatMgr) GetTableName() string {
	return "__all_virtual_tenant_parameter_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantParameterStatMgr) Reset() *_AllVirtualTenantParameterStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantParameterStatMgr) Get() (result AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantParameterStatMgr) Gets() (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantParameterStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithZone zone获取
func (obj *_AllVirtualTenantParameterStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithSvrType svr_type获取
func (obj *_AllVirtualTenantParameterStatMgr) WithSvrType(svrType string) Option {
	return optionFunc(func(o *options) { o.query["svr_type"] = svrType })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTenantParameterStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTenantParameterStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithName name获取
func (obj *_AllVirtualTenantParameterStatMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDataType data_type获取
func (obj *_AllVirtualTenantParameterStatMgr) WithDataType(dataType string) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_AllVirtualTenantParameterStatMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithInfo info获取
func (obj *_AllVirtualTenantParameterStatMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithSection section获取
func (obj *_AllVirtualTenantParameterStatMgr) WithSection(section string) Option {
	return optionFunc(func(o *options) { o.query["section"] = section })
}

// WithScope scope获取
func (obj *_AllVirtualTenantParameterStatMgr) WithScope(scope string) Option {
	return optionFunc(func(o *options) { o.query["scope"] = scope })
}

// WithSource source获取
func (obj *_AllVirtualTenantParameterStatMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithEditLevel edit_level获取
func (obj *_AllVirtualTenantParameterStatMgr) WithEditLevel(editLevel string) Option {
	return optionFunc(func(o *options) { o.query["edit_level"] = editLevel })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantParameterStatMgr) GetByOption(opts ...Option) (result AllVirtualTenantParameterStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantParameterStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantParameterStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantParameterStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantParameterStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where(options.query)
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

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualTenantParameterStatMgr) GetFromZone(zone string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualTenantParameterStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromSvrType 通过svr_type获取内容
func (obj *_AllVirtualTenantParameterStatMgr) GetFromSvrType(svrType string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`svr_type` = ?", svrType).Find(&results).Error

	return
}

// GetBatchFromSvrType 批量查找
func (obj *_AllVirtualTenantParameterStatMgr) GetBatchFromSvrType(svrTypes []string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`svr_type` IN (?)", svrTypes).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualTenantParameterStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTenantParameterStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTenantParameterStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTenantParameterStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualTenantParameterStatMgr) GetFromName(name string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualTenantParameterStatMgr) GetBatchFromName(names []string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllVirtualTenantParameterStatMgr) GetFromDataType(dataType string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllVirtualTenantParameterStatMgr) GetBatchFromDataType(dataTypes []string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllVirtualTenantParameterStatMgr) GetFromValue(value string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllVirtualTenantParameterStatMgr) GetBatchFromValue(values []string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllVirtualTenantParameterStatMgr) GetFromInfo(info string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllVirtualTenantParameterStatMgr) GetBatchFromInfo(infos []string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromSection 通过section获取内容
func (obj *_AllVirtualTenantParameterStatMgr) GetFromSection(section string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`section` = ?", section).Find(&results).Error

	return
}

// GetBatchFromSection 批量查找
func (obj *_AllVirtualTenantParameterStatMgr) GetBatchFromSection(sections []string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`section` IN (?)", sections).Find(&results).Error

	return
}

// GetFromScope 通过scope获取内容
func (obj *_AllVirtualTenantParameterStatMgr) GetFromScope(scope string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`scope` = ?", scope).Find(&results).Error

	return
}

// GetBatchFromScope 批量查找
func (obj *_AllVirtualTenantParameterStatMgr) GetBatchFromScope(scopes []string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`scope` IN (?)", scopes).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_AllVirtualTenantParameterStatMgr) GetFromSource(source string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_AllVirtualTenantParameterStatMgr) GetBatchFromSource(sources []string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromEditLevel 通过edit_level获取内容
func (obj *_AllVirtualTenantParameterStatMgr) GetFromEditLevel(editLevel string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`edit_level` = ?", editLevel).Find(&results).Error

	return
}

// GetBatchFromEditLevel 批量查找
func (obj *_AllVirtualTenantParameterStatMgr) GetBatchFromEditLevel(editLevels []string) (results []*AllVirtualTenantParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantParameterStat{}).Where("`edit_level` IN (?)", editLevels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
