package rpc

import (
	"github.com/sev-2/raiden"
)

type HandleNewUserParams struct {
}
type HandleNewUserResult interface{}

type HandleNewUser struct {
	raiden.RpcBase
	Params   *HandleNewUserParams `json:"-"`
	Return   HandleNewUserResult `json:"-"`
}

func (r *HandleNewUser) GetName() string {
	return "handle_new_user"
}

func (r *HandleNewUser) GetSecurity() raiden.RpcSecurityType {
	return raiden.RpcSecurityTypeDefiner
}

func (r *HandleNewUser) GetReturnType() raiden.RpcReturnDataType {
	return raiden.RpcReturnDataTypeTrigger
}

func (r *HandleNewUser) GetRawDefinition() string {
	return `DECLARE local_part TEXT; BEGIN local_part := split_part(NEW.email, '@', 1); INSERT INTO public.users (auth_id, full_name, email, role, password) VALUES (NEW.id, local_part, NEW.email, 'Patient', 'secret'); RETURN NEW; END;`
}