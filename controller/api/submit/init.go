package submit

import (
	"fmt"
	"github.com/ferriciron/GoJudgeWeb/common"
)

const judgeServerAddress="127.0.0.1"
const judgeServerPort="8080"
var addr=fmt.Sprintf("%s:%s",judgeServerAddress,judgeServerPort)
func openConnect()(*common.Socket,error){
	return common.Dial(addr)
}
func init(){

}
