package color

func GetTColors() (string, string, string, string) {
	greenColor := "\033[32m"
	blueColor := "\033[34m"
	redColor := "\033[31m"
	defaultColor := "\033[0m"
	return greenColor, blueColor, redColor, defaultColor
}
