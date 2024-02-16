package goodies

type Produk interface {
	Tabungan() string
	pinjaman() string
}

type Cabang struct {
	NamaPinjaman string
	NamaSimpanan string
}

type Unit struct {
	NamaPinjaman string
	NamaSimpanan string
	ProdukMikro  string
}

func (c Cabang) Tabungan() string {
	return (c.NamaSimpanan)
}

func (u Unit) Tabungan() string {
	return (u.NamaSimpanan + ", " + u.ProdukMikro)
}

func (c Cabang) pinjaman() string {
	return (c.NamaPinjaman)
}

func (u Unit) pinjaman() string {
	return (u.NamaPinjaman + ", " + u.ProdukMikro)
}
