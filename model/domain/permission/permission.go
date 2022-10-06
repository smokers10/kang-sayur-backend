package domain

type Permission struct {
	ID             string `json:"id" bson:"_id"`
	MajBarang      bool   `json:"manajemen_barang" bson:"manajemen_barang"`
	MajPenjualan   bool   `json:"manajemen_penjualan" bson:"manajemen_penjualan"`
	MajPengguna    bool   `json:"manajemen_pengguna" bson:"manajemen_pengguna"`
	MajPenggunaAdm bool   `json:"manajemen_pengguna_adm" bson:"manajemen_pengguna_adm"`
	MajKeuangan    bool   `json:"manajemen_keuangan" bson:"manajemen_keuangan"`
	SubAdminID     string `json:"sub_admin_id" bson:"sub_admin_id"`
}

type PermissionRepository interface {
	Upsert(data *Permission) error

	ReadOne(sub_admin_id string) (*Permission, error)
}
