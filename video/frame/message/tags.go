package message

type EventType int
type EventSubType int

type Tagser interface {
	AddValue(context IContext, event EventType)
	AddDayValue(context IContext, event EventType)
	AddValueRef(context IContext, event EventType, refID int32)
	AddDayValueRef(context IContext, event EventType, refID int32)
	AddSubValue(context IContext, event EventType, subEvent EventSubType)
	AddDaySubValue(context IContext, event EventType, subEvent EventSubType)
	AddSubValueRef(context IContext, event EventType, subEvent EventSubType, refID int32)
	AddDaySubValueRef(context IContext, event EventType, subEvent EventSubType, refID int32)

	SetValue(context IContext, event EventType, value int32)
	SetDayValue(context IContext, event EventType, value int32)
	SetValueRef(context IContext, event EventType, value int32, refID int32)
	SetDayValueRef(context IContext, event EventType, value int32, refID int32)
	SetSubValue(context IContext, event EventType, subEvent EventSubType, value int32)
	SetDaySubValue(context IContext, event EventType, subEvent EventSubType, value int32)
	SetSubValueRef(context IContext, event EventType, subEvent EventSubType, refID, value int32)
	SetDaySubValueRef(context IContext, event EventType, subEvent EventSubType, refID, value int32)

	Get(context IContext, event EventType) int32
	GetDay(context IContext, event EventType) int32
	GetRef(context IContext, event EventType, refID int32) int32
	GetDayRef(context IContext, event EventType, refID int32) int32
	GetSub(context IContext, event EventType, subEvent EventSubType) int32
	GetDaySub(context IContext, event EventType, subEvent EventSubType) int32
	GetSubRef(context IContext, event EventType, subEvent EventSubType, refID int32) int32
	GetDaySubRef(context IContext, event EventType, subEvent EventSubType, refID int32) int32
}
