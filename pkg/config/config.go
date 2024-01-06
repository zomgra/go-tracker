package config

import (
	"errors"
	"flag"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func MustSetConfig(s *ApplicationConfig) {
	configPath, err := FetchConfigPath()
	if err != nil {
		log.WithField("config-path", nil).Fatalf("Problem with fetching config, error: %s", err.Error())
	}
	err = IncludeEnv(configPath)
	if err != nil {
		log.WithField("config-path", configPath).Fatalf("Problem with including env from %s, error: %s", configPath, err.Error())
	}
	dbConfig := &PostgresConfig{}
	err = SetDbConfig(dbConfig)

	if dbConfig == nil {
		log.WithField("config-path", configPath).Fatalf("error: BloomConfig Is Null")
	}
	if err != nil {
		log.WithField("config-path", configPath).Fatalf("error: %s", err.Error())
	}
	bloomConfig := &BloomFilterConfig{}
	err = SetBloomConfig(bloomConfig)
	if err != nil {
		log.WithField("config-path", configPath).Fatalf("error: %s", err.Error())
	}

	if bloomConfig == nil {
		log.WithField("config-path", configPath).Fatalf("error: BloomConfig Is Null")
	}
	port := os.Getenv("TRACKER_APP_PORT")
	if port == "" {
		// os.Exit
		//port = defaultPort
		port = "8080"
	}

	s.Port = port
	s.LogLevel = os.Getenv("TRACKER_LOG_LEVEL")
	s.DbConfig = dbConfig
	s.BloomConfig = bloomConfig

	// Set log level
	logLevel, err := log.ParseLevel(s.LogLevel)
	if err != nil {
		log.SetLevel(log.WarnLevel)
	}
	log.SetLevel(logLevel)
}

func FetchConfigPath() (string, error) {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")

	flag.Parse()
	if res != "" {
		return res, nil
	} else {
		log.Errorf("config arg dont have config path, --config: %s", res)
		res = os.Getenv("TRACKER_CONFIG_PATH")
		log.Warn(res)
		if res == "" {
			log.Errorf("ENV `TRACKER_CONFIG_PATH` is havent path to config: %s", os.Getenv("TRACKER_CONFIG_PATH"))
			return "", errors.New("Can`t find config path")
		}
	}
	if res != "" {
		return res, nil
	}
	return "", errors.New("Can`t find config path")
}

func IncludeEnv(configPath string) error {
	log.Debugf("Config file path: %s", configPath)
	err := godotenv.Load(configPath)

	if err != nil {
		return err
	}
	return nil
}
func SetDbConfig(dbconf *PostgresConfig) error {
	connectionUrl := os.Getenv("TRACKER_CONNECTION_URI")
	maxIdleConn, err := strconv.Atoi(os.Getenv("TRACKER_MAX_IDLE_CONN"))

	if err != nil {
		return errors.New("Problem on parsing `TRACKER_MAX_IDLE_CONN` error: " + err.Error())
	}
	maxOpenConn, err := strconv.Atoi(os.Getenv("TRACKER_MAX_OPEN_CONN"))
	if err != nil {
		return errors.New("Problem on parsing `TRACKER_MAX_OPEN_CONN` error: " + err.Error())
	}
	dbconf.ConnectionUrl = connectionUrl
	dbconf.MaxIdleConn = maxIdleConn
	dbconf.MaxOpenConn = maxOpenConn

	return nil
}

func SetBloomConfig(bloom *BloomFilterConfig) error {
	bloomLimit, err := strconv.Atoi(os.Getenv("TRACKER_BLOOM_LIMIT"))

	if err != nil {
		return errors.New("Problem on parsing \"TRACKER_MAX_OPEN_CONN\" error: " + err.Error())
	}

	bloom.BloomLimit = uint(bloomLimit)

	return nil
}
