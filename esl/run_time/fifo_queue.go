package run_time

//default queue is cool_fifo@$${domain}
//define queue fifos@fifomember;fifos@comsumer;
//fifo member manage function.

type fifomember struct {
	fifoname  string
	fmstring  string
	fmsimo    string
	fmtimeout string
	fmlag     string
}

func FifoMemberAdd(fm *fifomember) (apicmd string, e error) {
	var err error
	var mycmd string
	return mycmd, err
}

func FifoMemberDel(fm *fifomember) (apicmd string, e error) {
	var err error
	var mycmd string
	return mycmd, err
}
