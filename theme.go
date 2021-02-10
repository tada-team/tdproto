package tdproto

type ButtonColors struct {
	BrandStatic   string `json:brand_static`
	BrandActive   string `json:brand_active`
	BrandDisable  string `json:brand_disable`
	SimpleStatic  string `json:simple_static`
	SimpleActive  string `json:simple_active`
	SimpleDisable string `json:simple_disable`
}

type InputColors struct {
	Static  string `json:static`
	Active  string `json:active`
	Disable string `json:disable`
	Error   string `json:error`
}
type SwitcherColors struct {
	On  string `json:on`
	Off string `json:off`
}

// Color theme
type Theme struct {
	// Web colors
	BgColor                string
	BgHoverColor           string
	TextColor              string
	MutedTextColor         string
	AccentColor            string
	AccentHoverColor       string
	TextOnAccentHoverColor string
	MainAccent             string
	MainAccentHover        string
	MainLightAccent        string
	MainLink               string

	// App colors
	AppAccentColor  string
	AppPrimaryColor string

	Brand          string `json:brand`
	BrandDark      string `json:brand_dark`
	BrandLight     string `json:brand_light`
	Text           string `json:text`
	Sub            string `json:sub`
	Back           string `json:back`
	BackLight      string `json:back_light`
	BackDark       string `json:back_dark`
	Error          string `json:error`
	ErrorLight     string `json:error_light`
	Success        string `json:success`
	SuccessLight   string `json:success_light`
	Backgroud      string `json:background`
	Attention      string `json:attention`
	AttentionLight string `json:attention_light`
	Button         ButtonColors
	Input          InputColors
	Switcher       SwitcherColors
}
