package models

import (
	"errors"
	"fmt"
)

type Bgjob struct {
	Juuid    string `db:"job_uuid" json:"uuid"`
	Jcmd     string `db:"job_cmd" json:"cmd"`
	Jcmdarg  string `db:"job_cmdarg" json:"cmdarg"`
	Jcontent string `db:"job_content" json:"content"`
}

func CreateBgjob(in *Bgjob) error {
	var err error

	job := in
	if len(job.Juuid) == 0 {
		err = errors.New("uuid not null")
	} else {
		query := fmt.Sprintf("insert into cc_bgjobs(job_uuid,job_cmd,job_cmdarg,job_content)values('%s','%s','%s','%s')", job.Juuid, job.Jcmd, job.Jcmdarg, job.Jcontent)
		db.MustExec(query)
	}

	return err
}
