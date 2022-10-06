package subadmin

type Create struct {
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	Status      string        `json:"status"`
	Position    string        `json:"position"`
	Permissions SetPermission `json:"permission"`
}

type Update struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Position string `json:"position"`
}

type Delete struct {
	ID string `json:"id"`
}

type UpdateStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type SetPermission struct {
	MajBarang      bool   `json:"manajemen_barang"`
	MajPenjualan   bool   `json:"manajemen_penjualan"`
	MajPengguna    bool   `json:"manajemen_pengguna"`
	MajPenggunaAdm bool   `json:"manajemen_pengguna_adm"`
	MajKeuangan    bool   `json:"manajemen_keuangan"`
	SubAdminID     string `json:"sub_admin_id"`
}
