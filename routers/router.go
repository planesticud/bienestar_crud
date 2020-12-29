// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/planesticud/bienestar_crud/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/convencion_odontograma",
			beego.NSInclude(
				&controllers.ConvencionOdontogramaController{},
			),
		),

		beego.NSNamespace("/tipo_antecedente",
			beego.NSInclude(
				&controllers.TipoAntecedenteController{},
			),
		),

		beego.NSNamespace("/soporte_aval",
			beego.NSInclude(
				&controllers.SoporteAvalController{},
			),
		),

		beego.NSNamespace("/hoja_historia",
			beego.NSInclude(
				&controllers.HojaHistoriaController{},
			),
		),

		beego.NSNamespace("/remision",
			beego.NSInclude(
				&controllers.RemisionController{},
			),
		),

		beego.NSNamespace("/signos_vitales",
			beego.NSInclude(
				&controllers.SignosVitalesController{},
			),
		),

		beego.NSNamespace("/orden",
			beego.NSInclude(
				&controllers.OrdenController{},
			),
		),

		beego.NSNamespace("/excusa_medica",
			beego.NSInclude(
				&controllers.ExcusaMedicaController{},
			),
		),

		beego.NSNamespace("/motivo_aval",
			beego.NSInclude(
				&controllers.MotivoAvalController{},
			),
		),

		beego.NSNamespace("/solicitud_servicio",
			beego.NSInclude(
				&controllers.SolicitudServicioController{},
			),
		),

		beego.NSNamespace("/dato_solicitud",
			beego.NSInclude(
				&controllers.DatoSolicitudController{},
			),
		),

		beego.NSNamespace("/acceso_historia",
			beego.NSInclude(
				&controllers.AccesoHistoriaController{},
			),
		),

		beego.NSNamespace("/tipo_antecedente_psicologico",
			beego.NSInclude(
				&controllers.TipoAntecedentePsicologicoController{},
			),
		),

		beego.NSNamespace("/diagnostico",
			beego.NSInclude(
				&controllers.DiagnosticoController{},
			),
		),

		beego.NSNamespace("/valoracion_interpersonal",
			beego.NSInclude(
				&controllers.ValoracionInterpersonalController{},
			),
		),

		beego.NSNamespace("/datos_emergencia",
			beego.NSInclude(
				&controllers.DatosEmergenciaController{},
			),
		),

		beego.NSNamespace("/comportamiento_consulta",
			beego.NSInclude(
				&controllers.ComportamientoConsultaController{},
			),
		),

		beego.NSNamespace("/limites",
			beego.NSInclude(
				&controllers.LimitesController{},
			),
		),

		beego.NSNamespace("/subtipo_servicio",
			beego.NSInclude(
				&controllers.SubtipoServicioController{},
			),
		),

		beego.NSNamespace("/antecedente",
			beego.NSInclude(
				&controllers.AntecedenteController{},
			),
		),

		beego.NSNamespace("/examen_dental",
			beego.NSInclude(
				&controllers.ExamenDentalController{},
			),
		),

		beego.NSNamespace("/antecedente_psicologico",
			beego.NSInclude(
				&controllers.AntecedentePsicologicoController{},
			),
		),

		beego.NSNamespace("/composicion_familiar",
			beego.NSInclude(
				&controllers.ComposicionFamiliarController{},
			),
		),

		beego.NSNamespace("/control_placa",
			beego.NSInclude(
				&controllers.ControlPlacaController{},
			),
		),

		beego.NSNamespace("/tipo_atencion",
			beego.NSInclude(
				&controllers.TipoAtencionController{},
			),
		),

		beego.NSNamespace("/tipo_servicio",
			beego.NSInclude(
				&controllers.TipoServicioController{},
			),
		),

		beego.NSNamespace("/persona",
			beego.NSInclude(
				&controllers.PersonaController{},
			),
		),

		beego.NSNamespace("/historia_clinica",
			beego.NSInclude(
				&controllers.HistoriaClinicaController{},
			),
		),

		beego.NSNamespace("/anamnesis",
			beego.NSInclude(
				&controllers.AnamnesisController{},
			),
		),

		beego.NSNamespace("/especialidad",
			beego.NSInclude(
				&controllers.EspecialidadController{},
			),
		),

		beego.NSNamespace("/odontograma",
			beego.NSInclude(
				&controllers.OdontogramaController{},
			),
		),

		beego.NSNamespace("/examen_estomatologico",
			beego.NSInclude(
				&controllers.ExamenEstomatologicoController{},
			),
		),

		beego.NSNamespace("/asignacion_solicitud",
			beego.NSInclude(
				&controllers.AsignacionSolicitudController{},
			),
		),

		beego.NSNamespace("/asignacion_atencion_medica",
			beego.NSInclude(
				&controllers.AsignacionAtencionMedicaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
