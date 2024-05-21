package certification

import (
	"encoding/hex"
	"github.com/aviate-labs/agent-go/principal"
	"github.com/fxamacker/cbor/v2"
	"testing"
)

const SampleCert = "d9d9f7a364747265658301830182045820250f5e26868d9c1ea7ab29cbe9c15bf1c47c0d7605e803e39e375a7fe09c6ebb830183024e726571756573745f7374617475738301820458204b268227774ec77ff2b37ecb12157329d54cf376694bdd59ded7803efd82386f83025820edad510eaaa08ed2acd4781324e6446269da6753ec17760f206bbe81c465ff528301830183024b72656a6563745f636f64658203410383024e72656a6563745f6d6573736167658203584443616e69737465722069766733372d71696161612d61616161622d61616167612d63616920686173206e6f20757064617465206d6574686f64202772656769737465722783024673746174757382034872656a65637465648204582097232f31f6ab7ca4fe53eb6568fc3e02bc22fe94ab31d010e5fb3c642301f1608301820458203a48d1fc213d49307103104f7d72c2b5930edba8787b90631f343b3aa68a5f0a83024474696d65820349e2dc939091c696eb16697369676e6174757265583089a2be21b5fa8ac9fab1527e041327ce899d7da971436a1f2165393947b4d942365bfe5488710e61a619ba48388a21b16a64656c65676174696f6ea2697375626e65745f6964581dd77b2a2f7199b9a8aec93fe6fb588661358cf12223e9a3af7b4ebac4026b6365727469666963617465590231d9d9f7a26474726565830182045820ae023f28c3b9d966c8fb09f9ed755c828aadb5152e00aaf700b18c9c067294b483018302467375626e6574830182045820e83bb025f6574c8f31233dc0fe289ff546dfa1e49bd6116dd6e8896d90a4946e830182045820e782619092d69d5bebf0924138bd4116b0156b5a95e25c358ea8cf7e7161a661830183018204582062513fa926c9a9ef803ac284d620f303189588e1d3904349ab63b6470856fc4883018204582060e9a344ced2c9c4a96a0197fd585f2d259dbd193e4eada56239cac26087f9c58302581dd77b2a2f7199b9a8aec93fe6fb588661358cf12223e9a3af7b4ebac402830183024f63616e69737465725f72616e6765738203581bd9d9f781824a000000000020000001014a00000000002fffff010183024a7075626c69635f6b657982035885308182301d060d2b0601040182dc7c0503010201060c2b0601040182dc7c050302010361009933e1f89e8a3c4d7fdcccdbd518089e2bd4d8180a261f18d9c247a52768ebce98dc7328a39814a8f911086a1dd50cbe015e2a53b7bf78b55288893daa15c346640e8831d72a12bdedd979d28470c34823b8d1c3f4795d9c3984a247132e94fe82045820996f17bb926be3315745dea7282005a793b58e76afeb5d43d1a28ce29d2d158583024474696d6582034995b8aac0e4eda2ea16697369676e61747572655830ace9fcdd9bc977e05d6328f889dc4e7c99114c737a494653cb27a1f55c06f4555e0f160980af5ead098acc195010b2f7"

func TestSampleCert(t *testing.T) {
	for _, s := range []string{
		"00000000002000000101",
		"000000000020000C0101",
		"00000000002FFFFF0101",
	} {
		t.Run(s, func(t *testing.T) {
			rawCertificate := hexToBytes(SampleCert)
			var certificate Certificate
			if err := cbor.Unmarshal(rawCertificate, &certificate); err != nil {
				t.Fatal(err)
			}
			if err := VerifyCertificate(
				certificate,
				principal.Principal{
					Raw: hexToBytes(s),
				},
				hexToBytes(RootKey),
			); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestSubnetDelegateTestVector(t *testing.T) {
	rawCertificate, err := hex.DecodeString("D9D9F7A3647472656583018301830183024863616E6973746572830183018301830183024A000000000010000001018301830183024E6365727469666965645F6461746182035820619D02453B55BA8EA01DA1D26DF7083644EBE84A97AFC8D4ECE7F82453548EAA82045820D598630C2C94E80F8EDD451F3B7E942ACE5680B60FB5897A96A466D2F8FCF6F882045820FD96014A4A0368DBD8BD4BD05806C0F8C6BDFBDFBE6F6182B9E4963F18ADD55B8204582072379FD63D4B0A8E7D0C9F87316613DC1ADED56B7311F83213FF213F8B802F1C820458208079C8D69F2C1813E63B2488CD1C7D0BAA73ECE24AEAAD48348D7A1F9B8E868F82045820F7251F6708258AA9E995EE4A923E615908D40E9D4C57DFA84640B77B76D016D0820458205D38A89DDE470A2252C2F4060C42AD6EBCBE5C689CA68241C2858D1149B13386820458200BDF772535F123B48C553D3AE0B6B464277B8E2DDF887CDFE5F54B595352BD658204582055A1F078E350F151DDDE43AEDFBAC8647AC4B392F7141917824EFB1A956B79DB8301820458205C9ABFBA1DFE4D188D30474AF0A1BB796D8EFA15353973A0EEC3427675789C1783024474696D6582034980EFCBB1A5FEF5E616697369676E61747572655830A754B5AE47F254B23420F17E71265CDAF64D34BA002BA88578FDF3CD9B2436FE5A3B7A2245E6DCF4A574169EB72DD4DF6A64656C65676174696F6EA2697375626E65745F6964581DCF9D54E35F653FD0FD4FF05E9C020923B719429C9A3139BEAAF4871B026B6365727469666963617465590199D9D9F7A26474726565830182045820E82F4E336F033E15D337975AF1617E1030F4C6F9F53FCF89A48E861F75C5C98483018302467375626E6574830182045820FB286DDA6CB7FE72AF261F5C896D51CBF82F233679907A1CB7E7299B44F5E23D8302581DCF9D54E35F653FD0FD4FF05E9C020923B719429C9A3139BEAAF4871B02830183024F63616E69737465725F72616E6765738203581BD9D9F781824A000000000010000001014A00000000001FFFFF010183024A7075626C69635F6B657982035885308182301D060D2B0601040182DC7C0503010201060C2B0601040182DC7C0503020103610088EB824FC43459023B08806C56AC224EDEA54AF7A656D96F6E909906B7442AC65A60D2C0B831425B0376674430E48F1D09658CD3F86BDD8607199401422C8B641C43F58740F52B497136E70B62522AEF12A6DB95ECBA58123D44D9B2E852B40883024474696D6582034987A5E7EB83F2E3E416697369676E6174757265583091EC641476446FFA0AB613BE624664BFEC32F7AC20B7E943EA7DACE1B7247101CA5B3CD6DFF38E6276BD6A7AF6C0587F")
	if err != nil {
		t.Fatal(err)
	}
	rootPubKey, err := hex.DecodeString("308182301D060D2B0601040182DC7C0503010201060C2B0601040182DC7C05030201036100923A67B791270CD8F5320212AE224377CF407D3A8A2F44F11FED5915A97EE67AD0E90BC382A44A3F14C363AD2006640417B4BBB3A304B97088EC6B4FC87A25558494FC239B47E129260232F79973945253F5036FD520DDABD1E2DE57ABFB40CB")
	if err != nil {
		t.Fatal(err)
	}
	certifiedData := []byte{
		97, 157, 2, 69, 59, 85, 186, 142, 160, 29, 161, 210, 109, 247, 8, 54, 68, 235, 232, 74,
		151, 175, 200, 212, 236, 231, 248, 36, 83, 84, 142, 170,
	}
	var certificate Certificate
	if err := cbor.Unmarshal(rawCertificate, &certificate); err != nil {
		t.Fatal(err)
	}
	if err := VerifyCertifiedData(
		certificate,
		principal.MustDecode("5v3p4-iyaaa-aaaaa-qaaaa-cai"),
		rootPubKey,
		certifiedData,
	); err != nil {
		t.Fatal(err)
	}
}

func hexToBytes(s string) []byte {
	b, _ := hex.DecodeString(s)
	return b
}
