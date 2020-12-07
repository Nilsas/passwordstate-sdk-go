package secret

type Secret struct {
	PasswordListID              int    `json:"PasswordListID,omitempty"`
	PasswordID                  int    `json:"PasswordID,omitempty"`
	APIKey                      string `json:"APIKey,omitempty"`
	Title                       string `json:"Title,omitempty"`
	Domain                      string `json:"Domain,omitempty"`
	HostName                    string `json:"HostName,omitempty"`
	UserName                    string `json:"UserName,omitempty"`
	Description                 string `json:"Description,omitempty"`
	GenericField1               string `json:"GenericField1,omitempty"`
	GenericField2               string `json:"GenericField2,omitempty"`
	GenericField3               string `json:"GenericField3,omitempty"`
	GenericField4               string `json:"GenericField4,omitempty"`
	GenericField5               string `json:"GenericField5,omitempty"`
	GenericField6               string `json:"GenericField6,omitempty"`
	GenericField7               string `json:"GenericField7,omitempty"`
	GenericField8               string `json:"GenericField8,omitempty"`
	GenericField9               string `json:"GenericField9,omitempty"`
	GenericField10              string `json:"GenericField10,omitempty"`
	AccountTypeID               int    `json:"AccountTypeID,omitempty"`
	AccountType                 string `json:"AccountType,omitempty"`
	Notes                       string `json:"Notes,omitempty"`
	URL                         string `json:"URL,omitempty"`
	Password                    string `json:"password,omitempty"`
	ExpiryDate                  string `json:"ExpiryDate,omitempty"`
	AllowExport                 bool   `json:"AllowExport,omitempty"`
	WebUserID                   string `json:"WebUser_ID,omitempty"`
	WebPasswordID               string `json:"WebPassword_ID,omitempty"`
	GeneratePassword            bool   `json:"GeneratePassword,omitempty"`
	GenerateGenFieldPassword    bool   `json:"GenerateGenFieldPassword,omitempty"`
	PasswordResetEnabled        bool   `json:"PasswordResetEnabled,omitempty"`
	EnablePasswordResetSchedule bool   `json:"EnablePasswordResetSchedule,omitempty"`
	PasswordResetSchedule       bool   `json:"PasswordResetSchedule,omitempty"`
	AddDaysToExpiryDate         int    `json:"AddDaysToExpiryDate,omitempty"`
	ScriptID                    int    `json:"ScriptID,omitempty"`
	PrivilegedAccountID         int    `json:"PrivilegedAccountID,omitempty"`
	HeartbeatEnabled            bool   `json:"HeartbeatEnabled,omitempty"`
	HeartbeatSchedule           string `json:"HeartbeatSchedule,omitempty"`
	ValidationScriptID          int    `json:"ValidationScriptID,omitempty"`
	ADDomainNetBIOS             string `json:"ADDomainNetBIOS,omitempty"`
	ValidatewithPrivAccount     bool   `json:"ValidatewithPrivAccount,omitempty"`
}

type SecretResponse struct {
	PasswordID      int    `json:"PasswordID,omitempty"`
	Title           string `json:"Title,omitempty"`
	Username        string `json:"Username,omitempty"`
	Description     string `json:"Description,omitempty"`
	GenericField1   string `json:"GenericField1,omitempty"`
	GenericField2   string `json:"GenericField2,omitempty"`
	GenericField3   string `json:"GenericField3,omitempty"`
	GenericField4   string `json:"GenericField4,omitempty"`
	GenericField5   string `json:"GenericField5,omitempty"`
	GenericField6   string `json:"GenericField6,omitempty"`
	GenericField7   string `json:"GenericField7,omitempty"`
	GenericField8   string `json:"GenericField8,omitempty"`
	GenericField9   string `json:"GenericField9,omitempty"`
	GenericField10  string `json:"GenericField10,omitempty"`
	AccountTypeID   int    `json:"AccountTypeID,omitempty"`
	Notes           string `json:"Notes,omitempty"`
	URL             string `json:"URL,omitempty"`
	Password        string `json:"Password,omitempty"`
	ExpiryDate      string `json:"ExpiryDate,omitempty"`
	AllowExport     bool   `json:"AllowExport,omitempty"`
	AccountType     string `json:"AccountType,omitempty"`
	Status          string `json:"Status,omitempty"`
	CurrentPassword string `json:"CurrentPassword,omitempty"`
	NewPassword     string `json:"NewPassword,omitempty"`
}
