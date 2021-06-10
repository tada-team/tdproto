package tdproto

// Color in ``#rrggbb`` format where ``rr``, ``gg``, ``bb`` are
// hexadecimal numbers from 00 to ff of red, green and blue channels
// correspondingly. (yellow would be ``#ffff00``)
type RGBColor = string

// Color theme
type Theme struct {
	// BgColor for web
	BgColor RGBColor

	// BgHoverColor for web
	BgHoverColor RGBColor

	// TextColor for web
	TextColor RGBColor

	// MutedTextColor for web
	MutedTextColor RGBColor

	// AccentColor for web
	AccentColor RGBColor

	// AccentHoverColor for web
	AccentHoverColor RGBColor

	// TextOnAccentHoverColor for web
	TextOnAccentHoverColor RGBColor

	// MainAccent for web
	MainAccent RGBColor

	// MainAccentHover for web
	MainAccentHover RGBColor

	// MainLightAccent for web
	MainLightAccent RGBColor

	// MainLink for web
	MainLink RGBColor

	// Brand color for app
	Brand RGBColor `json:"brand"`

	// BrandDark color for app
	BrandDark RGBColor `json:"brand_dark"`

	// Brand light color for app
	BrandLight RGBColor `json:"brand_light"`

	// Back light color for app
	Back RGBColor `json:"back"`

	// Back light color for app
	BackLight RGBColor `json:"back_light"`

	// Back dark color for app
	BackDark RGBColor `json:"back_dark"`

	// Success color for app
	Success RGBColor `json:"success"`

	// Success light color for app
	SuccessLight RGBColor `json:"success_light"`

	// Error color for app
	Error RGBColor `json:"error"`

	// Error light color for app
	ErrorLight RGBColor `json:"error_light"`

	// Background color for app
	Background RGBColor `json:"background"`

	// Tab background color for app
	TabBackground RGBColor `json:"tab_background"`

	// Chat input background color for app
	ChatInputBackground RGBColor `json:"chat_input_background"`

	// Substrate background color for app
	SubstrateBackground RGBColor `json:"substrate_background"`

	// Modal background color for app
	ModalBackground RGBColor `json:"modal_background"`

	// Title background color for app
	TitleBackground RGBColor `json:"title_background"`

	// Attention color for app
	Attention RGBColor `json:"attention"`

	// Attention light color for app
	AttentionLight RGBColor `json:"attention_light"`

	// Font colors for app
	Font *FontColors `json:"font"`

	// Message colors for app
	Message *MessageColors `json:"message"`

	// Switcher colors for app
	Switcher *SwitcherColors `json:"switcher"`

	// Button colors for app
	Button *ButtonColors `json:"button"`

	// Input colors for app
	Input *InputColors `json:"input"`

	// Icon colors for app
	Icon *IconColors `json:"ic"`

	// Deprecated
	AppAccentColor RGBColor

	// Deprecated
	AppPrimaryColor RGBColor
}

// Button colors for app
type ButtonColors struct {
	// Brand static color
	BrandStatic RGBColor `json:"brand_static"`

	// Brand active color
	BrandActive RGBColor `json:"brand_active"`

	// Brand disable color
	BrandDisable RGBColor `json:"brand_disable"`

	// Simple static color
	SimpleStatic RGBColor `json:"simple_static"`

	// Simple active color
	SimpleActive RGBColor `json:"simple_active"`

	// Simple disable color
	SimpleDisable RGBColor `json:"simple_disable"`
}

// Font colors for app
type FontColors struct {
	// Text color
	Text RGBColor `json:"text"`

	// Title color
	Title RGBColor `json:"title"`

	// Sub color
	Sub RGBColor `json:"sub"`

	// Brand button color
	BrandButton RGBColor `json:"brand_button"`

	// Simple button color
	SimpleButton RGBColor `json:"simple_button"`

	// Bubble sent color
	BubbleSent RGBColor `json:"bubble_sent"`

	// Bubble received color
	BubbleReceived RGBColor `json:"bubble_received"`
}

// Message colors for app
type MessageColors struct {
	// Bubble sent color
	BubbleSent RGBColor `json:"bubble_sent"`

	// Bubble received color
	BubbleReceived RGBColor `json:"bubble_received"`

	// Bubble important color
	BubbleImportant RGBColor `json:"bubble_important"`

	// Status feed color
	StatusFeed RGBColor `json:"status_feed"`

	// Status bubble color
	StatusBubble RGBColor `json:"status_bubble"`

	// Allocated color
	Allocated RGBColor `json:"allocated"`
}

// Input colors for app
type InputColors struct {
	// Static color
	Static RGBColor `json:"static"`

	// Active color
	Active RGBColor `json:"active"`

	// Disable color
	Disable RGBColor `json:"disable"`

	// Error color
	Error RGBColor `json:"error"`
}

// Switcher colors for app
type SwitcherColors struct {
	// On color
	On RGBColor `json:"on"`

	// Off color
	Off RGBColor `json:"off"`
}

// Icon colors for app
type IconColors struct {
	// Title color
	Title RGBColor `json:"title"`

	// Brand color
	Brand RGBColor `json:"brand"`

	// Other color
	Other RGBColor `json:"other"`
}
