// Copyright (c) 2017-2019 The Elastos Foundation
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.
// 

package types

import (
	"bytes"
	"math/rand"
	"strconv"
	"testing"

	"github.com/elastos/Elastos.ELA/auxpow"
	"github.com/elastos/Elastos.ELA/common"
	"github.com/elastos/Elastos.ELA/core/contract/program"
	"github.com/elastos/Elastos.ELA/core/types/outputpayload"
	"github.com/elastos/Elastos.ELA/core/types/payload"

	"github.com/stretchr/testify/suite"
)

type transactionSuite struct {
	suite.Suite

	InputNum   int
	OutputNum  int
	AttrNum    int
	ProgramNum int
}

func (s *transactionSuite) SetupSuite() {
	s.InputNum = 10
	s.OutputNum = 10
	s.AttrNum = 10
	s.ProgramNum = 10
}

func (s *transactionSuite) TestCoinbaseTransaction_SerializeDeserialize() {
	txn := randomOldVersionTransaction(true, byte(CoinBase), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.CoinBase{
		Content: []byte(strconv.FormatUint(rand.Uint64(), 10)),
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(true, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	s.True(bytes.Equal(txn.Payload.(*payload.CoinBase).Content, txn2.Payload.(*payload.CoinBase).Content))
}

func (s *transactionSuite) TestRegisterAssetTransaction_SerializeDeserialize() {
	txn := randomOldVersionTransaction(true, byte(RegisterAsset), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.RegisterAsset{
		Asset: payload.Asset{
			Name:        "test name",
			Description: "test desc",
			Precision:   byte(rand.Uint32()),
			AssetType:   payload.AssetType(rand.Uint32()),
			RecordType:  payload.AssetRecordType(rand.Uint32()),
		},
		Amount:     common.Fixed64(rand.Int63()),
		Controller: *randomUint168(),
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(true, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	p1 := txn.Payload.(*payload.RegisterAsset)
	p2 := txn2.Payload.(*payload.RegisterAsset)

	s.Equal(p1.Asset.Name, p2.Asset.Name)
	s.Equal(p1.Asset.Description, p2.Asset.Description)
	s.Equal(p1.Asset.Precision, p2.Asset.Precision)
	s.Equal(p1.Asset.AssetType, p2.Asset.AssetType)
	s.Equal(p1.Asset.RecordType, p2.Asset.RecordType)
	s.Equal(p1.Amount, p2.Amount)
	s.True(p1.Controller.IsEqual(p2.Controller))
}

func (s *transactionSuite) TestTransferAssert_SerializeDeserialize() {
	txn := randomOldVersionTransaction(true, byte(TransferAsset), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.TransferAsset{}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(true, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
}

func (s *transactionSuite) TestRecord_SerializeDeserialize() {
	txn := randomOldVersionTransaction(true, byte(Record), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.Record{
		Type:    "test record type",
		Content: []byte(strconv.FormatUint(rand.Uint64(), 10)),
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(true, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	p1 := txn.Payload.(*payload.Record)
	p2 := txn2.Payload.(*payload.Record)

	s.Equal(p1.Type, p2.Type)
	s.True(bytes.Equal(p1.Content, p2.Content))
}

func (s *transactionSuite) TestSideChainPow_SerializeDeserialize() {
	txn := randomOldVersionTransaction(true, byte(SideChainPow), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.SideChainPow{
		SideBlockHash:   *randomUint256(),
		SideGenesisHash: *randomUint256(),
		BlockHeight:     rand.Uint32(),
		Signature:       []byte(strconv.FormatUint(rand.Uint64(), 10)),
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(true, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	p1 := txn.Payload.(*payload.SideChainPow)
	p2 := txn2.Payload.(*payload.SideChainPow)

	s.True(p1.SideBlockHash.IsEqual(p2.SideBlockHash))
	s.True(p1.SideGenesisHash.IsEqual(p2.SideGenesisHash))
	s.Equal(p1.BlockHeight, p2.BlockHeight)
	s.True(bytes.Equal(p1.Signature, p2.Signature))
}

func (s *transactionSuite) TestWithdrawFromSideChain_SerializeDeserialize() {
	txn := randomOldVersionTransaction(true, byte(WithdrawFromSideChain), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.WithdrawFromSideChain{
		BlockHeight:         rand.Uint32(),
		GenesisBlockAddress: "test genesis block address",
		SideChainTransactionHashes: []common.Uint256{
			*randomUint256(),
			*randomUint256(),
		},
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(true, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	p1 := txn.Payload.(*payload.WithdrawFromSideChain)
	p2 := txn2.Payload.(*payload.WithdrawFromSideChain)

	s.Equal(p1.BlockHeight, p2.BlockHeight)
	s.Equal(p1.GenesisBlockAddress, p2.GenesisBlockAddress)
	s.Equal(len(p1.SideChainTransactionHashes), len(p2.SideChainTransactionHashes))
	for i := range p1.SideChainTransactionHashes {
		s.True(p1.SideChainTransactionHashes[i].IsEqual(p2.SideChainTransactionHashes[i]))
	}
}

func (s *transactionSuite) TestTransferCrossChainAsset_SerializeDeserialize() {
	txn := randomOldVersionTransaction(true, byte(TransferCrossChainAsset), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.TransferCrossChainAsset{
		CrossChainAddresses: []string{
			strconv.FormatUint(rand.Uint64(), 10),
			strconv.FormatUint(rand.Uint64(), 10),
		},
		OutputIndexes: []uint64{
			rand.Uint64(),
			rand.Uint64(),
		},
		CrossChainAmounts: []common.Fixed64{
			common.Fixed64(rand.Int63()),
			common.Fixed64(rand.Int63()),
		},
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(true, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	p1 := txn.Payload.(*payload.TransferCrossChainAsset)
	p2 := txn2.Payload.(*payload.TransferCrossChainAsset)

	s.Equal(len(p1.CrossChainAddresses), len(p2.CrossChainAddresses))
	s.Equal(len(p1.OutputIndexes), len(p2.OutputIndexes))
	s.Equal(len(p1.CrossChainAmounts), len(p2.CrossChainAmounts))
	s.Equal(len(p1.CrossChainAddresses), len(p2.OutputIndexes))
	s.Equal(len(p1.CrossChainAddresses), len(p1.CrossChainAmounts))
	for i := range p1.CrossChainAddresses {
		s.Equal(p1.CrossChainAddresses[i], p2.CrossChainAddresses[i])
	}
	for i := range p1.OutputIndexes {
		s.Equal(p1.OutputIndexes[i], p2.OutputIndexes[i])
	}
	for i := range p1.CrossChainAmounts {
		s.Equal(p1.CrossChainAmounts[i], p2.CrossChainAmounts[i])
	}
}

func (s *transactionSuite) TestRegisterProducer_SerializeDeserialize() {
	txn := randomOldVersionTransaction(false, byte(RegisterProducer), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.ProducerInfo{
		OwnerPublicKey: []byte(strconv.FormatUint(rand.Uint64(), 10)),
		NickName:       strconv.FormatUint(rand.Uint64(), 10),
		Url:            strconv.FormatUint(rand.Uint64(), 10),
		Location:       rand.Uint64(),
		NetAddress:     strconv.FormatUint(rand.Uint64(), 10),
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(false, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	p1 := txn.Payload.(*payload.ProducerInfo)
	p2 := txn2.Payload.(*payload.ProducerInfo)

	s.True(bytes.Equal(p1.OwnerPublicKey, p2.OwnerPublicKey))
	s.Equal(p1.NickName, p2.NickName)
	s.Equal(p1.Url, p2.Url)
	s.Equal(p1.Location, p2.Location)
	s.Equal(p1.NetAddress, p2.NetAddress)
}

func (s *transactionSuite) TestCancelProducer_SerializeDeserialize() {
	txn := randomOldVersionTransaction(false, byte(CancelProducer),
		s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.ProcessProducer{
		OwnerPublicKey: []byte(strconv.FormatUint(rand.Uint64(), 10)),
		Signature:      randomSignature(),
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(false, &s.Suite, txn, txn2, s.InputNum,
		s.OutputNum, s.AttrNum, s.ProgramNum)

	p1 := txn.Payload.(*payload.ProcessProducer)
	p2 := txn2.Payload.(*payload.ProcessProducer)

	s.True(bytes.Equal(p1.OwnerPublicKey, p2.OwnerPublicKey))
	s.True(bytes.Equal(p1.Signature, p2.Signature))
}

func (s *transactionSuite) TestActivateProducer_SerializeDeserialize() {
	txn := randomOldVersionTransaction(false, byte(ActivateProducer),
		s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.ActivateProducer{
		NodePublicKey: []byte(strconv.FormatUint(rand.Uint64(), 10)),
		Signature:     randomSignature(),
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(false, &s.Suite, txn, txn2, s.InputNum,
		s.OutputNum, s.AttrNum, s.ProgramNum)

	p1 := txn.Payload.(*payload.ActivateProducer)
	p2 := txn2.Payload.(*payload.ActivateProducer)

	s.True(bytes.Equal(p1.NodePublicKey, p2.NodePublicKey))
	s.True(bytes.Equal(p1.Signature, p2.Signature))
}

func (s *transactionSuite) TestUpdateProducer_SerializeDeserialize() {
	txn := randomOldVersionTransaction(false, byte(UpdateProducer), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.ProducerInfo{
		OwnerPublicKey: []byte(strconv.FormatUint(rand.Uint64(), 10)),
		NickName:       strconv.FormatUint(rand.Uint64(), 10),
		Url:            strconv.FormatUint(rand.Uint64(), 10),
		Location:       rand.Uint64(),
		NetAddress:     strconv.FormatUint(rand.Uint64(), 10),
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(false, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	p1 := txn.Payload.(*payload.ProducerInfo)
	p2 := txn2.Payload.(*payload.ProducerInfo)

	s.True(bytes.Equal(p1.OwnerPublicKey, p2.OwnerPublicKey))
	s.Equal(p1.NickName, p2.NickName)
	s.Equal(p1.Url, p2.Url)
	s.Equal(p1.Location, p2.Location)
	s.Equal(p1.NetAddress, p2.NetAddress)
}

func (s *transactionSuite) TestReturnDepositCoin_SerializeDeserialize() {
	txn := randomOldVersionTransaction(false, byte(ReturnDepositCoin), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.ReturnDepositCoin{}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(false, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
}

func (s *transactionSuite) randomPkBytes() []byte {
	pk := make([]byte, 33)
	rand.Read(pk)
	return pk
}

func (s *transactionSuite) TestIllegalProposalEvidence_SerializeDeserialize() {
	txn := randomOldVersionTransaction(false, byte(IllegalProposalEvidence), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.DPOSIllegalProposals{
		Evidence: payload.ProposalEvidence{
			Proposal: payload.DPOSProposal{
				Sponsor:    s.randomPkBytes(),
				BlockHash:  *randomUint256(),
				ViewOffset: rand.Uint32(),
				Sign:       []byte(strconv.FormatUint(rand.Uint64(), 10)),
			},
			BlockHeader: randomBlockHeaderBytes(),
			BlockHeight: rand.Uint32(),
		},
		CompareEvidence: payload.ProposalEvidence{
			Proposal: payload.DPOSProposal{
				Sponsor:    s.randomPkBytes(),
				BlockHash:  *randomUint256(),
				ViewOffset: rand.Uint32(),
				Sign:       []byte(strconv.FormatUint(rand.Uint64(), 10)),
			},
			BlockHeader: randomBlockHeaderBytes(),
			BlockHeight: rand.Uint32(),
		},
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(false, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	s.True(txn.Payload.(*payload.DPOSIllegalProposals).Hash().IsEqual(
		txn2.Payload.(*payload.DPOSIllegalProposals).Hash()))
}

func (s *transactionSuite) TestIllegalVoteEvidence_SerializeDeserialize() {
	txn := randomOldVersionTransaction(false, byte(IllegalVoteEvidence), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.DPOSIllegalVotes{
		Evidence: payload.VoteEvidence{
			ProposalEvidence: payload.ProposalEvidence{
				Proposal: payload.DPOSProposal{
					Sponsor:    s.randomPkBytes(),
					BlockHash:  *randomUint256(),
					ViewOffset: rand.Uint32(),
					Sign:       []byte(strconv.FormatUint(rand.Uint64(), 10)),
				},
				BlockHeader: randomBlockHeaderBytes(),
				BlockHeight: rand.Uint32(),
			},
			Vote: payload.DPOSProposalVote{
				ProposalHash: *randomUint256(),
				Signer:       s.randomPkBytes(),
				Accept:       true,
				Sign:         []byte(strconv.FormatUint(rand.Uint64(), 10)),
			},
		},
		CompareEvidence: payload.VoteEvidence{
			ProposalEvidence: payload.ProposalEvidence{
				Proposal: payload.DPOSProposal{
					Sponsor:    s.randomPkBytes(),
					BlockHash:  *randomUint256(),
					ViewOffset: rand.Uint32(),
					Sign:       []byte(strconv.FormatUint(rand.Uint64(), 10)),
				},
				BlockHeader: randomBlockHeaderBytes(),
				BlockHeight: rand.Uint32(),
			},
			Vote: payload.DPOSProposalVote{
				ProposalHash: *randomUint256(),
				Signer:       s.randomPkBytes(),
				Accept:       true,
				Sign:         []byte(strconv.FormatUint(rand.Uint64(), 10)),
			},
		},
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(false, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	s.True(txn.Payload.(*payload.DPOSIllegalVotes).Hash().IsEqual(
		txn2.Payload.(*payload.DPOSIllegalVotes).Hash()))
}

func (s *transactionSuite) TestIllegalBlockEvidence_SerializeDeserialize() {
	txn := randomOldVersionTransaction(false, byte(IllegalBlockEvidence), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	txn.Payload = &payload.DPOSIllegalBlocks{
		CoinType:    payload.CoinType(rand.Uint32()),
		BlockHeight: rand.Uint32(),
		Evidence: payload.BlockEvidence{
			Header:       []byte(strconv.FormatUint(rand.Uint64(), 10)),
			BlockConfirm: []byte(strconv.FormatUint(rand.Uint64(), 10)),
			Signers: [][]byte{
				[]byte(strconv.FormatUint(rand.Uint64(), 10)),
				[]byte(strconv.FormatUint(rand.Uint64(), 10)),
			},
		},
		CompareEvidence: payload.BlockEvidence{
			Header:       []byte(strconv.FormatUint(rand.Uint64(), 10)),
			BlockConfirm: []byte(strconv.FormatUint(rand.Uint64(), 10)),
			Signers: [][]byte{
				[]byte(strconv.FormatUint(rand.Uint64(), 10)),
				[]byte(strconv.FormatUint(rand.Uint64(), 10)),
			},
		},
	}

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(false, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	s.True(txn.Payload.(*payload.DPOSIllegalBlocks).Hash().IsEqual(
		txn2.Payload.(*payload.DPOSIllegalBlocks).Hash()))
}

func (s *transactionSuite) TestSidechainIllegalData_SerializeDeserialize() {
	txn := randomOldVersionTransaction(false, byte(IllegalSidechainEvidence), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	p := &payload.SidechainIllegalData{
		IllegalType:         payload.IllegalDataType(rand.Intn(6)),
		Height:              rand.Uint32(),
		IllegalSigner:       randomPublicKey(),
		GenesisBlockAddress: randomUint168().String(),
		Evidence: payload.SidechainIllegalEvidence{
			DataHash: *randomUint256(),
		},
		CompareEvidence: payload.SidechainIllegalEvidence{
			DataHash: *randomUint256(),
		},
	}
	p.Signs = make([][]byte, 0)
	for i := 0; i < 10; i++ {
		p.Signs = append(p.Signs, randomSignature())
	}
	txn.Payload = p

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(false, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	p2 := txn2.Payload.(*payload.SidechainIllegalData)
	s.Equal(p.IllegalType, p2.IllegalType)
	s.Equal(p.Height, p2.Height)
	s.True(bytes.Equal(p.IllegalSigner, p2.IllegalSigner))
	s.Equal(p.GenesisBlockAddress, p2.GenesisBlockAddress)
	s.Equal(p.Evidence.DataHash.String(), p2.Evidence.DataHash.String())
	s.Equal(p.CompareEvidence.DataHash.String(),
		p2.CompareEvidence.DataHash.String())
	for i := 0; i < 10; i++ {
		s.True(bytes.Equal(p.Signs[i], p2.Signs[i]))
	}
}

func (s *transactionSuite) TestInactiveArbitrators_SerializeDeserialize() {
	txn := randomOldVersionTransaction(false, byte(InactiveArbitrators), s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)
	p := &payload.InactiveArbitrators{
		Sponsor:     randomPublicKey(),
		BlockHeight: rand.Uint32(),
	}

	p.Arbitrators = make([][]byte, 0)
	for i := 0; i < 10; i++ {
		p.Arbitrators = append(p.Arbitrators, randomPublicKey())
	}
	txn.Payload = p

	serializedData := new(bytes.Buffer)
	txn.Serialize(serializedData)

	txn2 := &Transaction{}
	txn2.Deserialize(serializedData)

	assertOldVersionTxEqual(false, &s.Suite, txn, txn2, s.InputNum, s.OutputNum, s.AttrNum, s.ProgramNum)

	p2 := txn2.Payload.(*payload.InactiveArbitrators)
	s.True(bytes.Equal(p.Sponsor, p2.Sponsor))
	s.Equal(p.BlockHeight, p2.BlockHeight)
	for i := 0; i < 10; i++ {
		s.True(bytes.Equal(p.Arbitrators[i], p2.Arbitrators[i]))
	}
}

func (s *transactionSuite) TestTransaction_SpecificSample() {
	// update producer transaction deserialize sample
	byteReader := new(bytes.Buffer)
	updateProducerByteStr := "090b0021034f3a7d2f33ac7f4e30876080d359ce5f314c9eabddbaaca637676377f655e16c2103c77af162438d4b7140f8544ad6523b9734cca9c7a62476d54ed5d1bddc7a39c309656c615f74657374310d656c615f74657374312e6f726754b60100000000000931302e31302e302e3240920d8e769640b8494cfdf7c5581982c329485c8d3083db7d3079de104dc0dc650d8592a5e1d70f5c24f72b3f29b0625dc6348e375b13c3c48992d398d9f5d9ac000146f1d8002115ce89423ab29f26ede6ef1b813642cbf3c4c15b919b41d6d9f7760100ffffffff01b037db964a231458d2d6ffd5ea18944c4f90e63d547c5d3b9874df66a4ead0a300d414000000000000000000210d4109bf00e6d782db40ab183491c03cf4d6a37a00000000000141408fc0c3de6198c3ec4e9ab8a7748208d79a554c9d1ef84edec23c444495b651deb7a8796ac49e31f1d2598207d216c05496b35d7f75a22c55272223995781bb402321034f3a7d2f33ac7f4e30876080d359ce5f314c9eabddbaaca637676377f655e16cac"
	updateProducerByte, _ := common.HexStringToBytes(updateProducerByteStr)
	byteReader.Write(updateProducerByte)
	txn := &Transaction{}
	s.NoError(txn.Deserialize(byteReader))
}

func TestTransactionSuite(t *testing.T) {
	suite.Run(t, new(transactionSuite))
}

func assertOldVersionTxEqual(oldVersion bool, suite *suite.Suite, first, second *Transaction, inputNum, outputNum, attrNum, programNum int) {
	if oldVersion {
		suite.Equal(TxVersionDefault, second.Version)
	} else {
		suite.Equal(first.Version, second.Version)
	}
	suite.Equal(first.TxType, second.TxType)
	suite.Equal(first.PayloadVersion, second.PayloadVersion)
	suite.Equal(first.LockTime, second.LockTime)

	suite.Equal(inputNum, len(first.Inputs))
	suite.Equal(inputNum, len(second.Inputs))
	for i := 0; i < inputNum; i++ {
		suite.Equal(first.Inputs[i].Sequence, second.Inputs[i].Sequence)
		suite.Equal(first.Inputs[i].Previous.Index, second.Inputs[i].Previous.Index)
		suite.True(first.Inputs[i].Previous.TxID.IsEqual(second.Inputs[i].Previous.TxID))
	}

	suite.Equal(outputNum, len(first.Outputs))
	suite.Equal(outputNum, len(second.Outputs))
	for i := 0; i < outputNum; i++ {
		suite.True(first.Outputs[i].AssetID.IsEqual(second.Outputs[i].AssetID))
		suite.Equal(first.Outputs[i].Value, second.Outputs[i].Value)
		suite.Equal(first.Outputs[i].OutputLock, second.Outputs[i].OutputLock)
		suite.True(first.Outputs[i].ProgramHash.IsEqual(second.Outputs[i].ProgramHash))

		if !oldVersion {
			suite.Equal(first.Outputs[i].Type, second.Outputs[i].Type)
		}
	}

	suite.Equal(attrNum, len(first.Attributes))
	suite.Equal(attrNum, len(second.Attributes))
	for i := 0; i < attrNum; i++ {
		suite.Equal(first.Attributes[i].Usage, second.Attributes[i].Usage)
		suite.True(bytes.Equal(first.Attributes[i].Data, second.Attributes[i].Data))
	}

	suite.Equal(programNum, len(first.Programs))
	suite.Equal(programNum, len(second.Programs))
	for i := 0; i < programNum; i++ {
		suite.True(bytes.Equal(first.Programs[i].Code, second.Programs[i].Code))
		suite.True(bytes.Equal(first.Programs[i].Parameter, second.Programs[i].Parameter))
	}
}

func randomOldVersionTransaction(oldVersion bool, txType byte, inputNum, outputNum, attrNum, programNum int) *Transaction {
	txn := &Transaction{
		Version:        TransactionVersion(txType),
		TxType:         TxType(txType),
		PayloadVersion: byte(rand.Uint32()),
		LockTime:       rand.Uint32(),
		Inputs:         make([]*Input, 0),
		Outputs:        make([]*Output, 0),
		Attributes:     make([]*Attribute, 0),
		Programs:       make([]*program.Program, 0),
	}
	if !oldVersion {
		txn.Version = TxVersion09
	}

	for i := 0; i < inputNum; i++ {
		txn.Inputs = append(txn.Inputs, &Input{
			Sequence: rand.Uint32(),
			Previous: OutPoint{
				TxID:  *randomUint256(),
				Index: uint16(rand.Uint32()),
			},
		})
	}

	for i := 0; i < outputNum; i++ {
		output := &Output{
			AssetID:     *randomUint256(),
			Value:       common.Fixed64(rand.Int63()),
			OutputLock:  rand.Uint32(),
			ProgramHash: *randomUint168(),
			Type:        0,
			Payload:     nil,
		}
		if !oldVersion {
			output.Type = OTNone
			output.Payload = &outputpayload.DefaultOutput{}
		}
		txn.Outputs = append(txn.Outputs, output)
	}

	validAttrUsage := []AttributeUsage{Nonce, Script, Memo, Description, DescriptionUrl, Confirmations}
	for i := 0; i < attrNum; i++ {
		txn.Attributes = append(txn.Attributes, &Attribute{
			Usage: validAttrUsage[rand.Intn(len(validAttrUsage))],
			Data:  []byte(strconv.FormatUint(rand.Uint64(), 10)),
		})
	}

	for i := 0; i < programNum; i++ {
		txn.Programs = append(txn.Programs, &program.Program{
			Code:      []byte(strconv.FormatUint(rand.Uint64(), 10)),
			Parameter: []byte(strconv.FormatUint(rand.Uint64(), 10)),
		})
	}

	return txn
}

func randomBlockHeaderBytes() []byte {
	buf := new(bytes.Buffer)
	header := randomBlockHeader()
	header.Serialize(buf)

	return buf.Bytes()
}

func randomBlockHeader() *Header {
	return &Header{
		Version:    rand.Uint32(),
		Previous:   *randomUint256(),
		MerkleRoot: *randomUint256(),
		Timestamp:  rand.Uint32(),
		Bits:       rand.Uint32(),
		Nonce:      rand.Uint32(),
		Height:     rand.Uint32(),
		AuxPow: auxpow.AuxPow{
			AuxMerkleBranch: []common.Uint256{
				*randomUint256(),
				*randomUint256(),
			},
			AuxMerkleIndex: rand.Int(),
			ParCoinbaseTx: auxpow.BtcTx{
				Version: rand.Int31(),
				TxIn: []*auxpow.BtcTxIn{
					{
						PreviousOutPoint: auxpow.BtcOutPoint{
							Hash:  *randomUint256(),
							Index: rand.Uint32(),
						},
						SignatureScript: []byte(strconv.FormatUint(rand.Uint64(), 10)),
						Sequence:        rand.Uint32(),
					},
					{
						PreviousOutPoint: auxpow.BtcOutPoint{
							Hash:  *randomUint256(),
							Index: rand.Uint32(),
						},
						SignatureScript: []byte(strconv.FormatUint(rand.Uint64(), 10)),
						Sequence:        rand.Uint32(),
					},
				},
				TxOut: []*auxpow.BtcTxOut{
					{
						Value:    rand.Int63(),
						PkScript: []byte(strconv.FormatUint(rand.Uint64(), 10)),
					},
					{
						Value:    rand.Int63(),
						PkScript: []byte(strconv.FormatUint(rand.Uint64(), 10)),
					},
				},
				LockTime: rand.Uint32(),
			},
			ParCoinBaseMerkle: []common.Uint256{
				*randomUint256(),
				*randomUint256(),
			},
			ParMerkleIndex: rand.Int(),
			ParBlockHeader: auxpow.BtcHeader{
				Version:    rand.Uint32(),
				Previous:   *randomUint256(),
				MerkleRoot: *randomUint256(),
				Timestamp:  rand.Uint32(),
				Bits:       rand.Uint32(),
				Nonce:      rand.Uint32(),
			},
			ParentHash: *randomUint256(),
		},
	}
}

func randomUint256() *common.Uint256 {
	randBytes := make([]byte, 32)
	rand.Read(randBytes)

	result, _ := common.Uint256FromBytes(randBytes)
	return result
}

func randomUint168() *common.Uint168 {
	randBytes := make([]byte, 21)
	rand.Read(randBytes)
	result, _ := common.Uint168FromBytes(randBytes)

	return result
}

func randomSignature() []byte {
	randBytes := make([]byte, 64)
	rand.Read(randBytes)

	return randBytes
}

func randomPublicKey() []byte {
	randBytes := make([]byte, 33)
	rand.Read(randBytes)
	return randBytes
}
