package system

import (
	"fmt"
	"miniolearn/internal/styles"
	"miniolearn/internal/utils"

	"github.com/charmbracelet/lipgloss"
)

func OwnerBanner() {
	utils.ClearScreen()
	// Style for the border box

	// Title
	header := styles.TitleStyle.Render("Author & Project Information")

	// Compose info lines with consistent label width (15 chars)
	infoLines := []string{
		fmt.Sprintf("%s %s", styles.HeaderStyle.Render(fmt.Sprintf("%-15s:", "Author")), styles.ValueStyle.Render("Birat Aryal")),
		fmt.Sprintf("%s %s", styles.HeaderStyle.Render(fmt.Sprintf("%-15s:", "Email")), styles.ValueStyle.Render("birataryal2@gmail.com")),
		fmt.Sprintf("%s %s", styles.HeaderStyle.Render(fmt.Sprintf("%-15s:", "Website")), styles.ValueStyle.Render("https://birataryal.com.np")),
		fmt.Sprintf("%s %s", styles.HeaderStyle.Render(fmt.Sprintf("%-15s:", "Blog")), styles.ValueStyle.Render("https://birataryal.github.io")),
		fmt.Sprintf("%s %s", styles.HeaderStyle.Render(fmt.Sprintf("%-15s:", "Role")), styles.ValueStyle.Render("Infrastructure DevOps Engineer")),
		fmt.Sprintf("%s %s", styles.HeaderStyle.Render(fmt.Sprintf("%-15s:", "Project")), styles.ValueStyle.Render("MinIO Administration CLI Tool")),
		fmt.Sprintf("%s %s", styles.HeaderStyle.Render(fmt.Sprintf("%-15s:", "Version")), styles.ValueStyle.Render("v1.0.0")),
		fmt.Sprintf("%s %s", styles.HeaderStyle.Render(fmt.Sprintf("%-15s:", "License")), styles.ValueStyle.Render("MIT License")),
	}

	info := lipgloss.JoinVertical(lipgloss.Left, infoLines...)

	description := styles.DescriptionBoxStyle.Render(styles.DescriptionStyle.Render(
		"This CLI tool provides comprehensive administration capabilities for MinIO deployments, including user, bucket, and policy management, designed for ease of use and extensibility in production environments. Thank you for using this tool! For questions or feedback, please contact the author via email or website.",
	),
	)

	banner := styles.BorderStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left, header, info, description),
	)

	fmt.Println(banner)
}
