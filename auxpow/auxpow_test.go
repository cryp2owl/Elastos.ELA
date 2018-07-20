package auxpow

import (
	"bytes"
	"testing"

	"github.com/elastos/Elastos.ELA.Utility/common"
)

func TestAuxPow_Check(t *testing.T) {

	// ela auxpow data. chainid = 6. AuxMerkleBranch = 1
	const hashHex = "7926398947f332fe534b15c628ff0cd9dc6f7d3ea59c74801dc758ac65428e64"
	const auxPowHex = "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4b0313ee0904a880495b742f4254432e434f4d2ffabe6d6d9581ba0156314f1e92fd03430c6e4428a32bb3f1b9dc627102498e5cfbf26261020000004204cb9a010f32a00601000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf90000000014acac4ee8fdd8ca7e0b587b35fce8c996c70aefdf24c333038bdba7af531266000000000001ccc205f0e1cb435f50cc2f63edd53186b414fcb22b719da8c59eab066cf30bdb0000000000000020d1061d1e456cae488c063838b64c4911ce256549afadfc6a4736643359141b01551e4d94f9e8b6b03eec92bb6de1e478a0e913e5f733f5884857a7c2b965f53ca880495bffff7f20a880495b"

	// namecoin auxpow data. chainid = 6. AuxMerkleBranch = 1
	//const hashHex = "21187623de86cd62b4ce211cd8a74e88f80eda6cc12f279bf3cdb5c0d9539a9d"
	//const auxPowHex = "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4b039aff0904db044a5b742f4254432e434f4d2ffabe6d6d35ecfc5f5ca2971449ee78b7d810f280de7e3e7c407e3c0162ef8692df350ef8020000004204cb9a011fde202e00000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf9000000001d1879510258c5186e39cfcde4539c88686854b1ca640681dd38ed9527e635600000000000015f2f03802d61504f12e25d4b679b881ddb374cc04f240b6eb765d887679fb6360000000000000020a9f32bdb09d7777f3fa308fcd221e531393441f50e7f8b2d4ef63b2c3440940ec866338e7674b07d6a92269317f09f6c0fdb60ce7052e0211133e0015727ebb2db044a5bffff7f20db044a5b"

	// ela auxpow data. chainid = 6. AuxMerkleBranch = 0
	//const hashHex = "e7fe91999c21de47a64fe0e0bd1c82d7f35924c4c53c0a9ebab44ffbac0a05b3"
	//const auxPowHex = "01000000010000000000000000000000000000000000000000000000000000000000000000000000002cfabe6d6db3050aacfb4fb4ba9e0a3cc5c42459f3d7821cbde0e04fa647de219c9991fee70100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffff7f0000000000000000000000000000000000000000000000000000000000000000d55f1879a25b515208b06fdb1effd26478220fed3a284a5dd2d7e85cdafaa9575abd2f5b0000000002000000"

	var auxpow AuxPow

	buf, _ := common.HexStringToBytes(auxPowHex)
	if err := auxpow.Deserialize(bytes.NewReader(buf)); err != nil {
		t.Log("can not resolve auxpow parameter: ", auxPowHex)
	}

	hash, err := common.Uint256FromHexString(hashHex)
	if err != nil {
		t.Error(err)
	}
	if ok := auxpow.Check(hash, 6); !ok {
		t.Log("block check aux pow failed")
	}
}