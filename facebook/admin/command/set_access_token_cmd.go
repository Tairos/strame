package command

import (
	"net/http"
)

const cmd_id = "set_"

type SetAccesTokenCmd struct{}

func (cmd *SetAccesTokenCmd) ID() string {
  return cmd_id
}

func (cmd *SetAccesTokenCmd) Handle(w http.ResponseWriter, r *http.Request) {

}
