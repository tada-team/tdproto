package tdproto

type ButtonColors struct {
	BrandStatic   string `json:"brand_static"`
	BrandActive   string `json:"brand_active"`
	BrandDisable  string `json:"brand_disable"`
	SimpleStatic  string `json:"simple_static"`
	SimpleActive  string `json:"simple_active"`
	SimpleDisable string `json:"simple_disable"`
}

type FontColors struct {
	Text           string `json:"text"`
	Title          string `json:"title"`
	Sub            string `json:"sub"`
	BrandButton    string `json:"brand_button"`
	SimpleButton   string `json:"simple_button"`
	BubbleSent     string `json:"bubble_sent"`
	BubbleReceived string `json:"bubble_received"`
}

type MessageColors struct {
	BubbleSent      string `json:"bubble_sent"`
	BubbleReceived  string `json:"bubble_received"`
	BubbleImportant string `json:"bubble_important"`
	StatusFeed      string `json:"status_feed"`
	StatusBubble    string `json:"status_bubble"`
	Allocated       string `json:"allocated"`
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
	Brand               string          `json:"brand"`
	BrandDark           string          `json:"brand_dark"`
	BrandLight          string          `json:"brand_light"`
	Back                string          `json:"back"`
	BackLight           string          `json:"back_light"`
	BackDark            string          `json:"back_dark"`
	Success             string          `json:"success"`
	SuccessLight        string          `json:"success_light"`
	Error               string          `json:"error"`
	ErrorLight          string          `json:"error_light"`
	Background          string          `json:"background"`
	TabBackground       string          `json:"tab_background"`
	ChatInputBackground string          `json:"chat_input_background"`
	SubstrateBackground string          `json:"substrate_background"`
	ModalBackground     string          `json:"modal_background"`
	TitleBackground     string          `json:"title_background"`
	Attention           string          `json:"attention"`
	AttentionLight      string          `json:"attention_light"`
	Font                *FontColors     `json:"font"`
	Message             *MessageColors  `json:"message"`
	Switcher            *SwitcherColors `json:"switcher"`
	Button              *ButtonColors   `json:"button"`
	Input               *InputColors    `json:"input"`
	Icon                *IconColors     `json:"ic"`
}
