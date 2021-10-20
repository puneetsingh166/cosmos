package tx_test

import (
	gocontext "context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/puneetsingh166/tm-load-test/client"
	"github.com/puneetsingh166/tm-load-test/client/tx"
	"github.com/puneetsingh166/tm-load-test/crypto/hd"
	"github.com/puneetsingh166/tm-load-test/crypto/keyring"
	cryptotypes "github.com/puneetsingh166/tm-load-test/crypto/types"
	"github.com/puneetsingh166/tm-load-test/simapp"
	sdk "github.com/puneetsingh166/tm-load-test/types"
	txtypes "github.com/puneetsingh166/tm-load-test/types/tx"
	signingtypes "github.com/puneetsingh166/tm-load-test/types/tx/signing"
	"github.com/puneetsingh166/tm-load-test/x/auth/signing"
	banktypes "github.com/puneetsingh166/tm-load-test/x/bank/types"
)

func NewTestTxConfig() client.TxConfig {
	cfg := simapp.MakeTestEncodingConfig()
	return cfg.TxConfig
}

// mockContext is a mock client.Context to return abitrary simulation response, used to
// unit test CalculateGas.
type mockContext struct {
	gasUsed uint64
	wantErr bool
}

func (m mockContext) Invoke(grpcCtx gocontext.Context, method string, req, reply interface{}, opts ...grpc.CallOption) (err error) {
	if m.wantErr {
		return fmt.Errorf("mock err")
	}

	*(reply.(*txtypes.SimulateResponse)) = txtypes.SimulateResponse{
		GasInfo: &sdk.GasInfo{GasUsed: m.gasUsed, GasWanted: m.gasUsed},
		Result:  &sdk.Result{Data: []byte("tx data"), Log: "log"},
	}

	return nil
}
func (mockContext) NewStream(gocontext.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	panic("not implemented")
}

func TestCalculateGas(t *testing.T) {
	type args struct {
		mockGasUsed uint64
		mockWantErr bool
		adjustment  float64
	}

	testCases := []struct {
		name         string
		args         args
		wantEstimate uint64
		wantAdjusted uint64
		expPass      bool
	}{
		{"error", args{0, true, 1.2}, 0, 0, false},
		{"adjusted gas", args{10, false, 1.2}, 10, 12, true},
	}

	for _, tc := range testCases {
		stc := tc
		txCfg := NewTestTxConfig()

		txf := tx.Factory{}.
			WithChainID("test-chain").
			WithTxConfig(txCfg).WithSignMode(txCfg.SignModeHandler().DefaultMode())

		t.Run(stc.name, func(t *testing.T) {
			mockClientCtx := mockContext{
				gasUsed: tc.args.mockGasUsed,
				wantErr: tc.args.mockWantErr,
			}
			simRes, gotAdjusted, err := tx.CalculateGas(mockClientCtx, txf.WithGasAdjustment(stc.args.adjustment))
			if stc.expPass {
				require.NoError(t, err)
				require.Equal(t, simRes.GasInfo.GasUsed, stc.wantEstimate)
				require.Equal(t, gotAdjusted, stc.wantAdjusted)
				require.NotNil(t, simRes.Result)
			} else {
				require.Error(t, err)
				require.Nil(t, simRes)
			}
		})
	}
}

func TestBuildSimTx(t *testing.T) {
	txCfg := NewTestTxConfig()

	txf := tx.Factory{}.
		WithTxConfig(txCfg).
		WithAccountNumber(50).
		WithSequence(23).
		WithFees("50stake").
		WithMemo("memo").
		WithChainID("test-chain").
		WithSignMode(txCfg.SignModeHandler().DefaultMode())

	msg := banktypes.NewMsgSend(sdk.AccAddress("from"), sdk.AccAddress("to"), nil)
	bz, err := txf.BuildSimTx(msg)
	require.NoError(t, err)
	require.NotNil(t, bz)
}

func TestBuildUnsignedTx(t *testing.T) {
	txf := tx.Factory{}.
		WithTxConfig(NewTestTxConfig()).
		WithAccountNumber(50).
		WithSequence(23).
		WithFees("50stake").
		WithMemo("memo").
		WithChainID("test-chain")

	msg := banktypes.NewMsgSend(sdk.AccAddress("from"), sdk.AccAddress("to"), nil)
	tx, err := txf.BuildUnsignedTx(msg)
	require.NoError(t, err)
	require.NotNil(t, tx)

	sigs, err := tx.GetTx().(signing.SigVerifiableTx).GetSignaturesV2()
	require.NoError(t, err)
	require.Empty(t, sigs)
}

func TestSign(t *testing.T) {
	requireT := require.New(t)
	path := hd.CreateHDPath(118, 0, 0).String()
	encCfg := simapp.MakeTestEncodingConfig()
	kb, err := keyring.New(t.Name(), "test", t.TempDir(), nil, encCfg.Codec)
	requireT.NoError(err)

	var from1 = "test_key1"
	var from2 = "test_key2"

	// create a new key using a mnemonic generator and test if we can reuse seed to recreate that account
	_, seed, err := kb.NewMnemonic(from1, keyring.English, path, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	requireT.NoError(err)
	requireT.NoError(kb.Delete(from1))
	k1, _, err := kb.NewMnemonic(from1, keyring.English, path, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	requireT.NoError(err)

	k2, err := kb.NewAccount(from2, seed, "", path, hd.Secp256k1)
	requireT.NoError(err)

	pubKey1, err := k1.GetPubKey()
	requireT.NoError(err)
	pubKey2, err := k2.GetPubKey()
	requireT.NoError(err)
	requireT.NotEqual(pubKey1.Bytes(), pubKey2.Bytes())
	t.Log("Pub keys:", pubKey1, pubKey2)

	txfNoKeybase := tx.Factory{}.
		WithTxConfig(NewTestTxConfig()).
		WithAccountNumber(50).
		WithSequence(23).
		WithFees("50stake").
		WithMemo("memo").
		WithChainID("test-chain")
	txfDirect := txfNoKeybase.
		WithKeybase(kb).
		WithSignMode(signingtypes.SignMode_SIGN_MODE_DIRECT)
	txfAmino := txfDirect.
		WithSignMode(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON)
	addr1, err := k1.GetAddress()
	requireT.NoError(err)
	addr2, err := k2.GetAddress()
	requireT.NoError(err)
	msg1 := banktypes.NewMsgSend(addr1, sdk.AccAddress("to"), nil)
	msg2 := banktypes.NewMsgSend(addr2, sdk.AccAddress("to"), nil)
	txb, err := txfNoKeybase.BuildUnsignedTx(msg1, msg2)
	requireT.NoError(err)
	txb2, err := txfNoKeybase.BuildUnsignedTx(msg1, msg2)
	requireT.NoError(err)
	txbSimple, err := txfNoKeybase.BuildUnsignedTx(msg2)
	requireT.NoError(err)

	testCases := []struct {
		name         string
		txf          tx.Factory
		txb          client.TxBuilder
		from         string
		overwrite    bool
		expectedPKs  []cryptotypes.PubKey
		matchingSigs []int // if not nil, check matching signature against old ones.
	}{
		{"should fail if txf without keyring",
			txfNoKeybase, txb, from1, true, nil, nil},
		{"should fail for non existing key",
			txfAmino, txb, "unknown", true, nil, nil},
		{"amino: should succeed with keyring",
			txfAmino, txbSimple, from1, true, []cryptotypes.PubKey{pubKey1}, nil},
		{"direct: should succeed with keyring",
			txfDirect, txbSimple, from1, true, []cryptotypes.PubKey{pubKey1}, nil},

		/**** test double sign Amino mode ****/
		{"amino: should sign multi-signers tx",
			txfAmino, txb, from1, true, []cryptotypes.PubKey{pubKey1}, nil},
		{"amino: should append a second signature and not overwrite",
			txfAmino, txb, from2, false, []cryptotypes.PubKey{pubKey1, pubKey2}, []int{0, 0}},
		{"amino: should overwrite a signature",
			txfAmino, txb, from2, true, []cryptotypes.PubKey{pubKey2}, []int{1, 0}},

		/**** test double sign Direct mode
		  signing transaction with more than 2 signers should fail in DIRECT mode ****/
		{"direct: should fail to append a signature with different mode",
			txfDirect, txb, from1, false, []cryptotypes.PubKey{}, nil},
		{"direct: should fail to sign multi-signers tx",
			txfDirect, txb2, from1, false, []cryptotypes.PubKey{}, nil},
		{"direct: should fail to overwrite multi-signers tx",
			txfDirect, txb2, from1, true, []cryptotypes.PubKey{}, nil},
	}
	var prevSigs []signingtypes.SignatureV2
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = tx.Sign(tc.txf, tc.from, tc.txb, tc.overwrite)
			if len(tc.expectedPKs) == 0 {
				requireT.Error(err)
			} else {
				requireT.NoError(err)
				sigs := testSigners(requireT, tc.txb.GetTx(), tc.expectedPKs...)
				if tc.matchingSigs != nil {
					requireT.Equal(prevSigs[tc.matchingSigs[0]], sigs[tc.matchingSigs[1]])
				}
				prevSigs = sigs
			}
		})
	}
}

func testSigners(require *require.Assertions, tr signing.Tx, pks ...cryptotypes.PubKey) []signingtypes.SignatureV2 {
	sigs, err := tr.GetSignaturesV2()
	require.Len(sigs, len(pks))
	require.NoError(err)
	require.Len(sigs, len(pks))
	for i := range pks {
		require.True(sigs[i].PubKey.Equals(pks[i]), "Signature is signed with a wrong pubkey. Got: %s, expected: %s", sigs[i].PubKey, pks[i])
	}
	return sigs
}
