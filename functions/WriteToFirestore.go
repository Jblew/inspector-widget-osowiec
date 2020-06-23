package function

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Jblew/inspector-widget-osowiec/functions/util"
)

// WriteToFirestore publishes entry to firestore document
func WriteToFirestore(writer http.ResponseWriter, req *http.Request) {
	out, err := handleWriteToFirestore(writer, req)
	if err != nil {
		writer.WriteHeader(500)
		util.SendJSONError(writer, fmt.Sprintf("%v", err))
		return
	}

	fmt.Fprintln(writer, out)
}

func handleWriteToFirestore(w http.ResponseWriter, req *http.Request) (string, error) {
	log.Println(util.DebugHTTPRequest(req))

	contents, err := getRequestBody(req)
	if err != nil {
		return "", err
	}

	secretKey, err := getRequestSecret(req)
	if err != nil {
		return "", err
	}

	secretValid := application.CheckSecret(secretKey)
	if secretValid == false {
		return "", fmt.Errorf("Invalid request secret '%s'", secretKey)
	}

	err = application.WriteToFirestore(context.Background(), contents)
	if err != nil {
		return "", err
	}

	return "Success", nil
}

func getRequestBody(req *http.Request) (string, error) {
	entryBin, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", fmt.Errorf("Cannot read request body: %v", err)
	}

	entry := string(entryBin)
	if len(entry) == 0 {
		return "", fmt.Errorf("You must provide entry to publish as text in request body")
	}

	return entry, nil
}

func getRequestSecret(req *http.Request) (string, error) {
	path := req.URL.Path
	pathElems := strings.Split(path, "/")
	lastPathElem := pathElems[len(pathElems)-1]
	return lastPathElem, nil
}
