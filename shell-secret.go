package shellsecret

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/pbkdf2"
)

var (
	envKey         = "SECURE_SHELL_KEY"
	ErrEnvNotSetup = fmt.Errorf("environment not setup")

	pbkdf2Password = []byte("shell-secret")
)

type ShellSecret interface {
	Add(key string, value interface{}) error
	Delete(key string) error
	Get(key string, value interface{}) error
	List() ([]string, error)
}

func Setup() error {
	keyBytes := make([]byte, 32)
	_, err := rand.Read(keyBytes)
	if err != nil {
		return err
	}
	key := base64.RawStdEncoding.EncodeToString(keyBytes)
	printEnv(key)
	return nil
}

func New() (ShellSecret, error) {
	key := os.Getenv(envKey)
	if key == "" {
		return nil, ErrEnvNotSetup
	}

	keyBytes, err := base64.RawStdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return &shellSecret{
		secureFilePath: path.Join(os.TempDir(), fmt.Sprintf("secureshell%d.json", os.Getppid())),
		key:            keyBytes,
	}, nil
}

type shellSecret struct {
	secureFilePath string
	key            []byte
}

func (s *shellSecret) load() (map[string]interface{}, error) {
	res := make(map[string]interface{})
	contentEncrypted, err := ioutil.ReadFile(s.secureFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return res, nil
		}
	}

	var block cipher.Block
	if block, err = aes.NewCipher(s.key); err != nil {
		return nil, fmt.Errorf("couldn't initialize the dek: %w", err)
	}
	gcm, err := cipher.NewGCMWithNonceSize(block, aes.BlockSize)
	if err != nil {
		return nil, fmt.Errorf("failed to create gcm: %w", err)
	}

	derivedIV := pbkdf2.Key(pbkdf2Password, nil, 2000, aes.BlockSize, sha256.New)
	data, err := gcm.Open(nil, derivedIV, contentEncrypted, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open gcm data: %w", err)
	}
	return res, json.Unmarshal(data, &res)
}

func (s *shellSecret) save(input map[string]interface{}) error {
	b, err := json.Marshal(&input)
	if err != nil {
		return err
	}
	var block cipher.Block
	if block, err = aes.NewCipher(s.key); err != nil {
		return fmt.Errorf("couldn't initialize the dek: %w", err)
	}
	gcm, err := cipher.NewGCMWithNonceSize(block, aes.BlockSize)
	if err != nil {
		return fmt.Errorf("failed to create gcm: %w", err)
	}
	derivedIV := pbkdf2.Key(pbkdf2Password, nil, 2000, aes.BlockSize, sha256.New)

	contentEncrypted := gcm.Seal(nil, derivedIV, b, nil)

	return ioutil.WriteFile(s.secureFilePath, contentEncrypted, os.ModePerm)

}

func (s *shellSecret) Add(key string, value interface{}) error {
	m, err := s.load()
	if err != nil {
		return err
	}
	fmt.Printf("map %v\n", m)
	m[key] = value
	return s.save(m)
}
func (s *shellSecret) Delete(key string) error {
	m, err := s.load()
	if err != nil {
		return fmt.Errorf("failed to load: %w", err)
	}
	delete(m, key)
	return s.save(m)
}

func (s *shellSecret) Get(key string, value interface{}) error {
	m, err := s.load()
	if err != nil {
		return fmt.Errorf("failed to load: %w", err)
	}
	v, found := m[key]
	if !found {
		return fmt.Errorf("key not found: %s", key)
	}
	return mapstructure.Decode(v, value)
}

func (s *shellSecret) List() ([]string, error) {
	m, err := s.load()
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys, nil
}
