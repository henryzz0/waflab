package rule

type Testset struct {
	Meta  *Meta   `yaml:"meta"`
	Tests []*Test `yaml:"tests"`
}

type Meta struct {
	Author      string `yaml:"author"`
	Enabled     bool   `yaml:"enabled"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Test struct {
	Title  string            `yaml:"test_title"`
	Desc   string            `yaml:"desc"`
	Stages []*StageContainer `yaml:"stages"`
}

type StageContainer struct {
	Stage *Stage `yaml:"stage"`
}

type Stage struct {
	Input  *Input  `yaml:"input"`
	Output *Output `yaml:"output"`
}

type Input struct {
	DestAddr string            `yaml:"dest_addr"`
	Port     int               `yaml:"port"`
	Method   string            `yaml:"method"`
	Uri      string            `yaml:"uri"`
	Version  string            `yaml:"version"`
	Headers  map[string]string `yaml:"headers"`
}

type Output struct {
	LogContains   string `yaml:"log_contains"`
	NoLogContains string `yaml:"no_log_contains"`
}
