package api

import (
	// . "models"
	"samh_common_lib/base"
	"samh_common_lib/utils/log"
)

func subDiamond(uid int64, diamond int) (retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	// var (
	// err     error
	// exist   bool
	// nowTime int64 = time.Now().Unix()
	// )

	// _, err := ManhuaDB.Sql("UPDATE user_count SET Cdiamonds=Cdiamonds-? WHERE CUid=?",
	// diamond, uid).
	// Execute()
	// if err != nil {
	// log.Error(err)
	// retCode = base.SamhResponseCode_ServerError
	// return
	// }

	return
}
