package config

import (
	"fmt"

	"github.com/spf13/viper"
	_ "go.uber.org/zap"

	_ "gopkg.in/yaml.v3"
)

// Config представляет структуру конфигурации, соответствующую YAML файлу
type Config struct {
	Server struct {
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
		Protocol string `yaml:"protocol"`
	} `yaml:"server"`

	Database struct {
		Driver   string `yaml:"driver"`
		URL      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		MaxConn  int    `yaml:"max_connections"`
	} `yaml:"database"`

	Logging struct {
		Level        string `yaml:"level"`
		RuSitennIdel string `yaml:"ru_sitenn_idel"`
	} `yaml:"logging"`

	Security struct {
		Whitelist      []string `yaml:"whitelist"`
		AdminOnly      []string `yaml:"admin_only"`
		SuperAdminOnly []string `yaml:"super_admin_only"`
		JWT            struct {
			Secret     string `yaml:"secret"`
			Expiration int    `yaml:"expiration"`
		}
	} `yaml:"security"`

	Cors struct {
		AllowedOrigin        string `yaml:"allowedOrigin"`
		AllowedHeader        string `yaml:"allowedHeader"`
		AllowedMethod        string `yaml:"allowedMethod"`
		ConfigurationPattern string `yaml:"configurationPattern"`
	} `yaml:"cors"`

	Email struct {
		Host              string `yaml:"host"`
		Port              int    `yaml:"port"`
		Username          string `yaml:"username"`
		Password          string `yaml:"password"`
		TransportProtocol string `yaml:"transport_protocol"`
		SmtpAuth          bool   `yaml:"smtp_auth"`
		StartTLS          bool   `yaml:"starttls_enable"`
		Debug             bool   `yaml:"debug"`
	} `yaml:"email"`

	Feedback struct {
		Sender                string   `yaml:"sender"`
		Subject               string   `yaml:"subject"`
		EmailTemplateLocation string   `yaml:"email_template_location"`
		Languages             []string `yaml:"languages"`
	} `yaml:"feedback"`

	Newsletter struct {
		Sender                string   `yaml:"sender"`
		Languages             []string `yaml:"languages"`
		EmailTemplateLocation string   `yaml:"email_template_location"`
		UnsubscribeURLBase    string   `yaml:"unsubscribe_url_base"`
		SiteURLBase           string   `yaml:"site_url_base"`
	} `yaml:"newsletter"`

	Google struct {
		Recaptcha struct {
			URL            string  `yaml:"url"`
			Secret         string  `yaml:"secret"`
			ScoreThreshold float64 `yaml:"score_threshold"`
		} `yaml:"recaptcha"`
	} `yaml:"google"`

	SEO struct {
		Sitemap struct {
			Domain     string   `yaml:"domain"`
			StaticURLs []string `yaml:"static_urls"`
		} `yaml:"sitemap"`
		Product struct {
			URLPattern string `yaml:"url_pattern"`
		} `yaml:"product"`
		Category struct {
			URLPattern string `yaml:"url_pattern"`
		} `yaml:"category"`
		News struct {
			URLPattern string `yaml:"url_pattern"`
		} `yaml:"news"`
	} `yaml:"seo"`
}

func LoadConfig(configPath string) (*Config, error) {
	// Настроим viper
	viper.SetConfigFile(configPath) // Указание пути к конфигу
	viper.SetConfigType("yaml")     // Указание типа файла (YAML)
	viper.AutomaticEnv()            // Автоматически связываем переменные окружения с конфигом

	// Читаем конфигурационный файл
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("ошибка при чтении конфигурационного файла: %w", err)
	}

	// Прочитаем и десериализуем данные в структуру Config
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("ошибка при распаковке конфигурации: %w", err)
	}

	// Возвращаем конфигурацию
	return &config, nil
}
