package md

func CheckPermission(w string) bool {
	switch w {
	case "freely":
		return true
	case "editable":
		return true
	case "limited":
		return true
	case "locked":
		return true
	case "protected":
		return true
	case "private":
		return true
	default:
		return false
	}
}

func CheckCommentPermission(w string) bool {
	switch w {
	case "disable":
		return true
	case "forbidden":
		return true
	case "owners":
		return true
	case "signed_in_users":
		return true
	default:
		return false
	}
}

func CheckReadPermission(w string) bool {
	switch w {
	case "owner":
		return true
	case "sign_in":
		return true
	case "guest":
		return true
	default:
		return false
	}
}

func CheckWritePermission(w string) bool {
	switch w {
	case "owner":
		return true
	case "sign_in":
		return true
	case "guest":
		return true
	default:
		return false
	}
}
