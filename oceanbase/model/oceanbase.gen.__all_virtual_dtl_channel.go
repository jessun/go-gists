package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualDtlChannelMgr struct {
	*_BaseMgr
}

// AllVirtualDtlChannelMgr open func
func AllVirtualDtlChannelMgr(db *gorm.DB) *_AllVirtualDtlChannelMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDtlChannelMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDtlChannelMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_dtl_channel"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDtlChannelMgr) GetTableName() string {
	return "__all_virtual_dtl_channel"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDtlChannelMgr) Reset() *_AllVirtualDtlChannelMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDtlChannelMgr) Get() (result AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDtlChannelMgr) Gets() (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDtlChannelMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualDtlChannelMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualDtlChannelMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithChannelID channel_id获取
func (obj *_AllVirtualDtlChannelMgr) WithChannelID(channelID int64) Option {
	return optionFunc(func(o *options) { o.query["channel_id"] = channelID })
}

// WithOpID op_id获取
func (obj *_AllVirtualDtlChannelMgr) WithOpID(opID int64) Option {
	return optionFunc(func(o *options) { o.query["op_id"] = opID })
}

// WithPeerID peer_id获取
func (obj *_AllVirtualDtlChannelMgr) WithPeerID(peerID int64) Option {
	return optionFunc(func(o *options) { o.query["peer_id"] = peerID })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualDtlChannelMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIsLocal is_local获取
func (obj *_AllVirtualDtlChannelMgr) WithIsLocal(isLocal int8) Option {
	return optionFunc(func(o *options) { o.query["is_local"] = isLocal })
}

// WithIsData is_data获取
func (obj *_AllVirtualDtlChannelMgr) WithIsData(isData int8) Option {
	return optionFunc(func(o *options) { o.query["is_data"] = isData })
}

// WithIsTransmit is_transmit获取
func (obj *_AllVirtualDtlChannelMgr) WithIsTransmit(isTransmit int8) Option {
	return optionFunc(func(o *options) { o.query["is_transmit"] = isTransmit })
}

// WithAllocBufferCnt alloc_buffer_cnt获取
func (obj *_AllVirtualDtlChannelMgr) WithAllocBufferCnt(allocBufferCnt int64) Option {
	return optionFunc(func(o *options) { o.query["alloc_buffer_cnt"] = allocBufferCnt })
}

// WithFreeBufferCnt free_buffer_cnt获取
func (obj *_AllVirtualDtlChannelMgr) WithFreeBufferCnt(freeBufferCnt int64) Option {
	return optionFunc(func(o *options) { o.query["free_buffer_cnt"] = freeBufferCnt })
}

// WithSendBufferCnt send_buffer_cnt获取
func (obj *_AllVirtualDtlChannelMgr) WithSendBufferCnt(sendBufferCnt int64) Option {
	return optionFunc(func(o *options) { o.query["send_buffer_cnt"] = sendBufferCnt })
}

// WithRecvBufferCnt recv_buffer_cnt获取
func (obj *_AllVirtualDtlChannelMgr) WithRecvBufferCnt(recvBufferCnt int64) Option {
	return optionFunc(func(o *options) { o.query["recv_buffer_cnt"] = recvBufferCnt })
}

// WithProcessedBufferCnt processed_buffer_cnt获取
func (obj *_AllVirtualDtlChannelMgr) WithProcessedBufferCnt(processedBufferCnt int64) Option {
	return optionFunc(func(o *options) { o.query["processed_buffer_cnt"] = processedBufferCnt })
}

// WithSendBufferSize send_buffer_size获取
func (obj *_AllVirtualDtlChannelMgr) WithSendBufferSize(sendBufferSize int64) Option {
	return optionFunc(func(o *options) { o.query["send_buffer_size"] = sendBufferSize })
}

// WithHashVal hash_val获取
func (obj *_AllVirtualDtlChannelMgr) WithHashVal(hashVal int64) Option {
	return optionFunc(func(o *options) { o.query["hash_val"] = hashVal })
}

// WithBufferPoolID buffer_pool_id获取
func (obj *_AllVirtualDtlChannelMgr) WithBufferPoolID(bufferPoolID int64) Option {
	return optionFunc(func(o *options) { o.query["buffer_pool_id"] = bufferPoolID })
}

// WithPins pins获取
func (obj *_AllVirtualDtlChannelMgr) WithPins(pins int64) Option {
	return optionFunc(func(o *options) { o.query["pins"] = pins })
}

// WithFirstInTs first_in_ts获取
func (obj *_AllVirtualDtlChannelMgr) WithFirstInTs(firstInTs time.Time) Option {
	return optionFunc(func(o *options) { o.query["first_in_ts"] = firstInTs })
}

// WithFirstOutTs first_out_ts获取
func (obj *_AllVirtualDtlChannelMgr) WithFirstOutTs(firstOutTs time.Time) Option {
	return optionFunc(func(o *options) { o.query["first_out_ts"] = firstOutTs })
}

// WithLastIntTs last_int_ts获取
func (obj *_AllVirtualDtlChannelMgr) WithLastIntTs(lastIntTs time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_int_ts"] = lastIntTs })
}

// WithLastOutTs last_out_ts获取
func (obj *_AllVirtualDtlChannelMgr) WithLastOutTs(lastOutTs time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_out_ts"] = lastOutTs })
}

// WithStatus status获取
func (obj *_AllVirtualDtlChannelMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDtlChannelMgr) GetByOption(opts ...Option) (result AllVirtualDtlChannel, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDtlChannelMgr) GetByOptions(opts ...Option) (results []*AllVirtualDtlChannel, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDtlChannelMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDtlChannel, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where(options.query)
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
func (obj *_AllVirtualDtlChannelMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromChannelID 通过channel_id获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromChannelID(channelID int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`channel_id` = ?", channelID).Find(&results).Error

	return
}

// GetBatchFromChannelID 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromChannelID(channelIDs []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`channel_id` IN (?)", channelIDs).Find(&results).Error

	return
}

// GetFromOpID 通过op_id获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromOpID(opID int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`op_id` = ?", opID).Find(&results).Error

	return
}

// GetBatchFromOpID 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromOpID(opIDs []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`op_id` IN (?)", opIDs).Find(&results).Error

	return
}

// GetFromPeerID 通过peer_id获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromPeerID(peerID int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`peer_id` = ?", peerID).Find(&results).Error

	return
}

// GetBatchFromPeerID 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromPeerID(peerIDs []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`peer_id` IN (?)", peerIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIsLocal 通过is_local获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromIsLocal(isLocal int8) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`is_local` = ?", isLocal).Find(&results).Error

	return
}

// GetBatchFromIsLocal 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromIsLocal(isLocals []int8) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`is_local` IN (?)", isLocals).Find(&results).Error

	return
}

// GetFromIsData 通过is_data获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromIsData(isData int8) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`is_data` = ?", isData).Find(&results).Error

	return
}

// GetBatchFromIsData 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromIsData(isDatas []int8) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`is_data` IN (?)", isDatas).Find(&results).Error

	return
}

// GetFromIsTransmit 通过is_transmit获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromIsTransmit(isTransmit int8) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`is_transmit` = ?", isTransmit).Find(&results).Error

	return
}

// GetBatchFromIsTransmit 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromIsTransmit(isTransmits []int8) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`is_transmit` IN (?)", isTransmits).Find(&results).Error

	return
}

// GetFromAllocBufferCnt 通过alloc_buffer_cnt获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromAllocBufferCnt(allocBufferCnt int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`alloc_buffer_cnt` = ?", allocBufferCnt).Find(&results).Error

	return
}

// GetBatchFromAllocBufferCnt 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromAllocBufferCnt(allocBufferCnts []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`alloc_buffer_cnt` IN (?)", allocBufferCnts).Find(&results).Error

	return
}

// GetFromFreeBufferCnt 通过free_buffer_cnt获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromFreeBufferCnt(freeBufferCnt int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`free_buffer_cnt` = ?", freeBufferCnt).Find(&results).Error

	return
}

// GetBatchFromFreeBufferCnt 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromFreeBufferCnt(freeBufferCnts []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`free_buffer_cnt` IN (?)", freeBufferCnts).Find(&results).Error

	return
}

// GetFromSendBufferCnt 通过send_buffer_cnt获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromSendBufferCnt(sendBufferCnt int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`send_buffer_cnt` = ?", sendBufferCnt).Find(&results).Error

	return
}

// GetBatchFromSendBufferCnt 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromSendBufferCnt(sendBufferCnts []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`send_buffer_cnt` IN (?)", sendBufferCnts).Find(&results).Error

	return
}

// GetFromRecvBufferCnt 通过recv_buffer_cnt获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromRecvBufferCnt(recvBufferCnt int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`recv_buffer_cnt` = ?", recvBufferCnt).Find(&results).Error

	return
}

// GetBatchFromRecvBufferCnt 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromRecvBufferCnt(recvBufferCnts []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`recv_buffer_cnt` IN (?)", recvBufferCnts).Find(&results).Error

	return
}

// GetFromProcessedBufferCnt 通过processed_buffer_cnt获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromProcessedBufferCnt(processedBufferCnt int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`processed_buffer_cnt` = ?", processedBufferCnt).Find(&results).Error

	return
}

// GetBatchFromProcessedBufferCnt 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromProcessedBufferCnt(processedBufferCnts []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`processed_buffer_cnt` IN (?)", processedBufferCnts).Find(&results).Error

	return
}

// GetFromSendBufferSize 通过send_buffer_size获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromSendBufferSize(sendBufferSize int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`send_buffer_size` = ?", sendBufferSize).Find(&results).Error

	return
}

// GetBatchFromSendBufferSize 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromSendBufferSize(sendBufferSizes []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`send_buffer_size` IN (?)", sendBufferSizes).Find(&results).Error

	return
}

// GetFromHashVal 通过hash_val获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromHashVal(hashVal int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`hash_val` = ?", hashVal).Find(&results).Error

	return
}

// GetBatchFromHashVal 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromHashVal(hashVals []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`hash_val` IN (?)", hashVals).Find(&results).Error

	return
}

// GetFromBufferPoolID 通过buffer_pool_id获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromBufferPoolID(bufferPoolID int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`buffer_pool_id` = ?", bufferPoolID).Find(&results).Error

	return
}

// GetBatchFromBufferPoolID 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromBufferPoolID(bufferPoolIDs []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`buffer_pool_id` IN (?)", bufferPoolIDs).Find(&results).Error

	return
}

// GetFromPins 通过pins获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromPins(pins int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`pins` = ?", pins).Find(&results).Error

	return
}

// GetBatchFromPins 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromPins(pinss []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`pins` IN (?)", pinss).Find(&results).Error

	return
}

// GetFromFirstInTs 通过first_in_ts获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromFirstInTs(firstInTs time.Time) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`first_in_ts` = ?", firstInTs).Find(&results).Error

	return
}

// GetBatchFromFirstInTs 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromFirstInTs(firstInTss []time.Time) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`first_in_ts` IN (?)", firstInTss).Find(&results).Error

	return
}

// GetFromFirstOutTs 通过first_out_ts获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromFirstOutTs(firstOutTs time.Time) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`first_out_ts` = ?", firstOutTs).Find(&results).Error

	return
}

// GetBatchFromFirstOutTs 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromFirstOutTs(firstOutTss []time.Time) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`first_out_ts` IN (?)", firstOutTss).Find(&results).Error

	return
}

// GetFromLastIntTs 通过last_int_ts获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromLastIntTs(lastIntTs time.Time) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`last_int_ts` = ?", lastIntTs).Find(&results).Error

	return
}

// GetBatchFromLastIntTs 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromLastIntTs(lastIntTss []time.Time) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`last_int_ts` IN (?)", lastIntTss).Find(&results).Error

	return
}

// GetFromLastOutTs 通过last_out_ts获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromLastOutTs(lastOutTs time.Time) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`last_out_ts` = ?", lastOutTs).Find(&results).Error

	return
}

// GetBatchFromLastOutTs 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromLastOutTs(lastOutTss []time.Time) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`last_out_ts` IN (?)", lastOutTss).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualDtlChannelMgr) GetFromStatus(status int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualDtlChannelMgr) GetBatchFromStatus(statuss []int64) (results []*AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDtlChannelMgr) FetchByPrimaryKey(svrIP string, svrPort int64) (result AllVirtualDtlChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlChannel{}).Where("`svr_ip` = ? AND `svr_port` = ?", svrIP, svrPort).First(&result).Error

	return
}
