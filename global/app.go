package global

// 应用状态的相关切换

const (
	// AppStatus 记录 app 的status
	AppStatus = "app-status"
)

const (
	// AppRunning app's running
	AppRunning = iota + 1
	// AppPause app's pause
	AppPause
	// AppStop app's stop
	AppStop
)

// StartApplication start the application(just set the status to running)
func StartApplication() {
	Store(AppStatus, AppRunning)
}

// PauseApplication pause the application(just set the status to pause)
func PauseApplication() {
	Store(AppStatus, AppPause)
}

// IsApplicationRunning check the application status is running
func IsApplicationRunning() bool {
	v, ok := Load(AppStatus)
	if !ok {
		return false
	}
	return v.(int) == AppRunning
}
