// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package chainclient

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/Loopring/ringminer/types"
)

var _ = (*orderFilledEventMarshaling)(nil)

func (o OrderFilledEvent) MarshalJSON() ([]byte, error) {
	type OrderFilledEvent struct {
		RingIndex     *types.Big `json:"ringIndex" gencodec:"required"`
		Time          *types.Big `json:"time" gencodec:"required"`
		Blocknumber   *types.Big `json:"blockNumber"	gencodec:"required"`
		Ringhash      []byte     `json:"ringHash" gencodec:"required"`
		PreOrderHash  []byte     `json:"preOrderHash" gencodec:"required"`
		OrderHash     []byte     `json:"orderHash" gencodec:"required"`
		NextOrderHash []byte     `json:"nextOrderHash" gencodec:"required"`
		AmountS       *types.Big `json:"amountS" gencodec:"required"`
		AmountB       *types.Big `json:"amountB" gencodec:"required"`
		LrcReward     *types.Big `json:"lrcReward" gencodec:"required"`
		LrcFee        *types.Big `json:"lrcFee" gencodec:"required"`
	}
	var enc OrderFilledEvent
	enc.RingIndex = (*types.Big)(o.RingIndex)
	enc.Time = (*types.Big)(o.Time)
	enc.Blocknumber = (*types.Big)(o.Blocknumber)
	enc.Ringhash = o.Ringhash
	enc.PreOrderHash = o.PreOrderHash
	enc.OrderHash = o.OrderHash
	enc.NextOrderHash = o.NextOrderHash
	enc.AmountS = (*types.Big)(o.AmountS)
	enc.AmountB = (*types.Big)(o.AmountB)
	enc.LrcReward = (*types.Big)(o.LrcReward)
	enc.LrcFee = (*types.Big)(o.LrcFee)
	return json.Marshal(&enc)
}

func (o *OrderFilledEvent) UnmarshalJSON(input []byte) error {
	type OrderFilledEvent struct {
		RingIndex     *types.Big `json:"ringIndex" gencodec:"required"`
		Time          *types.Big `json:"time" gencodec:"required"`
		Blocknumber   *types.Big `json:"blockNumber"	gencodec:"required"`
		Ringhash      []byte     `json:"ringHash" gencodec:"required"`
		PreOrderHash  []byte     `json:"preOrderHash" gencodec:"required"`
		OrderHash     []byte     `json:"orderHash" gencodec:"required"`
		NextOrderHash []byte     `json:"nextOrderHash" gencodec:"required"`
		AmountS       *types.Big `json:"amountS" gencodec:"required"`
		AmountB       *types.Big `json:"amountB" gencodec:"required"`
		LrcReward     *types.Big `json:"lrcReward" gencodec:"required"`
		LrcFee        *types.Big `json:"lrcFee" gencodec:"required"`
	}
	var dec OrderFilledEvent
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.RingIndex == nil {
		return errors.New("missing required field 'ringIndex' for OrderFilledEvent")
	}
	o.RingIndex = (*big.Int)(dec.RingIndex)
	if dec.Time == nil {
		return errors.New("missing required field 'time' for OrderFilledEvent")
	}
	o.Time = (*big.Int)(dec.Time)
	if dec.Blocknumber != nil {
		o.Blocknumber = (*big.Int)(dec.Blocknumber)
	}
	if dec.Ringhash == nil {
		return errors.New("missing required field 'ringHash' for OrderFilledEvent")
	}
	o.Ringhash = dec.Ringhash
	if dec.PreOrderHash == nil {
		return errors.New("missing required field 'preOrderHash' for OrderFilledEvent")
	}
	o.PreOrderHash = dec.PreOrderHash
	if dec.OrderHash == nil {
		return errors.New("missing required field 'orderHash' for OrderFilledEvent")
	}
	o.OrderHash = dec.OrderHash
	if dec.NextOrderHash == nil {
		return errors.New("missing required field 'nextOrderHash' for OrderFilledEvent")
	}
	o.NextOrderHash = dec.NextOrderHash
	if dec.AmountS == nil {
		return errors.New("missing required field 'amountS' for OrderFilledEvent")
	}
	o.AmountS = (*big.Int)(dec.AmountS)
	if dec.AmountB == nil {
		return errors.New("missing required field 'amountB' for OrderFilledEvent")
	}
	o.AmountB = (*big.Int)(dec.AmountB)
	if dec.LrcReward == nil {
		return errors.New("missing required field 'lrcReward' for OrderFilledEvent")
	}
	o.LrcReward = (*big.Int)(dec.LrcReward)
	if dec.LrcFee == nil {
		return errors.New("missing required field 'lrcFee' for OrderFilledEvent")
	}
	o.LrcFee = (*big.Int)(dec.LrcFee)
	return nil
}
