package siap

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var mssqlconn = atdb.DBInfo{
	DBString: os.Getenv("MSSQLSTRING"),
}

func TestAtapi(t *testing.T) {
	db := atdb.MssqlConnect(mssqlconn)
	fmt.Println(mssqlconn)
	loginfo := GetLoginInfofromPhoneNumber("6281312000300", db)
	fmt.Println(loginfo)

}
