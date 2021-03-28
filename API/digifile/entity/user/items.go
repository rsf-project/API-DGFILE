package user

type (
	Folder struct {
		Folder_name string `json:"folder_name"`
		Directory   string `json:"directory"`
		Owner       string `json:"owner"`
	}
	Files struct {
		File_id      int    `json:"file_id"`
		File_name    string `json:"file_name"`
		Directory    string `json:"directory"`
		Size         string `json:"Size"`
		Trash_status bool   `json:"trash_status"`
		Owner        string `json:"owner"`
	}
	Items struct {
		Uid       string `json:"Uid"`
		Item_id   string `json:"item_id"`
		Parent_id string `json:"Parent_id"`
	}

	Download struct {
		Uid       string `json:"uid"`
		Folder_id string `json:"Folder_id"`
		File_name string `json:"File_name"`
	}
	Create_folder struct {
		Uid         string `json:"Iud"`
		Folder_name string `json:"Folder_name"`
		Parent_id   string `json:"Parent_id"`
	}
	Delete_file struct {
		File_id  string `json:"file_id"`
		Username string `json:"username"`
	}
	Delete_folder struct {
		Folder_id string `json:"folder_id"`
		Username  string `json:"username"`
	}
	Delete_trash struct {
		Uid     string `json:"Uid"`
		Item_id string `json:"id"`
	}
	Rename_file struct {
		Id       string `json:"id"`
		New_name string `json:"new_name"`
		Username string `json:"username"`
	}
	Rename_folder struct {
		Username      string `json:"username"`
		New_item_name string `json:"new_item_name"`
		Id            string `json:"id"`
	}
	Get_items struct {
		Uid         string `json:"uid"`
		Id_item     string `json:"id_item"`
		Curent_path string `json:"Curent_path"`
		Item_name   string `json:"item_name"`
	}
	Get_trash_folder struct {
		Id_item   string `json:"id"`
		Item_name string `json:"name"`
	}
	Get_trash_file struct {
		Id_item   string `json:"id"`
		Item_name string `json:"name"`
		Ekstensi  string `json:"ekstensi"`
	}
)
