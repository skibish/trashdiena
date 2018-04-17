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

// IFirebase is an interface that Firebase is implementing
type IFirebase interface {
	Set(path string, v interface{}) (err error)
	Get(path string) (result json.RawMessage, err error)
	FilterEqual(path, field string, value interface{}) (result json.RawMessage, err error)
	Delete(path string) (err error)
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

// Delete deletes record from the database
func (f *Firebase) Delete(path string) (err error) {
	ref, err := f.realFirebase.Ref(path)
	if err != nil {
		return
	}

	err = ref.Remove()

	return
}

// FilterEqual filter records with field equal to specific value
func (f *Firebase) FilterEqual(path, field string, value interface{}) (result json.RawMessage, err error) {
	ref, err := f.realFirebase.Ref(path)
	if err != nil {
		return
	}

	err = ref.OrderBy(field).EqualToValue(value).Value(&result)

	return
}
