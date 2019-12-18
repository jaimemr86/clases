package clases

type ClsCatalogo struct {
	IdCatalogoDeObras 		int64 `json:"IdCatalogoDeObras"`
	IdCodigoSql 			int64 `json:"IdCodigoSql"`
	IdCodigoNube 			int64 `json:"IdCodigoNube"`
	Codigo 					string `json:"Codigo"`
	CodigoSap 				string `json:"CodigoSap"`
	Descripcion 			string `json:"Descripcion"`
	DescripcionLarga 		string `json:"DescripcionLarga"`
	EsAgrupador 			bool `json:"EsAgrupador"`
	EsPorcentaje 			int64 `json:"EsPorcentaje"`
	IdFamilia 				int64 `json:"IdFamilia"`
	IdFichaTecnica 			int64 `json:"IdFichaTecnica"`
	IdImagen 				int64 `json:"IdImagen"`
	IdProcedimiento 		int64 `json:"IdProcedimiento"`
	IdProveedor 			int64 `json:"IdProveedor"`
	IdTipoInsumo 			int64 `json:"IdTipoInsumo"`
	IdUnidad 				int64 `json:"IdUnidad"`
	InsumoDescontinuado 	bool `json:"InsumoDescontinuado"`
	PorcentajeFondoGarantia float64 `json:"PorcentajeFondoGarantia"`
	Referencia 				string `json:"Referencia"`
	VolumenDefault 			float64 `json:"VolumenDefault"`
	NoActualizaCatalogo 	bool `json:"NoActualizaCatalogo"`
}

type ClsRegresaCatalogo struct{
	ListaIds  map[int64]int64 `json:"ListaIds"`
	Errores ClsErrores `json:"Errores"`
}

type ClsEnviaListaCatalogo struct {
	AccessToken ClsAccessToken    `json:"AccessToken"`
	Catalogo    []ClsCatalogo `json:"Catalogo"`
}