package connector

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

func BuildConfig() *mysql.Config {
	userConfig := mysql.NewConfig()

	// Config Variables
	// Assign Mandatory Values Attached in settings package
	if os.Getenv("User") != "" {
		userConfig.User = os.Getenv("User")
	} else {
		panic("User can not be empty to create db connection")
	}
	if os.Getenv("Passwd") != "" {
		userConfig.Passwd = os.Getenv("Passwd")
	} else {
		panic("Passwd can not be empty to create db connection")
	}
	if os.Getenv("Addr") != "" {
		userConfig.Addr = os.Getenv("Addr")
	} else {
		panic("Addr can not be empty to create db connection")
	}
	if os.Getenv("DBName") != "" {
		userConfig.DBName = os.Getenv("DBName")
	} else {
		panic("DBName can not be empty to create db connection")
	}
	if os.Getenv("Net") != "" {
		userConfig.Net = os.Getenv("Net")
	} else {
		panic("Net can not be empty to create db connection")
	}

	// Assign Values if specified in settings package else use default
	if os.Getenv("Collation") != "" {
		userConfig.Collation = os.Getenv("Collation")
	}
	if os.Getenv("Loc") != "" {
		location, err := time.LoadLocation(os.Getenv("Loc"))
		if err == nil {
			userConfig.Loc = location
		} else {
			panic("Invalid Loc assigned to create db connection")
		}
	}
	if os.Getenv("MaxAllowedPacket") != "" {
		MaxAllowedPacket, err := strconv.Atoi(os.Getenv("MaxAllowedPacket"))
		if err == nil {
			userConfig.MaxAllowedPacket = MaxAllowedPacket
		} else {
			panic("Invalid Value for MaxAllowedPacket")
		}
	}
	if os.Getenv("ServerPubKey") != "" {
		userConfig.ServerPubKey = os.Getenv("ServerPubKey")
	}
	if os.Getenv("TLSConfig") != "" {
		userConfig.TLSConfig = os.Getenv("TLSConfig")
	}
	if os.Getenv("Timeout") != "" {
		timeout, err := strconv.ParseInt(os.Getenv("Timeout"), 10, 64)
		if err != nil {
			panic("Invalid Timeout provided to connect to DB")
		}
		userConfig.Timeout = time.Duration(timeout) * time.Second
	}
	if os.Getenv("ReadTimeout") != "" {
		ReadTimeout, err := strconv.ParseInt(os.Getenv("ReadTimeout"), 10, 64)
		if err != nil {
			panic("Invalid ReadTimeout provided to connect to DB")
		}
		userConfig.ReadTimeout = time.Duration(ReadTimeout) * time.Second
	}
	if os.Getenv("WriteTimeout") != "" {
		WriteTimeout, err := strconv.ParseInt(os.Getenv("WriteTimeout"), 10, 64)
		if err != nil {
			panic("Invalid WriteTimeout provided to connect to DB")
		}
		userConfig.WriteTimeout = time.Duration(WriteTimeout) * time.Second
	}
	if os.Getenv("AllowAllFiles") != "" {
		AllowAllFiles, err := strconv.ParseBool(os.Getenv("AllowAllFiles"))
		if err != nil {
			panic("Invalid AllowAllFiles provided.")
		}
		userConfig.AllowAllFiles = AllowAllFiles
	}
	if os.Getenv("AllowCleartextPasswords") != "" {
		AllowCleartextPasswords, err := strconv.ParseBool(os.Getenv("AllowCleartextPasswords"))
		if err != nil {
			panic("Invalid AllowCleartextPasswords provided.")
		}
		userConfig.AllowCleartextPasswords = AllowCleartextPasswords
	}
	if os.Getenv("AllowNativePasswords") != "" {
		AllowNativePasswords, err := strconv.ParseBool(os.Getenv("AllowNativePasswords"))
		if err != nil {
			panic("Invalid AllowNativePasswords provided.")
		}
		userConfig.AllowNativePasswords = AllowNativePasswords
	}
	if os.Getenv("AllowOldPasswords") != "" {
		AllowOldPasswords, err := strconv.ParseBool(os.Getenv("AllowOldPasswords"))
		if err != nil {
			panic("Invalid AllowOldPasswords provided.")
		}
		userConfig.AllowOldPasswords = AllowOldPasswords
	}
	if os.Getenv("CheckConnLiveness") != "" {
		CheckConnLiveness, err := strconv.ParseBool(os.Getenv("CheckConnLiveness"))
		if err != nil {
			panic("Invalid CheckConnLiveness provided.")
		}
		userConfig.CheckConnLiveness = CheckConnLiveness
	}
	if os.Getenv("ClientFoundRows") != "" {
		ClientFoundRows, err := strconv.ParseBool(os.Getenv("ClientFoundRows"))
		if err != nil {
			panic("Invalid ClientFoundRows provided.")
		}
		userConfig.ClientFoundRows = ClientFoundRows
	}
	if os.Getenv("ColumnsWithAlias") != "" {
		ColumnsWithAlias, err := strconv.ParseBool(os.Getenv("ColumnsWithAlias"))
		if err != nil {
			panic("Invalid ColumnsWithAlias provided.")
		}
		userConfig.ColumnsWithAlias = ColumnsWithAlias
	}
	if os.Getenv("InterpolateParams") != "" {
		InterpolateParams, err := strconv.ParseBool(os.Getenv("InterpolateParams"))
		if err != nil {
			panic("Invalid InterpolateParams provided.")
		}
		userConfig.InterpolateParams = InterpolateParams
	}
	if os.Getenv("MultiStatements") != "" {
		MultiStatements, err := strconv.ParseBool(os.Getenv("MultiStatements"))
		if err != nil {
			panic("Invalid MultiStatements provided.")
		}
		userConfig.MultiStatements = MultiStatements
	}
	if os.Getenv("ParseTime") != "" {
		ParseTime, err := strconv.ParseBool(os.Getenv("ParseTime"))
		if err != nil {
			panic("Invalid ParseTime provided.")
		}
		userConfig.ParseTime = ParseTime
	}
	if os.Getenv("RejectReadOnly") != "" {
		RejectReadOnly, err := strconv.ParseBool(os.Getenv("RejectReadOnly"))
		if err != nil {
			panic("Invalid RejectReadOnly provided.")
		}
		userConfig.RejectReadOnly = RejectReadOnly
	}

	fmt.Print(userConfig)
	return userConfig
}
