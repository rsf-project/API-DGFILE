package responsegraph

type ResponseGenericGet struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Id      []int       `json:"id"`
}
type Info_storage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Used    string `json:"used"`
	Total   string `json:"total"`
}
type Folder_trash struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"Data"`
}
type File_trash struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"Data"`
}
type Getbool struct {
	Status string `json:"status"`
	Data   bool   `json:"data"`
}
type ResponseGenericGet2 struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Data1   interface{} `json:"data1"`
	Id      []int       `json:"id"`
}

type ResponseGenericIn struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseValidation struct {
	Status  string      `json:"status"`
	Message bool        `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseBytes struct {
	Status  string      `json:"status"`
	Message []byte      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseArrayMultiType struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}

type Data struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    int           `json:"data"`
	Data1   []interface{} `json:"data1"`
	Uid     string        `json:"uid"`
}

type Logs struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    int    `json:"data"`
}
type Login struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    int    `json:"data"`
	Uid     string `json:"uid"`
}

type Items struct {
	Status   string   `json:"status"`
	Message  string   `json:"message"`
	Data     []string `json:"data"`
	File     int      `json:"file"`
	Ekstensi []string `json:"ekstensi"`
	Id       []int    `json:"id"`
}

type Userinformation struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Username string `json:"Username"`
	Name     string `json:"Name"`
	Phone    string `json:"Phone"`
	Email    string `json:"Email"`
	Space    int    `json:"Space"`
}

type Name struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
