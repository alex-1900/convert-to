package sing_box

import "github.com/alex-1900/convert-to/config"

type Log struct {
	Disabled  bool   `json:"disabled"`
	Level     string `json:"level"`
	Output    string `json:"output"`
	Timestamp bool   `json:"timestamp"`
}

type DNSServer struct {
	Tag             string `json:"tag,omitempty"`
	Address         string `json:"address"`
	AddressResolver string `json:"address_resolver,omitempty"`
	AddressStrategy string `json:"address_strategy,omitempty"`
	Strategy        string `json:"strategy,omitempty"`
	Detour          string `json:"detour,omitempty"`
	ClientSubnet    string `json:"client_subnet,omitempty"`
}

type DNSRule struct {
	Inbound                  []string      `json:"inbound,omitempty"`
	IPVersion                int8          `json:"ip_version,omitempty"`
	QueryType                []interface{} `json:"query_type,omitempty"`
	Network                  string        `json:"network,omitempty"`
	AuthUser                 []string      `json:"auth_user,omitempty"`
	Protocol                 []string      `json:"protocol,omitempty"`
	Domain                   []string      `json:"domain,omitempty"`
	DomainSuffix             []string      `json:"domain_suffix,omitempty"`
	DomainKeyword            []string      `json:"domain_keyword,omitempty"`
	DomainRegex              []string      `json:"domain_regex,omitempty"`
	Geosite                  []string      `json:"geosite,omitempty"`
	SourceGeoIP              []string      `json:"source_geoip,omitempty"`
	GeoIP                    []string      `json:"geoip,omitempty"`
	SourceIPCidr             []string      `json:"source_ip_cidr,omitempty"`
	SourceIPIsPrivate        bool          `json:"source_ip_is_private,omitempty"`
	IPCidr                   []string      `json:"ip_cidr,omitempty"`
	IPIsPrivate              bool          `json:"ip_is_private,omitempty"`
	SourcePort               []int16       `json:"source_port,omitempty"`
	SourcePortRange          []string      `json:"source_port_range,omitempty"`
	Port                     []int16       `json:"port,omitempty"`
	PortRange                []string      `json:"port_range,omitempty"`
	ProcessName              []string      `json:"process_name,omitempty"`
	ProcessPath              []string      `json:"process_path,omitempty"`
	PackageName              []string      `json:"package_name,omitempty"`
	User                     []string      `json:"user,omitempty"`
	UserID                   []int         `json:"user_id,omitempty"`
	ClashMode                string        `json:"clash_mode,omitempty"`
	WifiSSID                 []string      `json:"wifi_ssid,omitempty"`
	WifiBssid                []string      `json:"wifi_bssid,omitempty"`
	RuleSet                  []string      `json:"rule_set,omitempty"`
	RuleSetIpCidrMatchSource bool          `json:"rule_set_ip_cidr_match_source,omitempty"`
	RuleSetIpCidrAcceptEmpty bool          `json:"rule_set_ip_cidr_accept_empty,omitempty"`
	Invert                   bool          `json:"invert,omitempty"`
	Outbound                 []string      `json:"outbound,omitempty"`
	Server                   string        `json:"server"`
	DisableCache             bool          `json:"disable_cache,omitempty"`
	RewriteTTL               int           `json:"rewrite_ttl,omitempty"`
	ClientSubnet             string        `json:"client_subnet,omitempty"`
	Type                     string        `json:"type,omitempty"`
	Mode                     string        `json:"mode,omitempty"`
}

type FakeIP struct {
	Enable     bool   `json:"enable,omitempty"`
	Inet4Range string `json:"inet4_range,omitempty"`
	Inet6Range string `json:"inet6_range,omitempty"`
}

type DNS struct {
	Servers          []DNSServer `json:"servers,omitempty"`
	Rules            []DNSRule   `json:"rules,omitempty"`
	Final            string      `json:"final,omitempty"`
	Strategy         string      `json:"strategy,omitempty"`
	DisableCache     bool        `json:"disable_cache,omitempty"`
	DisableExpire    bool        `json:"disable_expire,omitempty"`
	IndependentCache bool        `json:"independent,omitempty"`
	ReverseMapping   bool        `json:"reverse_mapping,omitempty"`
	ClientSubnet     string      `json:"client_subnet,omitempty"`
	FakeIP           FakeIP      `json:"fakeip,omitempty"`
}

type NTP struct {
}

type Inbound struct {
}

type Outbound struct {
}

type Route struct {
}

type Experimental struct {
}

type SingBox struct {
	content string

	Log          Log          `json:"log,omitempty"`
	DNS          DNS          `json:"dns,omitempty"`
	NTP          NTP          `json:"ntp,omitempty"`
	Inbound      Inbound      `json:"inbound,omitempty"`
	Outbound     Outbound     `json:"outbound"`
	Route        Route        `json:"route,omitempty"`
	Experimental Experimental `json:"experimental,omitempty"`
}

func (s SingBox) ProID() config.ProID {
	return config.ProIDSingBox
}

func (s SingBox) String() string {
	if s.content == "" {
		return "Sing-Box profile"
	}
	return s.content
}
