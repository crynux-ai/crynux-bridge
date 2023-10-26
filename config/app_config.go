package config

const (
	EnvProduction = "production"
	EnvDebug      = "debug"
	EnvTest       = "test"
)

type AppConfig struct {
	Environment string `mapstructure:"environment"`

	Db struct {
		Driver           string `mapstructure:"driver"`
		ConnectionString string `mapstructure:"connection"`
		Log              struct {
			Level       string `mapstructure:"level"`
			Output      string `mapstructure:"output"`
			MaxFileSize int    `mapstructure:"max_file_size"`
			MaxDays     int    `mapstructure:"max_days"`
			MaxFileNum  int    `mapstructure:"max_file_num"`
		} `mapstructure:"log"`
	} `mapstructure:"db"`

	Log struct {
		Level       string `mapstructure:"level"`
		Output      string `mapstructure:"output"`
		MaxFileSize int    `mapstructure:"max_file_size"`
		MaxDays     int    `mapstructure:"max_days"`
		MaxFileNum  int    `mapstructure:"max_file_num"`
	} `mapstructure:"log"`

	Http struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"http"`

	DataDir struct {
		InferenceTasks string `mapstructure:"inference_tasks"`
		ModelImages    string `mapstructure:"model_images"`
	} `mapstructure:"data_dir"`

	Blockchain struct {
		RpcEndpoint string `mapstructure:"rpc_endpoint"`

		Account struct {
			Address        string `mapstructure:"address"`
			PrivateKey     string `mapstructure:"private_key"`
			PrivateKeyFile string `mapstructure:"private_key_file"`
		} `mapstructure:"account"`

		Contracts struct {
			Task        string `mapstructure:"task"`
			Node        string `mapstructure:"node"`
			CrynuxToken string `mapstructure:"crynux_token"`
		} `mapstructure:"contracts"`

		GasLimit      uint64 `mapstructure:"gas_limit"`
		StartBlockNum uint64 `mapstructure:"start_block_num"`
	} `mapstructure:"blockchain"`

	Relay struct {
		BaseURL string `mapstructure:"base_url"`
	} `mapstructure:"relay"`

	TaskSchema struct {
		StableDiffusionInference string `mapstructure:"stable_diffusion_inference"`
	} `mapstructure:"task_schema"`

	Test struct {
		RootAddress    string `mapstructure:"root_address"`
		RootPrivateKey string `mapstructure:"root_private_key"`
	} `mapstructure:"test"`
}
