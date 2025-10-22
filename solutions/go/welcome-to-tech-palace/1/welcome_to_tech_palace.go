package techpalace
import "strings"


func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    var borderMessage = strings.Repeat("*", numStarsPerLine)
    
	return borderMessage + "\n" + welcomeMsg + "\n" + borderMessage
}

func CleanupMessage(oldMsg string) string {
	clean := strings.ReplaceAll(oldMsg, "*", "")
    
	return strings.TrimSpace(clean)
}
