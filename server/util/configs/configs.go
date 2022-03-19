package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joeshaw/envdecode"
)

type Configs struct {
	DbConf       *DbConf
	ServerConf   *ServerConf
	PasswordConf *PasswordConf
	JwtConf      *JwtConf
}

type ServerConf struct {
	Port int `env:"SERVER_PORT,required"`
}

type DbConf struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DbName   string `env:"DB_NAME,required"`
}

type PasswordConf struct {
	Time    uint32 `env:"ARGON_2_TIME,required"`
	Memory  uint32 `env:"ARGON_2_MEMORY,required"`
	Threads uint8  `env:"ARGON_2_THREADS,required"`
	KeyLen  uint32 `env:"ARGON_2_KEY_LENGTH,required"`
}

type JwtConf struct {
	SecretKey string `env:"JWT_SECRET_KEY,required"`
}

func LoadConfigs() (*Configs, error) {
	var serverConf ServerConf
	var passwordConf PasswordConf
	var jwtConf JwtConf
	var dbConf DbConf

	if err := envdecode.Decode(&serverConf); err != nil {
		return nil, err
	}

	if err := envdecode.Decode(&passwordConf); err != nil {
		return nil, err
	}

	if err := envdecode.Decode(&jwtConf); err != nil {
		return nil, err
	}

	if err := envdecode.Decode(&dbConf); err != nil {
		return nil, err
	}

	return &Configs{
		ServerConf:   &serverConf,
		PasswordConf: &passwordConf,
		JwtConf:      &jwtConf,
		DbConf:       &dbConf,
	}, nil
}

func ReadPasswordEnvConfigs() *PasswordConf {
	time, _ := strconv.ParseUint(os.Getenv("ARGON_2_TIME"), 10, 64)
	memory, _ := strconv.ParseUint(os.Getenv("ARGON_2_MEMORY"), 10, 64)
	threads, _ := strconv.ParseUint(os.Getenv("ARGON_2_THREADS"), 10, 64)
	keyLen, _ := strconv.ParseUint(os.Getenv("ARGON_2_KEY_LENGTH"), 10, 64)

	return &PasswordConf{
		Time:    uint32(time),
		Memory:  uint32(memory),
		Threads: uint8(threads),
		KeyLen:  uint32(keyLen),
	}
}

func (c *DbConf) ToDsnString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		c.Host, c.Username, c.Password, c.DbName, c.Port)
}
