package resolution

import (
	"encoding/json"
)

// contracts struct of contracts
type contracts map[string]struct {
	Address         string
	Implementation  string
	LegacyAddresses []string
	DeploymentBlock string
}

const (
	Mainnet string = "mainnet"
	Rinkeby string = "rinkeby"
	Polygon string = "polygon"
	Mumbai  string = "mumbai"
	Goerli  string = "goerli"
)

const (
	Layer1 string = "Layer 1"
	Layer2 string = "Layer 2"
)

type NetworkContracts map[string]contracts

var NetworkProviders = map[string]string{
	Mainnet: "https://eth-mainnet.g.alchemy.com/v2/RAQcwz7hhKhmwgoti6HYM_M_9nRJjEsQ",
	Rinkeby: "https://eth-goerli.g.alchemy.com/v2/r3soltxtCL_KhZGl7EwWsWax5ll_MFTN",
	Goerli:  "https://eth-goerli.g.alchemy.com/v2/r3soltxtCL_KhZGl7EwWsWax5ll_MFTN",
	Polygon: "https://polygon-mainnet.g.alchemy.com/v2/wh7r4O1amrfHhO-0-YiLa1Cg02JICqH2",
	Mumbai:  "https://polygon-mumbai.g.alchemy.com/v2/EP3SMW2f-2FMABuWuEdHdxKq_v1_ww82",
}

var NetworkNameToId = map[string]int{
	Mainnet: 1,
	Rinkeby: 4,
	Polygon: 137,
	Mumbai:  80001,
	Goerli:  5,
}

func parseContracts(data []byte) (contracts, error) {
	var contractsObject struct {
		Contracts contracts
	}
	err := json.Unmarshal(data, &contractsObject)
	if err != nil {
		return nil, err
	}
	return contractsObject.Contracts, nil
}

func newContracts() (NetworkContracts, error) {
	networks := make(NetworkContracts)
	var err error
	networks[Mainnet], err = parseContracts(unsMainnetConfigJSON)
	if err != nil {
		return nil, err
	}
	networks[Rinkeby], err = parseContracts(unsRinkebyConfigJSON)
	if err != nil {
		return nil, err
	}
	networks[Polygon], err = parseContracts(unsPolygonConfigJSON)
	if err != nil {
		return nil, err
	}
	networks[Mumbai], err = parseContracts(unsMumbaiConfigJSON)
	if err != nil {
		return nil, err
	}
	networks[Goerli], err = parseContracts(unsGoerliConfigJSON)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

var unsMainnetConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0x049aba7510f45BA5b64ea9E658E342F904DB358D",
			"implementation": "0xa715562307AA8AEDCba976b3793b3337F371c14a",
			"legacyAddresses": [],
			"deploymentBlock": "0xd62e9d",
			"forwarder": "0x049aba7510f45BA5b64ea9E658E342F904DB358D"
		},
		"CNSRegistry": {
			"address": "0xD1E5b0FF1287aA9f9A268759062E4Ab08b9Dacbe",
			"legacyAddresses": [],
			"deploymentBlock": "0x8a958b",
			"forwarder": "0x97B0E89fC1B7eD4A8B237D9d8Fcce9b234f25A37"
		},
		"MintingManager": {
			"address": "0x2a7084870bB724175a3C96Da8FaA55128fa3E19D",
			"implementation": "0x8caAeaD19aab5f54C94BB9F4be32e200E54AC8D7",
			"legacyAddresses": [],
			"deploymentBlock": "0xc2fee0",
			"forwarder": "0xb970fbCF52cd8111c76c379D4f2FE12E7f8AE7fb"
		},
		"ProxyAdmin": {
			"address": "0xAA16DA78110D9A9742c760a1a064F28654Ab93de",
			"legacyAddresses": [],
			"deploymentBlock": "0xc2fedc"
		},
		"SignatureController": {
			"address": "0x82EF94294C95aD0930055f31e53A34509227c5f7",
			"legacyAddresses": [],
			"deploymentBlock": "0x8a95a6"
		},
		"MintingController": {
			"address": "0xb0EE56339C3253361730F50c08d3d7817ecD60Ca",
			"legacyAddresses": [],
			"deploymentBlock": "0x8a95aa",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0xd3fF3377b0ceade1303dAF9Db04068ef8a650757",
			"legacyAddresses": [],
			"deploymentBlock": "0xa76ad3",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0x09B091492759737C03da9dB7eDF1CD6BCC3A9d91",
			"legacyAddresses": [],
			"deploymentBlock": "0x8a95ae",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0xeA70777e28E00E81f58b8921fC47F78B8a72eFE7",
			"legacyAddresses": [],
			"deploymentBlock": "0x98ca20",
			"deprecated": true
		},
		"Resolver": {
			"address": "0xb66DcE2DA6afAAa98F2013446dBCB0f4B0ab2842",
			"legacyAddresses": [
				"0xa1cac442be6673c49f8e74ffc7c4fd746f3cbd0d",
				"0x878bc2f3f717766ab69c0a5f9a6144931e61aed3"
			],
			"deploymentBlock": "0x960844",
			"forwarder": "0x486eb10E4F48C038513ECAf11585Ca2779768CF2"
		},
		"ProxyReader": {
			"address": "0x1BDc0fD4fbABeed3E611fd6195fCd5d41dcEF393",
			"legacyAddresses": [
				"0x58034A288D2E56B661c9056A0C27273E5460B63c",
				"0xc3C2BAB5e3e52DBF311b2aAcEf2e40344f19494E",
				"0xfEe4D4F0aDFF8D84c12170306507554bC7045878",
				"0xa6E7cEf2EDDEA66352Fd68E5915b60BDbb7309f5",
				"0x7ea9Ee21077F84339eDa9C80048ec6db678642B1"
			],
			"deploymentBlock": "0xde71cd"
		},
		"TwitterValidationOperator": {
			"address": "0x2F659766E3D08561CA3408FbAba7C0749ab2c402",
			"legacyAddresses": ["0xbb486C6E9cF1faA86a6E3eAAFE2e5665C0507855"],
			"deploymentBlock": "0xc300b5"
		},
		"FreeMinter": {
			"address": "0x1fC985cAc641ED5846b631f96F35d9b48Bc3b834",
			"legacyAddresses": [],
			"deploymentBlock": "0xacc390",
			"deprecated": true
		},
		"MintableERC721Predicate": {
			"address": "0x932532aA4c0174b8453839A6E44eE09Cc615F2b7",
			"legacyAddresses": [],
			"deploymentBlock": "0xa3cf69"
		},
		"RootChainManager": {
			"address": "0xA0c68C638235ee32657e8f720a23ceC1bFc77C77",
			"legacyAddresses": [],
			"deploymentBlock": "0xa3cf4d"
		}
	}
}`)

var unsRinkebyConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0x7fb83000B8eD59D3eAD22f0D584Df3a85fBC0086",
			"implementation": "0xc479D7A65243f7Eb1641F06a6C04E5F06cb5c4F7",
			"legacyAddresses": [],
			"deploymentBlock": "0x85e628",
			"forwarder": "0x7fb83000B8eD59D3eAD22f0D584Df3a85fBC0086"
		},
		"CNSRegistry": {
			"address": "0xAad76bea7CFEc82927239415BB18D2e93518ecBB",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232bc",
			"forwarder": "0xdf5CC97216785398D5C77348e68fc9461108f85d"
		},
		"MintingManager": {
			"address": "0xdAAf99A920D31F4f5720e4667b12b24e54A03070",
			"implementation": "0x38Fa95a0AC0E59D6e2845eFADBc17aF0FF9c7089",
			"legacyAddresses": [],
			"deploymentBlock": "0x85e629",
			"forwarder": "0xfB13e29C4D31a48B4Cd61131Cf3b681416e11681"
		},
		"ProxyAdmin": {
			"address": "0xaf9815005A208d1460b6fC60B4f90B9f2185E88c",
			"legacyAddresses": [],
			"deploymentBlock": "0x85e627"
		},
		"SignatureController": {
			"address": "0x66a5e3e2C27B4ce4F46BBd975270BE154748D164",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232be"
		},
		"MintingController": {
			"address": "0x51765307AeB3Df2E647014a2C501d5324212467c",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232bf",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0xbcB32f13f90978a9e059E8Cb40FaA9e6619d98e7",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232c6",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0xe1d2e4B9f0518CA5c803073C3dFa886470627237",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232c0",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0x6f8F96A566663C1d4fEe70edD37E9b62Fe39dE5D",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232c2",
			"deprecated": true
		},
		"Resolver": {
			"address": "0x95AE1515367aa64C462c71e87157771165B1287A",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232cf",
			"forwarder": "0xE172D8557d6F342b1b2976dE784F6Dff6ABC0a37"
		},
		"ProxyReader": {
			"address": "0xE6729D224D00b3dd4FC731C4Ee3274E35Da06578",
			"legacyAddresses": [
				"0x299974AeD8911bcbd2C61262605b89F591a53E83",
				"0x9F19473F6a98a715176291c930558E1954fd3D1e",
				"0x3A2e74CF832cbA3d77E72708d55370119E4323a6"
			],
			"deploymentBlock": "0x8dc79a"
		},
		"TwitterValidationOperator": {
			"address": "0x9ea4A63184ebE9CBA55CD1af473D98075Aa02b4C",
			"legacyAddresses": ["0x1CB337b3b208dc29a6AcE8d11Bb591b66c5Dd83d"],
			"deploymentBlock": "0x86935e"
		},
		"FreeMinter": {
			"address": "0x84214215904cDEbA9044ECf95F3eBF009185AAf4",
			"legacyAddresses": [],
			"deploymentBlock": "0x740d93",
			"deprecated": true
		}
	}
}`)

var unsPolygonConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0xa9a6A3626993D487d2Dbda3173cf58cA1a9D9e9f",
			"legacyAddresses": [],
			"deploymentBlock": "0x019d6188",
			"implementation": "0x5442953b0BFFf69FC945f5f1387cbFD2e2673447",
			"forwarder": "0xa9a6A3626993D487d2Dbda3173cf58cA1a9D9e9f"
		},
		"CNSRegistry": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"forwarder": "0x0000000000000000000000000000000000000000"
		},
		"MintingManager": {
			"address": "0x7be83293BeeDc9Eba1bd76c66A65F10F3efaeC26",
			"legacyAddresses": [],
			"deploymentBlock": "0x01272f41",
			"implementation": "0xBb45a6E10224Aa36EAcd812205F3763D353e9783",
			"forwarder": "0xC37d3c4326ab0E1D2b9D8b916bBdf5715f780fcF"
		},
		"ProxyAdmin": {
			"address": "0xe1D668052D52388F52b90f4d1798DB2b04bC3b88",
			"legacyAddresses": [],
			"deploymentBlock": "0x01272d15"
		},
		"SignatureController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"MintingController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"Resolver": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"forwarder": "0x0000000000000000000000000000000000000000"
		},
		"ProxyReader": {
			"address": "0x3E67b8c702a1292d1CEb025494C84367fcb12b45",
			"legacyAddresses": [
				"0x423F2531bd5d3C3D4EF7C318c2D1d9BEDE67c680",
				"0xA3f32c8cd786dc089Bd1fC175F2707223aeE5d00"
			],
			"deploymentBlock": "0x019d61a9"
		},
		"TwitterValidationOperator": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"FreeMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"MintableERC721Predicate": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"RootChainManager": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		}
	}
}`)

var unsMumbaiConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
			"legacyAddresses": [],
			"deploymentBlock": "0x0189f713",
			"implementation": "0xAc1a1F2136BfDe3a353a95C0676Cd0d55f311ee3",
			"forwarder": "0x2a93C52E7B6E7054870758e15A1446E769EdfB93"
		},
		"CNSRegistry": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"forwarder": "0x0000000000000000000000000000000000000000"
		},
		"MintingManager": {
			"address": "0x428189346bb3CC52f031A1092fd47C919AC30A9f",
			"legacyAddresses": [],
			"deploymentBlock": "0x01213f4a",
			"implementation": "0xCC17E698bA21bae4277579F22cA51135AaF00777",
			"forwarder": "0xEf3a491A8750BEC2Dff5339CF6Df94436d432C4d"
		},
		"ProxyAdmin": {
			"address": "0x460d63117c7Ab1624b7474C45BF46eC6702f57ce",
			"legacyAddresses": [],
			"deploymentBlock": "0x01213b22"
		},
		"SignatureController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"MintingController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"Resolver": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"forwarder": "0x0000000000000000000000000000000000000000"
		},
		"ProxyReader": {
			"address": "0x6fe7c857C1B0E54492C8762f27e0a45CA7ff264B",
			"legacyAddresses": [
				"0xbd9e01F6513E7C05f71Bf21d419a3bDF1EA9104b",
				"0x332A8191905fA8E6eeA7350B5799F225B8ed30a9"
			],
			"deploymentBlock": "0x0189f72d"
		},
		"TwitterValidationOperator": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"FreeMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"MintableERC721Predicate": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"RootChainManager": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		}
	}
}`)

var unsGoerliConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0x070e83FCed225184E67c86302493ffFCDB953f71",
			"implementation": "0x4473e84898E3F58feEFb7529dfF9E83Ff26CCae9",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57ea",
			"forwarder": "0x070e83FCed225184E67c86302493ffFCDB953f71"
		},
		"CNSRegistry": {
			"address": "0x801452cFAC27e79a11c6b185986fdE09e8637589",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57d7",
			"forwarder": "0x00443017FFaa4C840Caf5Dc7d3CB59147f363080"
		},
		"MintingManager": {
			"address": "0x9ee42D3EB042e06F8Cd241890C4fA0d51e4DA345",
			"implementation": "0xFB11410f3067BB6Db61bC335f0de23bE87A1767e",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57ec",
			"forwarder": "0x7F9F48cF94C69ce91D4b442DA186F31118ac0185"
		},
		"ProxyAdmin": {
			"address": "0xf4906E210523F9dA79E33811A44EE000441F4E04",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57e8"
		},
		"SignatureController": {
			"address": "0x5199dAE4B24B987ba18FcE1b64664D1B798d372B",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57d8"
		},
		"MintingController": {
			"address": "0xCEC41677be322049cC885c0DAe2fE0D52CA195ca",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57d9",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0x29465e3d2daA588E62375977bCe9b3f51406a794",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57da",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"Resolver": {
			"address": "0x0555344A5F440Bd1d8cb6B42db46c5e5D4070437",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57dc",
			"forwarder": "0xFCc1A95B7287Ae7a8B7cA813F12991dF5714d4C7"
		},
		"ProxyReader": {
			"address": "0xE3b961856C417d081a02cBa0161a051268F52677",
			"legacyAddresses": [
				"0x9A70ff906D422C2FD0F7B94244D6b36DB62Ee982",
				"0xFc5f608149f4D9e2Ed0733efFe9DD57ee24BCF68"
			],
			"deploymentBlock": "0x65bdfe"
		},
		"TwitterValidationOperator": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"FreeMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"MintableERC721Predicate": {
			"address": "0x56E14C4C1748a818a5564D33cF774c59EB3eDF59",
			"legacyAddresses": [],
			"deploymentBlock": "0x2fc240"
		},
		"RootChainManager": {
			"address": "0xBbD7cBFA79faee899Eaf900F13C9065bF03B1A74",
			"legacyAddresses": [],
			"deploymentBlock": "0x2dc9b9"
		}
	}
}
`)
