package userlib

import (
	"io/ioutil"
	"math/rand"
	"path"
	"time"
)

// We are declaring a function type to make the code below much cleaner.
type fileReader func(string, string)([]byte, error)

// These are constants which the autograder will be using to validate that your code is outputting the correct values.
// Make sure you call these instead of copying in the messages and codes since they could change!
const (
	FILEERRORCODE = 404
	FILEERRORMSG = "File Read Error"
	TIMEOUTERRORCODE = 408
	SUCCESSCODE = 200
	CapacityString = "Cache status:  # of entries (%v)\ntotal bytes occupied by entries [%v]\nmax allowed capacity {%v}\n"
	TimeoutString = "The file request timed out!\n"
	CacheCloseMessage = "The cache has been cleared!\n"
	ContextType = "Content-Type"
)

// We are setting a private function for the default functionality of the
var f fileReader = func(workingDir, filename string)(data []byte, err error){
	// We want to just append the working dir to the file name so we are just stripping out the leading current
	// directory marker (./) and combining the dir with filename.
	filepath := GetRealFilePath(workingDir, filename)
	// We are just emulating a slower access to disk to make it clear that we are caching data.
	time.Sleep(time.Duration(rand.Float64() * 3 + 1) * time.Second)
	// We finally do the file read which should not take too long.
	data, err = ioutil.ReadFile(filepath)
	return
}

// This is th function which you will be calling to read the file on disk.
func ReadFile(workingDir, filename string)(data []byte, err error){
	// We are calling the function f which is defined above. This is to make it easy for the autograder to hook with
	// your code. It also makes it much easier for you to write tests as well.
	data, err = f(workingDir, filename)
	return
}

// We use this function in our tests to set the read function to a new function. This will make it so subsequent calls
// to ReadFile will be made to the newFunc.
func ReplaceReadFile(newfunc func(string, string)([]byte, error)){
	f = newfunc
}

// This function abstracts away combining the current working dir for the files and the path to the file you want.
// Only the userlib should need to call this function.
func GetRealFilePath(workingDir, filename string) string {
	if filename[0:2] == "./" {
		filename = filename[2:]
	}
	if workingDir[len(workingDir) - 1:] != "/" {
		workingDir = workingDir + "/"
	}
	return workingDir + filename
}

// This function will get the context type based off of its extension. It is used to make sure the http response has
// the correct Context Type set.
func GetContentType(filename string) (string) {
	extension := path.Ext(filename)
	switch extension {
	case ".htm":
		fallthrough
	case ".html":
		return "text/html"
	case ".jpeg":
		fallthrough
	case ".jpg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".pdf":
		return "application/pdf"
	default:
		return "text/plain; charset=utf-8"
	}
}