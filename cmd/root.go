package cmd

import (
	"fmt"
	"miniolearn/internal/bucket"
	"miniolearn/internal/policy"
	"miniolearn/internal/prompt"
	"miniolearn/internal/styles"
	"miniolearn/internal/system"
	"miniolearn/internal/user"
	"miniolearn/internal/utils"
	"os"

	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func PrintMainMenu() {
	// ----- Styles -----

	// ----- Print Styled Menu -----
	fmt.Println(styles.HeaderStyle.Render("🛠️  MinIO Administration CLI - Main Menu"))
	userMgmt := lipgloss.JoinVertical(lipgloss.Left,
		styles.SectionStyle.Render("👤 User Management"),
		styles.ItemStyle.Render("  1) List Users"),
		styles.ItemStyle.Render("  2) Details of User"),
		styles.ItemStyle.Render("  3) Create User"),
		styles.ItemStyle.Render("  4) Set User Password"),
		styles.ItemStyle.Render("  5) Enable User"),
		styles.ItemStyle.Render("  6) Disable User"),
		styles.ItemStyle.Render("  7) Remove User"),
	)

	fmt.Println(styles.BoxStyle.Render(userMgmt))

	bucketMgmt := lipgloss.JoinVertical(lipgloss.Left,
		styles.SectionStyle.Render("🪣 Bucket Management"),
		styles.ItemStyle.Render("  8) List Buckets"),
		styles.ItemStyle.Render("  9) Create Bucket"),
		styles.ItemStyle.Render(" 10) Remove Bucket"),
	)
	fmt.Println(styles.BoxStyle.Render(bucketMgmt))

	policyMgmt := lipgloss.JoinVertical(lipgloss.Left,
		styles.SectionStyle.Render("📜 Policy Management"),
		styles.ItemStyle.Render(" 11) List Policies"),
		styles.ItemStyle.Render(" 12) Create Readonly Policy"),
		styles.ItemStyle.Render(" 13) Create Read-Write Policy"),
		styles.ItemStyle.Render(" 14) Create Full Access Policy"),
		styles.ItemStyle.Render(" 15) Assign Policy to User"),
		styles.ItemStyle.Render(" 16) Remove Policy"),
	)
	fmt.Println(styles.BoxStyle.Render(policyMgmt))

	systemConfig := lipgloss.JoinVertical(lipgloss.Left,
		styles.SectionStyle.Render("⚙️ System Configuration"),
		styles.ItemStyle.Render(" 17) Setup MinIO Alias"),
		styles.ItemStyle.Render(" 18) Prepare System"),
		styles.ItemStyle.Render(" 19) Verify Directories"),
		styles.ItemStyle.Render(" 20) Run Validation"),
	)
	fmt.Println(styles.BoxStyle.Render(systemConfig))

	miscMgmt := lipgloss.JoinVertical(lipgloss.Left,
		styles.SectionStyle.Render("🎨 Miscellaneous"),
		styles.ItemStyle.Render(" 21) Show Banner"),
	)
	fmt.Println(styles.BoxStyle.Render(miscMgmt))

	quitOption := lipgloss.JoinHorizontal(lipgloss.Center,
		styles.ItemStyle.Render("0) Q U I T"),
	)
	fmt.Println(styles.QuitBoxStyle.Render(quitOption))
	// ----- Get User Choice -----
	inputStr := prompt.PromptLine(styles.PromptStyle.Render("👉 Enter your choice (1–21): "))

	choice, err := strconv.Atoi(strings.TrimSpace(inputStr))
	if err != nil {
		fmt.Println("❌ Please enter a valid number.")
		return
	}

	fmt.Println()
	switch choice {
	case 1:
		lists := user.GetUserList()
		if len(lists) == 0 {
			fmt.Println("❌ No users found.")
		} else {
			prompt.PrintList("👤 Username", lists)
		}
	case 2:
		user.UserDetails()
	case 3:
		user.CreateUser()
	case 4:
		user.CreateUser()
	case 5:
		user.UserEnable()
	case 6:
		user.UserDisable()
	case 7:
		user.UserDelete()
	case 8:
		lists := bucket.Bucketlists()
		if len(lists) == 0 {
			fmt.Println("❌ No buckets found.")
		} else {
			prompt.PrintList("🪣 Available Buckets", lists)
		}
	case 9:
		bucket.BucketCreate()
	case 10:
		bucket.BucketDelete()
	case 11:
		lists := policy.GetPolicyList()
		if len(lists) == 0 {
			fmt.Println("❌ No policies found.")
		} else {
			prompt.PrintList("📜 Policies", lists)
		}
	case 21:
		system.OwnerBanner()
	case 0:
		utils.ClearScreen()
		fmt.Println(styles.QuitBoxStyle.Render("👋 Exiting... Goodbye!"))
		os.Exit(0)
	default:
		fmt.Println("⚠️ Only numbers 1–21 are supported.")
	}
}
