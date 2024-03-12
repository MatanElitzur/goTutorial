package knowledge

import (
	"log/slog"
	"math/rand"
	"os"
	"runtime"
	"strings"
)

type SlogUser struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// With this gunction we will print only the id of the user and not all the data (name, email, password)
func (su SlogUser) LogValue() slog.Value {
	return slog.Int64Value(su.ID)
}

// https://www.youtube.com/watch?v=ReWtK-HUbQQ
func Slog_log1() {
	su := SlogUser{
		ID:       1234,
		Name:     "Mike",
		Email:    "mike.elitzur@gmail.com",
		Password: "qwerty1234",
	}

	slog.Info("Golang rocks!", "version", runtime.Version())
	slog.Error("Gopher has stumbeld!")
	slog.Info("Gopher's getting dizzy!")
	slog.Debug("Debugging!")

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger.Info("Golang rocks!", "version", runtime.Version())
	logger.Error("Gopher has stumbeld!")
	logger.Info("Gopher's getting dizzy!")
	logger.Debug("Debugging!")

	logger1 := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger1)
	slog.Info("Golang rocks!", "version", runtime.Version())
	slog.Error("Gopher has stumbeld!")
	slog.Info("Gopher's getting dizzy!")
	slog.Debug("Debugging!")

	//Set log leverl to debug
	logger2 := slog.New(slog.NewTextHandler(os.Stderr,
		&slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger2)
	slog.Info("Golang rocks!", "version", runtime.Version())
	slog.Error("Gopher has stumbeld!")
	slog.Info("Gopher's getting dizzy!")
	slog.Debug("Debugging!")

	//Musk key values --> Some values we don't want to display in the log, like password
	replaceAttrFunc := func(groups []string, a slog.Attr) slog.Attr {
		keysToMask := []string{"password", "childLogger password"}
		exists := false
		keyToFind := strings.ToLower(a.Key)
		for _, key := range keysToMask {
			if key == keyToFind {
				exists = true
				break
			}
		}
		if exists { //this code will mask the passord value
			a.Value = slog.StringValue("<<MASKED>>")
		}
		return a
	}
	//Set log leverl to debug + JsonHandler --> the output will will be in json
	//All prints will have the app_version in them
	//AddSource will add the file name and the line number of the print to log
	// file, err := os.OpenFile("tmp/slog_demo.log", --> better to get from env variable
	//                          os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666) //0666 -> the file mode
	// if err != nil {
	// 	panic("Error opening the log file")
	// }
	// defer file.Close()
	// Will log only to a file
	// logger3 := slog.New(slog.NewJSONHandler(file,
	// 	&slog.HandlerOptions{
	// 		Level:       slog.LevelDebug,
	// 		ReplaceAttr: replaceAttrFunc,
	// 		AddSource:   true}).WithAttrs([]slog.Attr{slog.String("app_version", "v1.0.0")}))

	//Will log to the screen and to the log file
	//But there is a better way in linux, we eill keep only the print to the screen and
	//then invoke the command go run main.go >> /tmp/slog_demo.log 2>&1
	//This will print also to the file
	// multiWriter := io.MultiWriter(file, os.Stderr)
	// logger3 := slog.New(slog.NewJSONHandler(multiWriter,
	// 	&slog.HandlerOptions{
	// 		Level:       slog.LevelDebug,
	// 		ReplaceAttr: replaceAttrFunc,
	// 		AddSource:   true}).WithAttrs([]slog.Attr{slog.String("app_version", "v1.0.0")}))

	//Will log only to the screen
	logger3 := slog.New(slog.NewJSONHandler(os.Stderr,
		&slog.HandlerOptions{
			Level:       slog.LevelDebug,
			ReplaceAttr: replaceAttrFunc,
			AddSource:   true}).WithAttrs([]slog.Attr{slog.String("app_version", "v1.0.0")}))

	slog.SetDefault(logger3)
	slog.Info("Golang rocks!",
		slog.String("version", runtime.Version()),
		slog.Int("Random number", rand.Int()),
		slog.String("password", "qwerty$"),
		slog.Group("OS Info",
			slog.String("OS", runtime.GOOS),
			slog.Int("CPUs", runtime.NumCPU()),
			slog.String("arch", runtime.GOARCH),
		),
	)

	slog.Error("Gopher has stumbeld!")
	slog.Info("Gopher's getting dizzy!")
	slog.Debug("Debugging!")

	childLogger := slog.With(
		slog.Group("childLogger OS Info",
			slog.String("OS", runtime.GOOS),
			slog.Int("CPUs", runtime.NumCPU()),
			slog.String("arch", runtime.GOARCH),
		),
	)

	childLogger.Info("childLogger Golang rocks!",
		slog.String("childLogger version", runtime.Version()),
		//This password will not be masked casue the key starts with childLogger
		slog.String("childLogger password", "qwerty$"),
	)

	slog.Info("slog, User Info", "user", su)

}
