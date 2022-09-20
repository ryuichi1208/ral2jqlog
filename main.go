package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

type AuditLog struct {
	MessageType         string   `json:"messageType"`
	Owner               string   `json:"owner"`
	LogGroup            string   `json:"logGroup"`
	LogStream           string   `json:"logStream"`
	SubscriptionFilters []string `json:"subscriptionFilters"`
	LogEvents           []struct {
		ID        string `json:"id"`
		Timestamp int64  `json:"timestamp"`
		Message   string `json:"message"`
	} `json:"logEvents"`
}

type QueryLog struct {
	TimeStamp string `json:"timeStamp"`
	User      string `json:"user"`
	Host      string `json:"host"`
	Command   string `json:"command"`
	Query     string `json:"query"`
}

type options struct {
	File string `short:"f" long:"file" description:"audit log file" required:"true"`
	Type string `short:"t" long:"type" description:"File Content Type" required:"false"`
}

func message2CSV(csvString string, w io.Writer) error {
	arr := strings.Split(csvString, ",")

	var queryLog = QueryLog{}
	queryLog.TimeStamp = arr[0]
	queryLog.User = arr[2]
	queryLog.Host = arr[3]
	queryLog.Command = arr[3]
	s := strings.Join(arr[8:len(arr)-1], ",")[1:]
	queryLog.Query = s[:len(s)-1]

	outputJson, err := json.Marshal(&queryLog)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "%s\n", string(outputJson))

	return nil
}

func auditLog2Json(jsonString string) error {
	var auditlog AuditLog
	if err := json.Unmarshal([]byte(jsonString), &auditlog); err != nil {
		fmt.Println(err)
		return err
	}
	for _, v := range auditlog.LogEvents {
		if err := message2CSV(v.Message, os.Stdout); err != nil {
			return err
		}
	}

	return nil
}

func readGzip(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	r, err := gzip.NewReader(br)
	if err != nil {
		return err
	}
	defer r.Close()

	for {
		r.Multistream(false)
		if data, err := ioutil.ReadAll(r); err == nil {
			auditLog2Json(string(data))
		}

		if err := r.Reset(br); err != nil {
			break
		}
	}

	return err
}

func DetectFileContentType(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	buffer := make([]byte, 512)
	f.Read(buffer)
	contentType := http.DetectContentType(buffer)
	f.Seek(0, 0)

	return contentType, err
}

func Do(opts options) error {
	var err error
	switch opts.Type {
	case "gzip":
		err = readGzip(opts.File)
	default:
		t, err := DetectFileContentType(opts.File)
		if err != nil {
			return err
		}
		switch t {
		case "application/x-gzip":
			err = readGzip(opts.File)
		default:
			return fmt.Errorf("No Support File Type")
		}
	}
	return err
}

func main() {
	var opts options
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(1)
	}

	if Do(opts) != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
