package channel

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/state"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mosaicnetworks/chatterboxhub/x/ibc/03-connection"
	"github.com/mosaicnetworks/chatterboxhub/x/ibc/23-commitment"
)

// Manager is unrestricted
type Manager struct {
	protocol state.Mapping

	connection connection.Manager

	counterparty CounterpartyManager
}

type CounterpartyManager struct {
	protocol commitment.Mapping

	connection connection.CounterpartyManager
}

func NewManager(protocol state.Base, connection connection.Manager) Manager {
	return Manager{
		protocol:     state.NewMapping(protocol, []byte("/connection/")),
		connection:   connection,
		counterparty: NewCounterpartyManager(protocol.Cdc()),
	}
}

func NewCounterpartyManager(cdc *codec.Codec) CounterpartyManager {
	protocol := commitment.NewBase(cdc)

	return CounterpartyManager{
		protocol:   commitment.NewMapping(protocol, []byte("/connection/")),
		connection: connection.NewCounterpartyManager(cdc),
	}
}

// CONTRACT: connection and counterparty must be filled by the caller
func (man Manager) Object(connid, chanid string) Object {
	key := connid + "/channels/" + chanid
	return Object{
		chanid:      chanid,
		channel:     man.protocol.Value([]byte(key)),
		state:       state.NewEnum(man.protocol.Value([]byte(key + "/state"))),
		nexttimeout: state.NewInteger(man.protocol.Value([]byte(key+"/timeout")), state.Dec),

		// TODO: remove length functionality from state.Indeer(will be handled manually)
		Seqsend: state.NewInteger(man.protocol.Value([]byte(key+"/nextSequenceSend")), state.Dec),
		Seqrecv: state.NewInteger(man.protocol.Value([]byte(key+"/nextSequenceRecv")), state.Dec),
		Packets: state.NewIndexer(man.protocol.Prefix([]byte(key+"/Packets/")), state.Dec),
	}
}

func (man CounterpartyManager) object(connid, chanid string) CounterObject {
	key := connid + "/channels/" + chanid
	return CounterObject{
		chanid:      chanid,
		channel:     man.protocol.Value([]byte(key)),
		state:       commitment.NewEnum(man.protocol.Value([]byte(key + "/state"))),
		nexttimeout: commitment.NewInteger(man.protocol.Value([]byte(key+"/timeout")), state.Dec),

		Seqsend: commitment.NewInteger(man.protocol.Value([]byte(key+"/nextSequenceSend")), state.Dec),
		Seqrecv: commitment.NewInteger(man.protocol.Value([]byte(key+"/nextSequenceRecv")), state.Dec),
		Packets: commitment.NewIndexer(man.protocol.Prefix([]byte(key+"/Packets")), state.Dec),
	}
}

func (man Manager) Create(ctx sdk.Context, connid, chanid string, channel Channel) (obj Object, err error) {
	obj = man.Object(connid, chanid)
	if obj.exists(ctx) {
		err = errors.New("channel already exists for the provided id")
		return
	}
	obj.connection, err = man.connection.Query(ctx, connid)
	if err != nil {
		return
	}
	if obj.connection.State(ctx) != connection.Open {
		err = errors.New("connection exists but not opened")
		return
	}

	counterconnid := obj.connection.Value(ctx).Counterparty
	obj.counterparty = man.counterparty.object(counterconnid, channel.Counterparty)
	obj.counterparty.connection = man.counterparty.connection.Query(counterconnid)

	obj.channel.Set(ctx, channel)

	return
}

func (man Manager) Query(ctx sdk.Context, connid, chanid string) (obj Object, err error) {
	obj = man.Object(connid, chanid)
	if !obj.exists(ctx) {
		err = errors.New("channel not exists for the provided id")
		return
	}
	obj.connection, err = man.connection.Query(ctx, connid)
	if err != nil {
		return
	}
	if obj.connection.State(ctx) != connection.Open {
		err = errors.New("connection exists but not opened")
		return
	}

	channel := obj.Value(ctx)
	counterconnid := obj.connection.Value(ctx).Counterparty
	obj.counterparty = man.counterparty.object(counterconnid, channel.Counterparty)
	obj.counterparty.connection = man.counterparty.connection.Query(counterconnid)

	return
}

func (man Manager) Module(module string, chanid func(string) bool) ModuleManager {
	return ModuleManager{
		man:    man,
		module: module,
		chanid: chanid,
	}
}

// ModuleManage is module specific
type ModuleManager struct {
	man    Manager
	module string
	chanid func(string) bool
}

func (man ModuleManager) Create(ctx sdk.Context, connid, chanid string, channel Channel) (Object, error) {
	if !man.chanid(chanid) {
		return Object{}, errors.New("invalid channel id")
	}

	if channel.Module != man.module {
		return Object{}, errors.New("invalid module")
	}

	return man.man.Create(ctx, connid, chanid, channel)
}

func (man ModuleManager) Query(ctx sdk.Context, connid, chanid string) (Object, error) {
	if !man.chanid(chanid) {
		return Object{}, errors.New("invalid channel id")
	}

	obj, err := man.man.Query(ctx, connid, chanid)
	if err != nil {
		return Object{}, err
	}

	if obj.Value(ctx).Module != man.module {
		return Object{}, errors.New("invalid module")
	}

	return obj, nil
}

type Object struct {
	chanid      string
	channel     state.Value
	state       state.Enum
	nexttimeout state.Integer

	// TODO: must be unexported
	Seqsend state.Integer
	Seqrecv state.Integer
	Packets state.Indexer

	connection connection.Object

	counterparty CounterObject
}

type CounterObject struct {
	chanid      string
	channel     commitment.Value
	state       commitment.Enum
	nexttimeout commitment.Integer

	Seqsend commitment.Integer
	Seqrecv commitment.Integer
	Packets commitment.Indexer

	connection connection.CounterObject
}

func (obj Object) ChanID() string {
	return obj.chanid
}

func (obj Object) Value(ctx sdk.Context) (res Channel) {
	obj.channel.Get(ctx, &res)
	return
}

func (obj Object) exists(ctx sdk.Context) bool {
	return obj.channel.Exists(ctx)
}

func assertTimeout(ctx sdk.Context, timeoutHeight uint64) error {
	if uint64(ctx.BlockHeight()) > timeoutHeight {
		return errors.New("timeout")
	}
	return nil
}

/*
// TODO: ocapify callingModule
func (obj Object) OpenInit(ctx sdk.Context) error {
	// CONTRACT: OpenInit() should be called after man.Create(), not man.Query(),
	// which will ensure
	// assert(get() === null) and
	// set() and
	// connection.state == open

	// getCallingModule() === channel.moduleIdentifier is ensured by ModuleManager

	if !obj.state.Transit(ctx, Idle, Init) {
		return errors.New("init on non-idle channel")
	}

	obj.Seqsend.Set(ctx, 0)
	obj.Seqrecv.Set(ctx, 0)

	return nil
}

/*
func (obj Object) OpenTry(ctx sdk.Context, timeoutHeight, nextTimeoutHeight uint64) error {
	if !obj.state.Transit(ctx, Idle, OpenTry) {
		return errors.New("opentry on non-idle channel")
	}

	err := assertTimeout(ctx, timeoutHeight)
	if err != nil {
		return err
	}

	// XXX
}
*/

func (obj Object) Open(ctx sdk.Context) error {
	if !obj.state.Transit(ctx, Idle, Open) {
		return errors.New("open on non-idle channel")
	}

	return nil
}

func (obj Object) Send(ctx sdk.Context, packet Packet) error {
	if obj.state.Get(ctx) != Open {
		return errors.New("send on non-open channel")
	}

	if obj.connection.State(ctx) != Open {
		return errors.New("send on non-open connection")
	}

	if obj.connection.Client(ctx).GetHeight() >= packet.Timeout() {
		return errors.New("timeout height higher than the latest known")
	}

	obj.Packets.Set(ctx, obj.Seqsend.Incr(ctx), packet)

	return nil
}

func (obj Object) Receive(ctx sdk.Context, packet Packet) error {
	if obj.state.Get(ctx) != Open {
		return errors.New("send on non-open channel")
	}

	if obj.connection.State(ctx) != Open {
		return errors.New("send on non-open connection")
	}

	err := assertTimeout(ctx, packet.Timeout())
	if err != nil {
		return err
	}

	// XXX: increment should happen before verification, reflect on the spec
	if !obj.counterparty.Packets.Value(obj.Seqrecv.Incr(ctx)).Is(ctx, packet.Commit()) {
		return errors.New("verification failed")
	}

	return nil
}
