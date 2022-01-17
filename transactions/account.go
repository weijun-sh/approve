package transactions

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsn-dev/fsn-go-sdk/efsn/accounts/keystore"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/crypto"
)

type keystoreStruct struct {
	ADDRESS string
}

func getKeystoreAddress(keyFile string) (string, error) {
	jsonString, err := getKeystoreJson(keyFile)
	if err == nil {
		var s keystoreStruct
		err = json.Unmarshal([]byte(strings.ToUpper(jsonString)), &s)
		if err == nil {
			ret := fmt.Sprintf("0x%v", strings.ToLower(s.ADDRESS))
			return ret, nil
		}
		//fmt.Printf("getKeystoreAddress(), json.Unmarshal kefile: %v err: %v\n", keyFile, err)
	}
	return "", err
}

func getKeystoreJson(keyFile string) (string, error) {
	kjson, err := ioutil.ReadFile(keyFile)
	if err == nil {
		return string(kjson), nil
		//fmt.Printf("getKeystoreJson(), ioutil.ReadFile keystore: %v err: %v\n", keyFile, err)
	}
	return "", err
}

func getKeyFromKeystore(keyFile, pwdfile string) (*keystore.Key, []byte, error) {
	var (
		kjson    []byte
		kwrapper *keystore.Key
		err      error
	)
	password, err := getFileString(pwdfile)
	if err != nil {
		fmt.Printf("getFileString file: %v fail, err: %v\n", pwdfile, err)
		return nil, []byte(""), err
	}
	kjson, err = ioutil.ReadFile(keyFile)
	if err != nil {
		fmt.Println("Read keystore fail", err)
		return nil, []byte(""), err
	}
	kwrapper, err = keystore.DecryptKey(kjson, password)
	if err != nil {
		fmt.Printf("Key decrypt error: %v\n", err)
		return nil, []byte(""), err
	}
	fmt.Printf("[ OK ] DecryptKey account: %v\n", kwrapper.Address.String())
	return kwrapper, kjson, nil
}

func getFileString(file string) (string, error) {
	dir, exist := isDir(file)
	if exist != true {
		return "", errors.New("--password file not exist!")
	}
	if dir == false { //file
		kjson, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Read keystore %v err: %v\n", file, err)
			return "", errors.New("ReadFile error!")
		}
		passwdString := strings.Replace(string(kjson), "\n", "", -1)
		return passwdString, nil
	} else {
		return "", errors.New("--password is dir, error!")
	}
}

// FileExist checks if a file exists at filePath.
func getFile(path string) ([]string, error) {
	filelist := []string{}
	dir, exist := isDir(path)
	if exist != true {
		rets := fmt.Sprintf("file '%v' not exist", path)
		return filelist, errors.New(rets)
	}
	if dir == false { //file
		filelist = append(filelist, path)
		return filelist, nil
	}
	// dir
	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		if file.IsDir() {
			dDir := filepath.Join(path, file.Name())
			getDirFile, err := getFile(dDir)
			if err == nil {
				for _, fd := range getDirFile {
					filelist = append(filelist, fd)
				}
			}
		} else {
			fl := filepath.Join(path, file.Name())
			filelist = append(filelist, fl)
		}
	}
	return filelist, nil
}

// FileExist checks if a file exists at filePath.
func isDir(path string) (bool, bool) { //parms: (dir, fileExist)
	fi, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false, false
	}
	if !fi.IsDir() {
		return false, true
	}

	return true, true
}

func CheckAccount(keyFile, password_args string) error {
	_, _, err := getKeyFromKeystore(keyFile, password_args)
	return err
}

func getPrivAddr(keyFile, passFile string) (*ecdsa.PrivateKey, common.Address, error) {
        password, err := getFileString(passFile)
        if err != nil {
                fmt.Printf("getFileString file: %v fail, err: %v\n", passFile, err)
                return nil, common.Address{}, err
        }
        keyjson, err := getFileString(keyFile)
        if err != nil {
                fmt.Printf("Read keystore fail, err: %v\n", err)
                return nil, common.Address{}, err
        }
        keyWrapper, err := keystore.DecryptKey([]byte(keyjson), password)
        if err != nil {
                fmt.Printf("Key decrypt err: %v\n", err)
                return nil, common.Address{}, err
        }
        privateKey := keyWrapper.PrivateKey

        publicKey := privateKey.Public()
        publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
        if !ok {
                fmt.Printf("get publicKey err: %v\n", err)
                return nil, common.Address{}, err
        }

        address := crypto.PubkeyToAddress(*publicKeyECDSA)
        fmt.Printf("from: %v\n", address.String())
        return privateKey, address, nil
}
