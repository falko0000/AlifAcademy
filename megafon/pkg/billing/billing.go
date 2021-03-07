package billing

// M new 1000
func Calculate1000(traffic int) int {
	subscriptionFee := 3500
	costAdditionalPackage := 6
	if traffic <= 1000 {
		return subscriptionFee
	}
	result := subscriptionFee + (traffic-1000)*costAdditionalPackage
	return result
}

// M new 2000
func Calculate2000(traffic int) int {
	subscriptionFee := 5500
	costAdditionalPackage := 6
	if traffic <= 2000 {
		return subscriptionFee
	}
	result := subscriptionFee + (traffic-2000)*costAdditionalPackage
	return result
}

// M new 3000
func Calculate3000(traffic int) int {
	subscriptionFee := 7000
	costAdditionalPackage := 6
	if traffic <= 3000 {
		return subscriptionFee
	}
	result := subscriptionFee + (traffic-3000)*costAdditionalPackage
	return result
}

// M new 5000
func Calculate5000(traffic int) int {
	subscriptionFee := 9500
	costAdditionalPackage := 6
	if traffic <= 5000 {
		return subscriptionFee
	}
	result := subscriptionFee + (traffic-5000)*costAdditionalPackage
	return result
}

// M new 10000
func Calculate10000(traffic int) int {
	subscriptionFee := 17000
	costAdditionalPackage := 6
	if traffic <= 10000 {
		return subscriptionFee
	}
	result := subscriptionFee + (traffic-10000)*costAdditionalPackage
	return result
}
