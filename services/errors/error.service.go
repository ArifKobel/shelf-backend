package errors_service

type ErrorMessages struct {
	FileNotFound        string
	DbConnectionFailed  string
	FailedToSaveFile    string
	FileUploadedSuccess string
	FailedToSaveUser    string
	EmailAlreadyExists  string
	Unauthorized        string
	UserNotFound        string
}

func GetErrorMessages() ErrorMessages {
	return ErrorMessages{
		FileNotFound:        "File not found",
		DbConnectionFailed:  "Database connection failed",
		FailedToSaveFile:    "Failed to save file",
		FileUploadedSuccess: "File uploaded successfully",
		FailedToSaveUser:    "Failed to save user",
		EmailAlreadyExists:  "Email already exists",
		Unauthorized:        "Unauthorized",
		UserNotFound:        "User not found",
	}
}
