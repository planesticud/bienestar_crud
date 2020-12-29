package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Anamnesis struct {
	Id                     int              `orm:"column(id);pk"`
	HistoriaClinicaId      *HistoriaClinica `orm:"column(historia_clinica_id);rel(fk)"`
	Tratamiento            string           `orm:"column(tratamiento)"`
	Medicamentos           string           `orm:"column(medicamentos)"`
	Alergias               string           `orm:"column(alergias)"`
	Hemorragias            string           `orm:"column(hemorragias)"`
	Irradiaciones          string           `orm:"column(irradiaciones)"`
	Sinusitis              string           `orm:"column(sinusitis)"`
	EnfermedadRespiratoria string           `orm:"column(enfermedad_respiratoria)"`
	Cardiopatias           string           `orm:"column(cardiopatias)"`
	Diabetes               string           `orm:"column(diabetes)"`
	FiebreReumatica        string           `orm:"column(fiebre_reumatica)"`
	Hepatitis              string           `orm:"column(hepatitis)"`
	Hipertension           string           `orm:"column(hipertension)"`
	AntecedenteFamiliar    string           `orm:"column(antecedente_familiar)"`
	Cepillado              int              `orm:"column(cepillado)"`
	Seda                   int              `orm:"column(seda)"`
	Enjuague               int              `orm:"column(enjuague)"`
	Dulces                 string           `orm:"column(dulces)"`
	Fuma                   string           `orm:"column(fuma)"`
	Chicle                 string           `orm:"column(chicle)"`
	Otras                  string           `orm:"column(otras);null"`
	UltimaVisita           time.Time        `orm:"column(ultima_visita);type(date)"`
}

func (t *Anamnesis) TableName() string {
	return "anamnesis"
}

func init() {
	orm.RegisterModel(new(Anamnesis))
}

// AddAnamnesis insert a new Anamnesis into database and returns
// last inserted Id on success.
func AddAnamnesis(m *Anamnesis) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAnamnesisById retrieves Anamnesis by Id. Returns error if
// Id doesn't exist
func GetAnamnesisById(id int) (v *Anamnesis, err error) {
	o := orm.NewOrm()
	v = &Anamnesis{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAnamnesis retrieves all Anamnesis matches certain condition. Returns empty list if
// no records exist
func GetAllAnamnesis(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Anamnesis))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Anamnesis
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateAnamnesis updates Anamnesis by Id and returns error if
// the record to be updated doesn't exist
func UpdateAnamnesisById(m *Anamnesis) (err error) {
	o := orm.NewOrm()
	v := Anamnesis{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAnamnesis deletes Anamnesis by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAnamnesis(id int) (err error) {
	o := orm.NewOrm()
	v := Anamnesis{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Anamnesis{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
