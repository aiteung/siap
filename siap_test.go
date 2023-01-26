package siap

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/whatsauth"
)

var mssqlconn = atdb.DBInfo{
	DBString: os.Getenv("MSSQLSTRING"),
}

var usertables []whatsauth.LoginInfo
var usertable1 = whatsauth.LoginInfo{
	Userid:   "id",
	Password: "Password",
	Phone:    "Phone",
	Username: "Nama",
	Uuid:     "dbo.Pass",
	Login:    "md5",
}

func TestAtapi(t *testing.T) {
	usertables = append(usertables, usertable1)
	db := atdb.MssqlConnect(mssqlconn)
	fmt.Println(mssqlconn)
	loginfo := GetLoginInfofromPhoneNumber("6281312000300", usertables, db)
	fmt.Println(loginfo)

}
