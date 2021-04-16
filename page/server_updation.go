package page

type ServerUpdation struct {
	ServerInsertion
	ServerDeletion
}

func (upd ServerUpdation) ExecuteServer() (err error) {
	if err = upd.ServerDeletion.ExecuteServer(); nil != err {return}
	return upd.ServerInsertion.ExecuteServer()
}

func (upd ServerUpdation) ExecuteTags() (err error) {
	if err = upd.ServerDeletion.ExecuteTags(); nil != err {return}
	return upd.ServerInsertion.ExecuteTags()
}