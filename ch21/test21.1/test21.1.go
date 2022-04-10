func Print(args ...interfaceP{}) string {
	for _, arg := range args {
		switch f := arg.(type) {
		case bool:
			var := arg.(bool)
		
		case float64:
			var := arg.(float64)
		case int:
			var := arg.(int)
		}
	}
}