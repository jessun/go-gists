package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualRaidStatMgr struct {
	*_BaseMgr
}

// AllVirtualRaidStatMgr open func
func AllVirtualRaidStatMgr(db *gorm.DB) *_AllVirtualRaidStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRaidStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRaidStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_raid_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRaidStatMgr) GetTableName() string {
	return "__all_virtual_raid_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRaidStatMgr) Reset() *_AllVirtualRaidStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRaidStatMgr) Get() (result AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRaidStatMgr) Gets() (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRaidStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualRaidStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualRaidStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithDiskIDx disk_idx获取
func (obj *_AllVirtualRaidStatMgr) WithDiskIDx(diskIDx int64) Option {
	return optionFunc(func(o *options) { o.query["disk_idx"] = diskIDx })
}

// WithInstallSeq install_seq获取
func (obj *_AllVirtualRaidStatMgr) WithInstallSeq(installSeq int64) Option {
	return optionFunc(func(o *options) { o.query["install_seq"] = installSeq })
}

// WithDataNum data_num获取
func (obj *_AllVirtualRaidStatMgr) WithDataNum(dataNum int64) Option {
	return optionFunc(func(o *options) { o.query["data_num"] = dataNum })
}

// WithParityNum parity_num获取
func (obj *_AllVirtualRaidStatMgr) WithParityNum(parityNum int64) Option {
	return optionFunc(func(o *options) { o.query["parity_num"] = parityNum })
}

// WithCreateTs create_ts获取
func (obj *_AllVirtualRaidStatMgr) WithCreateTs(createTs int64) Option {
	return optionFunc(func(o *options) { o.query["create_ts"] = createTs })
}

// WithFinishTs finish_ts获取
func (obj *_AllVirtualRaidStatMgr) WithFinishTs(finishTs int64) Option {
	return optionFunc(func(o *options) { o.query["finish_ts"] = finishTs })
}

// WithAliasName alias_name获取
func (obj *_AllVirtualRaidStatMgr) WithAliasName(aliasName string) Option {
	return optionFunc(func(o *options) { o.query["alias_name"] = aliasName })
}

// WithStatus status获取
func (obj *_AllVirtualRaidStatMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithPercent percent获取
func (obj *_AllVirtualRaidStatMgr) WithPercent(percent int64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRaidStatMgr) GetByOption(opts ...Option) (result AllVirtualRaidStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRaidStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualRaidStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRaidStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRaidStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where(options.query)
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
func (obj *_AllVirtualRaidStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualRaidStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualRaidStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualRaidStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromDiskIDx 通过disk_idx获取内容
func (obj *_AllVirtualRaidStatMgr) GetFromDiskIDx(diskIDx int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`disk_idx` = ?", diskIDx).Find(&results).Error

	return
}

// GetBatchFromDiskIDx 批量查找
func (obj *_AllVirtualRaidStatMgr) GetBatchFromDiskIDx(diskIDxs []int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`disk_idx` IN (?)", diskIDxs).Find(&results).Error

	return
}

// GetFromInstallSeq 通过install_seq获取内容
func (obj *_AllVirtualRaidStatMgr) GetFromInstallSeq(installSeq int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`install_seq` = ?", installSeq).Find(&results).Error

	return
}

// GetBatchFromInstallSeq 批量查找
func (obj *_AllVirtualRaidStatMgr) GetBatchFromInstallSeq(installSeqs []int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`install_seq` IN (?)", installSeqs).Find(&results).Error

	return
}

// GetFromDataNum 通过data_num获取内容
func (obj *_AllVirtualRaidStatMgr) GetFromDataNum(dataNum int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`data_num` = ?", dataNum).Find(&results).Error

	return
}

// GetBatchFromDataNum 批量查找
func (obj *_AllVirtualRaidStatMgr) GetBatchFromDataNum(dataNums []int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`data_num` IN (?)", dataNums).Find(&results).Error

	return
}

// GetFromParityNum 通过parity_num获取内容
func (obj *_AllVirtualRaidStatMgr) GetFromParityNum(parityNum int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`parity_num` = ?", parityNum).Find(&results).Error

	return
}

// GetBatchFromParityNum 批量查找
func (obj *_AllVirtualRaidStatMgr) GetBatchFromParityNum(parityNums []int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`parity_num` IN (?)", parityNums).Find(&results).Error

	return
}

// GetFromCreateTs 通过create_ts获取内容
func (obj *_AllVirtualRaidStatMgr) GetFromCreateTs(createTs int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`create_ts` = ?", createTs).Find(&results).Error

	return
}

// GetBatchFromCreateTs 批量查找
func (obj *_AllVirtualRaidStatMgr) GetBatchFromCreateTs(createTss []int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`create_ts` IN (?)", createTss).Find(&results).Error

	return
}

// GetFromFinishTs 通过finish_ts获取内容
func (obj *_AllVirtualRaidStatMgr) GetFromFinishTs(finishTs int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`finish_ts` = ?", finishTs).Find(&results).Error

	return
}

// GetBatchFromFinishTs 批量查找
func (obj *_AllVirtualRaidStatMgr) GetBatchFromFinishTs(finishTss []int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`finish_ts` IN (?)", finishTss).Find(&results).Error

	return
}

// GetFromAliasName 通过alias_name获取内容
func (obj *_AllVirtualRaidStatMgr) GetFromAliasName(aliasName string) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`alias_name` = ?", aliasName).Find(&results).Error

	return
}

// GetBatchFromAliasName 批量查找
func (obj *_AllVirtualRaidStatMgr) GetBatchFromAliasName(aliasNames []string) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`alias_name` IN (?)", aliasNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualRaidStatMgr) GetFromStatus(status string) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualRaidStatMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容
func (obj *_AllVirtualRaidStatMgr) GetFromPercent(percent int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找
func (obj *_AllVirtualRaidStatMgr) GetBatchFromPercent(percents []int64) (results []*AllVirtualRaidStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRaidStat{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
