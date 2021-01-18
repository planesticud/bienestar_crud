package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AccesoHistoriaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AccesoHistoriaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AccesoHistoriaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AccesoHistoriaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AccesoHistoriaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AccesoHistoriaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AccesoHistoriaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AccesoHistoriaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AccesoHistoriaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AccesoHistoriaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AnamnesisController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AnamnesisController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AnamnesisController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AnamnesisController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AnamnesisController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AnamnesisController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AnamnesisController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AnamnesisController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AnamnesisController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AnamnesisController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedenteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedenteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedenteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedenteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedenteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedentePsicologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedentePsicologicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedentePsicologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedentePsicologicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedentePsicologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedentePsicologicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedentePsicologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedentePsicologicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedentePsicologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AntecedentePsicologicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionAtencionMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionAtencionMedicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionAtencionMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionAtencionMedicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionAtencionMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionAtencionMedicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionAtencionMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionAtencionMedicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionAtencionMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionAtencionMedicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:AsignacionSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComportamientoConsultaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComportamientoConsultaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComportamientoConsultaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComportamientoConsultaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComportamientoConsultaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComportamientoConsultaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComportamientoConsultaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComportamientoConsultaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComportamientoConsultaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComportamientoConsultaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComposicionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComposicionFamiliarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComposicionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComposicionFamiliarController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComposicionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComposicionFamiliarController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComposicionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComposicionFamiliarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComposicionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ComposicionFamiliarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ControlPlacaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ControlPlacaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ControlPlacaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ControlPlacaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ControlPlacaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ControlPlacaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ControlPlacaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ControlPlacaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ControlPlacaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ControlPlacaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ConvencionOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ConvencionOdontogramaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ConvencionOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ConvencionOdontogramaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ConvencionOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ConvencionOdontogramaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ConvencionOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ConvencionOdontogramaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ConvencionOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ConvencionOdontogramaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatosEmergenciaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatosEmergenciaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatosEmergenciaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatosEmergenciaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatosEmergenciaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatosEmergenciaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatosEmergenciaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatosEmergenciaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatosEmergenciaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DatosEmergenciaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DiagnosticoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DiagnosticoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DiagnosticoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DiagnosticoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:DiagnosticoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:EspecialidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:EspecialidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:EspecialidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:EspecialidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:EspecialidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenDentalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenDentalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenDentalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenDentalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenDentalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenDentalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenDentalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenDentalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenDentalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenDentalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenEstomatologicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenEstomatologicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenEstomatologicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenEstomatologicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExamenEstomatologicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExcusaMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExcusaMedicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExcusaMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExcusaMedicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExcusaMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExcusaMedicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExcusaMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExcusaMedicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExcusaMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ExcusaMedicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HistoriaClinicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HistoriaClinicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HistoriaClinicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HistoriaClinicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HistoriaClinicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HojaHistoriaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HojaHistoriaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HojaHistoriaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HojaHistoriaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HojaHistoriaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HojaHistoriaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HojaHistoriaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HojaHistoriaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HojaHistoriaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:HojaHistoriaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:LimitesController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:LimitesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:LimitesController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:LimitesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:LimitesController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:LimitesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:LimitesController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:LimitesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:LimitesController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:LimitesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:MotivoAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:MotivoAvalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:MotivoAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:MotivoAvalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:MotivoAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:MotivoAvalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:MotivoAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:MotivoAvalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:MotivoAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:MotivoAvalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OdontogramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OdontogramaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OdontogramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OdontogramaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OdontogramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OdontogramaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OdontogramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OdontogramaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OdontogramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OdontogramaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OrdenController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OrdenController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OrdenController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OrdenController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OrdenController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OrdenController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OrdenController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OrdenController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OrdenController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:OrdenController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:RemisionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:RemisionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:RemisionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:RemisionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:RemisionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:RemisionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:RemisionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:RemisionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:RemisionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:RemisionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SignosVitalesController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SignosVitalesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SignosVitalesController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SignosVitalesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SignosVitalesController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SignosVitalesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SignosVitalesController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SignosVitalesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SignosVitalesController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SignosVitalesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SolicitudServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SolicitudServicioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SolicitudServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SolicitudServicioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SolicitudServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SolicitudServicioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SolicitudServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SolicitudServicioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SolicitudServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SolicitudServicioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SoporteAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SoporteAvalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SoporteAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SoporteAvalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SoporteAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SoporteAvalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SoporteAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SoporteAvalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SoporteAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SoporteAvalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SubtipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SubtipoServicioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SubtipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SubtipoServicioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SubtipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SubtipoServicioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SubtipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SubtipoServicioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SubtipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:SubtipoServicioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedenteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedenteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedenteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedenteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedenteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedentePsicologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedentePsicologicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedentePsicologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedentePsicologicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedentePsicologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedentePsicologicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedentePsicologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedentePsicologicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedentePsicologicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAntecedentePsicologicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAtencionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAtencionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAtencionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAtencionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAtencionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAtencionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAtencionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAtencionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAtencionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoAtencionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoServicioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoServicioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoServicioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoServicioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:TipoServicioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ValoracionInterpersonalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ValoracionInterpersonalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ValoracionInterpersonalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ValoracionInterpersonalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_crud/controllers:ValoracionInterpersonalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
