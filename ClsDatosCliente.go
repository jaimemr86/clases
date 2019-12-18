package clases

import (
	"cloud.google.com/go/civil"
)

type ClsDatosCliente struct {
	AccessToken 		string `json:"AccessToken"`
	RefreshToken 		string `json:"RefreshToken"`
	Name 				string `json:"Name"`
	Email 				string `json:"Email"`
	Picture 			string `json:"Picture"`
	ExpiresIn 			string `json:"ExpiresIn"`
	OauthExpireIn 		string `json:"OauthExpireIn"`
	FechaVigencia 		string `json:"FechaVigencia"`
	NumeroLicencia 		int64 `json:"NumeroLicencia"`
	RazonSocial 		string `json:"RazonSocial"`
	Sesion 				int64 `json:"Sesion"`
	OtraSesionActiva 	bool `json:"OtraSesionActiva"`
	CaducoSesion 		bool `json:"CaducoSesion"`
	TokenCaducado 		bool `json:"TokenCaducado"`
	NoTieneVigencia 	bool `json:"NoTieneVigencia"`
	CambioIp 			bool `json:"CambioIp"`
	IdUsuario 			int64 `json:"IdUsuario"`
	IdUsuarioAdmin 		int64 `json:"IdUsuarioAdmin"`
	EmailAdmin 			string `json:"EmailAdmin"`
	IdEmpresa 			int64 `json:"IdEmpresa"`
	LicenciaEstudiantil bool `json:"LicenciaEstudiantil"`
	Errores 			ClsErrores `json:"Errores"`
}

type ClsDatosLicenciaCliente struct {
	IdLicencia          int64      `json:"idLicencia"`
	IdSistema           int64      `json:"idSistema"`
	IdUsuario           int64      `json:"idUsuario"`
	FechaVigencia       civil.Date `json:"fechaVigencia"`
	NumeroLicencia      int64      `json:"numeroLicencia"`
	RazonSocial         string     `json:"razonSocial"`
	RefreshToken        string     `json:"refreshToken"`
	LicenciaEstudiantil bool       `json:"LicenciaEstudiantil"`
	Errores             ClsErrores `json:"Errores"`
}

type ClsSesion struct {
	IdSesionActiva int64 `json:"IdSesionActiva"`
	TieneActiva bool `json:"TieneActiva"`
	IpPublica string `json:"IpPublica"`
	Errores ClsErrores `json:"Errores"`
}

type ClsUsuarioAdmin struct {
	IdUsuario           int64      `json:"IdUsuario"`
	Email               string     `json:"Email"`
	IdUsuarioAdmin      int64      `json:"IdUsuarioAdmin"`
	EmailAdmin          string     `json:"EmailAdmin"`
	IdEmpresa           int64      `json:"IdEmpresa"`
	LicenciaEstudiantil bool       `json:"LicenciaEstudiantil"`
	Errores             ClsErrores `json:"Errores"`
}

