// Code generated by "stringer -type=Capability"; DO NOT EDIT.

package providers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CanAutoDNSSEC-0]
	_ = x[CanConcur-1]
	_ = x[CanGetZones-2]
	_ = x[CanUseAKAMAICDN-3]
	_ = x[CanUseAlias-4]
	_ = x[CanUseAzureAlias-5]
	_ = x[CanUseCAA-6]
	_ = x[CanUseDHCID-7]
	_ = x[CanUseDNAME-8]
	_ = x[CanUseDS-9]
	_ = x[CanUseDSForChildren-10]
	_ = x[CanUseHTTPS-11]
	_ = x[CanUseLOC-12]
	_ = x[CanUseNAPTR-13]
	_ = x[CanUseOPENPGPKEY-14]
	_ = x[CanUsePTR-15]
	_ = x[CanUseRoute53Alias-16]
	_ = x[CanUseSOA-17]
	_ = x[CanUseSRV-18]
	_ = x[CanUseSSHFP-19]
	_ = x[CanUseSVCB-20]
	_ = x[CanUseTLSA-21]
	_ = x[CanUseDNSKEY-22]
	_ = x[DocCreateDomains-23]
	_ = x[DocDualHost-24]
	_ = x[DocOfficiallySupported-25]
}

const _Capability_name = "CanAutoDNSSECCanConcurCanGetZonesCanUseAKAMAICDNCanUseAliasCanUseAzureAliasCanUseCAACanUseDHCIDCanUseDNAMECanUseDSCanUseDSForChildrenCanUseHTTPSCanUseLOCCanUseNAPTRCanUseOPENPGPKEYCanUsePTRCanUseRoute53AliasCanUseSOACanUseSRVCanUseSSHFPCanUseSVCBCanUseTLSACanUseDNSKEYDocCreateDomainsDocDualHostDocOfficiallySupported"

var _Capability_index = [...]uint16{0, 13, 22, 33, 48, 59, 75, 84, 95, 106, 114, 133, 144, 153, 164, 180, 189, 207, 216, 225, 236, 246, 256, 268, 284, 295, 317}

func (i Capability) String() string {
	if i >= Capability(len(_Capability_index)-1) {
		return "Capability(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Capability_name[_Capability_index[i]:_Capability_index[i+1]]
}
