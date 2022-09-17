package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualSysParameterStatMgr struct {
	*_BaseMgr
}

// AllVirtualSysParameterStatMgr open func
func AllVirtualSysParameterStatMgr(db *gorm.DB) *_AllVirtualSysParameterStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSysParameterStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSysParameterStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sys_parameter_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSysParameterStatMgr) GetTableName() string {
	return "__all_virtual_sys_parameter_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSysParameterStatMgr) Reset() *_AllVirtualSysParameterStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSysParameterStatMgr) Get() (result AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSysParameterStatMgr) Gets() (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSysParameterStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithZone zone获取
func (obj *_AllVirtualSysParameterStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithSvrType svr_type获取
func (obj *_AllVirtualSysParameterStatMgr) WithSvrType(svrType string) Option {
	return optionFunc(func(o *options) { o.query["svr_type"] = svrType })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualSysParameterStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualSysParameterStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithName name获取
func (obj *_AllVirtualSysParameterStatMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDataType data_type获取
func (obj *_AllVirtualSysParameterStatMgr) WithDataType(dataType string) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_AllVirtualSysParameterStatMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithValueStrict value_strict获取
func (obj *_AllVirtualSysParameterStatMgr) WithValueStrict(valueStrict string) Option {
	return optionFunc(func(o *options) { o.query["value_strict"] = valueStrict })
}

// WithInfo info获取
func (obj *_AllVirtualSysParameterStatMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithNeedReboot need_reboot获取
func (obj *_AllVirtualSysParameterStatMgr) WithNeedReboot(needReboot int64) Option {
	return optionFunc(func(o *options) { o.query["need_reboot"] = needReboot })
}

// WithSection section获取
func (obj *_AllVirtualSysParameterStatMgr) WithSection(section string) Option {
	return optionFunc(func(o *options) { o.query["section"] = section })
}

// WithVisibleLevel visible_level获取
func (obj *_AllVirtualSysParameterStatMgr) WithVisibleLevel(visibleLevel string) Option {
	return optionFunc(func(o *options) { o.query["visible_level"] = visibleLevel })
}

// WithScope scope获取
func (obj *_AllVirtualSysParameterStatMgr) WithScope(scope string) Option {
	return optionFunc(func(o *options) { o.query["scope"] = scope })
}

// WithSource source获取
func (obj *_AllVirtualSysParameterStatMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithEditLevel edit_level获取
func (obj *_AllVirtualSysParameterStatMgr) WithEditLevel(editLevel string) Option {
	return optionFunc(func(o *options) { o.query["edit_level"] = editLevel })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSysParameterStatMgr) GetByOption(opts ...Option) (result AllVirtualSysParameterStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSysParameterStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualSysParameterStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSysParameterStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSysParameterStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where(options.query)
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
func (obj *_AllVirtualSysParameterStatMgr) GetFromZone(zone string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromSvrType 通过svr_type获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromSvrType(svrType string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`svr_type` = ?", svrType).Find(&results).Error

	return
}

// GetBatchFromSvrType 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromSvrType(svrTypes []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`svr_type` IN (?)", svrTypes).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromName(name string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromName(names []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromDataType(dataType string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromDataType(dataTypes []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromValue(value string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromValue(values []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromValueStrict 通过value_strict获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromValueStrict(valueStrict string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`value_strict` = ?", valueStrict).Find(&results).Error

	return
}

// GetBatchFromValueStrict 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromValueStrict(valueStricts []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`value_strict` IN (?)", valueStricts).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromInfo(info string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromInfo(infos []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromNeedReboot 通过need_reboot获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromNeedReboot(needReboot int64) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`need_reboot` = ?", needReboot).Find(&results).Error

	return
}

// GetBatchFromNeedReboot 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromNeedReboot(needReboots []int64) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`need_reboot` IN (?)", needReboots).Find(&results).Error

	return
}

// GetFromSection 通过section获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromSection(section string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`section` = ?", section).Find(&results).Error

	return
}

// GetBatchFromSection 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromSection(sections []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`section` IN (?)", sections).Find(&results).Error

	return
}

// GetFromVisibleLevel 通过visible_level获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromVisibleLevel(visibleLevel string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`visible_level` = ?", visibleLevel).Find(&results).Error

	return
}

// GetBatchFromVisibleLevel 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromVisibleLevel(visibleLevels []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`visible_level` IN (?)", visibleLevels).Find(&results).Error

	return
}

// GetFromScope 通过scope获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromScope(scope string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`scope` = ?", scope).Find(&results).Error

	return
}

// GetBatchFromScope 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromScope(scopes []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`scope` IN (?)", scopes).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromSource(source string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromSource(sources []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromEditLevel 通过edit_level获取内容
func (obj *_AllVirtualSysParameterStatMgr) GetFromEditLevel(editLevel string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`edit_level` = ?", editLevel).Find(&results).Error

	return
}

// GetBatchFromEditLevel 批量查找
func (obj *_AllVirtualSysParameterStatMgr) GetBatchFromEditLevel(editLevels []string) (results []*AllVirtualSysParameterStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysParameterStat{}).Where("`edit_level` IN (?)", editLevels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
