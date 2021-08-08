package model

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/gofiber_svelte/db"
)

type Mcompany struct {
	IdCompany         string    `json:"idcompany"`
	Idcurr            string    `json:"idcurr"`
	NmCompany         string    `json:"nmcompany"`
	NmOwner           string    `json:"nmowner"`
	Status            string    `json:"statuscompany"`
	CreateCompany     string    `json:"createcompany"`
	CreateDateCompany time.Time `json:"createdatecompany"`
	UpdateCompany     string    `json:"updatecompany"`
	UpdateDateCompany time.Time `json:"updatedatecompany"`
}

func FetchAll_mcompany() (Response, error) {
	var obj Mcompany
	var arraobj []Mcompany
	var res Response
	con := db.CreateCon()

	sql := "SELECT * FROM tbl_mst_company"

	rows, err := con.Query(sql)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&obj.IdCompany,
			&obj.Idcurr,
			&obj.NmCompany,
			&obj.NmOwner,
			&obj.Status,
			&obj.CreateCompany,
			&obj.CreateDateCompany,
			&obj.UpdateCompany,
			&obj.UpdateDateCompany)
		if err != nil {
			return res, err
		}
		arraobj = append(arraobj, obj)
	}
	res.Status = fiber.StatusOK
	res.Message = "Success"
	res.Record = arraobj

	return res, nil
}
