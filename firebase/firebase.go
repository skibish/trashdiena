package firebase

import (
	"encoding/json"
	"io/ioutil"

	"github.com/zabawaba99/firego"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Firebase describes firebase structure
type Firebase struct {
	realFirebase *firego.Firebase
}

// New return new instance of the Firebase
func New(firebaseDB, secretsFilePath string) (fbase *Firebase, err error) {
	d, err := ioutil.ReadFile(secretsFilePath)
	if err != nil {
		return
	}

	conf, err := google.JWTConfigFromJSON(d, "https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/firebase.database")
	if err != nil {
		return
	}

	fbase = &Firebase{
		realFirebase: firego.New(firebaseDB, conf.Client(oauth2.NoContext)),
	}

	return
}

// Set sets data to the database
func (f *Firebase) Set(path string, v interface{}) (err error) {
	ref, err := f.realFirebase.Ref(path)
	if err != nil {
		return
	}
	err = ref.Set(v)

	return
}

// Get return value by path from the database
func (f *Firebase) Get(path string) (result json.RawMessage, err error) {

	ref, err := f.realFirebase.Ref(path)
	if err != nil {
		return
	}

	err = ref.Value(&result)

	return
}
