package tdproto

type ButtonColors struct {
	BrandStatic   string `json:"brand_static"`
	BrandActive   string `json:"brand_active"`
	BrandDisable  string `json:"brand_disable"`
	SimpleStatic  string `json:"simple_static"`
	SimpleActive  string `json:"simple_active"`
	SimpleDisable string `json:"simple_disable"`
}

type InputColors struct {
	Static  string `json:"static"`
	Active  string `json:"active"`
	Disable string `json:"disable"`
	Error   string `json:"error"`
}

type SwitcherColors struct {
	On  string `json:"on"`
	Off string `json:"off"`
}

type IconColors struct {
	Title string `json:"title"`
	Brand string `json:"brand"`
	Other string `json:"other"`
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

	// Deprecated
	AppAccentColor string

	// Deprecated
	AppPrimaryColor string

	// App colors
	Brand      string        `json:"brand"`
	Text       string        `json:"text"`
	Title      string        `json:"title"`
	Sub        string        `json:"sub"`
	BackLight  string        `json:"back_light"`
	Error      string        `json:"error"`
	Background string        `json:"background"`
	Attention  string        `json:"attention"`
	Button     *ButtonColors `json:"button"`
	Input      *InputColors  `json:"input"`
	Ic         *IconColors   `json:"ic"`

	// TODO: Not used, auxiliary colors, reserved for future use
	// BrandDark           string          `json:"brand_dark"`
	// BrandLight          string          `json:"brand_light"`
	// BackDark            string          `json:"back_dark"`
	// Back                string        `json:"back"`
	// ErrorLight          string          `json:"error_light"`
	// Success             string          `json:"success"`
	// SuccessLight        string          `json:"success_light"`
	// AttentionLight      string          `json:"attention_light"`
	// TabBackground       string          `json:"tab_background"`
	// ChatInputBackground string        `json:"chat_input_background"`
	// Switcher            *SwitcherColors `json:"switcher"`
}
