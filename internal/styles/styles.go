package styles

import "github.com/charmbracelet/lipgloss"

var (
	SectionStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFA500")).
			Bold(true).
			Blink(true)

	ItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A1E3D8"))

	OrderedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#A1E3D8")).
				Bold(true)

	PromptStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFFF")).
			Bold(true)

	BorderStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 4). // adds left/right padding inside the border
			Width(70).     // slightly wider to accommodate long values
			Align(lipgloss.Left)

	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF75B5")).
			Bold(true).
			Underline(true).
			MarginBottom(1).
			Align(lipgloss.Left)

	HeaderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A1E3D8")).
			Bold(true)

	ValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#C0C0C0"))

	DescriptionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#D4D4D4")).
				MarginTop(1).
				MarginBottom(1).
				Align(lipgloss.Left)

	DescriptionBoxStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#FFA500")). // a bright orange border
				Foreground(lipgloss.Color("#00FFFF")).       // neon cyan text
				Background(lipgloss.Color("#1A1A1A")).       // darker background for contrast
				Padding(1, 3, 1, 3).
				Margin(1, 0).
				Width(60).
				Bold(true)

	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFA500")). // a bright orange border
			Foreground(lipgloss.Color("#00FFFF")).       // neon cyan text
			Background(lipgloss.Color("#1A1A1A")).       // darker background for contrast
			Padding(1, 1, 1, 1).                         // top, right, bottom, left
			Margin(1, 0).
			Width(40).
			Bold(true).
			Blink(false)

	QuitBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF0000")). // intense red border
			Foreground(lipgloss.Color("#FF4C4C")).       // bright red text
			Background(lipgloss.Color("#1A1A1A")).       // deep blood-dark background
			Padding(1, 2).
			Margin(1, 1).
			Width(40).
			Bold(true).
			Blink(true) // blinking makes it feel urgent

	ErrorStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF0000")). // intense red border
			Foreground(lipgloss.Color("#FF4C4C")).       // bright red text
			Background(lipgloss.Color("#1A1A1A")).       // deep blood-dark background
			Padding(1, 2).
			Margin(1, 1).
			Width(40).
			Bold(true).
			Blink(true) // blinking makes it feel urgent
	SuccessStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00FF00")). // neon green border
			Foreground(lipgloss.Color("#ADFF2F")).       // green-yellow (bright and readable)
			Background(lipgloss.Color("#1A1A1A")).       // same dark background as ErrorStyle
			Padding(1, 2).
			Margin(1, 1).
			Width(40).
			Bold(true)
)
