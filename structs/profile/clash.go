package profile

type Profile struct {
	StoreSelected bool `json:"store-selected"`
	StoreFakeIP   bool `json:"store-fake-ip"`
}

type ClashGroup struct {
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	Proxies       []string `json:"proxies"`
	Url           string   `json:"url,omitempty"`
	Interval      int16    `json:"interval,omitempty"`
	Strategy      string   `json:"strategy,omitempty"`
	InterfaceName string   `json:"interfaceName,omitempty"`
	RoutingMark   int      `json:"routingMark,omitempty"`
	Use           []string `json:"use,omitempty"`
}

type PluginOpts struct {
	Mode           string            `json:"mode"`
	Host           string            `json:"host,omitempty"`
	TLS            bool              `json:"tls,omitempty"`
	SkipCertVerify bool              `json:"skip-cert-verify,omitempty"`
	Path           string            `json:"path,omitempty"`
	MUX            bool              `json:"mux,omitempty"`
	Headers        map[string]string `json:"headers,omitempty"`
}

type H2Opts struct {
	Path string   `json:"path,omitempty"`
	Host []string `json:"host,omitempty"`
}

type WSOpts struct {
	Path                string            `json:"path"`
	Headers             map[string]string `json:"headers,omitempty"`
	MaxEarlyData        int               `json:"max-early-data,omitempty"`
	EarlyDataHeaderName string            `json:"early-data-header-name,omitempty"`
}

type HttpOpts struct {
	Method  string                 `json:"method"`
	Path    []string               `json:"path"`
	Headers map[string]interface{} `json:"headers,omitempty"`
}

type GrpcOpts struct {
	GrpcServiceName string `json:"grpc-service-name"`
}

type ObfsOpts struct {
	Mode string `json:"mode"`
	Host string `json:"host,omitempty"`
}

type ClashProxy struct {
	Name           string     `json:"name,omitempty"`
	Type           string     `json:"type,omitempty"`
	InterfaceName  string     `json:"interface-name,omitempty"`
	RoutingMark    string     `json:"routing-mark,omitempty"`
	Server         string     `json:"server,omitempty"`
	Port           int16      `json:"port,omitempty"`
	Cipher         string     `json:"cipher,omitempty"`
	Username       string     `json:"username,omitempty"`
	Password       string     `json:"password"`
	UDP            bool       `json:"udp"`
	Plugin         string     `json:"plugin,omitempty"`
	PluginOpts     PluginOpts `json:"plugin_opts,omitempty"`
	Network        string     `json:"network,omitempty"`
	WSOpts         WSOpts     `json:"ws-opts,omitempty"`
	SkipCertVerify bool       `json:"skip-cert-verify,omitempty"`
	AlterId        int        `json:"alterId,omitempty"`
	Uuid           string     `json:"uuid,omitempty"`
	TLS            bool       `json:"tls,omitempty"`
	Servername     string     `json:"servername,omitempty"`
	H2Opts         H2Opts     `json:"h2-opts,omitempty"`
	HttpOpts       HttpOpts   `json:"http-opts"`
	GrpcOpts       GrpcOpts   `json:"grpc-opts,omitempty"`
	SNI            string     `json:"sni,omitempty"`
	PSK            string     `json:"psk,omitempty"`
	Obfs           string     `json:"obfs,omitempty"`
	ObfsOpts       ObfsOpts   `json:"obfs-opts,omitempty"`
	ObfsParam      string     `json:"obfs-param,omitempty"`
	Alpn           []string   `json:"alpn,omitempty"`
	Protocol       string     `json:"protocol,omitempty"`
	ProtocolParam  string     `json:"protocol-param,omitempty"`
}

type FallbackFilter struct {
	GeoIP     bool     `json:"geoip,omitempty"`
	GeoIPCode int      `json:"geoip-code,omitempty"`
	IPcidr    []string `json:"ipcidr,omitempty"`
	Domain    []string `json:"domain,omitempty"`
}

type DNS struct {
	Enable            bool              `json:"enable"`
	Listen            string            `json:"listen,omitempty"`
	IPv6              bool              `json:"ipv6"`
	DefaultNameserver []string          `json:"default-nameserver,omitempty"`
	EnhancedMode      string            `json:"enhanced-mode,omitempty"`
	FakeIPRange       string            `json:"fake-ip-range,omitempty"`
	UseHosts          bool              `json:"use-hosts,omitempty"`
	SearchDomains     []string          `json:"search-domains,omitempty"`
	FakeIPFilter      []string          `json:"fake-ip-filter,omitempty"`
	Nameserver        []string          `json:"nameserver,omitempty"`
	Fallback          []string          `json:"fallback,omitempty"`
	FallbackFilter    FallbackFilter    `json:"fallback-filter,omitempty"`
	NameserverPolicy  map[string]string `json:"nameserver-policy,omitempty"`
}

type HealthCheck struct {
	Enable   bool   `json:"enable"`
	Interval int    `json:"interval"`
	URL      string `json:"url"`
	Lazy     bool   `json:"lazy,omitempty"`
}

type ProxyProvider struct {
	Type        string      `json:"type"`
	URL         string      `json:"url,omitempty"`
	Interval    int         `json:"interval,omitempty"`
	Path        string      `json:"path"`
	HealthCheck HealthCheck `json:"health-check"`
}

type Clash struct {
	Port               int16                    `json:"port"`
	SocksPort          int16                    `json:"socks-port"`
	RedirPort          int16                    `json:"redir-port"`
	MixedPort          int16                    `json:"mixed-port"`
	TProxyPrt          int16                    `json:"tproxy-prt"`
	AllowLan           bool                     `json:"allow-lan"`
	Authentication     []string                 `json:"authentication,omitempty"`
	BindAddress        string                   `json:"bind-address"`
	Mode               string                   `json:"mode,omitempty"`
	LogLevel           string                   `json:"log-level,omitempty"`
	IPv6               bool                     `json:"ipv6"`
	ExternalController string                   `json:"external-controller,omitempty"`
	ExternalUI         string                   `json:"external-ui,omitempty"`
	Secret             string                   `json:"secret"`
	RoutingMark        int16                    `json:"routing-mark"`
	Hosts              map[string]string        `json:"hosts,omitempty"`
	Profile            Profile                  `json:"profile,omitempty"`
	InterfaceName      string                   `json:"interface-name"`
	DNS                DNS                      `json:"dns,omitempty"`
	Proxies            []ClashProxy             `json:"proxies"`
	ProxyGroups        []ClashGroup             `json:"proxy-groups"`
	Rules              []string                 `json:"rules"`
	ProxyProviders     map[string]ProxyProvider `json:"proxy-providers,omitempty"`
	Tunnels            []interface{}            `json:"tunnels,omitempty"`
}
