package bubbletea

import (
	"fmt"
	"miniolearn/config"
	"miniolearn/internal/bucket"
	"miniolearn/internal/firstrun"
	"miniolearn/internal/helper"
	"miniolearn/internal/policy"
	"miniolearn/internal/prompt"
	"miniolearn/internal/system"
	"miniolearn/internal/user"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

// ───── Model ─────
type model struct {
	Tabs       []string
	TabContent map[string][]string
	activeTab  int
	subCursor  int
	selected   string
	actionMap  map[string]func()
}

func (m model) Init() tea.Cmd {
	return nil
}

// ───── Update ─────
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "right", "l", "d", "tab":
			m.activeTab = min(m.activeTab+1, len(m.Tabs)-1)
			m.subCursor = 0
			m.selected = ""
			return m, nil
		case "left", "h", "a", "shift+tab":
			m.activeTab = max(m.activeTab-1, 0)
			m.subCursor = 0
			m.selected = ""
			return m, nil
		case "up", "w", "i":
			if m.subCursor > 0 {
				m.subCursor--
			}
		case "down", "s", "k":
			tabKey := m.Tabs[m.activeTab]
			if m.subCursor < len(m.TabContent[tabKey])-1 {
				m.subCursor++
			}
		case "enter", " ":
			tabKey := m.Tabs[m.activeTab]
			if items := m.TabContent[tabKey]; len(items) > 0 {
				selectedItem := items[m.subCursor]
				m.selected = selectedItem
				if _, ok := m.actionMap[selectedItem]; ok {
					return m, tea.Quit // ⬅ Exit Bubbletea when item selected
				}
			}
		}
	}
	return m, nil
}

// ───── Styling ─────
func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Border(activeTabBorder, true).Foreground(lipgloss.Color("229")).BorderForeground(highlightColor).Background(lipgloss.Color("57"))
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(1, 2).Align(lipgloss.Left).Border(lipgloss.NormalBorder()).UnsetBorderTop()
	highlightStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("229")).Background(lipgloss.Color("#7D56F4")).Bold(true)
	outputStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("36"))
)

// ───── View ─────
func (m model) View() string {
	var doc strings.Builder

	// 🔹 Show current MinIO alias if set
	if config.MinioAlias != "" {
		header := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			Background(lipgloss.Color("236")).
			Padding(0, 1).
			Render(fmt.Sprintf("🔗 Connected to MinIO Alias: %s", config.MinioAlias))

		doc.WriteString(header + "\n\n")
	}

	var renderedTabs []string
	for i, t := range m.Tabs {
		style := inactiveTabStyle
		if i == m.activeTab {
			style = activeTabStyle
		}
		border, _, _, _, _ := style.GetBorder()
		isFirst := i == 0
		isLast := i == len(m.Tabs)-1
		if isFirst && i == m.activeTab {
			border.BottomLeft = "│"
		} else if isFirst {
			border.BottomLeft = "├"
		}
		if isLast && i == m.activeTab {
			border.BottomRight = "│"
		} else if isLast {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}
	tabsRow := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(tabsRow + "\n")

	currentTab := m.Tabs[m.activeTab]
	items := m.TabContent[currentTab]
	var contentLines []string

	for i, item := range items {
		if i == m.subCursor {
			contentLines = append(contentLines, highlightStyle.Render("=> "+item))
		} else {
			contentLines = append(contentLines, "  "+item)
		}
	}

	if m.selected != "" {
		contentLines = append(contentLines, "")
		contentLines = append(contentLines, lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("42")).Render(m.selected))
	}

	contentBox := windowStyle.Width(lipgloss.Width(tabsRow)).Render(strings.Join(contentLines, "\n"))
	doc.WriteString(contentBox)

	return docStyle.Render(doc.String())
}

func BubbleCall() func() {
	tabs := []string{"User Management", "Bucket Management", "Policy Management", "System Configuration", "Misc"}
	tabContent := map[string][]string{
		"User Management":      {"List Users", "Details of User", "Create User", "Set User Password", "Enable User", "Disable User", "Remove User"},
		"Bucket Management":    {"List Buckets", "Create Bucket", "Remove Bucket"},
		"Policy Management":    {"List Policies", "Create Readonly Policy", "Create Read-Write Policy", "Create Full Access Policy", "Create All above policies", "Assign Policy to User", "Remove Policy"},
		"System Configuration": {"Setup MinIO Alias", "Prepare System", "Verify Directories", "Run Validation"},
		"Misc":                 {"Show Banner", "Add Server"},
	}
	actions := map[string]func(){
		"List Users": func() {
			users := user.GetUserList()
			if len(users) == 0 {
				fmt.Println("❌ No users found.")
			} else {
				prompt.PrintList("👤 Username", users)
			}
		},
		"List Buckets": func() {
			lists := bucket.Bucketlists()
			if len(lists) == 0 {
				fmt.Println("❌ No buckets found.")
			} else {
				prompt.PrintList("🪣 Available Buckets", lists)
			}
		},
		"List Policies": func() {
			lists := policy.GetPolicyList()
			if len(lists) == 0 {
				fmt.Println("❌ No policies found.")
			} else {
				prompt.PrintList("📜 Policies", lists)
			}
		},
		// User Management
		"Details of User":   user.UserDetails,
		"Create User":       user.CreateUser,
		"Set User Password": user.CreateUser,
		"Enable User":       user.UserEnable,
		"Disable User":      user.UserDisable,
		"Remove User":       user.UserDelete,
		// Bucket Management
		"Create Bucket": bucket.BucketCreate,
		"Remove Bucket": bucket.BucketDelete,
		// Policy Management
		"Create Readonly Policy": func() {
			buckets := bucket.Bucketlists()
			if len(buckets) == 0 {
				fmt.Println("❌ No buckets found.")
				return
			}
			selectedBucket := prompt.PromptSelectFromList("🪣 Available Buckets", buckets)
			helper.PolicyToBucket(selectedBucket, "readonly")
		},
		"Create Read-Write Policy": func() {
			buckets := bucket.Bucketlists()
			if len(buckets) == 0 {
				fmt.Println("❌ No buckets found.")
				return
			}
			selectedBucket := prompt.PromptSelectFromList("🪣 Available Buckets", buckets)
			helper.PolicyToBucket(selectedBucket, "readwrite")
		},
		"Create Full Access Policy": func() {
			buckets := bucket.Bucketlists()
			if len(buckets) == 0 {
				fmt.Println("❌ No buckets found.")
				return
			}
			selectedBucket := prompt.PromptSelectFromList("🪣 Available Buckets", buckets)
			helper.PolicyToBucket(selectedBucket, "readwritedelete")
		},
		"Create All above policies": func() {
			buckets := bucket.Bucketlists()
			if len(buckets) == 0 {
				fmt.Println("❌ No buckets found.")
				return
			}
			selectedBucket := prompt.PromptSelectFromList("🪣 Available Buckets", buckets)
			helper.PolicyToBucket(selectedBucket, "all")
		},
		"Assign Policy to User": func() {
			users := user.GetUserList()
			policies := policy.GetPolicyList()
			if len(users) == 0 {
				fmt.Println("❌ No users found.")
				return
			}
			if len(policies) == 0 {
				fmt.Println("❌ No policies found.")
				return
			}
			selectedUser := prompt.PromptSelectFromList("👤 Username", users)
			selectedPolicy := prompt.PromptSelectFromList("📜 Policies", policies)
			helper.PolicyToUser(selectedPolicy, selectedUser)
		},
		// System Configuration
		"Setup MinIO Alias":  firstrun.InitialSetup,
		"Prepare System":     firstrun.Directories,
		"Verify Directories": firstrun.McDirCheck,
		// Misc
		"Show Banner": system.OwnerBanner,
		"Add Server":  firstrun.InitialSetup,
	}
	m := model{
		Tabs:       tabs,
		TabContent: tabContent,
		activeTab:  0,
		subCursor:  0,
		selected:   "",
		actionMap:  actions,
	}
	prog, err := tea.NewProgram(m).Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	if final, ok := prog.(model); ok {
		if fn, exists := final.actionMap[final.selected]; exists {
			return fn
		}
	}
	// Fallback if nothing selected
	return func() {
		fmt.Println("No valid action selected.")
	}
}

// ───── Utilities ─────
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
