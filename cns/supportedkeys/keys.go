package supportedkeys

import (
	"bytes"
	"github.com/spf13/viper"
)

func NewConfig() (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigType("json")
	err := config.ReadConfig(bytes.NewBuffer(rawJson))
	if err != nil {
		return nil, err
	}
	return config, nil
}

var rawJson = []byte(`
{
    "version": "1.1.0",
    "keys": {
        "crypto.BTC.address": {
            "deprecatedKeyName": "BTC"
        },
        "crypto.ETH.address": {
            "deprecatedKeyName": "ETH"
        },
        "crypto.ZIL.address": {
            "deprecatedKeyName": "ZIL"
        },
        "crypto.LTC.address": {
            "deprecatedKeyName": "LTC"
        },
        "crypto.ETC.address": {
            "deprecatedKeyName": "ETC",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.EQL.address": {
            "deprecatedKeyName": "EQL",
            "validationRegex": "^bnb[0-9a-z]{39}$"
        },
        "crypto.LINK.address": {
            "deprecatedKeyName": "LINK",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.USDC.address": {
            "deprecatedKeyName": "USDC",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.BAT.address": {
            "deprecatedKeyName": "BAT",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.REP.address": {
            "deprecatedKeyName": "REP",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.ZRX.address": {
            "deprecatedKeyName": "ZRX",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.DAI.address": {
            "deprecatedKeyName": "DAI",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.BCH.address": {
            "deprecatedKeyName": "BCH",
            "validationRegex": "^[13][a-km-zA-HJ-NP-Z1-9]{33}$|^((bitcoincash|bchreg|bchtest):)?(q|p)[a-z0-9]{41}$|^((BITCOINCASH:)?(Q|P)[A-Z0-9]{41})$"
        },
        "crypto.XMR.address": {
            "deprecatedKeyName": "XMR",
            "validationRegex": "^[48]{1}[0-9AB][1-9A-HJ-NP-Za-km-z]{93}$"
        },
        "crypto.DASH.address": {
            "deprecatedKeyName": "DASH",
            "validationRegex": "^X[1-9A-HJ-NP-Za-km-z]{33}$"
        },
        "crypto.NEO.address": {
            "deprecatedKeyName": "NEO",
            "validationRegex": "^A[0-9a-zA-Z]{33}$"
        },
        "crypto.SWTH.address": {
            "deprecatedKeyName": "SWTH",
            "validationRegex": "^A[0-9a-zA-Z]{33}$"
        },
        "crypto.DOGE.address": {
            "deprecatedKeyName": "DOGE",
            "validationRegex": "^D[5-9A-HJ-NP-U]{1}[1-9A-HJ-NP-Za-km-z]{32}$"
        },
        "crypto.XRP.address": {
            "deprecatedKeyName": "XRP",
            "validationRegex": "^r[1-9a-km-zA-HJ-NP-Z]{24,34}$"
        },
        "crypto.ZEC.address": {
            "deprecatedKeyName": "ZEC",
            "validationRegex": "^z([a-zA-Z0-9]){94}$|^zs([a-zA-Z0-9]){75}$|^t([a-zA-Z0-9]){34}$"
        },
        "crypto.ADA.address": {
            "deprecatedKeyName": "ADA",
            "validationRegex": "^[1-9a-km-zA-HJ-NP-Z]{104}$|^A[1-9A-HJ-NP-Za-km-z]{58}$|^addr[0-9a-zA-Z]{99}$"
        },
        "crypto.EOS.address": {
            "deprecatedKeyName": "EOS",
            "validationRegex": "^[a-z][a-z1-5.]{10}[a-z1-5]$"
        },
        "crypto.XLM.address": {
            "deprecatedKeyName": "XLM",
            "validationRegex": "^G[A-Z2-7]{55}$"
        },
        "crypto.BNB.address": {
            "deprecatedKeyName": "BNB",
            "validationRegex": "^bnb[0-9a-z]{39}$"
        },
        "crypto.BTG.address": {
            "deprecatedKeyName": "BTG",
            "validationRegex": "^[GA][a-km-zA-HJ-NP-Z1-9]{33}$"
        },
        "crypto.NANO.address": {
            "deprecatedKeyName": "NANO",
            "validationRegex": "^nano_[1-9a-z]{60}$"
        },
        "crypto.WAVES.address": {
            "deprecatedKeyName": "WAVES",
            "validationRegex": "^3[a-km-zA-HJ-NP-Z1-9]{34}$"
        },
        "crypto.KMD.address": {
            "deprecatedKeyName": "KMD",
            "validationRegex": "^R[a-km-zA-Z1-9]{33}$"
        },
        "crypto.AE.address": {
            "deprecatedKeyName": "AE",
            "validationRegex": "^ak_[a-km-zA-Z1-9]{48,52}$"
        },
        "crypto.RSK.address": {
            "deprecatedKeyName": "RSK",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.WAN.address": {
            "deprecatedKeyName": "WAN",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.STRAT.address": {
            "deprecatedKeyName": "STRAT",
            "validationRegex": "^S[a-km-zA-HJ-NP-Z1-9]{33}$"
        },
        "crypto.UBQ.address": {
            "deprecatedKeyName": "UBQ",
            "validationRegex": "^0x[a-km-zA-HJ-NP-Z0-9]{40}$"
        },
        "crypto.XTZ.address": {
            "deprecatedKeyName": "XTZ",
            "validationRegex": "^(tz|KT)[a-km-zA-HJ-NP-Z1-9]{34}$"
        },
        "crypto.IOTA.address": {
            "deprecatedKeyName": "IOTA",
            "validationRegex": "^[A-Z0-9]{90}$"
        },
        "crypto.VET.address": {
            "deprecatedKeyName": "VET",
            "validationRegex": "^0x[a-km-zA-HJ-NP-Z0-9]{40}$"
        },
        "crypto.QTUM.address": {
            "deprecatedKeyName": "QTUM",
            "validationRegex": "^Q[a-km-zA-HJ-NP-Z1-9]{33}$"
        },
        "crypto.ICX.address": {
            "deprecatedKeyName": "ICX",
            "validationRegex": "^[a-km-zA-HJ-NP-Z0-9]{42}$"
        },
        "crypto.DGB.address": {
            "deprecatedKeyName": "DGB",
            "validationRegex": "^[a-km-zA-HJ-NP-Z1-9]{34}$|^[a-zA-Z1-9]{42}$"
        },
        "crypto.XZC.address": {
            "deprecatedKeyName": "XZC",
            "validationRegex": "^[a-km-zA-HJ-NP-Z1-9]{34}$"
        },
        "crypto.BURST.address": {
            "deprecatedKeyName": "BURST",
            "validationRegex": "^BURST-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{5}"
        },
        "crypto.DCR.address": {
            "deprecatedKeyName": "DCR"
        },
        "crypto.XEM.address": {
            "deprecatedKeyName": "XEM",
            "validationRegex": "^N[ABCDEFGHIJKLMNOPQRSTUVWXYZ234567]{39}$"
        },
        "crypto.LSK.address": {
            "deprecatedKeyName": "LSK",
            "validationRegex": "^\\d{1,21}[L]$"
        },
        "crypto.ATOM.address": {
            "deprecatedKeyName": "ATOM",
            "validationRegex": "^(cosmos)1([qpzry9x8gf2tvdw0s3jn54khce6mua7l]+)$"
        },
        "crypto.ONG.address": {
            "deprecatedKeyName": "ONG",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.ONT.address": {
            "deprecatedKeyName": "ONT",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.SMART.address": {
            "deprecatedKeyName": "SMART",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.TPAY.address": {
            "deprecatedKeyName": "TPAY",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.GRS.address": {
            "deprecatedKeyName": "GRS",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.BSV.address": {
            "deprecatedKeyName": "BSV",
            "validationRegex": "^bitcoincash:[a-zA-Z0-9]{42}$"
        },
        "crypto.GAS.address": {
            "deprecatedKeyName": "GAS"
        },
        "crypto.TRX.address": {
            "deprecatedKeyName": "TRX",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.VTHO.address": {
            "deprecatedKeyName": "VTHO",
            "validationRegex": "^[a-zA-Z0-9]{42}$"
        },
        "crypto.BCD.address": {
            "deprecatedKeyName": "BCD",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.BTT.address": {
            "deprecatedKeyName": "BTT",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.KIN.address": {
            "deprecatedKeyName": "KIN",
            "validationRegex": "^[a-zA-Z0-9]{56}$"
        },
        "crypto.RVN.address": {
            "deprecatedKeyName": "RVN",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.ARK.address": {
            "deprecatedKeyName": "ARK",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.XVG.address": {
            "deprecatedKeyName": "XVG",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.ALGO.address": {
            "deprecatedKeyName": "ALGO",
            "validationRegex": "^[a-zA-Z0-9]{58}$"
        },
        "crypto.NEBL.address": {
            "deprecatedKeyName": "NEBL",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.XPM.address": {
            "deprecatedKeyName": "XPM",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.ONE.address": {
            "deprecatedKeyName": "ONE",
            "validationRegex": "^one[a-zA-Z0-9]{39}$"
        },
        "crypto.BNTY.address": {
            "deprecatedKeyName": "BNTY",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.CRO.address": {
            "deprecatedKeyName": "CRO",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.TWT.address": {
            "deprecatedKeyName": "TWT",
            "validationRegex": "^bnb[0-9a-z]{39}$"
        },
        "crypto.SIERRA.address": {
            "deprecatedKeyName": "SIERRA",
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.VSYS.address": {
            "deprecatedKeyName": "VSYS",
            "validationRegex": "^[a-zA-Z0-9]{35}$"
        },
        "crypto.HIVE.address": {
            "deprecatedKeyName": "HIVE",
            "validationRegex": "^(?!s*$).+"
        },
        "crypto.HT.address": {
            "deprecatedKeyName": "HT",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.ENJ.address": {
            "deprecatedKeyName": "ENJ",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.YFI.address": {
            "deprecatedKeyName": "YFI",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.MTA.address": {
            "deprecatedKeyName": "MTA",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.COMP.address": {
            "deprecatedKeyName": "COMP",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.BAL.address": {
            "deprecatedKeyName": "BAL",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.AMPL.address": {
            "deprecatedKeyName": "AMPL",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.LEND.address": {
            "deprecatedKeyName": "LEND",
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.TLOS.address": {
            "deprecatedKeyName": "TLOS",
            "validationRegex": "^[a-z][a-z1-5.]{10}[a-z1-5]$"
        },
        "crypto.USDT.version.ERC20.address": {
            "deprecatedKeyName": "USDT_ERC20"
        },
        "crypto.USDT.version.TRON.address": {
            "deprecatedKeyName": "USDT_TRON"
        },
        "crypto.USDT.version.EOS.address": {
            "deprecatedKeyName": "USDT_EOS",
            "validationRegex": "^[a-z][a-z1-5.]{10}[a-z1-5]$"
        },
        "crypto.USDT.version.OMNI.address": {
            "deprecatedKeyName": "USDT_OMNI"
        },
        "social.payid.name": {
            "deprecatedKeyName": "payid",
            "validationRegex": "^[0-9a-zA-Z]+\\$[0-9a-zA-Z]+\\.[0-9a-zA-Z]+$"
        },
        "whois.email.value": {
            "deprecatedKeyName": "email",
            "validationRegex": "^[^@]+@[^\\.]+\\..+$"
        },
        "whois.for_sale.value": {
            "deprecatedKeyName": "for_sale",
            "validationRegex": "(true)|(false)"
        },
        "ipfs.html.value": {
            "deprecatedKeyName": "html",
            "validationRegex": ".{0,100}"
        },
        "ipfs.redirect_domain.value": {
            "deprecatedKeyName": "redirect_domain",
            "validationRegex": ".{0,253}"
        },
        "gundb.username.value": {
            "deprecatedKeyName": "gundb_username"
        },
        "gundb.public_key.value": {
            "deprecatedKeyName": "gundb_public_key"
        },
        "social.image.value": {
            "deprecatedKeyName": "image"
        },
        "social.twitter.username": {
            "deprecatedKeyName": "twitter_username"
        },
        "validation.social.twitter.username": {
            "deprecatedKeyName": "validation_twitter_username"
        },
        "dns.ttl": {
            "deprecatedKeyName": "ttl"
        },
        "dns.A": {
            "deprecatedKeyName": "A"
        },
        "dns.A.ttl": {
            "deprecatedKeyName": "A_ttl"
        },
        "dns.AAAA": {
            "deprecatedKeyName": "AAAA"
        },
        "dns.AAAA.ttl": {
            "deprecatedKeyName": "AAAA_ttl"
        },
        "dns.AFSDB": {
            "deprecatedKeyName": "AFSDB"
        },
        "dns.AFSDB.ttl": {
            "deprecatedKeyName": "AFSDB_ttl"
        },
        "dns.APL": {
            "deprecatedKeyName": "APL"
        },
        "dns.APL.ttl": {
            "deprecatedKeyName": "APL_ttl"
        },
        "dns.CAA": {
            "deprecatedKeyName": "CAA"
        },
        "dns.CAA.ttl": {
            "deprecatedKeyName": "CAA_ttl"
        },
        "dns.CDNSKEY": {
            "deprecatedKeyName": "CDNSKEY"
        },
        "dns.CDNSKEY.ttl": {
            "deprecatedKeyName": "CDNSKEY_ttl"
        },
        "dns.CDS": {
            "deprecatedKeyName": "CDS"
        },
        "dns.CDS.ttl": {
            "deprecatedKeyName": "CDS_ttl"
        },
        "dns.CERT": {
            "deprecatedKeyName": "CERT"
        },
        "dns.CERT.ttl": {
            "deprecatedKeyName": "CERT_ttl"
        },
        "dns.CNAME": {
            "deprecatedKeyName": "CNAME"
        },
        "dns.CNAME.ttl": {
            "deprecatedKeyName": "CNAME_ttl"
        },
        "dns.CSYNC": {
            "deprecatedKeyName": "CSYNC"
        },
        "dns.CSYNC.ttl": {
            "deprecatedKeyName": "CSYNC_ttl"
        },
        "dns.DHCID": {
            "deprecatedKeyName": "DHCID"
        },
        "dns.DHCID.ttl": {
            "deprecatedKeyName": "DHCID_ttl"
        },
        "dns.DLV": {
            "deprecatedKeyName": "DLV"
        },
        "dns.DLV.ttl": {
            "deprecatedKeyName": "DLV_ttl"
        },
        "dns.DNAME": {
            "deprecatedKeyName": "DNAME"
        },
        "dns.DNAME.ttl": {
            "deprecatedKeyName": "DNAME_ttl"
        },
        "dns.DNSKEY": {
            "deprecatedKeyName": "DNSKEY"
        },
        "dns.DNSKEY.ttl": {
            "deprecatedKeyName": "DNSKEY_ttl"
        },
        "dns.DS": {
            "deprecatedKeyName": "DS"
        },
        "dns.DS.ttl": {
            "deprecatedKeyName": "DS_ttl"
        },
        "dns.EUI48": {
            "deprecatedKeyName": "EUI48"
        },
        "dns.EUI48.ttl": {
            "deprecatedKeyName": "EUI48_ttl"
        },
        "dns.EUI64": {
            "deprecatedKeyName": "EUI64"
        },
        "dns.EUI64.ttl": {
            "deprecatedKeyName": "EUI64_ttl"
        },
        "dns.HINFO": {
            "deprecatedKeyName": "HINFO"
        },
        "dns.HINFO.ttl": {
            "deprecatedKeyName": "HINFO_ttl"
        },
        "dns.HIP": {
            "deprecatedKeyName": "HIP"
        },
        "dns.HIP.ttl": {
            "deprecatedKeyName": "HIP_ttl"
        },
        "dns.HTTPS": {
            "deprecatedKeyName": "HTTPS"
        },
        "dns.HTTPS.ttl": {
            "deprecatedKeyName": "HTTPS_ttl"
        },
        "dns.IPSECKEY": {
            "deprecatedKeyName": "IPSECKEY"
        },
        "dns.IPSECKEY.ttl": {
            "deprecatedKeyName": "IPSECKEY_ttl"
        },
        "dns.KEY": {
            "deprecatedKeyName": "KEY"
        },
        "dns.KEY.ttl": {
            "deprecatedKeyName": "KEY_ttl"
        },
        "dns.KX": {
            "deprecatedKeyName": "KX"
        },
        "dns.KX.ttl": {
            "deprecatedKeyName": "KX_ttl"
        },
        "dns.LOC": {
            "deprecatedKeyName": "LOC"
        },
        "dns.LOC.ttl": {
            "deprecatedKeyName": "LOC_ttl"
        },
        "dns.MX": {
            "deprecatedKeyName": "MX"
        },
        "dns.MX.ttl": {
            "deprecatedKeyName": "MX_ttl"
        },
        "dns.NAPTR": {
            "deprecatedKeyName": "NAPTR"
        },
        "dns.NAPTR.ttl": {
            "deprecatedKeyName": "NAPTR_ttl"
        },
        "dns.NS": {
            "deprecatedKeyName": "NS"
        },
        "dns.NS.ttl": {
            "deprecatedKeyName": "NS_ttl"
        },
        "dns.NSEC": {
            "deprecatedKeyName": "NSEC"
        },
        "dns.NSEC.ttl": {
            "deprecatedKeyName": "NSEC_ttl"
        },
        "dns.NSEC3": {
            "deprecatedKeyName": "NSEC3"
        },
        "dns.NSEC3.ttl": {
            "deprecatedKeyName": "NSEC3_ttl"
        },
        "dns.NSEC3PARAM": {
            "deprecatedKeyName": "NSEC3PARAM"
        },
        "dns.NSEC3PARAM.ttl": {
            "deprecatedKeyName": "NSEC3PARAM_ttl"
        },
        "dns.OPENPGPKEY": {
            "deprecatedKeyName": "OPENPGPKEY"
        },
        "dns.OPENPGPKEY.ttl": {
            "deprecatedKeyName": "OPENPGPKEY_ttl"
        },
        "dns.PTR": {
            "deprecatedKeyName": "PTR"
        },
        "dns.PTR.ttl": {
            "deprecatedKeyName": "PTR_ttl"
        },
        "dns.RP": {
            "deprecatedKeyName": "RP"
        },
        "dns.RP.ttl": {
            "deprecatedKeyName": "RP_ttl"
        },
        "dns.RRSIG": {
            "deprecatedKeyName": "RRSIG"
        },
        "dns.RRSIG.ttl": {
            "deprecatedKeyName": "RRSIG_ttl"
        },
        "dns.SIG": {
            "deprecatedKeyName": "SIG"
        },
        "dns.SIG.ttl": {
            "deprecatedKeyName": "SIG_ttl"
        },
        "dns.SMIMEA": {
            "deprecatedKeyName": "SMIMEA"
        },
        "dns.SMIMEA.ttl": {
            "deprecatedKeyName": "SMIMEA_ttl"
        },
        "dns.SOA": {
            "deprecatedKeyName": "SOA"
        },
        "dns.SOA.ttl": {
            "deprecatedKeyName": "SOA_ttl"
        },
        "dns.SRV": {
            "deprecatedKeyName": "SRV"
        },
        "dns.SRV.ttl": {
            "deprecatedKeyName": "SRV_ttl"
        },
        "dns.SSHFP": {
            "deprecatedKeyName": "SSHFP"
        },
        "dns.SSHFP.ttl": {
            "deprecatedKeyName": "SSHFP_ttl"
        },
        "dns.SVCB": {
            "deprecatedKeyName": "SVCB"
        },
        "dns.SVCB.ttl": {
            "deprecatedKeyName": "SVCB_ttl"
        },
        "dns.TA": {
            "deprecatedKeyName": "TA"
        },
        "dns.TA.ttl": {
            "deprecatedKeyName": "TA_ttl"
        },
        "dns.TKEY": {
            "deprecatedKeyName": "TKEY"
        },
        "dns.TKEY.ttl": {
            "deprecatedKeyName": "TKEY_ttl"
        },
        "dns.TLSA": {
            "deprecatedKeyName": "TLSA"
        },
        "dns.TLSA.ttl": {
            "deprecatedKeyName": "TLSA_ttl"
        },
        "dns.TSIG": {
            "deprecatedKeyName": "TSIG"
        },
        "dns.TSIG.ttl": {
            "deprecatedKeyName": "TSIG_ttl"
        },
        "dns.TXT": {
            "deprecatedKeyName": "TXT"
        },
        "dns.TXT.ttl": {
            "deprecatedKeyName": "TXT_ttl"
        },
        "dns.URI": {
            "deprecatedKeyName": "URI"
        },
        "dns.URI.ttl": {
            "deprecatedKeyName": "URI_ttl"
        },
        "dns.ZONEMD": {
            "deprecatedKeyName": "ZONEMD"
        },
        "dns.ZONEMD.ttl": {
            "deprecatedKeyName": "ZONEMD_ttl"
        }
    }
}
`)
