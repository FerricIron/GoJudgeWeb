package api
func CheckPrivilege(user,privilege int)(bool){
	if user>=privilege{
		return true
	}
	return false
}
